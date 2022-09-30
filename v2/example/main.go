package main

import (
	"fmt"

	"github.com/wallet-manager/treemap/v2"
)

func main() {
	tr := treemap.New[int, string]()
	tr.Set(1, "World")
	tr.Set(0, "Hello")
	for it := tr.Iterator(); it.Valid(); it.Next() {
		fmt.Println(it.Key(), it.Value())
	}
}
