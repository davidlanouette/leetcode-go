package leetcode

/* https://leetcode.com/problems/jump-game/description/


Given an array of non-negative integers, you are initially positioned at the first index of the array.
Each element in the array represents your maximum jump length at that position.
Determine if you are able to reach the last index.

Example 1:
```
Input: [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
```

Example 2:
```
Input: [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum
             jump length is 0, which makes it impossible to reach the last index.
```
*/

func canJump(nums []int) bool {
	curPos := 0
	lastPos := 0
	target := len(nums) - 1

	for curPos < target {
		lastPos = curPos
		curPos += nums[curPos]

		// if we get to a spot where we are not moving anymore, we are done.
		if lastPos == curPos {
			return false
		}
	}
	return true
}
