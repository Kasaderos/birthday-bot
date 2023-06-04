package notifier

type Notifier interface {
	Send(Message) error
}