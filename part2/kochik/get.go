package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Message struct {
	FirstName string
	LastName  string
}

type PostMessage struct {
	Email, FirstName, LastName string
}

func setMyName() {
	message := PostMessage{Email: "parsa_ghorbani@outlook.com", FirstName: "Parsa", LastName: "Ghorbani"}
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(message)
	http.Post("http://207.154.236.19:8000/name", "application/json;charset=utf-8", buffer)
}

func getMyName() {
	req, _ := http.NewRequest("GET", "http://207.154.236.19:8000/name", nil)
	query := req.URL.Query()
	query.Add("email", "parsa_ghorbani@outlook.com")
	req.URL.RawQuery = query.Encode()
	timeout := time.Duration(5 + time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	resp, _ := client.Do(req)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	json_body := string(body)
	var message Message
	json.Unmarshal([]byte(json_body), &message)
	fmt.Println(message)
}

func main() {
	getMyName()
	setMyName()
	getMyName()
}
