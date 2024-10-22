package contract

type Publish interface {
	Publish(event string, message string)
}

type Consume interface {
	Consume(event string) <-chan string
}
