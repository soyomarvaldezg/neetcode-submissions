func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    countS := make(map[byte]int)
    countT := make(map[byte]int)

    for i := 0; i < len(s); i++ {
            // s[i] and t[i] give the byte at position i
            countS[s[i]] = countS[s[i]] + 1
            countT[t[i]] = countT[t[i]] + 1
    }

    for ch, count := range countS {
        if countT[ch] != count {
            return false
        }
    }

    return true

}
