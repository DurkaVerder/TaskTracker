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

		c.JSON(http.StatusOK, tasks)

		return nil
	})

	e.POST("/addTask", func(c echo.Context) error {

		return nil
	})

	e.PUT("/changeTask", func(c echo.Context) error {

		return nil
	})

	e.DELETE("/deleteTask", func(c echo.Context) error {

		return nil
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
