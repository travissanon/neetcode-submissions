func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    set := make(map[rune]int)

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
