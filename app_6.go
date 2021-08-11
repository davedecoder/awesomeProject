package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func gracefullyCloseFile(file *os.File, label string){
	fmt.Printf("Closing %s\n", label)
	err := file.Close()
	if err != nil {
		panic(fmt.Sprintf("Error Closing Gracefully: %s", err))
	}
}

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	defer fmt.Println(i)
	i++
	return
}

func c() (i int){
	defer func(){ i++ }()
	return 1
}

func CopyFile(srcName, dstName string) (written int64, err error){
	src, err := os.Open(srcName)
	if err != nil {
		return 0, err
	}
	defer gracefullyCloseFile(src, "Source")

	dst, err := os.Create(dstName)
	if err != nil {
		return 0, err
	}
	defer gracefullyCloseFile(dst, "Destiny")

	return io.Copy(dst, src)
}

func main() {
	a()
	c()
	fmt.Println("Copying")
	if len(os.Args) < 3{
		log.Fatalln("Specify Source and Destiny File")
	}
	nBytes, e := CopyFile(os.Args[1], os.Args[2])
	if e != nil {
		log.Fatalf("Error copying %s\n", e)
	}
	fmt.Printf("Succesfully copied %d bytes\n", nBytes)
}
