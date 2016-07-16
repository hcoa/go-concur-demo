package main

import (
	"fmt"
	"net/http"
	"time"
)

const (
	LogFormatString string = "%d|%s|%s|%d|%d\n"
	Hostname        string = "127.0.0.1"
)

func sequentialHandler(w http.ResponseWriter, r *http.Request) {
	response := searchOnFacebook() +  ", " + searchOnGoogle() + ", " + searchOnTwitter()

	fmt.Fprintf(w, "%s", response)
}

func searchOnFacebook() string {
	time.Sleep(100 * time.Millisecond)
	fmt.Printf(LogFormatString, time.Now().Unix(), Hostname, "/facebook", 200, 2048)
	return "from facebook"
}

func searchOnGoogle() string {
	time.Sleep(3300 * time.Millisecond)
	fmt.Printf(LogFormatString, time.Now().Unix(), Hostname, "/google", 200, 2048)
	return "from google"
}

func searchOnTwitter() string {
	time.Sleep(2000 * time.Millisecond)
	fmt.Printf(LogFormatString, time.Now().Unix(), Hostname, "/twitter", 200, 2048)
	return "from twitter"
}

func main() {
	http.HandleFunc("/sequential", sequentialHandler)
	http.ListenAndServe(":8081", nil)
}
