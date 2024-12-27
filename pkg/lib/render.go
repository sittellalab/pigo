package lib

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

// Render renders zero to many templ.Component objects into an HTML response with the given status code.
//
// It sequentially renders each component to an internal buffer and sends the result as an HTML response.
// If any component fails to render, the function returns the error immediately.
func Render(c echo.Context, status int, cmps ...templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	for _, cmp := range cmps {
		if err := cmp.Render(c.Request().Context(), buf); err != nil {
			return err
		}
	}
	return c.HTML(status, buf.String())
}

// IsHTMX checks if the request is an HTMX request.
//
// It returns true if the request contains the "HX-Request" header, otherwise, it returns false.
func IsHTMX(c echo.Context) bool {
	return c.Request().Header.Get("HX-Request") == "true"
}

// HxRender conditionally renders a component based on whether the request is an HTMX request.
//
// If the request is an HTMX request, it renders the `cmp` component. Otherwise, it renders the `page` component.
func HxRender(c echo.Context, status int, page, cmp templ.Component) error {
	if IsHTMX(c) {
		return Render(c, status, cmp)
	}
	return Render(c, status, page)
}

// Redirect performs an HTTP redirect, with special handling for HTMX requests.
//
// For HTMX requests, the "HX-Redirect" header is set with the route, and a 200 OK status is used.
// HTMX does not handle 3xx redirects, it only sees the final 2xx, 4xx, or 5xx response.
// For non-HTMX requests, a standard HTTP redirect is performed with the provided status.
//
// See github.com/bigskysoftware/htmx/issues/2052#issuecomment-1979805051 for more details.
func Redirect(c echo.Context, status int, route string) error {
	if IsHTMX(c) {
		c.Response().Header().Add("HX-Redirect", route)
		status = http.StatusOK
	}
	return c.Redirect(status, route)
}
