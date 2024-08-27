package handler

import (
	"database/sql"
	"io"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdateAt    time.Time `json:"updateAt"`
}

type PageData struct {
	Title string `json:"title"`
	Tasks []Task `json:"tasks"`
}

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Run(e *echo.Echo, db *sql.DB) {

	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	e.Renderer = t
	e.Static("/static", "static")
	e.Use(middleware.Logger())

	e.GET("/tasks", func(c echo.Context) error {
		tasks, err := allTasks(db)
		if err != nil {
			log.Println("Error in func allTask: ", err)
			return c.String(http.StatusInternalServerError, "Error get tasks")
		}

		data := PageData{
			Title: "TaskTracker",
			Tasks: tasks,
		}

		return c.Render(http.StatusOK, "layout.html", data)
	})

	e.POST("/tasks/add", func(c echo.Context) error {
		if err := addTask(c, db); err != nil {
			log.Println("Error add task: ", err)
			return c.String(http.StatusInternalServerError, "Error add task")
		}

		return c.Redirect(http.StatusSeeOther, "/tasks")
	})

	e.POST("/tasks/updateDescription/:id", func(c echo.Context) error {
		if err := updateTask(c, db); err != nil {
			log.Println("Error update Task: ", err)
			return c.String(echo.ErrInternalServerError.Code, "Error update task")
		}
		return c.Redirect(http.StatusSeeOther, "/tasks")
	})

	e.POST("/tasks/updateStatus/:id", func(c echo.Context) error {
		if err := changeStatus(c, db); err != nil {
			log.Println("Error change status: ", err)
			return c.String(echo.ErrInternalServerError.Code, "Error change status")
		}

		return c.Redirect(http.StatusSeeOther, "/tasks")
	})

	e.POST("/tasks/delete/:id", func(c echo.Context) error {
		if err := deleteTask(c, db); err != nil {
			log.Println("Error delete task: ", err)
			return c.String(echo.ErrInternalServerError.Code, "Error delete task")
		}
		return c.Redirect(http.StatusSeeOther, "/tasks")
	})

	e.Start(":2222")

}

func allTasks(db *sql.DB) ([]Task, error) {
	req := "SELECT * FROM tasks"

	rows, err := db.Query(req)
	if err != nil {
		return nil, err
	}

	tasks := []Task{}

	for rows.Next() {
		task := Task{}
		if err = rows.Scan(&task.Id, &task.Description, &task.Status, &task.CreatedAt, &task.UpdateAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func addTask(c echo.Context, db *sql.DB) error {
	req := "INSERT INTO tasks (description, status) VALUES ($1, $2)"

	des := c.FormValue("description")
	status := c.FormValue("status")

	if _, err := db.Exec(req, des, status); err != nil {
		return err
	}

	return nil
}

func updateTask(c echo.Context, db *sql.DB) error {
	idTask := c.Param("id")
	newDes := c.FormValue("description")
	newTime := time.Now()
	req := "UPDATE tasks SET description = $1, updateAt = $2 WHERE id = $3"
	if _, err := db.Exec(req, newDes, newTime, idTask); err != nil {
		return err
	}

	return nil
}

func changeStatus(c echo.Context, db *sql.DB) error {
	idTask := c.Param("id")
	newStatus := c.FormValue("status")
	req := "UPDATE tasks SET status = $1 WHERE id = $2"

	if _, err := db.Exec(req, newStatus, idTask); err != nil {
		return err
	}

	return nil
}

func deleteTask(c echo.Context, db *sql.DB) error {
	idTask := c.Param("id")
	req := "DELETE FROM tasks WHERE id = $1"
	if _, err := db.Exec(req, idTask); err != nil {
		return err
	}

	return nil
}
