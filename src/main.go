package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var awsToken string

func main() {
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	// load secret via env
	awsToken = os.Getenv("AWS_TOKEN")
	if awsToken != "" {
		fmt.Println("AWS Token loaded successfully:", awsToken)
	}

	// Function that potentially introduces security vulnerability (using strings.Contains)
	if strings.Contains("Hello, World!", "World") {
		fmt.Println("Substring 'World' found")
	}

	// HTTP server
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func GetAWSToken() string {
	return awsToken
}
