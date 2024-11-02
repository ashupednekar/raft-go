package main

import "github.com/ashupednekar/raft-go/internal"

func main(){
  go internal.Serve()

  internal.AppendEntries()

  select{}
}
