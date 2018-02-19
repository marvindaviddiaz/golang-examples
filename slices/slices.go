package slices

import "fmt"

func Init() {
	a := [] int {1, 2, 3, 4, 5, 6}
	b := a[2:4]
	c := a[:3]
	d := a[3:]

	fmt.Println("Slices \na:", a, "\nb:", b, "\nc:", c, "\nd:", d)

	fmt.Println("Capacity of b:", cap(b))
	fmt.Println("What b actually sees:", b[:cap(b)])

	// change the value in slices [B,C,D] affect the original slice [A]
	b[0] = 10
	fmt.Println(a)

	// the previous behavior changes if you uses 'copy' function
	bCopy := make([]int, 2)
	copy(bCopy, a[2:4])
	bCopy[0] = 99 // the value in original array doesn't change
	fmt.Println(a)
}