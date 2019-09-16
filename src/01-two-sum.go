// time O(n), space O(n)
func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
    for i, num := range nums {
        index, ok := m[target-num]
        if ok {
            return []int{index, i}
        } else {
            m[num] = i
        }
    }
    return []int{}
}