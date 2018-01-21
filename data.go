package main

type Record struct {
  Data      string    `json:"data"`
}

type Records []Record
type Rows    []string

type Block struct {
  PreviousBlockHash string  `json:"previous_block_hash"`
  Rows              Rows    `json:"rows"`
  Timestamp         int64   `json:"timestamp"`
  BlockHash         string  `json:"block_hash"`
}

type Blocks []Block
