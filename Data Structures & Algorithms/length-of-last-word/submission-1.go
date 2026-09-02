func lengthOfLastWord(s string) int {
    i := len(s) - 1        // 1. start from last index
    for s[i] == ' ' {      // 2. skip trailing spaces
        i--                // 3. move left (same as i -= 1)
    }
    length := 0            
    for i >= 0 && s[i] != ' ' {  // 4. count the word
        length++            // 5. increment length
        i--                 // 6. move left
    }
    return length
}
