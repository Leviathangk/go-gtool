package main

import (
	"fmt"
	"gtool/pathlib"
	"log"
)

func main() {
	p := pathlib.NewPath("D:\\Download")

	err := p.Iter(func(path *pathlib.Path, err error) error {
		fmt.Println(path)
		return err
	})

	if err != nil {
		log.Fatalln(err)
	}
}
