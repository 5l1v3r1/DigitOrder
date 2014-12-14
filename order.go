package main

import (
	"fmt"
	"math/big"
	"sync"
)

var out chan []int
var wg sync.WaitGroup

func CheckCombination(current []int, base int) bool {
	// Figure out the full number
	number := big.NewInt(int64(current[0]))
	bigBase := big.NewInt(int64(base))
	temp := big.NewInt(0)
	zero := big.NewInt(0)
	for i := 1; i < len(current); i++ {
		if i != 0 {
			temp.Mul(number, bigBase)
			number.Add(temp, big.NewInt(int64(current[i])))
		}
	}
	divisor := big.NewInt(int64(len(current)))
	temp.Mod(number, divisor)
	return temp.Cmp(zero) == 0
}

func BestNumber(current []int, remaining []int, base int) {
	if len(current) == 1 {
		defer wg.Done()
	}
	if !CheckCombination(current, base) {
		return
	}
	if len(remaining) == 0 {
		out <- current
		return
	}
	for i, x := range remaining {
		newCur := append(DupList(current), x)
		newRem := DupList(remaining)
		newRem[i] = newRem[len(newRem)-1]
		newRem = newRem[0 : len(newRem)-1]
		BestNumber(newCur, newRem, base)
	}
}

func DupList(list []int) []int {
	res := make([]int, len(list))
	copy(res, list)
	return res
}

func main() {
	for i := 2; i < 30; i++ {
		fmt.Println("running base", i)
		out = make(chan []int)
		// Run a goroutine for every starting digit
		for j := 1; j < i; j++ {
			remaining := make([]int, 0, i-1)
			current := make([]int, 1)
			current[0] = j
			for x := 1; x < i; x++ {
				if x != j {
					remaining = append(remaining, x)
				}
			}
			wg.Add(1)
			go BestNumber(current, remaining, i)
		}
		go func() {
			wg.Wait()
			close(out)
		}()
		for {
			if val, ok := <-out; ok {
				fmt.Println("result", append(DupList(val), 0))
			} else {
				break
			}
		}
	}
}
