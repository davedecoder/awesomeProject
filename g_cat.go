package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func SafeFileClose(file *os.File) {
	if err := file.Close(); err != nil {
		println(fmt.Sprintln("Error Closing File", err))
	}
}

func main() {
	if len(os.Args) < 2{
		log.Fatalln("No File specified")
	}
	f, e := os.Open(os.Args[1])
	if e != nil {
		log.Fatalln(e)
	}
	defer SafeFileClose(f)

	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		_, _ = os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}
