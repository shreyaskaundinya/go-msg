package subscriber

func CreateSubscriber(hostname string, port int) Subscriber {
	return Subscriber{Hostname: hostname, Port: port}
}