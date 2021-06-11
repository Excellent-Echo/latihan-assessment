package main

import (
	"books-project/config"
	"fmt"
)

func main() {
	config.Connection()
	fmt.Print("Sukses konek database")
}
