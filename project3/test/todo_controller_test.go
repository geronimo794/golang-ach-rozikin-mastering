package test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/controller"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/model"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/model/web"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/repository"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project3/service"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func setUpTodoTestRouterController() (controller.TodoController, *gorm.DB) {
	validate := validator.New()
	db := NewDatabaseTest()
	// Auth API
	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepository, db, validate)
	todoController := controller.NewTodoController(todoService, validate)

	ClearTodosData(db)
	CreateSampleTodoData(db, todoRepository)

	return todoController, db
}

// Struct test case
type TestCaseRequestTodo struct {
	Req web.RequestTodo
	Exp ExpectationResultTest
}
type TestCaseRequestParameterTodo struct {
	Req web.RequestParameterTodo
	Exp ExpectationResultTest
}
type TestCaseRequestId struct {
	ReqId int
	Exp   ExpectationResultTest
}
type TestCaseRequestIdUpdateTodo struct {
	ReqId   int
	ReqData model.Todo
	Exp     ExpectationResultTest
}

// Testing Todo Controller
// Create
func TestTodoCreate(t *testing.T) {
	// Setup authentification controller
	todoController, _ := setUpTodoTestRouterController()

	// Set up table test
	f := make(url.Values)
	var testCase = []TestCaseRequestTodo{
		/**
		* Success test case
		**/
		// Check content data
		{
			Req: web.RequestTodo{
				Name:     "Create todo success",
				Priority: "high",
			},
			Exp: ExpectationResultTest{
				ExpectedCode:        http.StatusCreated,
				ExpectedContainData: "Create todo success",
			},
		},
		// Check priority data
		{
			Req: web.RequestTodo{
				Name:     "Create todo ok",
				Priority: "low",
			},
			Exp: ExpectationResultTest{
				ExpectedCode:        http.StatusCreated,
				ExpectedContainData: "low",
			},
		},
		/**
		* Failed test case
		**/
		// Empty all data
		{
			Req: web.RequestTodo{},
			Exp: ExpectationResultTest{
				ExpectedCode: http.StatusBadRequest,
			},
		},
		// Empty priority data
		{
			Req: web.RequestTodo{
				Name:     "Create todo ok",
				Priority: "",
			},
			Exp: ExpectationResultTest{
				ExpectedCode:        http.StatusBadRequest,
				ExpectedContainData: "priority",
			},
		},
		// Empty name data
		{
			Req: web.RequestTodo{
				Name:     "",
				Priority: "high",
			},
			Exp: ExpectationResultTest{
				ExpectedCode:        http.StatusBadRequest,
				ExpectedContainData: "name",
			},
		},
		// Wrong Priority data
		{
			Req: web.RequestTodo{
				Name:     "Create todo ok",
				Priority: "don't know",
			},
			Exp: ExpectationResultTest{
				ExpectedCode:        http.StatusBadRequest,
				ExpectedContainData: "priority",
			},
		},
	}

	// Doing test with table test
	for _, v := range testCase {
		f.Set("name", v.Req.Name)
		f.Set("priority", v.Req.Priority)
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(f.Encode()))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Assertions
		if assert.NoError(t, todoController.Create(c)) {
			assert.Equal(t, v.Exp.ExpectedCode, rec.Code)
			// Test for checking content data
			if len(v.Exp.ExpectedContainData) > 0 {
				assert.Equal(t, true, strings.Contains(rec.Body.String(), v.Exp.ExpectedContainData))
			}
		}
	}

}

// Find All
func TestTodoFindAll(t *testing.T) {
	// Setup authentification controller
	todoController, db := setUpTodoTestRouterController()

	// Create test case
	f := make(url.Values)
	var testCase = []TestCaseRequestParameterTodo{
		/**
		* Success test case
		**/
		// Get without request parameter
		{
			Req: web.RequestParameterTodo{},
			Exp: ExpectationResultTest{
				ExpectedCode: http.StatusOK,
			},
		},
		// Get with request parameter status
		{
			Req: web.RequestParameterTodo{
				IsDone: "true",
			},
			Exp: ExpectationResultTest{
				ExpectedCode:           http.StatusOK,
				ExpectedContainData:    `"is_done":true`,
				NotExpectedContainData: `"is_done":false`,
			},
		},
		{
			Req: web.RequestParameterTodo{
				IsDone: "false",
			},
			Exp: ExpectationResultTest{
				ExpectedCode:           http.StatusOK,
				ExpectedContainData:    `"is_done":false`,
				NotExpectedContainData: `"is_done":true`,
			},
		},
		// Get with request parameter keyowrd
		// Keyword found on sample
		{
			Req: web.RequestParameterTodo{
				Keyword: "Sample",
			},
			Exp: ExpectationResultTest{
				ExpectedCode:        http.StatusOK,
				ExpectedContainData: "Sample",
			},
		},
		// Keyword not found on sample
		{
			Req: web.RequestParameterTodo{
				Keyword: "This data not available on data sample",
			},
			Exp: ExpectationResultTest{
				ExpectedCode: http.StatusNotFound,
			},
		},
	}

	// Doing test with table test
	for _, v := range testCase {
		f.Set("keyword", v.Req.Keyword)
		f.Set("is_done", v.Req.IsDone)

		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/?"+f.Encode(), nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Assertions
		if assert.NoError(t, todoController.FindAll(c)) {
			assert.Equal(t, v.Exp.ExpectedCode, rec.Code)

			// Check expected data
			if len(v.Exp.ExpectedContainData) > 0 {
				assert.Contains(t, rec.Body.String(), v.Exp.ExpectedContainData)
			}

			// Check Not expected data
			if len(v.Exp.NotExpectedContainData) > 0 {
				assert.NotContains(t, rec.Body.String(), v.Exp.NotExpectedContainData)
			}
		}
	}

	/**
	*	Test if data not found or empty data
	**/
	// Clear all the data first
	ClearTodosData(db)
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/?"+f.Encode(), nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, todoController.FindAll(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}

// Find By Id
func TestTodoFindById(t *testing.T) {
	// Setup authentification controller
	todoController, _ := setUpTodoTestRouterController()

	// Create table test case
	var testCase = []TestCaseRequestId{
		{
			ReqId: 1,
			Exp: ExpectationResultTest{
				ExpectedCode:        http.StatusOK,
				ExpectedContainData: "Sample", // Check sample data is exist
			},
		},
		{
			Exp: ExpectationResultTest{
				ExpectedCode: http.StatusNotFound,
			},
		},
	}

	for _, v := range testCase {
		e := echo.New()

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/todo/:id")
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(v.ReqId))

		// Assertions
		if assert.NoError(t, todoController.FindById(c)) {
			assert.Equal(t, v.Exp.ExpectedCode, rec.Code)

			// Check expected data
			if len(v.Exp.ExpectedContainData) > 0 {
				assert.Equal(t, true, strings.Contains(rec.Body.String(), v.Exp.ExpectedContainData))
			}
		}
	}

}

