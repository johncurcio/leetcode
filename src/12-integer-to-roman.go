func romanize(num, pos int) string {
    one, five, ten := "","",""
    switch pos {
        case 0: one, five, ten = "I","V","X"
        case 1: one, five, ten = "X","L","C"
        case 2: one, five, ten = "C","D","M"
        case 3: one, five, ten = "M","",""
    }
    
    switch {
        case num > 0 && num < 4 : return strings.Repeat(one, num)
        case num > 5 && num < 9 : return five + strings.Repeat(one, num-5)
        case num == 4 : return one + five
        case num == 5 : return five
        case num == 9 : return one + ten
    }
    return ""
}

// O(n) time, O(1) space
func intToRoman(num int) string {
    roman := ""
    digits := int(math.Log10(float64(num)))
    for i := digits; i >= 0; i-- {
        roman = romanize(num % 10, digits-i) + roman
        num = num / 10
    }
    return roman
}