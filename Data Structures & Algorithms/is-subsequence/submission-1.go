func isSubsequence(s string, t string) bool {
    i, j := 0, 0                    // i points into s, j points into t
    for i < len(s) && j < len(t) {  // Continue until we exhaust either string
        if s[i] == t[j] {           // If current characters match
            i++                     // Move to next char in s (consumed a match)
        }
        j++                         // Always advance pointer in t (characters we skip don't matter)
    }
    return i == len(s)              // We matched all of s? True = subsequence
}