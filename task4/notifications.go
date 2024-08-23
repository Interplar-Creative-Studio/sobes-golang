package task4

type Channel interface {
	Send(message string) error
}

type NotificationSender struct {
	// TODO
}

func (s *NotificationSender) Send(message string, channels ...string) error {
	// TODO
}
