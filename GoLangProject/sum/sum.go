package sum

import "fmt"

// Sum calculates the sum of a slice of integers.
func Sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	// Example usage
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("Sum:", Sum(nums)) // Output: Sum: 15
}
