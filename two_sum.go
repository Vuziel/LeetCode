package main

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := mp[target-nums[i]]; ok {
			return []int{mp[target-nums[i]], i}
		}

		mp[nums[i]] = i
	}

	return nil
}
