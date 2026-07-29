func twoSum(nums []int, target int) []int {
    cache := make(map[int]int)

    for i, item := range nums {
        remainder := target - item

        if j, ok := cache[remainder]; ok {
            return []int{j, i}
        }

        cache[item] = i
    }

    return nil
}
