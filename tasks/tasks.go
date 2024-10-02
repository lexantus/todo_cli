package tasks

import (
	"fmt"
	"log"
)

type Id string
type Status int

const (
	WAITING Status = iota
	IN_PROGRESS
	CANCELED
	DONE
)

var TaskStatus = map[Status]string{
	WAITING:     "WAITING",
	IN_PROGRESS: "IN_PROGRESS",
	CANCELED:    "CANCELED",
	DONE:        "DONE",
}

func (s Status) String() string {
	return TaskStatus[s]
}

type Task struct {
	id       Id
	desc     string
	status   Status
	progress uint8 // [0-100] percents
}

func (t Task) String() string {
	var statusIndicator = " "
	switch t.status {
	case WAITING:
		statusIndicator = "_"
	case IN_PROGRESS:
		statusIndicator = "👨🏻‍💻"
	case CANCELED:
		statusIndicator = "❌"
	case DONE:
		statusIndicator = "✅"
	}
	return fmt.Sprintf("[%s] %s [%d%%]", statusIndicator, t.desc, t.progress)
}

var nextId = getNextId()

func NewTask(desc string) Task {
	id := nextId()
	task := Task{
		id:     id,
		desc:   desc,
		status: WAITING,
	}
	tasks[id] = task
	return task
}

type RemovedTask struct {
	Task
	reason string
}

var tasks map[Id]Task = make(map[Id]Task)
var removedTasks map[Id]RemovedTask = make(map[Id]RemovedTask)

func getNextId() func() Id {
	idCounter := 0
	return func() Id {
		idCounter++
		return Id(fmt.Sprintf("id%d", idCounter))
	}
}

func (t Task) getNotExistTaskError(id Id) error {
	return fmt.Errorf("task with id %s does not exist", id)
}

func (t Task) getTaskById(id Id) (error, Task) {
	task, exists := tasks[id]
	if !exists {
		return t.getNotExistTaskError(id), Task{}
	}
	return nil, task
}

func (t Task) Remove(id Id, reason string) Task {
	err, task := t.getTaskById(id)
	if err != nil {
		log.Fatalf(err.Error())
		return Task{}
	}
	task.status = CANCELED
	taskToRemove := RemovedTask{
		Task:   task,
		reason: reason,
	}
	removedTasks[task.id] = taskToRemove
	delete(tasks, id)
	return task
}

func (t Task) Complete(id Id) Task {
	err, task := t.getTaskById(id)
	if err != nil {
		log.Fatalf(err.Error())
		return Task{}
	}
	task.progress = 100
	task.status = DONE
	return task
}

func (t Task) Progress(progress uint8) (Task, error) {
	if progress > 100 {
		return Task{}, fmt.Errorf("progress more than 100%% - %d%%", progress)
	}
	t.status = IN_PROGRESS
	t.progress = progress
	return t, nil
}
