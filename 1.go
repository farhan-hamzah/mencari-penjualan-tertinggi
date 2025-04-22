package main
import "fmt"
const NMAX int = 1024
type arrPenjualan[NMAX]int
func main(){
	var A arrPenjualan
	var bilBul, i, size, idMax, j int
	i = 0
	size = 1
	fmt.Scan(&bilBul)
	for bilBul != - 1{
		A[i] = bilBul
		i+= 1
		size+= 1
		fmt.Scan(&bilBul)
	}
	idMax = nilaiMax(A, size)
	j = puncak(A, size)
	fmt.Print(j, " ")
	fmt.Print(idMax)
}
func nilaiMax(A arrPenjualan, size int)int{
	var max, i int
	max = A[0]
	for i = 0; i < size; i++{
		if A[i] > max{
			max = A[i]
		}
	}
	return max
}
func puncak(A arrPenjualan, size int)int{
	var i, j, max int
	max = A[0]
	for i = 0; i < size; i++{
		if A[i] > max{
			j = i+1
		}
	}
	return j
}