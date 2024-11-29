func minimumSubarrayLength(nums []int, k int) int {
	result := 220000
	n := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] >= k {
			return 1
		}
		n = n | nums[i]
		if n >= k {
			t := nums[i]
			j := i
			for t < k {
				j--
				t = t | nums[j]
			}
			if i-j+1 < result {
				result = i - j + 1
			}
			j++
			n = nums[j]
			i = j
		}
	}
	if result < 220000 {
		return result
	}
	return -1
}