package main


import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
)


func sendRequest(url string, ch chan<- string, printResponse int) {
	start := time.Now()
	client := &http.Client{}
	resp, err := client.Get(url)

	if err != nil {
		log.Fatal(err)
		return
	}

	defer resp.Body.Close()
	ms := time.Since(start).Milliseconds()
	if printResponse == 1 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("Response:", string(body))
	}
	statusCode := resp.StatusCode
	if statusCode == 200 {
		ch <- fmt.Sprintf("Response Time: %dms | Status Code: %d", ms, statusCode)
	} else {
		ch <- fmt.Sprintf("Server can't handle the requests: %dms | Status Code: %d", ms, statusCode)
	}
}


func main() {
	if len(os.Args) != 4 {
		fmt.Printf("Usage: %s <url> <parallel requests> <print response>\n ", os.Args[0])
		fmt.Printf("Example for 1 request printing response: %s https://www.google.com 1 1\n ", os.Args[0])
		fmt.Printf("Example for 100 parallel requests not printing response: %s https://www.google.com 100 0\n ", os.Args[0])
		os.Exit(1)
	}
	url := os.Args[1]
	tries, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Type the number of parralel requests")
		os.Exit(1)
	}
	printResponse, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Type 1 for print the response")
		os.Exit(1)
	}

	ch := make(chan string)
	for i := 1; i <= tries; i++ {
		go sendRequest(url, ch, printResponse)
	}

	for i := 1; i <= tries; i++ {
		fmt.Println(<-ch)
	}
	os.Exit(0)
}