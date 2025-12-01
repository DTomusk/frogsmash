package container

type WorkerContainer struct {
	*Container
}

func NewWorkerContainer(c *Container) *WorkerContainer {
	return &WorkerContainer{
		Container: c,
	}
}
