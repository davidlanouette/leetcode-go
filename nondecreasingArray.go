package leetcode

func checkPossibility(nums []int) bool {
	// fmt.Printf("--------\n%v\n", nums)

	for i, n := range nums[:len(nums)-1] {
		next := nums[i+1]
		// fmt.Printf("n=%v, next=%v\n", n, next)
		if n > next {

			origN := n

			// set n equal to next, and check if that "fixes" the list
			nums[i] = next
			if isSorted(nums) {
				return true
			}

			// if that didn't do it, use the original N for both n and next
			nums[i], nums[i+1] = origN, origN
			if isSorted(nums) {
				return true
			}

			// if that didn't do it, it will take more than one swap, so return false
			return false
		}
	}

	return true
}

func isSorted(list []int) bool {
	last := list[0]
	for _, i := range list[1:] {
		if last > i {
			return false
		}
		last = i
	}
	return true
}
