package test

import (
	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/controller"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/model/web"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/repository"
	"github.com/geronimo794/golang-ach-rozikin-mastering/project2/service"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func setUpTodoTestRouterController() controller.TodoController {
	validate := validator.New()
	db := NewDatabaseTest()

	// Auth API
	todoRepository := repository.NewTodoRepository()
	todoService := service.NewTodoService(todoRepository, db, validate)
	todoController := controller.NewTodoController(todoService, validate)
	return todoController
}
func setUpDatabaseTest(db *gorm.DB, todoService service.TodoService) {
	// Truncate table
	db.Exec("TRUNCATE TABLE users")

	// Create single todo data
	todoData := web.RequestTodo{
		Name:     "Example content todo data",
		Priority: "high",
	}
	todoService.Create(db.Statement.Context, todoData)

}

// type fields struct {
// 	TodoService service.TodoService
// 	Validate    *validator.Validate
// }
// type args struct {
// 	e echo.Context
// }
// tests := []struct {
// 	name    string
// 	fields  fields
// 	args    args
// 	wantErr bool
// }{
// 	// TODO: Add test cases.
// }
// for _, tt := range tests {
// 	t.Run(tt.name, func(t *testing.T) {
// 		controller := &TodoControllerImpl{
// 			TodoService: tt.fields.TodoService,
// 			Validate:    tt.fields.Validate,
// 		}
// 		if err := controller.Update(tt.args.e); (err != nil) != tt.wantErr {
// 			t.Errorf("TodoControllerImpl.Update() error = %v, wantErr %v", err, tt.wantErr)
// 		}
// 	})
// }
