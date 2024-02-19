package atq

import (
	"testing"
)
func TestTaskStateString(t *testing.T) {
	tst := map[TaskState]string {
		TaskStateActive: "active",
		TaskStatePending: "pending",
		TaskStateScheduled: "scheduled",
		TaskStateRetry: "retry",
		TaskStateCompleted: "completed",
		TaskStateArchived:"archived",
		TaskStateAggregating: "aggregating",
	}
	for k, v := range tst {
		if taskStateString[k] != v {
			t.Errorf("expected %+v, returned %+v", v, taskStateString[k])
		}
	}
}