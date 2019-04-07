/**
 * @Author: Joakim Olsson <joakss>
 * @Date:   2019-04-07T13:36:17+02:00
 * @Last modified by:   joakss
 * @Last modified time: 2019-04-07T13:40:55+02:00
 */

// Problem statement: https://tour.golang.org/moretypes/23

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)

	for i := range words {
		m[words[i]]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
