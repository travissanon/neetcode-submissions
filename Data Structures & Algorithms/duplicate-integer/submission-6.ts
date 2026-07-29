class Solution {
    /**
     * @param {number[]} nums
     * @return {boolean}
     */
    hasDuplicate(nums: number[]): boolean {
		const store = {};

		for (let i  = 0; i < nums.length; i++) {
			if (store[nums[i]] === 0) {
				return true
			}
			store[nums[i]] = 0
		}

		return false
	}
}
