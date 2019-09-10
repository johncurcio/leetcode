// o(n) time and o(1) space
func isPalindrome(x int) bool {
    if x < 0 {
        return false
    } else {
        digits := int(math.Log10(float64(x)))
        y := x
        stop := digits/2
        for digits > stop {
            exp := int(math.Pow(10, float64(digits)))
            left := y / exp
            y = y % exp
            right := x % 10
            x = x / 10
            digits--
            if left != right {
                return false
            }
        }
    }
    return true
}