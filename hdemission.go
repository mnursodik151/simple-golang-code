package main

import "fmt"


func YnInt(X int, res int) int {
	var val int
	if (X >= 0) {
		fmt.Scanf("%d", &val)
		if (val > 0) {
			res += val*val
		}
		res = YnInt(X-1,res)
	}
	return res
}

func NCase(N int, result[]int) []int {
	var X int
	if (X == 0) {
		fmt.Scanf("%d", &X)
	}
	if (N > 0) {
		N -= 1
		fmt.Scanf("%d", &X)
		result[N] = YnInt(X, 0)
		NCase(N,result)
	}
	return result
}

func printRes(result []int, length int) {
	if (length > 0) {
		length -= 1
		fmt.Println(result[length])
		printRes(result, length)
	} 
}

func main() {
	var N int
	fmt.Println("Input Begins")
	fmt.Scanf("%d", &N)
	result := make([]int, N, N)
	result = NCase(N, result)
	fmt.Println("Output Begins")
	printRes(result, len(result))
}
