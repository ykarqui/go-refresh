package main

import (
	"fmt"
	"go-refresh/src/github.com/ykarqui/go_crash_course/03_pckgs/strutil"
	"math"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(2.7))
	fmt.Println(strutil.Reverse("topaet"))
}
