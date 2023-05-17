package worker

import (
	"fmt"
	"log"
	"net/http"

	http3 "github.com/quic-go/quic-go/http3"
)

func CreateWorker(hostname string, port int) Worker {
	return Worker{Hostname: hostname, Port: port}
}


func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func (w *Worker) Serve() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", pingHandler)
	
	addr := fmt.Sprintf("localhost:%d", w.Port)

	err := http3.ListenAndServeQUIC(addr, "quic.pem", "quic.key", mux)

	if (err != nil) {
		log.Fatal(err.Error())
	}

	fmt.Printf("Serving at %s", addr)
}