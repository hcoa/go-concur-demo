package main

import (
	"net/http"
	"fmt"
	"time"
)

const LogFormatString string = "%d|%s|%s|%d|%d\n"

func concurrentHandler(w http.ResponseWriter, r *http.Request) {
	receiver := make(chan string)

	go searchOnFacebook(receiver)
	go searchOnGoogle(receiver)
	go searchOnTwitter(receiver)

	response := <-receiver + ", " + <-receiver + ", " + <-receiver
	fmt.Fprintf(w, response)
}

func searchOnFacebook(ch chan string)  {
	time.Sleep(100 * time.Millisecond)
	fmt.Printf(LogFormatString, time.Now().Unix(), "127.0.0.1", "/facebook", 200, 2048)
	ch <- "from facebook"
}

func searchOnGoogle(ch chan string)  {
	time.Sleep(3300 * time.Millisecond)
	fmt.Printf(LogFormatString, time.Now().Unix(), "127.0.0.1", "/google", 200, 2048)
	ch <- "from google"
}

func searchOnTwitter(ch chan string)  {
	time.Sleep(2000 * time.Millisecond)
	fmt.Printf(LogFormatString, time.Now().Unix(), "127.0.0.1", "/twitter", 200, 2048)
	ch <- "from twitter"
}

func main() {
	http.HandleFunc("/concurrent", concurrentHandler)
	http.ListenAndServe(":8080", nil)
}
