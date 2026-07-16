package echo

import "io"

// RebindBody allows re-reading the request body after middleware
func RebindBody(c *Context, body []byte) {
    c.Request().Body = io.NopCloser(strings.NewReader(string(body)))
}
