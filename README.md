Source: https://github.com/go-monk/todo-rest-api

In this post, I build a simple REST API server using only Go and its standard library packages. The application is a to-do list with the following endpoints:

* `POST /task` – Add a task and return its ID
* `GET /tasks` – Return all existing tasks
* `DELETE /task/{id}` – Delete the task with the given ID
* `GET /task/{id}` – Return the task with the given ID

All data, both incoming and outgoing, is JSON-encoded. Here's how it works in practice:

```sh
$ curl localhost:8080/task --json '{ "text": "Learn Bash" }'
{"id":0}
$ curl localhost:8080/task --json '{ "text": "Learn Go" }'
{"id":1}
$ curl localhost:8080/tasks
[{"Id":0,"Text":"Learn Bash"},{"Id":1,"Text":"Learn Go"}]
$ curl localhost:8080/task/0 -X DELETE
$ curl localhost:8080/task/0
task with id=0 not found
```

First, I define a type to represent to-do list tasks, along with an in-memory data store:

```go
type Task struct {
	Id   int
	Text string
}

type TaskStore struct {
	sync.Mutex

	tasks  map[int]Task
	nextId int
}
```

In a production application, the task store would typically be a database. Here, tasks are stored in memory only, so they’re lost when the server stops. Since the store can be accessed concurrently (each request is handled in a separate goroutine), I use a mutex lock to protect it.

Next, I implement the operations for working with tasks. These are defined as methods on the `TaskStore`:

```go
func (ts *TaskStore) CreateTask(text string) int
func (ts *TaskStore) GetTasks() []Task
func (ts *TaskStore) GetTask(id int) (Task, error)
func (ts *TaskStore) DeleteTask(id int) error
```

At this point, I have the data structure to hold the tasks and the methods to operate on them. Since this is a REST API, I now need to map HTTP methods (POST, GET, DELETE) and paths to handler functions. These handlers process incoming HTTP requests:

```go
func main() {
	mux := http.NewServeMux()
	handler := handler.NewTaskHandler()

	mux.HandleFunc("POST /task", handler.AddTask)
	mux.HandleFunc("GET /tasks", handler.GetTasks)
	mux.HandleFunc("GET /task/{id}", handler.GetTask)
	mux.HandleFunc("DELETE /task/{id}", handler.DeleteTask)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
```

This is based on: https://eli.thegreenplace.net/2021/rest-servers-in-go-part-1-standard-library