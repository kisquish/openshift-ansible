package main

import (
	//	"bytes"
	"encoding/json"
	"fmt"
)

type node struct {
	Name      string `json:"name"`
	PrivateIP string `json:"private IP"`
	PublicIP  string `json:"public IP"`
}

type nodeList []node

type nodeGroup map[string]nodeList

type cluster map[string]nodeGroup

type AnsibleResponse struct {
	Msg map[string]cluster `json:"msg"`
}

func (a AnsibleResponse) String() string {
	json, _ := json.Marshal(a)
	return string(json)
}

func main() {
	fmt.Println("Hello")
}

func New(str string) {
	res := AnsibleResponse{}
	if err := json.Unmarshal([]byte(str), &res); err != nil {
		panic(err)
	}
	fmt.Println(res)
}
