// UVa 10154 - Weights and Measures

package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

type (
	turtle  struct{ weight, strength int }
	turtles []turtle
)

func (t turtles) Len() int { return len(t) }

func (t turtles) Less(i, j int) bool { return t[i].strength < t[j].strength }

func (t turtles) Swap(i, j int) { t[i], t[j] = t[j], t[i] }

func lis(ts turtles) int {
	sort.Sort(ts)
	dp := make([]int, len(ts)+1)
	for i := 1; i <= len(ts); i++ {
		dp[i] = math.MaxInt32
	}
	maxLen := 0
	for i := 1; i <= len(ts); i++ {
		for j := len(ts); j >= 1; j-- {
			if ts[i-1].strength-ts[i-1].weight >= dp[j-1] {
				if dp[j]-dp[j-1] >= ts[i-1].weight {
					dp[j] = dp[j-1] + ts[i-1].weight
					if j > maxLen {
						maxLen = j
					}
				}
			}
		}
	}
	return maxLen
}

func main() {
	in, _ := os.Open("10154.in")
	defer in.Close()
	out, _ := os.Create("10154.out")
	defer out.Close()

	var ts turtles
	var w, s int
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &w, &s); err != nil {
			break
		}
		ts = append(ts, turtle{w, s})
	}
	fmt.Fprintln(out, lis(ts))
}
