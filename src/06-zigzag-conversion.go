func min (a, b int) int {
    if a < b {
        return a
    }
    return b
}

func ternary(cond bool, t int, f int) int {
    if cond {
        return t   
    }
    return f
}

func convert(s string, numRows int) string {
    if numRows <= 1 {
        return s
    }
    
    rows := min(len(s), numRows)
    list := make([]string, rows)
    for i := 0; i < rows; i++ {
        list[i] = ""
    }
    
    currRow := 0
    isGoingDown := false
    for _, c := range s {
        list[currRow] += string(c)
        if currRow == 0 || currRow == numRows - 1 {
            isGoingDown = !isGoingDown
        }
        currRow += ternary(isGoingDown, 1, -1)
    }
    
    res := ""
    for i := 0; i < rows; i++ {
        res += list[i]
    } 
    return res
}