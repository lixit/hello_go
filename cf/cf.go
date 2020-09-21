package cf

import (
    "fmt"
    "os"
    "strconv"

    "github.com/lixit/hello_go/tempconv"
)

func Main() {
    for _, arg := range os.Args[1:] {
        t, err := strconv.ParseFloat(arg, 64)
        if err != nil {
            fmt.Fprintf(os.Stderr, "cf: %v\n", err)
            os.Exit(1)
        }
        f := tempconv.Fahrenheit(t)
        c := tempconv.Celsius(t)
        fmt.Printf("%s = %s, %s = %s\n",
        f, tempconv.FToC(f), c, tempconv.CToF(c))
    }
}
