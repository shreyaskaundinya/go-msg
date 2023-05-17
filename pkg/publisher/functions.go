package publisher

func CreatePublisher(hostname string, port int) Publisher {
	return Publisher{Hostname: hostname, Port: port}
}