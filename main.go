package main

import (
	"fmt"
	"github.com/DvorakDavid7/dynamic-array/dynarr"
	"log"
)

func main() {
	arr := dynarr.NewDynArray()

	for i := 0; i < 100; i++ {
		arr.Append(dynarr.Int(i))
	}

	for i := 0; i < arr.Size(); i++ {
		val, err := arr.Get(i)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(val)
	}
}
