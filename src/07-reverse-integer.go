func overflows(val int) bool {
   return val > math.MaxInt32 || val < math.MinInt32
}

func sign(res, sign int) int {
    if sign < 0 {
        return -res
    }
    return res
}

func reverse(x int) int {
    y := int(math.Abs(float64(x)))
    digits := int(math.Log10(float64(y)))
    multiplier := int(math.Pow(10, float64(digits)))
    res := 0
    for y > 0 {
        lastDigit := y % 10
        y = y / 10
        res += multiplier * lastDigit
        multiplier = multiplier / 10
    }
    
    if overflows(x) || overflows(res) {
        return 0
    }
    
    return sign(res, x)
}