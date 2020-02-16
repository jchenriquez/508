package twosum

// TwoSum will return two numbers that sum to a target
func TwoSum(nums []int, target int) []int {
	mp := make(map[int]map[int]bool, len(nums))

	for index, num := range nums {
		mi, in := mp[num]

		if !in {
			mi = make(map[int]bool)
		}
		mi[index] = true

		mp[num] = mi
	}

	for index, num := range nums {
		targetN := target - num
		mi, _ := mp[num]
		delete(mi, index)

		mix, in := mp[targetN]

		if in && len(mix) > 0 {
			for ind := range mix {
				return []int{index, ind}
			}
		}

		mi[index] = true
	}

	return []int{}
}
