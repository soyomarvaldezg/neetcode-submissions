impl Solution {
    pub fn is_subsequence(s: String, t: String) -> bool {
        let s_bytes = s.as_bytes();  // Convert s to byte array for efficient indexing
        let t_bytes = t.as_bytes();  // Convert t to byte array
        let mut i = 0;               // Pointer into s
        let mut j = 0;               // Pointer into t

        while i < s_bytes.len() && j < t_bytes.len() {
            if s_bytes[i] == t_bytes[j] {
                i += 1;              // Matched, advance in s
            }
            j += 1;                  // Always advance in t
        }

        i == s_bytes.len()           // Did we match all of s?
    }
}
