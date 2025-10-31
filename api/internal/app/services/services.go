package services

type EventsRepo interface {
	LogEvent(winnerId, loserId string) error
}

type EventsService struct {
	Repo EventsRepo
}

func NewEventsService(repo EventsRepo) *EventsService {
	return &EventsService{Repo: repo}
}

func (s *EventsService) LogEvent(winnerId, loserId string) error {
	return s.Repo.LogEvent(winnerId, loserId)
}
