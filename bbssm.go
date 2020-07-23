package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
    "net/http"
    "bytes"
    "io/ioutil"
)


func main() {
    s := bufio.NewScanner(os.Stdin)
    for s.Scan() {
        var linetext = s.Text()
        log.Println("line", linetext)
        
        if os.Args[1] == "project" {
            addProject(linetext)
        }
        if os.Args[1] == "scopeline" {
            addScopeLine(linetext, os.Args[2])
        }        
    }
}


func addProject(codename string) {
    url := "http://bbs-scopemanager-service:7000/api/project"
    var jsonStrStart = []byte(`{"code":"`)
    var jsonStrEnd = []byte(`"}`)
    var part1 = append(jsonStrStart, codename...)
    var completeValue = append(part1, jsonStrEnd...)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(completeValue))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    //fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
    
}

func addScopeLine(scopeline string, project_id string) {
    url := "http://bbs-scopemanager-service:7000/api/scope_line"
    var jsonStrStart = []byte(`{"lineitem":"`)
    var jsonStrNext = []byte(`", "project_id":"`)
    var jsonStrEnd = []byte(`"}`)
    var part1 = append(jsonStrStart, scopeline...)
    var part2 = append(part1, jsonStrNext...)
    var part3 = append(part2, project_id...)
    var completeValue= append(part3, jsonStrEnd...)
    
    
    
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(completeValue))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    //fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
    
}

