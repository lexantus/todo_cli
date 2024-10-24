package tasks

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"time"
)

//go:generate go run github.com/dmarkham/enumer -type=Status -output=status_enumer.go

type Id string
type Status int

const (
	WAITING Status = iota
	IN_PROGRESS
	CANCELED
	DONE
)

type Task struct {
	Id       Id     `toml:"id"`
	Desc     string `toml:"desc"`
	Status   Status `toml:"status"`
	Progress uint8  `toml:"progress"` // [0-100] percents
}

func (t Task) String() string {
	var statusIndicator = " "
	switch t.Status {
	case WAITING:
		statusIndicator = "_"
	case IN_PROGRESS:
		statusIndicator = "👨🏻‍💻"
	case CANCELED:
		statusIndicator = "❌"
	case DONE:
		statusIndicator = "✅"
	}
	return fmt.Sprintf("[%s] %s [%d%%]", statusIndicator, t.Desc, t.Progress)
}

func generateUniqueID() Id {
	// Get the current timestamp
	timestamp := time.Now().UnixNano()

	// Generate a random number
	randomBytes := make([]byte, 4)
	_, err := rand.Read(randomBytes)
	if err != nil {
		panic(err)
	}

	// Combine timestamp and random bytes
	id := fmt.Sprintf("%d-%s", timestamp, hex.EncodeToString(randomBytes))
	return Id(id)
}

func NewTask(desc string) Task {
	id := generateUniqueID()
	task := Task{
		Id:     id,
		Desc:   desc,
		Status: WAITING,
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
	task.Status = CANCELED
	taskToRemove := RemovedTask{
		Task:   task,
		reason: reason,
	}
	removedTasks[task.Id] = taskToRemove
	delete(tasks, id)
	return task
}

func (t Task) Complete(id Id) Task {
	err, task := t.getTaskById(id)
	if err != nil {
		log.Fatalf(err.Error())
		return Task{}
	}
	task.Progress = 100
	task.Status = DONE
	return task
}

func (t Task) ProgressOn(progress uint8) (Task, error) {
	if progress > 100 {
		return Task{}, fmt.Errorf("progress more than 100%% - %d%%", progress)
	}
	t.Status = IN_PROGRESS
	t.Progress = progress
	return t, nil
}
