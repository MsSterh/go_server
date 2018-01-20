package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "io"
  "io/ioutil"

  "github.com/gorilla/mux"
)

func AllBlocksShow(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)

  if err := json.NewEncoder(w).Encode(blocks); err != nil {
    panic(err)
  }
}

func AllRecordsShow(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)

  if err := json.NewEncoder(w).Encode(records); err != nil {
    panic(err)
  }
}

func LastBlockShow(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  blockId := vars["blockId"]
  fmt.Fprintln(w, "Last block show:", blockId)

  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusOK)

  if err := json.NewEncoder(w).Encode(todos); err != nil {
    panic(err)
  }
}

func AddData(w http.ResponseWriter, r *http.Request) {
  var record Record
  body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

  if err != nil {
    panic(err)
  }

  if err := r.Body.Close(); err != nil {
    panic(err)
  }

  if err := json.Unmarshal(body, &record); err != nil {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(422) // unprocessable entity

    if err := json.NewEncoder(w).Encode(err); err != nil {
      panic(err)
    }
  }

  t := RepoCreateRecord(record)
  w.Header().Set("Content-Type", "application/json; charset=UTF-8")
  w.WriteHeader(http.StatusCreated)

  if err := json.NewEncoder(w).Encode(t); err != nil {
    panic(err)
  }
}
