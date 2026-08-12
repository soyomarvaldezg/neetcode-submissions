impl Solution {
    pub fn is_anagram(s: String, t: String) -> bool {
        // Early exit if lengths differ
        if s.len() != t.len() {
            return false;
        }
    
        // Create a counting array for 26 letters
        let mut store = [0i32; 26];
        
        // Get bytes for efficient access
        let s_bytes = s.as_bytes();
        let t_bytes = t.as_bytes();
        
        // Single pass through both strings
        for i in 0..s.len() {
            // Add 1 for character from s
            store[(s_bytes[i] - b'a') as usize] += 1;
            // Subtract 1 for character from t
            store[(t_bytes[i] - b'a') as usize] -= 1;
        }
        
        // If all counts are 0, they're anagrams
        store.iter().all(|&n| n == 0)

 }
}
