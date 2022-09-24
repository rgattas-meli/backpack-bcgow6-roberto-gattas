package main
import "fmt"


var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

func main() {
	fmt.Println(employees["Benjamin"])
	count:= 0
	for _,s := range employees	{
		if s > 21 {
            count++
        }
	}
	fmt.Println("Cantidad mayor a 21 es:", count)
    employees["Federico"] = 25
	delete(employees,"Pedro")
	fmt.Println(employees)


}