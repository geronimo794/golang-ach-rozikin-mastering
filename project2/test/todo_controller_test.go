package test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/controller"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/helper"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/model"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/model/web"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/repository"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/service"
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

	clearTodosData(db)
	createSampleTodoData(db, todoRepository)

	return todoController, db
}
func createSampleTodoData(db *gorm.DB, todoRepository repository.TodoRepository) {
	todoData := model.Todo{
		Name:     "Sample data low",
		Priority: "low",
		Status:   1,
	}
	todoRepository.Create(db.Statement.Context, db, todoData)

	todoData = model.Todo{
		Name:     "Sample data medium",
		Priority: "medium",
		Status:   0,
	}
	todoRepository.Create(db.Statement.Context, db, todoData)

	todoData = model.Todo{
		Name:     "Sample data high",
		Priority: "high",
		Status:   1,
	}
	todoRepository.Create(db.Statement.Context, db, todoData)

}
func clearTodosData(db *gorm.DB) {
	// Truncate table
	err := db.Exec("TRUNCATE TABLE todos").Error
	helper.PanicIfError(err)
}

// Struct test case
type TestCaseRequestTodo struct {
	Req          web.RequestTodo
	ExpectedCode int
	ExpectedData string
}
type TestCaseRequestParameterTodo struct {
	Req             web.RequestParameterTodo
	ExpectedCode    int
	ExpectedData    string
	NotExpectedData string
}
type TestCaseRequestId struct {
	ReqId        int
	ExpectedCode int
	ExpectedData string
}
type TestCaseRequestIdUpdateTodo struct {
	ReqId        int
	ReqData      model.Todo
	ExpectedCode int
	ExpectedData string
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
			ExpectedCode: http.StatusCreated,
			ExpectedData: "Create todo success",
		},
		// Check priority data
		{
			Req: web.RequestTodo{
				Name:     "Create todo ok",
				Priority: "low",
			},
			ExpectedCode: http.StatusCreated,
			ExpectedData: "low",
		},
		/**
		* Failed test case
		**/
		// Empty all data
		{
			Req:          web.RequestTodo{},
			ExpectedCode: http.StatusBadRequest,
		},
		// Empty priority data
		{
			Req: web.RequestTodo{
				Name:     "Create todo ok",
				Priority: "",
			},
			ExpectedCode: http.StatusBadRequest,
			ExpectedData: "priority",
		},
		// Empty name data
		{
			Req: web.RequestTodo{
				Name:     "",
				Priority: "high",
			},
			ExpectedCode: http.StatusBadRequest,
			ExpectedData: "name",
		},
		// Wrong Priority data
		{
			Req: web.RequestTodo{
				Name:     "Create todo ok",
				Priority: "don't know",
			},
			ExpectedCode: http.StatusBadRequest,
			ExpectedData: "priority",
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
			assert.Equal(t, v.ExpectedCode, rec.Code)
			// Test for checking content data
			if len(v.ExpectedData) > 0 {
				assert.Equal(t, true, strings.Contains(rec.Body.String(), v.ExpectedData))
			}
		}
	}

}

// Find All
func TestTodoFindAll(t *testing.T) {
	// Setup authentification controller
	todoController, _ := setUpTodoTestRouterController()

	// Create test case
	f := make(url.Values)
	var testCase = []TestCaseRequestParameterTodo{
		/**
		* Success test case
		**/
		// Get without request parameter
		{
			Req:          web.RequestParameterTodo{},
			ExpectedCode: http.StatusOK,
		},
		// Get with request parameter status
		{
			Req: web.RequestParameterTodo{
				Status: 1,
			},
			ExpectedCode:    http.StatusOK,
			ExpectedData:    "\"status\":1",
			NotExpectedData: "\"status\":0",
		},
		{
			Req: web.RequestParameterTodo{
				Status: 0,
			},
			ExpectedCode:    http.StatusOK,
			ExpectedData:    "\"status\":0",
			NotExpectedData: "\"status\":1",
		},
		// Get with request parameter keyowrd
		// Keyword found on sample
		{
			Req: web.RequestParameterTodo{
				Keyword: "Sample",
			},
			ExpectedCode: http.StatusOK,
			ExpectedData: "Sample",
		},
		// Keyword not found on sample
		{
			Req: web.RequestParameterTodo{
				Keyword: "This data not available on data sample",
			},
			ExpectedCode: http.StatusNotFound,
		},
	}

	// Doing test with table test
	for _, v := range testCase {
		f.Set("keyword", v.Req.Keyword)
		f.Set("status", strconv.Itoa(int(v.Req.Status)))
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/?"+f.Encode(), nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Assertions
		if assert.NoError(t, todoController.FindAll(c)) {
			assert.Equal(t, v.ExpectedCode, rec.Code)

			// Check expected data
			if len(v.ExpectedData) > 0 {
				assert.Equal(t, true, strings.Contains(rec.Body.String(), v.ExpectedData))
			}

			// Check Not expected data
			if len(v.NotExpectedData) > 0 {
				assert.Equal(t, false, strings.Contains(rec.Body.String(), v.NotExpectedData))
			}
		}
	}

}

// Find By Id
func TestTodoFindById(t *testing.T) {
}

// Update
func TestTodoUpdate(t *testing.T) {
}

// Delete
func TestTodoDelete(t *testing.T) {
}

// Reverse Status
func TestTodoReverseStatus(t *testing.T) {
}
