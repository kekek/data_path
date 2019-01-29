package main

import (
	"fmt"
	"github.com/kekek/data_path/testdata"
)

func main() {
	l  := testdata.Path("ca.pem")
	fmt.Println(l)
}