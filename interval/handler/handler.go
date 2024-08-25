package handler

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type Task struct {
	Id          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdateAt    time.Time `json:"updateAt"`
}

func Run(e *echo.Echo, db *sql.DB) {

	e.GET("/tasks", func(c echo.Context) error {
		tasks, err := allTasks(db)
		if err != nil {
			log.Println("Error in func allTask: ", err)
			return c.String(http.StatusInternalServerError, "Error get tasks")
		}

		err = c.JSON(http.StatusOK, tasks)
		if err != nil {
			log.Println("Error json tasks")
			return c.String(echo.ErrInternalServerError.Code, "Error get tasks")
		}

		return c.String(http.StatusOK, "")
	})

	e.POST("/addTask", func(c echo.Context) error {
		if err := addTask(c, db); err != nil {
			log.Println("Error add task: ", err)
			return c.String(http.StatusInternalServerError, "Error add task")
		}
		
		return c.String(http.StatusCreated, "Add task successful")
	})

	e.PUT("/changeTask/:id", func(c echo.Context) error {
		if err := updateTask(c, db); err != nil {
			log.Println("Error update Task: ", err)
			return c.String(echo.ErrInternalServerError.Code, "Error update task")
		}
		return c.String(http.StatusOK, "Change successful")
	})

	e.PUT("/changeStatus/:id", func(c echo.Context) error {
		if err := changeStatus(c, db); err != nil {
			log.Println("Error change status: ", err)
			return c.String(echo.ErrInternalServerError.Code, "Error change status")
		}
		
		return c.String(http.StatusOK, "Change successful")
	})

	e.DELETE("/deleteTask/:id", func(c echo.Context) error {
		if err := deleteTask(c, db); err != nil {
			log.Println("Error delete task: ", err)
			return c.String(echo.ErrInternalServerError.Code, "Error delete task")
		}
		return c.String(http.StatusOK, "Delete successful")
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
	task := Task{}
	if err := c.Bind(&task); err != nil {
		return err
	}

	if _, err := db.Exec(req, task.Description, task.Status); err != nil {
		return err
	}

	return nil
}

func updateTask(c echo.Context, db *sql.DB) error {
	task := Task{}
	idTask := c.Param("id")
	if err := c.Bind(&task); err != nil {
		return err
	}
	
	req := "UPDATE tasks SET description = $1 WHERE id = $2"
	if _, err := db.Exec(req, task.Description, idTask); err != nil {
		return err
	}

	return nil
}

func changeStatus(c echo.Context, db *sql.DB) error {
	task := Task{}
	idTask := c.Param("id")
	req := "UPDATE tasks SET status = $1 WHERE id = $2"
	if err := c.Bind(&task); err != nil {
		return err
	}

	if _, err := db.Exec(req, task.Status, idTask); err != nil {
		return nil
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