/**
 * @Author: Joakim Olsson <joakss>
 * @Date:   2019-04-07T13:28:02+02:00
 * @Last modified by:   joakss
 * @Last modified time: 2019-04-07T13:47:02+02:00
 */

//Problem statement: https://tour.golang.org/flowcontrol/8

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {
	z := 1.0
	counter := 0
	for {
		p := z - (z*z-x)/(2*z)
		if math.Abs(z-p) < 0.000001 {
			break
		}
		counter++
		z = p
	}
	return z, counter
}

// Example of squaring the number 251
func main() {
	fmt.Print("Approx root with number of iterations: ")
	fmt.Println(Sqrt(251))
	fmt.Println("Exact root:", math.Sqrt(251))
}
