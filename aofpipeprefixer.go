package main

import (
	"flag"
	"fmt"
	"github.com/Prokopiev/aof"
	"io"
	"os"
	"strings"
)

// Create a new type for a list of Strings
type stringList []string

// Implement the flag.Value interface
func (s *stringList) String() string {
	return fmt.Sprintf("%v", *s)
}

func (s *stringList) Set(value string) error {
	*s = strings.Split(value, ",")
	return nil
}

var baseList stringList

func main() {
	if len(os.Args) < 2 {
		fmt.Println("list of bases is required")
		os.Exit(1)
	}

	flag.Var(&baseList, "bases", "Comma separated list of bases (required)")
	flag.Parse()

	reader := aof.NewBufioReader(os.Stdin)
	for {
		op1, err := reader.ReadOperation()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		valid := false
		for _, element := range baseList {
			if strings.HasPrefix(op1.Key, element+":") {
				valid = true
				break
			}
		}

		if !valid {
			continue
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