// Update
func TestTodoUpdate(t *testing.T) {
	// Setup authentification controller
	todoController, _ := setUpTodoTestRouterController()

	// Create table test
	f := make(url.Values)
	var testCase = []TestCaseRequestIdUpdateTodo{
		// Success update
		{
			ReqId: 1,
			ReqData: model.Todo{
				Name:     "Updated todo",
				Priority: "high",
			},
			Exp: ExpectationResultTest{
				ExpectedCode:        http.StatusOK,
				ExpectedContainData: "Updated todo",
			},
		},
		// Failed update
		{
			ReqId: 1,
			ReqData: model.Todo{
				Name:     "Updated todo",
				Priority: "None",
			},
			Exp: ExpectationResultTest{
				ExpectedCode:        http.StatusBadRequest,
				ExpectedContainData: "priority",
			},
		},
		{
			ReqId: 1,
			ReqData: model.Todo{
				Name:     "Updated todo",
				Priority: "Unknown Priority",
			},
			Exp: ExpectationResultTest{
				ExpectedCode:        http.StatusBadRequest,
				ExpectedContainData: "priority",
			},
		},
		{
			ReqId: 1,
			Exp: ExpectationResultTest{
				ExpectedCode: http.StatusBadRequest,
			},
		},
		{
			ReqData: model.Todo{
				Name:     "Updated todo",
				Priority: "high",
			},
			Exp: ExpectationResultTest{
				ExpectedCode: http.StatusNotFound,
			},
		},
	}

	// Doing test with table test
	for _, v := range testCase {
		f.Set("name", v.ReqData.Name)
		f.Set("priority", v.ReqData.Priority)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(f.Encode()))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/todo/:id/")
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(v.ReqId))

		// Assertions
		if assert.NoError(t, todoController.Update(c)) {
			assert.Equal(t, v.Exp.ExpectedCode, rec.Code)

			// Check expected data
			if len(v.Exp.ExpectedContainData) > 0 {
				assert.Equal(t, true, strings.Contains(rec.Body.String(), v.Exp.ExpectedContainData))
			}
		}
	}

}

// Delete
func TestTodoDelete(t *testing.T) {
	// Setup authentification controller
	todoController, _ := setUpTodoTestRouterController()

	// Create table test case
	var testCase = []TestCaseRequestId{
		{
			ReqId: 1,
			Exp: ExpectationResultTest{
				ExpectedCode: http.StatusOK,
			},
		},
		{
			ReqId: 1,
			Exp: ExpectationResultTest{
				ExpectedCode: http.StatusNotFound,
			},
		},
		{
			Exp: ExpectationResultTest{
				ExpectedCode: http.StatusNotFound,
			},
		},
	}

	for _, v := range testCase {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/todo/:id")
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(v.ReqId))

		// Assertions
		if assert.NoError(t, todoController.Delete(c)) {
			assert.Equal(t, v.Exp.ExpectedCode, rec.Code)

			// Check expected data
			if len(v.Exp.ExpectedContainData) > 0 {
				assert.Equal(t, true, strings.Contains(rec.Body.String(), v.Exp.ExpectedContainData))
			}
		}
	}

}

// Reverse Is Done
func TestTodoReverseIsDone(t *testing.T) {
	// Setup authentification controller
	todoController, _ := setUpTodoTestRouterController()

	// Create table test case
	var testCase = []TestCaseRequestId{
		{
			ReqId: 1,
			Exp: ExpectationResultTest{
				ExpectedCode:        http.StatusOK,
				ExpectedContainData: `"is_done":true`,
			},
		},
		{
			ReqId: 1,
			Exp: ExpectationResultTest{
				ExpectedCode:        http.StatusOK,
				ExpectedContainData: `"is_done":false`,
			},
		},
		{
			Exp: ExpectationResultTest{
				ExpectedCode: http.StatusNotFound,
			},
		},
	}

	for _, v := range testCase {
		e := echo.New()

		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/todo/:id/reverse-status")
		c.SetParamNames("id")
		c.SetParamValues(strconv.Itoa(v.ReqId))

		// Assertions
		if assert.NoError(t, todoController.ReverseIsDone(c)) {
			assert.Equal(t, v.Exp.ExpectedCode, rec.Code)

			// Check expected data
			if len(v.Exp.ExpectedContainData) > 0 {
				assert.Contains(t, rec.Body.String(), v.Exp.ExpectedContainData)
			}
		}
	}

}
