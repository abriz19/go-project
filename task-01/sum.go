// Fundamentals of Go Tasks
// Task: Sum of Numbers
// Write a Go function that takes a slice of integers as input and returns the sum of all the numbers. If the slice is empty, the function should return 0.
// [Optional]: Write a test for your function

package main

import "fmt"

func Sum(nums []int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

func main() {
fmt.Println(Sum([]int{12, 34}))
}