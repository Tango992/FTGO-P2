package main

import (
	"avengers-chat/controller"
	"avengers-chat/handler"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	t := &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Renderer = t

	e.GET("/", handler.Menu)
	e.GET("/inventory", handler.InventoryChatRoom)
	e.GET("/crime", handler.CrimesChatRoom)
	e.GET("/random", handler.RandomChatRoom)

	e.GET("/ws/inventory", controller.InventoryWebsocket)
	e.GET("/ws/crime", controller.CrimeWebsocket)
	e.GET("/ws/random", controller.RandomWebsocket)

	e.Logger.Fatal(e.Start(":1323"))
}
