package main

import (
	"fmt"

	"github.com/lixit/hello_go/cf"
	"github.com/lixit/hello_go/hello"
	"github.com/lixit/hello_go/tempconv"
)

func main() {
	hello.Main()
	fmt.Printf("Brrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
	cf.Main()
}
