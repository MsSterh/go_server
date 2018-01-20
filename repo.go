package main

import (
  "time"
)

var records Records
var blocks Blocks

const Dimension = 5

// Give us some seed data
func init() {
  RepoCreateRecord(Record{Data: "First transaction"})
  RepoCreateRecord(Record{Data: "Second transaction"})
  RepoCreateRecord(Record{Data: "Third transaction"})
}

func RepoCreateRecord(r Record) Record {
  records = append(records, r)

  if len(records) % Dimension == 0 {
    RepoCreateBlock(
      Block{
        PreviousBlockHash: "1",
        Rows: records[len(records) - Dimension:],
        Timestamp: time.Now().Unix(),
        BlockHash: "2",
      },
    )
  }

  return r
}

func RepoCreateBlock(b Block) Block {
  blocks = append(blocks, b)
  return b
}
