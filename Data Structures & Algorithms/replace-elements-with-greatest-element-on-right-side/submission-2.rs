impl Solution {
    pub fn replace_elements(arr: Vec<i32>) -> Vec<i32> {
        let n = arr.len();                    // Get length
        let mut ans = vec![0; n];             // Create output vector filled with 0s
        let mut right_max = -1;               // Track max to the right, starts at -1
        
        // Loop from last index down to 0
        // (0..n).rev() gives us: n-1, n-2, ..., 1, 0
        for i in (0..n).rev() {
            ans[i] = right_max;               // Put current right_max into answer
            right_max = right_max.max(arr[i]); // Update right_max with current element
        }
        
        ans                                     // Return answer (no "return" needed)
    }
}