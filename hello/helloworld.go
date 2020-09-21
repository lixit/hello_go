package hello

import (
	"fmt"

	"github.com/lixit/hello_go/tempconv"
)

func Main() {
	fmt.Println("Hello, 世界")
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
}
