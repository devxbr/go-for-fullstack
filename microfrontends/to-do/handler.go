package todo

import (
	"fmt"
)

// in memory database
var todoList []*Todo

type Todo struct {
	Title       string
	Description string
	Status      string
}

func Add(todo Todo) {
	todoList = append(todoList, &Todo{Title: todo.Title, Description: todo.Description, Status: todo.Status})
}

func buildHeader() string {
	return `<div class="row">
            <form id="todoForm" onsubmit="submitTodo(event)">
            <div class="input-field">
                <input id="title" type="text" name="title" required>
                <label for="title">Title</label>
            </div>
            <div class="input-field">
                <input id="description" type="text" name="description" required>
                <label for="description">Description</label>
            </div>
            <div class="input-field">
                <select id="status" name="status" required>
                    <option value="In Progress">In Progress</option>
                    <option value="Waiting">Waiting</option>
                    <option value="Completed">Completed</option>
                </select>
                <label for="status">Status</label>
            </div>
            <button class="btn waves-effect waves-light" type="submit">Add Todo</button>
        </form>
        </div>
    
    <div class=row>
    <table class="highlight centered"><thead><tr><th>Title</th><th>Description</th><th>Status</th></tr></thead><tbody>`
}

func buildFooter() string {
	return `</tbody></table></div>`
}

func Render() string {
	component := buildHeader()
	for i := range todoList {
		component += fmt.Sprintf(`<tr><td>%s</td><td>%s</td><td>%s</td></tr>`,
			todoList[i].Title, todoList[i].Description, todoList[i].Status)

	}
	component += buildFooter()
	return component
}
