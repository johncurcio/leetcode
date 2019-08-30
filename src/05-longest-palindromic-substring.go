// time: O(n^2), space: O(1)
func longestPalindrome(s string) string {
    longest := ""
    start, end := 0, 0
    for center := 0; center < len(s); center++ {
        len1 := longestPalindromeWithCenter(s, center, center)
        len2 := longestPalindromeWithCenter(s, center, center+1)
        l := 0
        if len1 > len2 {
            l = len1
        } else {
            l = len2
        }
        if l > end-start {
            start = center - (l-1)/2
            end = start + l - 1
            longest = s[start:end+1]
        }
    }
    return longest
}

func longestPalindromeWithCenter(s string, i int, j int) int {
    beg, en := i, j
    for beg >= 0 && en < len(s) && s[beg] == s[en] {
        beg--
        en++
    }
    return en - beg - 1
}