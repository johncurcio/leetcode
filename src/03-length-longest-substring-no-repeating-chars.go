// O(n) space complexity, because of map. Something close to O(n^2) time complexity, because of delete
func lengthOfLongestSubstring(s string) int {
    start, end := 0, 0
    maxLen := 0
    set := make(map[string]bool)
    for start < len(s) && end < len(s) {
        str := string(s[end])
        if set[str] {
            delete(set, string(s[start]))
            start = start + 1
        } else {
            set[str] = true
            end = end + 1
        }
        if len(set) > maxLen {
            maxLen = len(set)   
        }
    }
    return maxLen
}
