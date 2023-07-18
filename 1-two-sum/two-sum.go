func twoSum(nums []int, target int) []int {
    var results []int
    for i, num := range nums {
        for j, num2 := range nums {
            if i == j {
                continue
            }
            if num+num2 == target {
                results = append(results, i)
                results = append(results, j)
                return results
            }
        }
    }

    return results
}