package container

type APIContainer struct {
	*Container
}

func NewAPIContainer(c *Container) *APIContainer {
	return &APIContainer{
		Container: c,
	}
}
