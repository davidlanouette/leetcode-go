package leetcode

/*

https://leetcode.com/problems/search-in-rotated-sorted-array-ii

Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,0,1,2,2,5,6] might become [2,5,6,0,0,1,2]).

You are given a target value to search. If found in the array return true, otherwise return false.

Example 1:

Input: nums = [2,5,6,0,0,1,2], target = 0
Output: true

Example 2:

Input: nums = [2,5,6,0,0,1,2], target = 3
Output: false

*/

/*
1) everytime check if target == nums[mid], if so, we find it.
2) otherwise, we check if the first half is in order (i.e. nums[left]<=nums[mid])
  and if so, go to step 3), otherwise, the second half is in order,   go to step 4)
3) check if target in the range of [left, mid-1] (i.e. nums[left]<=target < nums[mid]), if so, do search in the first half, i.e. right = mid-1; otherwise, search in the second half left = mid+1;
4)  check if target in the range of [mid+1, right] (i.e. nums[mid]<target <= nums[right]), if so, do search in the second half, i.e. left = mid+1; otherwise search in the first half right = mid-1;
*/
func search2(nums []int, target int) bool {

	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return (nums[0] == target)
	}

	mid := len(nums) / 2
	if nums[mid] == target {
		return true
	}

	if search2(nums[0:mid], target) {
		return true
	}
	if search2(nums[mid+1:], target) {
		return true
	}
	return false
}

// search is a brute force attempt that works 100% of the time.
// But, it's O(n) compute complexity, and O(1) space complexity.  i.e., long lists will take a long time to solve.
func search(nums []int, target int) bool {
	for _, i := range nums {
		if i == target {
			return true
		}
	}
	return false
}
