package main

var records Records
var blocks RecordsBlocks

// Give us some seed data
func init() {
  RepoCreateRecord(Record{Data: "First transaction"})
  RepoCreateRecord(Record{Data: "Second transaction"})
  RepoCreateRecord(Record{Data: "Third transaction"})
}

func RepoCreateRecord(r Record) Record {
  records = append(records, r)
  return r
}
