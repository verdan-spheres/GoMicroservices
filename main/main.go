package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := &http.Server{  // use this address for pod binding
		Addr: ":3000", // binding to this ports
		Handler: http.HandlerFunc(basicHandler),
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("failed to listen to server", err)
	}
}

// *http.Requst is restrict inbound http request 

func basicHandler (w http.ResponseWriter , r *http.Request){
	 w.Write([]byte("Hello , World !"))
}