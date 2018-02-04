package twoSum

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, num := range nums {
		if idx, ok := mp[target-num]; ok {
			return []int{idx, i}
		}
		mp[num] = i
	}
	return nil
}
