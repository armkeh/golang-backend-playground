package routes

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_pong(t *testing.T) {
	const pong = "pong"

	t.Run("Test ping response", func(t *testing.T) {
		// Test setup
		gin.SetMode(gin.TestMode)

		w := httptest.NewRecorder()
		ctx, _ := gin.CreateTestContext(w)
		ctx.Request = &http.Request{
			Method: http.MethodGet,
		}

		ctx.Request.Method = "GET"
		pingHandler(ctx)

		// The test
		r := w.Body.String()

		// Result verification
		if r != pong {
			t.Fatalf(`Ping response was "%s", expected "%s"`, r, pong)
		}
	})
}
