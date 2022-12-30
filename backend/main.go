package main

import (
    "os"
    "fmt"
    "net/http"
    "encoding/json"
    "log"
	"github.com/joho/godotenv"
    "io/ioutil"
    "strings"
) 

type Search struct {
    Query string `json:"searchTerm"`
    MaxResults int `json:"maxResults"`
}

var apiKey string

func (s Search) String() string {
    return fmt.Sprintf("Query: %s, MaxResults: %d", s.Query, s.MaxResults)
}

func enableCors(w *http.ResponseWriter) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func authenticatedRequestReddit(body Search) {
    json_body, _ := json.Marshal(body)
    fmt.Println(string(json_body))

    godotenv.Load()
    apiKey := os.Getenv("REDDIT_SECRET")

    client := &http.Client{}
    req, _ := http.NewRequest("GET", "https://oauth.reddit.com/r/all/search", strings.NewReader(string(json_body)))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer " + apiKey)
    res, _ := client.Do(req)

    response_body, _ := ioutil.ReadAll(res.Body)
    fmt.Println(string(response_body))
}

// Send an authenticated request to reddit


func createSearch(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
    w.Header().Set("Content-Type", "application/json")
    if r.Method == "POST" {
        var s Search
        err := json.NewDecoder(r.Body).Decode(&s)
        s.MaxResults = 10
        if err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        redditURI := fmt.Sprintf("https://www.reddit.com/search?q=%s&sort=relevance&t=all", s.Query)
        // print the response body of a GET request send to redditURI

        fmt.Println(redditURI)

        // respond with the search
        fmt.Fprintf(w, s.String())
    }
}


func main() {
    godotenv.Load()
    apiKey = os.Getenv("REDDIT_SECRET")
    mux := http.NewServeMux()
    mux.HandleFunc("/search", createSearch)

    err := http.ListenAndServe(":8080", mux)
    fmt.Println("Server started on port 8080")
    if err != nil {
        log.Fatal(err)
    }
}

