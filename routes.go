package main

import "net/http"

type Route struct {
  Name        string
  Method      string
  Pattern     string
  HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
  Route{
    "Blocks",
    "GET",
    "/",
    AllBlocksShow,
  },
  Route{
    "Records",
    "GET",
    "/all_data",
    AllRecordsShow,
  },
  Route{
    "AddData",
    "POST",
    "/add_data",
    AddData,
  },
  Route{
    "LastBlock",
    "GET",
    "/last_blocks/{blockId}",
    LastBlockShow,
  },
}
