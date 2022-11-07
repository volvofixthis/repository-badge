package main

import (
	"fmt"
	"os"
)

var counter int

func init() {
	counter = 99
}

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       string            `json:"body,omitempty"`
}

func Main(in Request) (*Response, error) {
	counter += 1
	if in.Name == "" {
		in.Name = "stranger"
	}

	return &Response{
		Body: fmt.Sprintf("Hello %s number %d! %s", in.Name, counter, os.Getenv("WELCOME_GREETING")),
	}, nil
}
