/**
 * @Author: Joakim Olsson <joakss>
 * @Date:   2019-04-07T13:39:56+02:00
 * @Last modified by:   joakss
 * @Last modified time: 2019-04-07T13:47:29+02:00
 */

// Problem statement: https://tour.golang.org/moretypes/18

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
 	picture := make([][]uint8, dy)
 	for i := range picture {
 		picture[i] = make([]uint8, dx)
 		for y := range picture {
 			picture[i][y] = uint8(i^y)
 		}
 	}
 	return picture
 }

// Shows the picture
func main() {
 	pic.Show(Pic)
 }
