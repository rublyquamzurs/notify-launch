package core

var (
	senders    map[string]chan string
	registers  chan string
	unregister chan string
)

func initDebug() {
	senders = make(map[string]chan string)
	registers = make(chan string)
	unregister = make(chan string)
}

func newChannel(key string) chan string {
	c := make(chan string)
	senders[key] = c
	return c
}

func getChannel(key string) chan string {
	if c, ok := senders[key]; ok {
		return c
	}
	return nil
}

func freeChannel(key string) {
	if c, ok := senders[key]; ok {
		close(c)
	}
}
