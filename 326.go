package main

func main() {
}

func isPowerOfThree(n int) bool {
	if n > 0 && 1162261467%n == 0 {
		return true
	}
	return false
}
