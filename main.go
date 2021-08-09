package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	port      = "25565"
	directory = "./root"
)

func init() {
	err := os.MkdirAll(directory, 0755)
	check(err)
}

func main() {

	startServer()
}

func startServer() {
	http.Handle("/", http.FileServer(http.Dir(directory)))

	log.Printf("Serving folder '%s' from %s:%s\n", directory, getIP(), port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getIP() string {
	res, err := http.Get("https://api.ipify.org")
	check(err)

	ip, err := ioutil.ReadAll(res.Body)
	check(err)

	return string(ip)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
