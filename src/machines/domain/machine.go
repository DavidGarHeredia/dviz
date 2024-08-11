package domain

type machineData struct {
	Name string `json:"name"`
}

type Machine struct {
	machineData
}

func NewMachine(name string) Machine {
	return Machine{
		machineData{Name: name},
	}
}

func (m *Machine) Name() string {
	return m.machineData.Name
}
