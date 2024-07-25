package main

import "fmt"

func main() {
	//nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	//fmt.Println(nums)
	//nums = nums[:5]
	//fmt.Println(nums)
	removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
}

func removeElement(nums []int, val int) int {
	var k int
	fmt.Println("nums before: ", nums)

	var i int
	for i < len(nums) {
		if nums[i] == val {
			//indexes = append(indexes, i)
			nums = append(nums[:i], nums[i+1:]...)
			i--
			k++
		}
		i++
	}

	fmt.Println("nums after: ", nums)
	fmt.Println("k: ", k)
	return k
}
