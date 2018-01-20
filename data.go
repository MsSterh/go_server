package main

import "time"

type Record struct {
  Data      string    `json:"data"`
}

type Records []Record

type RecordsBlock struct {
  PreviousBlockHash string    `json:"previous_block_hash"`
  Rows              Records   `json:"rows"`
  Timestamp         time.Time `json:"timestamp"`
  BlockHash         string    `json:"block_hash"`
}

type RecordsBlocks []RecordsBlock
