package main

import "fmt"

func main() {
	// slices are made using the make function
	// it accepts three arguments; type, length, and capacity
	// capacity v length:
	// capacity is effectively the length of the underlying array backing the slice
	// slices don't actually store data, it's stored in the array.
	// courses := make([]string, 5, 10)
	// courses[0] = "History"
	// courses[1] = "Math"
	// courses[2] = "English"
	// fmt.Printf("Length of slice is %d and capacity is %d\n", len(courses), cap(courses))
	// fmt.Println(courses)
	courses := []string{"History", "Math", "English", "Geography", "Science"}
	fmt.Printf("Length of slice is %d and capacity is %d\n", len(courses), cap(courses))
	for _, i := range courses {
		fmt.Println(i)
	}
	sliceOfSlice := courses[0:2]
	fmt.Println(sliceOfSlice)
	// when beginning index is omitted, 0 is implied.
	fmt.Println(courses[:2])
	// when end index is omitted, end of slice is implied.
	fmt.Println(courses[2:])

	// when length of slice is equal to capacity of slice, any attempt to append to slice results in expanding of underlying array.
	// specifically, go simply doubles the size of the underlying array.
	initSlice := make([]int, 1, 3)
	fmt.Printf("Length of slice is %d and capacity is %d\n", len(initSlice), cap(initSlice))
	for i := 1; i < 8; i++ {
		initSlice = append(initSlice, i)
		// Length of slice is 1 and capacity is 3
		// Length of slice is 2 and capacity is 3
		// Length of slice is 3 and capacity is 3
		// Length of slice is 4 and capacity is 6
		// Length of slice is 5 and capacity is 6
		// Length of slice is 6 and capacity is 6
		// Length of slice is 7 and capacity is 12
		// Length of slice is 8 and capacity is 12
		fmt.Printf("Length of slice is %d and capacity is %d\n", len(initSlice), cap(initSlice))
	}

	// slices can be appended to other slices as so;
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	fmt.Println(append(slice1, slice2...))
}
