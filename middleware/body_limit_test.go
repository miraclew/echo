package middleware

import (
	"io/ioutil"
	"net/http"
	"testing"

	"bytes"

	"github.com/labstack/echo"
	"github.com/labstack/echo/test"
	"github.com/stretchr/testify/assert"
)

func TestBodyLimit(t *testing.T) {
	e := echo.New()
	req := test.NewRequest(echo.POST, "/", bytes.NewReader([]byte("Hello, World!")))
	rec := test.NewResponseRecorder()
	c := e.NewContext(req, rec)
	h := func(c echo.Context) error {
		body, _ := ioutil.ReadAll(c.Request().Body())
		return c.String(http.StatusOK, string(body))
	}

	// Within limit
	BodyLimit("2M")(h)(c)
	assert.Equal(t, http.StatusOK, rec.Status())
	assert.Equal(t, "Hello, World!", rec.Body.String())

	// Overlimit
	// BodyLimit("2B")(h)(c)
	// assert.Equal(t, "Hello, World!", rec.Body.String())
}
