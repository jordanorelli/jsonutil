package main

import (
    "encoding/json"
    "fmt"
    "github.com/jordanorelli/jsonutil"
    "log"
    "io"
    "os"
)

func main() {
    c, e := make(chan *json.RawMessage), make(chan error)
    go jsonutil.Split(os.Stdin, c, e)
    for {
        select {
        case raw := <-c:
            jsonutil.PrettyPrint(raw)
            fmt.Println("")
        case err := <-e:
            if err == io.EOF {
                fmt.Println("")
                os.Exit(0)
            }
            log.Fatal(err)
        }
    }
}
