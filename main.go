package main

import (
	"bytes"
	"errors"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"strings"
)

func main() {
	s := martini.Classic()
	s.Use(render.Renderer())
	s.Use(martini.Static("static", martini.StaticOptions{Prefix: "static"}))

	s.Get("/", func(r render.Render) {
		r.HTML(200, "index", "")
	})

	s.Post("/eval", func(req *http.Request) string {
		buf := make([]byte, 1024)
		_, _ = req.Body.Read(buf)
		buf = bytes.Trim(buf, "\x00")
		code := strings.TrimSpace(string(buf))
		err, result := run(code)

		if err == nil {
			return result
		}
		return "error" + errors.Unwrap(err).Error()
	})

	s.Run()
}
