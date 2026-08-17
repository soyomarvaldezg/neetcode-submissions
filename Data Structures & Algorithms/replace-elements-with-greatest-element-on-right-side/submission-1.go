func replaceElements(arr []int) []int {
    // n = length of the array (how many elements)
    n := len(arr)
    
    // Create a new array called "ans" with the same length as arr.
    // It starts filled with 0s, but we'll fill in the real values.
    ans := make([]int, n)

    // currentMax tracks the biggest number we've seen to the RIGHT.
    // We start at -1 because the last element has nothing to its right.
    currentMax := -1

    // Loop from the END of the array (index n-1) to the START (index 0).
    // i-- means we go backwards: n-1, n-2, ..., 1, 0
    for i := n - 1; i >= 0; i-- {
        // Put the current "largest to the right" into this position.
        // Example: at the last position, this is -1.
        ans[i] = currentMax
        
        // Now update currentMax for the NEXT iteration (which looks further left).
        // If arr[i] is bigger than what we had, use arr[i] instead.
        if arr[i] > currentMax {
            currentMax = arr[i]
        }
    }
    
    // Give back the finished answer array
    return ans
}