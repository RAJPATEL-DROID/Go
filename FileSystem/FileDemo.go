package FileSystem

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func FileDemo() {
	//p := fmt.Println
	var newFile *os.File
	fmt.Printf("%T\n", newFile)

	var err error
	//newFile, err = os.Create("test.txt")

	//if err != nil {
	//	//fmt.Println(err)
	//	//os.Exit(1)
	//
	//	log.Fatal(err)
	//}

	//os.Truncate(newFile.Name(), 0) // 0 means completely empty the file,
	// n -> everything past n bytes will be cleared

	//file, err := os.Open("test.txt") // Readonly mode

	file, err := os.OpenFile("test.txt", os.O_APPEND|os.O_WRONLY, 0655)
	//if err != nil {
	//	log.Fatal(err)
	//}
	defer file.Close() // This get called once every other function in surrounding(main()) get returns from stack

	//var fileInfo os.FileInfo
	//fileInfo, err = os.Stat("test.txt")
	//
	//p("File NAme", fileInfo.Name())
	//p("Size", fileInfo.Size())
	//p("Mode", fileInfo.Mode())
	//p("ModTime", fileInfo.ModTime())
	//p("Permissions", fileInfo.Mode())

	//fileInfo, err = os.Stat("test.txt")
	//if err != nil {
	//	if os.IsNotExist(err) {
	//		log.Fatal("File does not exist")
	//	}
	//}

	//err = os.Rename("test.txt", "b.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}

	//err = os.Remove("b.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}

	byteSlice := []byte("I am doing Golang")
	//byteWritten, err := file.Write(byteSlice)
	//if err != nil {
	//	log.Fatal(err)
	//}

	//log.Printf("Bytes written %d\n", byteWritten)

	//bs := []byte("Go is cool")
	////err = ioutil.WriteFile("test.txt", bs, 0644) // Deprecated, use os.WriteFile()
	//err = os.WriteFile("test1.txt", bs, 0644)
	//if err != nil {
	//	log.Fatal(err)
	//}
	bufferedWriter := bufio.NewWriter(file)
	bs := []byte{99, 77, 91}
	bytesWritten, err := bufferedWriter.Write(bs)
	if err != nil {
		fmt.Println(err)
	}
	log.Printf("bytes written to buffer(Not file): %d", bytesWritten)

	bytesAvailable := bufferedWriter.Available()
	log.Printf("bytes available in buffer: %d", bytesAvailable)

	bytesWritten, err = bufferedWriter.WriteString("\n Random String")
	if err != nil {
		fmt.Println(err)
	}

	unflushedBufferedSize := bufferedWriter.Buffered()
	log.Printf("Buffered pushed :%d\n", unflushedBufferedSize)

	// Flush all data to file
	//bufferedWriter.Flush()

	// Clear the all bytes in the buffer
	bufferedWriter.Reset(bufferedWriter)

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("data read : %d", data)

}
