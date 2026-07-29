import (
	"slices"
)

func hasDuplicate(nums []int) bool {
	var bank []int

	for count := 0; count < len(nums); count++ {
		if slices.Contains(bank, nums[count]) {
			return true
		} else {
			bank = append(bank, nums[count])
		}
	}

	return false
}
