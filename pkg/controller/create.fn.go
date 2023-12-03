package controller

func CreateController(hostname string, port int) Controller {
	return Controller{Hostname: hostname, Port: port}
}