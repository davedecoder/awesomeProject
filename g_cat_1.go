package main

import (
	u "awesomeProject/utils"
	"io"
	"log"
	"os"
)


func getFile(name string) (*os.File, func(), error){
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}
	return file, func() {
		u.SafeFileClose(file)
	}, nil
}


func main() {
	if len(os.Args) < 2{
		log.Fatalln("No File specified")
	}
	f, closerFile, e := getFile(os.Args[1])
	if e != nil {
		log.Fatalln(e)
	}
	defer closerFile()

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
