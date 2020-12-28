package main

import "fmt"

type figure struct {
	height, width, area int
	name                string
}

func main() {
	var area = func(f figure) figure {
		f.area = f.height * f.width
		return f
	}

	var inc = func(f figure) figure {
		f.height += 15
		f.width += 15
		return f
	}

	fmt.Println("With increment", side(inc))
	fmt.Println("With area and increment:", area(side(inc)))

}

func side(fn func(figure) figure) figure {
	var height, width int
	var name string
	fmt.Println("Enter name")
	fmt.Scan(&name)
	fmt.Println("Enter height")
	fmt.Scan(&height)
	fmt.Println("Enter width")
	fmt.Scan(&width)

	return fn(figure{name: name, height: height, width: width})
}
