package reader

import (
	"strings"
	"testing"
	"time"
)

func TestGetJsonContent(t *testing.T) {
	reader := createReader()
	content := getJsonContent(reader)
	tasks := parseTasks(content)
	machines := parseMachines(content)

	if len(machines) != 1 {
		t.Errorf("Expected 1 machine, got %d", len(machines))
	}
	if machines[0].Name() != "Machine1" {
		t.Errorf("Expected Machine1, got %s", machines[0].Name())
	}
	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}
	if tasks[0].Name() != "Task1" {
		t.Errorf("Expected Task1, got %s", tasks[0].Name())
	}
	if tasks[0].Start() != time.Date(2020, 1, 1, 1, 0, 0, 0, time.UTC) {
		t.Errorf("Expected 2020-01-01T01:00:00Z, got %s", tasks[0].Start())
	}
	if tasks[0].End() != time.Date(2020, 1, 1, 2, 0, 0, 0, time.UTC) {
		t.Errorf("Expected 2020-01-01T02:00:00Z, got %s", tasks[0].End())
	}

}

func createReader() *strings.Reader {
	reader := strings.NewReader(`
	{"Tasks": 
		[
			{
				"name": "Task1", 
				"start": "2020-01-01T01:00:00Z", 
				"end": "2020-01-01T02:00:00Z"
			}
		], 
	 "Machines": [{"name": "Machine1"}]
	 }`)
	return reader
}
