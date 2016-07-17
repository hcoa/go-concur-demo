package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const RandomUpperThreshold int = 5

func concurrentHandler(w http.ResponseWriter, r *http.Request) {
	receiver := make(chan string)
	rand.Seed(time.Now().UnixNano())

	go searchOnFacebook(receiver)
	go searchOnGoogle(receiver)
	go searchOnTwitter(receiver)

	response := <-receiver + ", " + <-receiver + ", " + <-receiver
	fmt.Fprintf(w, response)
}

func searchOnFacebook(ch chan string) {
	time.Sleep(time.Duration(rand.Intn(RandomUpperThreshold) * int(time.Second)))
	fmt.Printf("%d|%s|%s|%d|%d|1|FF0000\n", time.Now().Unix(), "127.0.0.1", "/facebook", 200, 2048)
	ch <- "from facebook"
}

func searchOnGoogle(ch chan string) {
	time.Sleep(time.Duration(rand.Intn(RandomUpperThreshold) * int(time.Second)))
	fmt.Printf("%d|%s|%s|%d|%d|1|00FF00\n", time.Now().Unix(), "127.0.0.1", "/google", 200, 2048)
	ch <- "from google"
}

func searchOnTwitter(ch chan string) {
	time.Sleep(time.Duration(rand.Intn(RandomUpperThreshold) * int(time.Second)))
	fmt.Printf("%d|%s|%s|%d|%d|1|0000FF\n", time.Now().Unix(), "127.0.0.1", "/twitter", 200, 2048)
	ch <- "from twitter"
}

func main() {
	http.HandleFunc("/concurrent", concurrentHandler)
	http.ListenAndServe(":8080", nil)
}
