func hasDuplicate(nums []int) bool {
    seen := make(map[int]bool)
    // To check if a value exists:
    for _, num := range nums {
        if seen[num] {
            return true
        }
        seen[num] = true
    }
    return false
}
