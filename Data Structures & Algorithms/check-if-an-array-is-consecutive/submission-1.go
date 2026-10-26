func isConsecutive(nums []int) bool {
    low := nums[0]
    high := nums[0]
    seen := make(map[int]bool)

    for _, x := range nums {
        if seen[x] {
            return false
        }

        seen[x] = true

        if x < low {
            low = x
        }

        if x > high {
            high = x
        }
    }

    return high-low+1 == len(nums)
}