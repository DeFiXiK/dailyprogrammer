package main

import (
	"fmt"
	"math/big"
	"sort"
	"strconv"
)

var sign = []rune{'+', '-', ' '}

func calc(operations []rune) int {
	var nums []int
	var opers []rune
	nums = append(nums, 1)
	for i := range operations {
		if operations[i] == '+' {
			nums = append(nums, i+2)
			opers = append(opers, '+')
		}
		if operations[i] == '-' {
			nums = append(nums, i+2)
			opers = append(opers, '-')
		}
		if operations[i] == ' ' {
			nums[len(nums)-1] = (nums[len(nums)-1] * 10) + (i + 2)
		}
	}
	var result int
	result = nums[0]
	for i := 0; i < len(opers); i++ {
		if opers[i] == '+' {
			result += nums[i+1]
		}
		if opers[i] == '-' {
			result -= nums[i+1]
		}
	}
	return result
}

func next(numbers int) [][]rune {
	var out [][]rune
	if numbers == 2 {
		for s := range sign {
			out = append(out, []rune{sign[s]})
		}
		return out
	}
	input := next(numbers - 1)
	for i := range sign {
		for j := range input {
			var temp []rune
			temp = append(temp, input[j]...)
			temp = append(temp, sign[i])
			out = append(out, temp)
		}
	}
	return out
}

func forCucle(f []int) int {
	var acum int
	for _, n := range f {
		acum += n
	}
	return acum
}

func whileCucle(f []int) int {
	var acum int
	var i int
	for {
		acum += f[i]
		i++
		if len(f) <= i {
			break
		}
	}
	return acum
}

func recSum(f []int, acum int) int {
	if len(f) == 0 {
		return acum
	}
	acum += f[0]
	return recSum(f[1:], acum)
}

func oneWithTwo(one []string, two []int) []string {
	var result []string

	for i := range one {
		result = append(result, one[i])
		result = append(result, strconv.Itoa(two[i]))
	}
	return result
}

func fibSum() []big.Int {
	var fib []big.Int
	var temp *big.Int
	first := big.NewInt(0)
	second := big.NewInt(1)
	for i := 0; i < 100; i++ {
		fib = append(fib, *first)
		temp = add(first, second)
		first = second
		second = temp
	}
	return fib
}

func add(x, y *big.Int) *big.Int {
	return big.NewInt(0).Add(x, y)
}

func main() {
	// Problem 1
	l := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(forCucle(l))
	fmt.Println(whileCucle(l))
	fmt.Println(recSum(l, 0))
	fmt.Println("Solved")
	// Problem 2
	one := []string{"a", "b", "c"}
	two := []int{1, 2, 3}
	fmt.Println(oneWithTwo(one, two))
	fmt.Println("Solved")
	// Problem 3
	elems := fibSum()
	for i := range elems {
		fmt.Println(fmt.Sprintf("%v %v", i+1, elems[i].String()))
	}
	fmt.Println("Solved")
	// Problem 4
	elements := []int{50, 2, 1, 9}
	sort.Slice(elements, func(i, j int) bool {
		first := strconv.Itoa(elements[i])
		second := strconv.Itoa(elements[j])
		if first+second >= second+first {
			return true
		} else {
			return false
		}
	})
	var output string
	for _, elem := range elements {
		output += strconv.Itoa(elem)
	}
	fmt.Println(output)
	fmt.Println("Solved")
	// Problem 5
	out := next(9)
	for i := range out {
		if calc(out[i]) == 100 {
			for j := range out[i] {
				fmt.Printf("%v", j+1)
				if out[i][j] != ' ' {
					fmt.Printf(" %c ", out[i][j])
				}
			}
			fmt.Printf("9")
			fmt.Println(" =", calc(out[i]))
		}
	}
	fmt.Println("Solved")
}
