package recover //nolint:predeclared // TODO: Rename to some non-builtin

import (
	"net/http/httptest"
	"testing"

	"github.com/hakantaskin/fiber"
	"github.com/hakantaskin/fiber/utils"
)

// go test -run Test_Recover
func Test_Recover(t *testing.T) {
	t.Parallel()
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			utils.AssertEqual(t, "Hi, I'm an error!", err.Error())
			return c.SendStatus(fiber.StatusTeapot)
		},
	})

	app.Use(New())

	app.Get("/panic", func(c *fiber.Ctx) error {
		panic("Hi, I'm an error!")
	})

	resp, err := app.Test(httptest.NewRequest(fiber.MethodGet, "/panic", nil))
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusTeapot, resp.StatusCode)
}

// go test -run Test_Recover_Next
func Test_Recover_Next(t *testing.T) {
	t.Parallel()
	app := fiber.New()
	app.Use(New(Config{
		Next: func(_ *fiber.Ctx) bool {
			return true
		},
	}))

	resp, err := app.Test(httptest.NewRequest(fiber.MethodGet, "/", nil))
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusNotFound, resp.StatusCode)
}

func Test_Recover_EnableStackTrace(t *testing.T) {
	t.Parallel()
	app := fiber.New()
	app.Use(New(Config{
		EnableStackTrace: true,
	}))

	app.Get("/panic", func(c *fiber.Ctx) error {
		panic("Hi, I'm an error!")
	})

	resp, err := app.Test(httptest.NewRequest(fiber.MethodGet, "/panic", nil))
	utils.AssertEqual(t, nil, err)
	utils.AssertEqual(t, fiber.StatusInternalServerError, resp.StatusCode)
}
