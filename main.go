package main

import (
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type List struct {
	Todos []*Todo
}

type Todo struct {
	Name        string
	Description string
	Date        string
}

func index(c *fiber.Ctx, todos []*Todo) error {
	return c.Render("index", fiber.Map{
		"Todos": todos,
	}, "main")
}

func todo(c *fiber.Ctx, todos []*Todo) error {
	return c.Render("partials/todos", fiber.Map{
		"Todos": todos,
	})
}

func main() {
	list := List{
		Todos: []*Todo{
			{
				Name:        "clara bday",
				Description: "this is description",
				Date:        "12-11-23",
			},
			{
				Name:        "feed the dawg",
				Description: "",
				Date:        "09-11-23",
			},
			{
				Name:        "do homework",
				Description: "chemistry class",
				Date:        "01-09-09",
			},
		},
	}

	// Setup app
	app := fiber.New(fiber.Config{
		Views: html.New("./views", ".html"),
	})

	// Controllers
	app.Get("/", func(c *fiber.Ctx) error {
		return index(c, list.Todos)
	})

	app.Post("/add", func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		description := c.FormValue("description")

		add := &Todo{
			Name:        name,
			Description: description,
			Date:        time.Now().Format("01-02-06"),
		}

		list.Todos = append(list.Todos, add)

		return todo(c, list.Todos)
	})

	app.Post("/delete", func(c *fiber.Ctx) error {
		name := c.FormValue("name")

		for i := 0; i < len(list.Todos); i++ {
			if list.Todos[i].Name == name {
				list.Todos = append(list.Todos[:i], list.Todos[i+1:]...)
				break
			}
		}

		return todo(c, list.Todos)
	})

	app.Post("/search", func(c *fiber.Ctx) error {
		search := c.FormValue("search")
		result := []*Todo{}

		for _, todo := range list.Todos {
			if strings.Contains(todo.Name, search) {
				result = append(result, todo)
			}
		}

		return todo(c, result)
	})

	// Start
	app.Listen(":5000")
}
