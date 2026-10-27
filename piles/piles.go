package piles

var A []int
var B []int

func GetA() []int {
	return append([]int(nil), A...)
}

func GetB() []int {
	return append([]int(nil), B...)
}
