/*
Big O
- N: the length of the given []int `nums`
- Time complexity: O(N)
  - Since we iterate over `nums`, the time complexity might increase linearly
- Space complexity: O(N)
  - `hashmap` could store up to N elements in worst case
*/

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	res := make([]int, 0, 2)
	for i, num := range nums {
		if j, ok := hashmap[target-num]; ok {
			res = append(res, i, j)
			break
		} else {
			hashmap[num] = i
		}
	}
	return res
}