// recursive, complexity is exponential
func isMatch(s string, p string) bool {
    if len(p) == 0 {
        return len(s) == 0
    }
    fmatch := len(s) > 0 && (s[0] == p[0] || p[0] == '.')
    if len(p) >= 2 && p[1] == '*' {
        return isMatch(s, p[2:len(p)]) || (fmatch && isMatch(s[1:len(s)], p))
    } else {
        return fmatch && isMatch(s[1:len(s)], p[1:len(p)])        
    }
}

