package main

import (
	"github.com/Prokopiev/aof"
	"io"
	"os"
)

func main() {
	var currentDatabase = "0"
	reader := aof.NewBufioReader(os.Stdin)
	for {
		op1, err := reader.ReadOperation()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		if op1.Command == "SELECT" {
			currentDatabase = op1.Arguments[0]
			//Databases are not supported - skipping command
			continue
		}
		if op1.Key != "" {
			op1.Key = currentDatabase + ":" + op1.Key
		}
		//Databases are not supported - dangerous commands
		if op1.Command == "FLUSHDB" || op1.Command == "FLUSHALL" {
			continue
		}
		err = op1.ToAof(os.Stdout)
		if err != nil {
			panic(err)
		}
	}
}
