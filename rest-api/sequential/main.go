package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const RandomUpperThreshold int = 5

func sequentialHandler(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	response := searchOnFacebook() + ", " + searchOnGoogle() + ", " + searchOnTwitter()

	fmt.Fprintf(w, "%s", response)
}

func searchOnFacebook() string {
	time.Sleep(time.Duration(rand.Intn(RandomUpperThreshold) * int(time.Second)))
	fmt.Printf("%d|%s|%s|%d|%d|1|FF0000\n", time.Now().Unix(), "127.0.0.1", "/facebook", 200, 2048)
	return "from facebook"
}

func searchOnGoogle() string {
	time.Sleep(time.Duration(rand.Intn(RandomUpperThreshold) * int(time.Second)))
	fmt.Printf("%d|%s|%s|%d|%d|1|00FF00\n", time.Now().Unix(), "127.0.0.1", "/google", 200, 2048)
	return "from google"
}

func searchOnTwitter() string {
	time.Sleep(time.Duration(rand.Intn(RandomUpperThreshold) * int(time.Second)))
	fmt.Printf("%d|%s|%s|%d|%d|1|0000FF\n", time.Now().Unix(), "127.0.0.1", "/twitter", 200, 2048)
	return "from twitter"
}

func main() {
	http.HandleFunc("/sequential", sequentialHandler)
	http.ListenAndServe(":8081", nil)
}
