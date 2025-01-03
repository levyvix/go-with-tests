package main

import "fmt"

func Ola(name string) string {
	return fmt.Sprintf("Ola, %s", name)
}
func main() {
	fmt.Println(Ola("Levy"))
}
