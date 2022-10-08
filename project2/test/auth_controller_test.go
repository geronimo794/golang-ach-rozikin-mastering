package test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/controller"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/model/web"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func setUpTestAuthController() controller.AuthController {
	validate := validator.New()

	// Auth Controller
	authService := service.NewAuthService()
	authController := controller.NewAuthController(authService, validate)
	return authController
}

type TestCaseAuth struct {
	Req          web.RequestAuth
	ExpectedCode int
	ExpectedData string
}

func TestAuth(t *testing.T) {
	// Setup authentification controller
	authController := setUpTestAuthController()

	// Setup table test
	f := make(url.Values)
	var testCase = []TestCaseAuth{
		// Username password success
		{Req: web.RequestAuth{
			Username: "admin",
			Password: "admin",
		},
			ExpectedCode: http.StatusCreated,
			ExpectedData: "\"token\":\""},
		// Username password wrong
		{Req: web.RequestAuth{
			Username: "salah",
			Password: "salah",
		},
			ExpectedCode: http.StatusUnauthorized},
		// Empty username and password
		{Req: web.RequestAuth{
			Username: "",
			Password: "",
		},
			ExpectedCode: http.StatusBadRequest},
	}

	// Doing test with table test
	for _, v := range testCase {
		f.Set("username", v.Req.Username)
		f.Set("password", v.Req.Password)
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(f.Encode()))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Assertions
		if assert.NoError(t, authController.Authenticate(c)) {
			assert.Equal(t, v.ExpectedCode, rec.Code)
			// Test for checking content data
			if len(v.ExpectedData) > 0 {
				assert.Equal(t, strings.Contains(rec.Body.String(), v.ExpectedData), true)
			}
		}
	}
}
