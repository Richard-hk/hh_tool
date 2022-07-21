package arrays

func twoSum(nums []int, target int) []int {
	numMap := map[int]int{}
	for index, val := range nums {
		mapIndex, ok := numMap[target-val]
		if ok {
			return []int{mapIndex, index}
		}
		numMap[val] = index
	}
	return []int{}
}
