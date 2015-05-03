package main

import (
    "fmt"
    "flag"
    "github.com/dotcloud/docker/pkg/namesgenerator"
)

func main() {
    n := flag.Int("n", 1, "Number of generate names")
    flag.Parse()

    for i := 0; i < *n; i++ {
      fmt.Println(namesgenerator.GetRandomName(0))
    }
}
