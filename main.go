package main

import (
	"fmt"
	"gtool/file"
)

func main() {
	fmt.Println("Hello world!")
	info, err := file.Info("D:\\Windows.iso")
	if err != nil {
		return
	}

	fmt.Println(info.Name())
	fmt.Println(info.ModTime())
	fmt.Println(info.Size())
}
