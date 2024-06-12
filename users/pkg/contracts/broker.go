package contracts

type Topic int

const (
	UserCreated Topic = iota
	UserUpdated
	UserDeleted
)

func (s Topic) String() string {
	switch s {
	case UserCreated:
		return "user-created"
	case UserUpdated:
		return "user-updated"
	case UserDeleted:
		return "user-deleted"
	default:
		return "unknown"
	}
}

type Broker interface {
	Publish(topic Topic, message []byte) error
}
