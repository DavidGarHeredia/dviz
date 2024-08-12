package reader

import (
	"encoding/json"

	"github.com/DavidGarHeredia/dviz/src/machines/domain"
)

func parseMachines(content map[string]json.RawMessage) []domain.Machine {
	machines := []domain.Machine{}
	err := json.Unmarshal(content["Machines"], &machines)
	if err != nil {
		panic(err)
	}
	return machines
}

func parseTasks(content map[string]json.RawMessage) []domain.Task {
	tasks := []domain.Task{}
	err := json.Unmarshal(content["Tasks"], &tasks)
	if err != nil {
		panic(err)
	}
	return tasks
}
