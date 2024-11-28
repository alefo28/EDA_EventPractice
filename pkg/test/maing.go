package main

import "fmt"

func main() {
	evento := []string{"test", "test1", "test2", "test3"}
	evento = evento[2:]
	fmt.Println(evento)
}
