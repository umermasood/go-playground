package main

import "fmt"

func main() {
	// * Slice 1 - Creating an empty slice with zero length
	slc1 := make([]string, 0)
	fmt.Println("Emp :", slc1)
	fmt.Println("slc1 length :", len(slc1))

	// Adding values in zero length slc1
	slc1 = append(slc1, "abacus")
	slc1 = append(slc1, "charlie")
	slc1 = append(slc1, "zeta")
	fmt.Println("Emp :", slc1)
	fmt.Println("slc1 length :", len(slc1))

	fmt.Println("-------------------------")

	// * Slice 2 - Creating another empty slice with non-zero length, i.e. 3
	slc2 := make([]string, 3) // Empty slice of length 3
	fmt.Println("Emp :", slc2)
	fmt.Println("slc2 length :", len(slc2))

	// Adding values in non-zero length slc2
	slc2 = append(slc2, "uno")
	slc2 = append(slc2, "dos")
	slc2 = append(slc2, "tres")
	fmt.Println("Emp :", slc2)
	fmt.Println("slc2 length :", len(slc2)) // length becomes six

	fmt.Println("-------------------------")

	// * Slice 3 - One liner slice declaration and initialization
	num_slice := []int{1, 2, 3, 4, 5}
	fmt.Println("num_slice :", num_slice)

	fmt.Println("-------------------------")

	// * Slice 4 - Using the slice operator
	slice_optr := num_slice[1:3]
	fmt.Println("new slice using operator :", slice_optr)

	fmt.Println("-------------------------")

	// * Slice 5 - Multi Dimensional Slices
	slice2D := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		slice2D[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			slice2D[i][j] = i + j
		}
	}
	fmt.Println("2d: ", slice2D)
}
