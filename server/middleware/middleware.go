package middleware

import (
	"go-webapp-mark1-showcase/gojwt"

	"github.com/gofiber/fiber/v2"
)

// MIDDLEWARE

func Auth(c *fiber.Ctx) error {
	tokenString := c.Cookies("authToken")
	if tokenString == "" {
		switch c.Path() {
		case "/profile":
			return c.Redirect("/", fiber.StatusSeeOther)

		case "/profile/edit":
			return c.Redirect("/", fiber.StatusSeeOther)

		case "/profile/edit/edit-submit":
			return c.Redirect("/", fiber.StatusSeeOther)

		case "/logout":
			return c.Redirect("/", fiber.StatusSeeOther)
		}

		return c.Next()
	}

	_, err := gojwt.VerifyJWT(tokenString)
	if err != nil {
		switch c.Path() {
		case "/profile":
			return c.Redirect("/", fiber.StatusSeeOther)

		case "/profile/edit":
			return c.Redirect("/", fiber.StatusSeeOther)

		case "/profile/edit/edit-submit":
			return c.Redirect("/", fiber.StatusSeeOther)

		case "/logout":
			return c.Redirect("/", fiber.StatusSeeOther)
		}

		return c.Next()
	}

	switch c.Path() {
	case "/login":
		return c.Redirect("/", fiber.StatusSeeOther)

	case "/login/login-submit":
		return c.Redirect("/", fiber.StatusSeeOther)

	case "/registration":
		return c.Redirect("/", fiber.StatusSeeOther)

	case "/registration/registration-submit":
		return c.Redirect("/", fiber.StatusSeeOther)
	}

	return c.Next()
}
