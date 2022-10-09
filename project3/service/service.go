package service

// type TodoService struct{}

// func (t *TodoService) FindAll() []*model.Todo {
// 	todos := []*model.Todo{}

// 	database.DB.Find(&todos)

// 	return todos
// }

// func (t *TodoService) FindById(id string) (*model.Todo, error) {
// 	todo := model.Todo{}

// 	result := database.DB.First(&todo, "id = ?", id)

// 	if result.RowsAffected == 0 {
// 		return &todo, errors.New("todo not found")
// 	}
// 	return &todo, nil
// }

// func (t *TodoService) Create(input model.TodoInput) model.Todo {
// 	var newTodo = model.Todo{
// 		ID:       fmt.Sprintf("%d", rand.Int()),
// 		Name:     input.Name,
// 		Priority: input.Priority,
// 		Status:   0,
// 	}
// 	database.DB.Create(&newTodo)

// 	return newTodo
// }
