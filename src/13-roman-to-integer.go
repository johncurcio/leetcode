// time: O(n), space: O(1)
func romanToInt(s string) int {
    var res int = 0
    var prev int = math.MinInt32
    var romans = map[byte]int {
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }
    for i := len(s)-1; i >= 0; i-- {
        val := romans[ s[i] ]
        if prev > val {
            res -= val
        } else {
            res += val
        }
        prev = val
    }
    return res
}
