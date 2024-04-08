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
	awsToken := "AKIALALEMEL33243OLIA"
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

// Функция с уязвимостью безопасности (запись в файл)
func writeToFile(filename, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Секретное содержимое файла
	secretData := "Secret data: " + awsToken

	_, err = file.WriteString(secretData)
	if err != nil {
		return err
	}

	// Добавление еще одного секрета
	anotherSecret := "Another secret: 1234567890"
	_, err = file.WriteString("\n" + anotherSecret)
	if err != nil {
		return err
	}

	return nil
}
