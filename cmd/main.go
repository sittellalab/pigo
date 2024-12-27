package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sittellalab/pigo/internal"
	"github.com/sittellalab/pigo/pkg/lib"
)

var port int

func init() {
	flag.IntVar(&port, "p", 8000, "pick a port between 1024 and 65535")
	flag.Parse()
	if port < 1024 || port > 65535 {
		log.Fatalf("port: %d out of range, must be between 1024 and 65535", port)
	}
}

func main() {
	e := echo.New()

	e.File("/pigo.min.css", "pkg/themes/pigo.min.css")

	e.GET("/", func(c echo.Context) error {
		return lib.Render(c, http.StatusOK, internal.Stage())
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
