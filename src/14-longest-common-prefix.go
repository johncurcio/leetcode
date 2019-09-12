
// time: O(n*m), where n = len(strs) and m = sum of the lens of each string 
func longestCommonPrefix(strs []string) string {
    if len(strs) < 1 {
        return ""
    }
    
    prefix := strs[0]
    for _, str := range strs {
        length := 0
        for length < len(str) && length < len(prefix) && prefix[length] == str[length] {
            length++
        }
        prefix = prefix[0:length]
    }
    
    return prefix
}