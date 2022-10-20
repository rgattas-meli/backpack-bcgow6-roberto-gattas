package calculadora
import "sort"

//a := []int{5, 3, 4, 7, 8, 9}
func ordenar(a []int) []int{
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	return a

}