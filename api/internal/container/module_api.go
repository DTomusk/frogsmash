package container

type APIContainer struct {
	*BaseContainer
	*Auth
}

func NewAPIContainer(c *BaseContainer) *APIContainer {
	auth := NewAuth(c.Config, c.User.UserService, c.InfraServices.MessageProducer)

	return &APIContainer{
		BaseContainer: c,
		Auth:          auth,
	}
}
