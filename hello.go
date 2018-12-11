package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	fmt.Print("Hello, world!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/test", test)

	log.Fatal(http.ListenAndServe(":8099", nil))

}

func test(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(body))
	var item item
	err = json.Unmarshal(body, &item)
	if err != nil {
		panic(err)
	}

	//log.Println(item.Request.Command)
	//log.Println(item.Session.SkillId)
	//log.Println(item.Version)
	//
	//log.Println(item.Request.Nlu.Tokens[0])
	//
	//log.Println(item.Request.Nlu.Entities[0].Type)
}

type item struct {
	Request request `json:"request"`
	Session session `json:"session"`
	Version string  `json:"version"`
}

type request struct {
	Command           string `json:"command"`
	OriginalUtterance string `json:"original_utterance"`
	Type              string `json:"type"`
	Nlu               nlu    `json:"nlu"`
}

type session struct {
	Version   string `json:"version"`
	New       bool   `json:"new"`
	MessageId int    `json:"message_id"`
	SessionId string `json:"session_id"`
	SkillId   string `json:"skill_id"`
	UserId    string `json:"user_id"`
}

type nlu struct {
	Tokens   []string `json:"tokens"`
	Entities []entity `json:"entities"`
}

type entity struct {
	Type string `json:"type"`
}

