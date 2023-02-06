package main

import (
	"fmt"
	"gtool/pathlib"
	"log"
)

func main() {
	p := pathlib.Path("D:\\Download")

	err := p.Iter(func(path *pathlib.Handler, err error) error {
		fmt.Println(path)
		return err
	})

	if err != nil {
		log.Fatalln(err)
	}
}
