package arrays

// https://leetcode.cn/problems/container-with-most-water/
// func maxArea(height []int) int {
// 	len := len(height)
// 	var maxArea int
// 	for i := 0; i < len; i++ {
// 		for j := i + 1; j < len; j++ {
// 			tmpArea := (j - i) * compareNum(height[i], height[j])
// 			if tmpArea > maxArea {
// 				maxArea = tmpArea
// 			}

// 		}
// 	}
// 	return maxArea
// }

// func compareNum(x, y int) int {
// 	if x > y {
// 		return y * y
// 	}
// 	return x * x
// }

func maxArea(height []int) int {
	var indexMap map[int]int
	var maxRes int
	for index, val := range height {
		indexMap[index] = val
	}

	for i := 0; i < len(height); i++ {

	}
	return maxRes
}

func compareNum(x, y int) int {
	if x > y {
		return y
	}
	return x
}
