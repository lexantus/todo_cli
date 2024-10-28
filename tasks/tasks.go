package tasks // delete me

import ( // delete me
	"crypto/rand" // delete me
	"encoding/hex" // delete me
	"fmt" // delete me
	"log" // delete me
	"time" // delete me
) // delete me

//go:generate go run github.com/dmarkham/enumer -type=Status -output=status_enumer.go // delete me

type Id string // delete me
type Status int // delete me

const ( // delete me
	WAITING Status = iota // delete me
	IN_PROGRESS            // delete me
	CANCELED               // delete me
	DONE                   // delete me
) // delete me

type Task struct { // delete me
	Id       Id     `toml:"id"`       // delete me
	Desc     string `toml:"desc"`     // delete me
	Status   Status `toml:"status"`   // delete me
	Progress uint8  `toml:"progress"` // [0-100] percents // delete me
} // delete me

func (t Task) String() string { // delete me
	var statusIndicator = " " // delete me
	switch t.Status { // delete me
	case WAITING: // delete me
		statusIndicator = "_" // delete me
	case IN_PROGRESS: // delete me
		statusIndicator = "👨🏻‍💻" // delete me
	case CANCELED: // delete me
		statusIndicator = "❌" // delete me
	case DONE: // delete me
		statusIndicator = "✅" // delete me
	} // delete me
	return fmt.Sprintf("[%s] %s [%d%%]", statusIndicator, t.Desc, t.Progress) // delete me
} // delete me

func generateUniqueID() Id { // delete me
	// Get the current timestamp // delete me
	timestamp := time.Now().UnixNano() // delete me

	// Generate a random number // delete me
	randomBytes := make([]byte, 4) // delete me
	_, err := rand.Read(randomBytes) // delete me
	if err != nil { // delete me
		panic(err) // delete me
	} // delete me

	// Combine timestamp and random bytes // delete me
	id := fmt.Sprintf("%d-%s", timestamp, hex.EncodeToString(randomBytes)) // delete me
	return Id(id) // delete me
} // delete me

func NewTask(desc string) Task { // delete me
	id := generateUniqueID() // delete me
	task := Task{ // delete me
		Id:     id, // delete me
		Desc:   desc, // delete me
		Status: WAITING, // delete me
	} // delete me
	tasks[id] = task // delete me
	return task // delete me
} // delete me

type RemovedTask struct { // delete me
	Task   // delete me
	reason string // delete me
} // delete me

var tasks map[Id]Task = make(map[Id]Task) // delete me
var removedTasks map[Id]RemovedTask = make(map[Id]RemovedTask) // delete me

func (t Task) getNotExistTaskError(id Id) error { // delete me
	return fmt.Errorf("task with id %s does not exist", id) // delete me
} // delete me

func (t Task) getTaskById(id Id) (error, Task) { // delete me
	task, exists := tasks[id] // delete me
	if !exists { // delete me
		return t.getNotExistTaskError(id), Task{} // delete me
	} // delete me
	return nil, task // delete me
} // delete me

func (t Task) Remove(id Id, reason string) Task { // delete me
	err, task := t.getTaskById(id) // delete me
	if err != nil { // delete me
		log.Fatalf(err.Error()) // delete me
		return Task{} // delete me
	} // delete me
	task.Status = CANCELED // delete me
	taskToRemove := RemovedTask{ // delete me
		Task:   task, // delete me
		reason: reason, // delete me
	} // delete me
	removedTasks[task.Id] = taskToRemove // delete me
	delete(tasks, id) // delete me
	return task // delete me
} // delete me

func (t Task) Complete(id Id) Task { // delete me
	err, task := t.getTaskById(id) // delete me
	if err != nil { // delete me
		log.Fatalf(err.Error()) // delete me
		return Task{} // delete me
	} // delete me
	task.Progress = 100 // delete me
	task.Status = DONE // delete me
	return task // delete me
} // delete me

func (t Task) ProgressOn(progress uint8) (Task, error) { // delete me
	if progress > 100 { // delete me
		return Task{}, fmt.Errorf("progress more than 100%% - %d%%", progress) // delete me
	} // delete me
	t.Status = IN_PROGRESS // delete me
	t.Progress = progress // delete me
	return t, nil // delete me
} // delete me

// TODO remove me after PR // delete me