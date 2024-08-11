package domain

type Problem struct {
	tasks    []Task
	machines []Machine
}

func NewProblem(tasks []Task, machines []Machine) Problem {
	return Problem{
		tasks:    tasks,
		machines: machines,
	}
}
