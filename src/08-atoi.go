func ternary(cond bool, t,f int) int {
    if cond {
        return t
    }
    return f
}

// o(n) time and o(1) space
func myAtoi(str string) int {
    str = strings.TrimSpace(str) // remove trailing whitespaces
    res := 0
    
    if len(str) > 0 {
        sign := str[0] // get the sign if it exists
        for index := ternary(sign == '+' || sign == '-', 1, 0) ; 
            index < len(str) && str[index]-'0' >= 0 && str[index]-'0' <= 9 && res < math.MaxInt32+1;
            index++ {
                digit := int(str[index]-'0')
                if !(res == 0 && str[index]-'0' == 0) { // removes trailing zeroes
                   res = 10 * res + digit
                }
        }
        
        if sign == '-' {
            res = -res
        }
    }
    
    if res > math.MaxInt32 {
        return math.MaxInt32
    } else if res < math.MinInt32 {
        return math.MinInt32
    }
    
    return res
}