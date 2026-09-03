impl Solution {
    pub fn length_of_last_word(s: String) -> i32 {
        let bytes = s.as_bytes();
        let mut i = bytes.len();

        // Skip trailing spaces
        while i > 0 && bytes[i - 1] == b' ' {
            i -= 1;
        }

        let mut length = 0;

        // Count the last word
        while i > 0 && bytes[i - 1] != b' ' {
            length += 1;
            i -= 1;
        }

        length as i32
    }
}
