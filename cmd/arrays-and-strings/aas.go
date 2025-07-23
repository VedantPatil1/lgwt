package aas

func Sum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func SumAll(coll_nums ...[]int) []int {
	var results []int
	for _, coll := range coll_nums {
		results = append(results, Sum(coll))
	}
	return results 
}

func SumAllTail(coll_nums ...[]int) []int {
	var results []int
	for _, coll := range coll_nums {
		results = append(results, Sum(coll[1:]))
	}
	return results
}
