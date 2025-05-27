package todo

import (
	"fmt"
	"sync"
)

type Task struct {
	Id   int
	Text string
}

type TaskStore struct {
	sync.Mutex

	tasks  map[int]Task
	nextId int
}

func NewTaskStore() *TaskStore {
	ts := &TaskStore{}
	ts.tasks = make(map[int]Task)
	ts.nextId = 0
	return ts
}

func (ts *TaskStore) CreateTask(text string) int {
	ts.Lock()
	defer ts.Unlock()

	task := Task{
		Id:   ts.nextId,
		Text: text,
	}

	ts.tasks[ts.nextId] = task
	ts.nextId++
	return task.Id
}

func (ts *TaskStore) GetTasks() []Task {
	ts.Lock()
	defer ts.Unlock()

	var tasks []Task
	for _, t := range ts.tasks {
		tasks = append(tasks, t)
	}
	return tasks
}

func (ts *TaskStore) GetTask(id int) (Task, error) {
	ts.Lock()
	defer ts.Unlock()

	t, ok := ts.tasks[id]
	if !ok {
		return Task{}, fmt.Errorf("task with id=%d not found", id)
	}
	return t, nil
}

func (ts *TaskStore) DeleteTask(id int) error {
	ts.Lock()
	defer ts.Unlock()

	if _, ok := ts.tasks[id]; !ok {
		return fmt.Errorf("task with id=%d not found", id)
	}

	delete(ts.tasks, id)
	return nil
}
