package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s) // len=10 cap=10 [10 20 30 40 50 60 70 80 90 100]

	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])          // len=0 cap=10 []
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:5]), cap(s[2:5]), s[2:5])       // len=3 cap=8 [30 40 50]
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])          // len=8 cap=8 [30 40 50 60 70 80 90 100]
	fmt.Printf("len=%d cap=%d %v\n", len(s[:5]), cap(s[:5]), s[:5])          // len=5 cap=10 [10 20 30 40 50]
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:5:5]), cap(s[2:5:5]), s[2:5:5]) // len=3 cap=3 [30 40 50]
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:5:6]), cap(s[2:5:6]), s[2:5:6]) // len=3 cap=4 [30 40 50]
	s = append(s, 110)

	sa := []int{20, 10, 21, 39, 50, 60, 70, 80, 90, 100}

	s1 := sa[2:7:10]
	println(len(s1)) // 5
	println(cap(s1)) // 8

	s2 := sa[:7:8]
	println(len(s2)) // 7
	println(cap(s2)) // 8
	s3 := sa[2:3:8]
	println(len(s3)) // 1
	println(cap(s3)) // 6
}
