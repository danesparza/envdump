package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	//	Prep the dump file:
	f, err := os.Create("envdump.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//	Dump the environment variables to the
	f.WriteString("Environment variables:\r\n")

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		stringtowrite := fmt.Sprintf("%s = %s\r\n", pair[0], pair[1])
		f.WriteString(stringtowrite)
		fmt.Printf(stringtowrite)
	}
}
