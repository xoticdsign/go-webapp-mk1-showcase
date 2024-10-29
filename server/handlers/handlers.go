package handlers

import (
	"go-webapp-mark1-showcase/gojwt"
	"go-webapp-mark1-showcase/gorm"

	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// SUPPORTING FUNCTIONS

func renderHTML(c *fiber.Ctx) error {
	switch c.Path() {
	case "/":
		_, role, _ := getRole(c)

		err := c.Render("index", fiber.Map{
			"Role":                   role,
			"RootName":               "Home",
			"LoginName":              "Login",
			"LoginSubmitName":        "Submit",
			"RegistrationName":       "Registration",
			"RegistrationSubmitName": "Submit",
			"LogoutName":             "Logout",
			"ProfileName":            "Profile",
			"ProfileEditName":        "Edit Profile",
			"ProfileEditSubmitName":  "Submit",
		})
		if err != nil {
			return err
		}

	case "/login":
		err := c.Render("login", fiber.Map{})
		if err != nil {
			return err
		}

	case "/registration":
		err := c.Render("registration", fiber.Map{})
		if err != nil {
			return err
		}

	case "/profile":
		token, role, _ := getRole(c)

		userDets, err := gorm.SelectProfile(token.Claims.(*gojwt.CustomClaims).Subject)
		if err != nil {
			return err
		}

		err = c.Render("profile", fiber.Map{
			"Role":                   role,
			"RootName":               "Home",
			"LoginName":              "Login",
			"LoginSubmitName":        "Submit",
			"RegistrationName":       "Registration",
			"RegistrationSubmitName": "Submit",
			"LogoutName":             "Logout",
			"ProfileName":            "Profile",
			"ProfileEditName":        "Edit Profile",
			"ProfileEditSubmitName":  "Submit",

			"Nickname": userDets.Nickname,
			"Bio":      userDets.Bio,
		})
		if err != nil {
			return err
		}

	case "/profile/edit":
		_, role, _ := getRole(c)

		err := c.Render("profile_edit", fiber.Map{
			"Role":                   role,
			"RootName":               "Home",
			"LoginName":              "Login",
			"LoginSubmitName":        "Submit",
			"RegistrationName":       "Registration",
			"RegistrationSubmitName": "Submit",
			"LogoutName":             "Logout",
			"ProfileName":            "Profile",
			"ProfileEditName":        "Edit Profile",
			"ProfileEditSubmitName":  "Submit",
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func getRole(c *fiber.Ctx) (*jwt.Token, string, error) {
	var role string

	token, err := getTokenFromCookie(c)
	if err != nil {
		role = "visitor"

		return &jwt.Token{}, role, err
	}

	aud, err := token.Claims.GetAudience()
	if err != nil {
		role = "visitor"

		return &jwt.Token{}, role, err
	}

	for _, val := range aud {
		role = val
	}

	return token, role, nil
}

func setTokenCookie(c *fiber.Ctx, tokenString string) {
	cookie := new(fiber.Cookie)
	cookie.Name = "authToken"
	cookie.Value = tokenString
	cookie.Path = "/"
	cookie.Expires = time.Now().Add(time.Hour * 24)
	cookie.HTTPOnly = true
	c.Cookie(cookie)
}

func deleteTokenCookie(c *fiber.Ctx) {
	cookie := new(fiber.Cookie)
	cookie.Name = "authToken"
	cookie.Value = ""
	cookie.Path = "/"
	cookie.Expires = time.Now()
	cookie.HTTPOnly = true
	c.Cookie(cookie)
}

func getTokenFromCookie(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := c.Cookies("authToken")

	token, err := gojwt.VerifyJWT(tokenString)
	if err != nil {
		return &jwt.Token{}, err
	}
	return token, nil
}

// HANDLERS

func MainPage(c *fiber.Ctx) error {
	err := renderHTML(c)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "error 404: "+err.Error())
	}
	return nil
}

func Login(c *fiber.Ctx) error {
	err := renderHTML(c)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "error 404: "+err.Error())
	}
	return nil
}

func LoginSubmit(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	userCreds, err := gorm.SelectUser(username, password)
	if err != nil {
		return c.Redirect("/login", fiber.StatusSeeOther)
	}

	tokenString, err := gojwt.ConfigJWT(userCreds)
	if err != nil {
		return c.Redirect("/login", fiber.StatusSeeOther)
	}

	setTokenCookie(c, tokenString)

	return c.Redirect("/profile", fiber.StatusSeeOther)
}

func Registration(c *fiber.Ctx) error {
	err := renderHTML(c)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "error 404: "+err.Error())
	}
	return nil
}

func RegistrationSubmit(c *fiber.Ctx) error {
	username := c.FormValue("username")
	nickname := c.FormValue("nickname")
	bio := c.FormValue("bio")
	password := c.FormValue("password")
	passwordVerify := c.FormValue("passwordVerify")
	email := c.FormValue("email")

	if password != passwordVerify {
		return c.Redirect("/registration")
	}

	userCreds, err := gorm.CreateUser(username, nickname, bio, password, email)
	if err != nil {
		fmt.Println(err)
		return c.Redirect("/registration", fiber.StatusSeeOther)
	}

	tokenString, err := gojwt.ConfigJWT(userCreds)
	if err != nil {
		fmt.Println(err)
		return c.Redirect("/registration", fiber.StatusSeeOther)
	}

	setTokenCookie(c, tokenString)

	return c.Redirect("/profile", fiber.StatusSeeOther)
}

func Logout(c *fiber.Ctx) error {
	deleteTokenCookie(c)

	return c.Redirect("/", fiber.StatusSeeOther)
}

func Profile(c *fiber.Ctx) error {
	err := renderHTML(c)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "error 404: "+err.Error())
	}
	return nil
}

func ProfileEdit(c *fiber.Ctx) error {
	err := renderHTML(c)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "error 404: "+err.Error())
	}
	return nil
}

func EditSubmit(c *fiber.Ctx) error {
	token, err := getTokenFromCookie(c)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "error 401: "+err.Error())
	}

	username := token.Claims.(*gojwt.CustomClaims).Subject

	userDets, err := gorm.SelectProfile(username)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "error 500: "+err.Error())
	}

	nickname := c.FormValue("nickname", userDets.Nickname)
	bio := c.FormValue("bio", userDets.Bio)

	err = gorm.UpdateUser(username, nickname, bio)
	if err != nil {
		c.Redirect("/profile/edit", fiber.StatusSeeOther)
	}

	return c.Redirect("/profile/edit", fiber.StatusSeeOther)
}
