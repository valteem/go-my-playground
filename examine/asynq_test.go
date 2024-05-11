package examine

import (
	"testing"

	"github.com/hibiken/asynq"
)

func TestAsynq(t *testing.T) {

	ttype := "task"
	task := asynq.NewTask(ttype, []byte("payload"))

	if task.Type() != "task" {
		t.Errorf("task type shoud be '%+v', is '%v'", ttype, task.Type())
	}
}
