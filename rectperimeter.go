<<<<<<< HEAD
package main

import "fmt"

func RectPerimeter(w, h int) int {
	if w < 0 || h < 0 {
		return -1
	}
	perimeter := (w + h) * 2
	return perimeter
}

func main() {
	fmt.Println(RectPerimeter(10, 2))
	fmt.Println(RectPerimeter(434343, 898989))
	fmt.Println(RectPerimeter(10, -2))
}
=======
package main

import "fmt"

func RectPerimeter(w, h int) int {
	if w < 0 || h < 0 {
		return -1
	}
	perimeter := (w + h) * 2
	return perimeter
}

func main() {
	fmt.Println(RectPerimeter(10, 2))
	fmt.Println(RectPerimeter(434343, 898989))
	fmt.Println(RectPerimeter(10, -2))
}
>>>>>>> 2262bbcac982957294399b14cf2d87d0af9024cf
