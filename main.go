package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	port      string
	directory string
	useTLS    bool
)

func init() {
	flag.StringVar(&port, "p", "25565", "Port to serve on")
	flag.StringVar(&directory, "d", "./root", "Directory to serve")
	flag.BoolVar(&useTLS, "tls", false, "Set to use TLS (Requires cert and key file)")
	flag.Parse()

	err := os.MkdirAll(directory, 0755)
	check(err)
}

func main() {

	startServer()
}

func startServer() {
	http.Handle("/", http.FileServer(http.Dir(directory)))
	log.Printf("Serving folder '%s'\n", directory)

	if useTLS {
		fmt.Printf("https://%s:%s\nhttps://localhost%s\n", getIP(), port, port)
		log.Fatal(http.ListenAndServeTLS(":"+port, "cert.pem", "key.pem", nil))
	} else {
		fmt.Printf("http://%s:%s\nhttp://localhost%s\n", getIP(), port, port)
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}
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
