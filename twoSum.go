package leetcode

func twoSum(nums []int, target int) []int {
	// fmt.Println("nums =", nums, ", target =", target)
	numMap := make(map[int]int)
	for n, v := range nums {
		numMap[v] = n
	}
	// fmt.Println("numMap =", numMap)
	for i := 0; i < len(nums); i++ {
		j, ok := numMap[target-nums[i]]
		// fmt.Println("i =", i, ", nums[i] =", nums[i], ", target-nums[i] =", target-nums[i], ", j =", j, ", ok =", ok)
		if ok && i != j {
			return []int{i, j}
		}
	}
	// couldn't find a match.  return empty slice
	return []int{}
}
