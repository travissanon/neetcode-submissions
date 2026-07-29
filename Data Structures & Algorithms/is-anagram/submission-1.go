func isAnagram(s string, t string) bool {
    set := make(map[rune]int, len(s))

    if len(s) != len(t) {
        return false
    }
    
    for _, char := range s {
        set[char]++
    }

    for _, char := range t {
        if set[char] <= 0 {
            return false
        }
        set[char]--
    }

    return true
}
