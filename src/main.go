package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"todoapp/src/model"
)

const dataSource  = "root@tcp(127.0.0.1:3314)/todo_app"

func main() {
	router := gin.Default()
	router.Static("/assets", "assets")
	router.LoadHTMLGlob("template/*.html")

	// 初期表示処理.
	router.GET("/", func(ctx *gin.Context){
		unCompletedTodo := getUncompletedTodos()
		completedTodo := getCompletedTodos()

		ctx.HTML(200, "index.html", gin.H{
			"unCompleted": unCompletedTodo,
			"completed": completedTodo,
		})
	})

	// タスク新規登録API.
	router.POST("/api/todo", func(ctx *gin.Context) {
		name := ctx.PostForm("name")
		deadline := ctx.PostForm("deadline")

		db, err := sql.Open("mysql", dataSource)
		if err != nil {
			fmt.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, err.Error())
		} else {
			fmt.Println("DB connected")
		}

		defer db.Close()

		row, err := db.Exec("INSERT INTO todos (task_name, deadline) VALUES (?, ?)", name, deadline)
		if err != nil {
			fmt.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, err.Error())
		} else {
			id, err := row.LastInsertId()
			if err != nil {
				fmt.Println(err.Error())
				ctx.JSON(http.StatusInternalServerError, err.Error())
			} else {
				ctx.JSON(http.StatusOK, id)
			}
		}
	})

	// ステータス更新API.
	router.POST("/api/todo/:id/complete", func(ctx *gin.Context) {
		id := ctx.Param("id")

		db, err := sql.Open("mysql", dataSource)
		if err != nil {
			fmt.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, err.Error())
		} else {
			fmt.Println("DB connected")
		}

		defer db.Close()

		row, err := db.Exec("UPDATE todos SET status = ? WHERE id = ?", 1, id)
		if err != nil {
			fmt.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, err.Error())
		} else {
			ctx.JSON(http.StatusOK, row)
		}
	})

	router.Run()
}

// 未完了タスク取得.
func getUncompletedTodos() []model.Todo {
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB connected")
	}

	defer db.Close()

	rows, err := db.Query("SELECT id, task_name, deadline FROM todos WHERE status = 0")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	var todos []model.Todo
	for rows.Next() {
		todo := model.Todo{}
		err := rows.Scan(&todo.ID, &todo.Name, &todo.Deadline)

		if err != nil {
			panic(err.Error())
		}

		todos = append(todos, todo)
	}

	return todos
}

// 完了タスク取得.
func getCompletedTodos() []model.Todo {
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("DB connected")
	}

	defer db.Close()

	rows, err := db.Query("SELECT id, task_name, deadline FROM todos WHERE status = 1")
	defer rows.Close()
	if err != nil {
		panic(err.Error())
	}

	var todos []model.Todo
	for rows.Next() {
		todo := model.Todo{}
		err := rows.Scan(&todo.ID, &todo.Name, &todo.Deadline)

		if err != nil {
			panic(err.Error())
		}

		todos = append(todos, todo)
	}

	return todos
}