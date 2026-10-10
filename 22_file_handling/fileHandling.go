package main

import (
	// "bufio"
	"fmt"
	"os"
)

func main() {
	// f, err := os.Open("./examples.txt")
	// if err != nil{
	// 	panic(err)
	// } 
	// fileInfo, err := f.Stat()
	// if err != nil{
	// 	panic(err)
	// } 
	// fmt.Println("FileName=", fileInfo.Name())
	// fmt.Println("File or folder=", fileInfo.IsDir())
	// fmt.Println("File Size=", fileInfo.Size())
	// fmt.Println("File Permissions=", fileInfo.Mode())
	// fmt.Println("File Last Modified at=", fileInfo.ModTime())


	//read file
	// f, err := os.Open("examples.txt")
	// if err != nil{
	// 	panic(err)
	// }

	// defer f.Close()

	// buf := make([]byte, 12)

	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("data", d)
	// for i := range buf{
	// 	fmt.Print(string(buf[i]))
	// }
	// fmt.Println()
	// data, err := os.ReadFile("examples.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(data))

	//read folders

	// dir, err := os.Open("../")
	// if err != nil{
	// 	panic(err)
	// }
	// defer dir.Close()
	// folderInfo, err := dir.ReadDir(-1)
	// if err != nil{
	// 	panic(err)
	// }
	// for _, fi:= range folderInfo {
	// 	fmt.Println(fi.Name())
	// }

	//create file

	// fmt.Println()
	// f, err := os.Create("example2.txt")
	// if err != nil{
	// 	panic(err)
	// }
	// defer f.Close()

	// // f.WriteString("hi go")
	// // f.WriteString("nice lang")
	
	// bytes := []byte("Hello Golang")

	// f.Write(bytes)


	//read from one file and write to another file
	// sourceFile, err := os.Open("./examples.txt")
	// if err != nil{
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	// destFile, err := os.Create("./example2.txt")
	// if err != nil{
	// 	panic(err)
	// }

	// defer destFile.Close()

	// reader := bufio.NewReader(sourceFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil{
	// 		if err.Error() != "EOF"{
	// 			panic(err)
	// 		}
	// 		break
	// 	}
	// 	e := writer.WriteByte(b)
	// 	if e != nil{
	// 		panic(e)
	// 	}
	// }
	// if err := writer.Flush(); err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Written to destination file from source file successful.")

	// delete a file
	// sourceFile, err := os.Open("./example2.txt")
	// if err != nil{
	// 	panic(err)
	// }

	// defer sourceFile.Close()

	err:= os.Remove("example2.txt")
	if err != nil{
		panic(err)
	}
	fmt.Println("File Deleted Successfully")
}
