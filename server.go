package main

import (
  "fmt"
  "net/http"
  "strings"
  "log"
  "github.com/gorilla/mux"
)


// POST / 
// {data: 'somadata'}

// GET /last_blocks/10
// 

func Index(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  fmt.Println(r.Form)
  fmt.Println("path", r.URL.Path)
  fmt.Println("scheme", r.URL.Scheme)
  fmt.Println(r.Form["url_long"])

  for k, v := range r.Form {
      fmt.Println("key:", k)
      fmt.Println("val:", strings.Join(v, ""))
  }
  fmt.Fprintf(w, "Hello Maksim!")
}

func AddData(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  fmt.Fprintf(w, "AddData")
}

func LastBlocks(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  fmt.Fprintf(w, "LastBlocks")
}

func main() {
  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/", Index)
  router.HandleFunc("/add_data", AddData)
  router.HandleFunc("/last_blocks/{count}", LastBlocks)

  err := http.ListenAndServe(":9000", router)

  if err != nil {
      log.Fatal("ListenAndServe: ", err)
  }
}
