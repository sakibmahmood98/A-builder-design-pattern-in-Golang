package main

import "fmt"

func main() {
    normalBuilder := getBuilder("normal")


    director := newDirector(normalBuilder)
    normalSomething := director.buildSomething()

    fmt.Printf("Normal Something Data: %d\n", normalSomething.data1)
    fmt.Printf("Normal Something add(): %d\n", normalSomething.data2)
    fmt.Printf("Normal Something delete(): %d\n", normalSomething.data3)
}
