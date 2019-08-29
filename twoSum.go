package leetcode

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

	Given nums = [2, 7, 11, 15], target = 9,

	Because nums[0] + nums[1] = 2 + 7 = 9,
	return [0, 1].

*/
func twoSum(nums []int, target int) []int {
	// fmt.Println("nums =", nums, ", target =", target)
	numMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		// fmt.Println("numMap =", numMap)
		j, ok := numMap[target-nums[i]]
		// fmt.Println("i =", i, ", nums[i] =", nums[i], ", target-nums[i] =", target-nums[i], ", j =", j, ", ok =", ok)
		if ok && i != j {
			return []int{j, i}
		}
		numMap[nums[i]] = i
	}
	// couldn't find a match.  return empty slice
	return []int{}
}
