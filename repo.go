package main

import (
  "time"
  "crypto/sha256"
  "encoding/hex"
  "strconv"
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

func CalculatePrevHash() string {
  if len(blocks) > 0 {
    return blocks[len(blocks)-1].BlockHash
  } else {
    return "0"
  }
}

func CalculateTimestamp() int64 {
  return time.Now().Unix()
}

func CalculateRows() Records {
  return records[len(records) - Dimension:]
}

func CalculateBlockHash(prevHash string, rows Records, timestamp int64) string {
  hasher := sha256.New()

  hasher.Write([]byte(prevHash))

  for _, row := range rows {
    hasher.Write([]byte(row.Data))
  }

  hasher.Write([]byte(strconv.FormatInt(timestamp, 10)))

  return hex.EncodeToString(hasher.Sum(nil))
}

func RepoCreateRecord(r Record) Record {
  records = append(records, r)

  if len(records) % Dimension == 0 {
    prevHash := CalculatePrevHash()
    rows := CalculateRows()
    timestamp := CalculateTimestamp()
    blockHash := CalculateBlockHash(prevHash, rows, timestamp)

    RepoCreateBlock(
      Block{
        PreviousBlockHash: prevHash,
        Rows: rows,
        Timestamp: timestamp,
        BlockHash: blockHash,
      },
    )
  }

  return r
}

func RepoCreateBlock(b Block) Block {
  blocks = append(blocks, b)
  return b
}
