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
	return `
                <script>
                document.addEventListener('DOMContentLoaded', function() {
                    var elems = document.querySelectorAll('select');
                    M.FormSelect.init(elems);
                });

                function renderTodoItem(title, description, status) {
                    const todoList = document.getElementById('todo-list');
                    const row = document.createElement('tr');
                    row.innerHTML = "<td>" + title + "</td>" +
                                    "<td>" + description + "</td>" +
                                    "<td>" + status + "</td>";
                    todoList.appendChild(row);
                }           

                function submitTodo(event) {
                    event.preventDefault();
                    const form = document.getElementById('todoForm');
                    const data = {
                        title: form.title.value,
                        description: form.description.value,
                        status: form.status.value
                    };

                    fetch('http://localhost:8080/add-to-do', {
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/json'
                        },
                        body: JSON.stringify(data)
                    })
                    .then()
                    .then((data) => {
                        alert('Tarefa adicionada com sucesso!');
                        renderTodoItem(form.title.value, form.description.value, form.status.value);
                        form.reset();
                    })
                    .catch((error) => {
                        console.error('Error:', error);
                        alert('Erro ao adicionar tarefa.');
                    });
                }

            </script>
    
        <div class="row">
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
    <table class="highlight centered"><thead><tr><th>Title</th><th>Description</th><th>Status</th></tr></thead><tbody id="todo-list">`
}

func buildFooter() string {
	return "</tbody></table></div>"
}

func buildTodoList() string {
	component := ""
	for i := range todoList {
		component += fmt.Sprintf(`<tr><td>%s</td><td>%s</td><td>%s</td></tr>`,
			todoList[i].Title, todoList[i].Description, todoList[i].Status)

	}
	return component
}

func Render() string {
	component := buildHeader()
	component += buildTodoList()
	component += buildFooter()
	return component
}
