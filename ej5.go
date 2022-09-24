package main
import "fmt"
import "strings"
var palabra string

func main() {
	palabra = "Roberto"
	s := strings.Split(palabra, "")
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(palabra)
}