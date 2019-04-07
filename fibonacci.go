/**
 * @Author: Joakim Olsson <joakss>
 * @Date:   2019-04-07T13:42:43+02:00
 * @Last modified by:   joakss
 * @Last modified time: 2019-04-07T13:46:34+02:00
 */

// Problem statement: https://tour.golang.org/moretypes/26

package main

import "fmt"

func fibonacci() func() int {
	x := 0
	y := 1
	return func() int {
		y, x = x, y+x
		return y
	}
}

// Prints the first 10 numbers of the fibonacci sequence
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
