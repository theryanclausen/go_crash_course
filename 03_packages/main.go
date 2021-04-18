package main

import (
	"fmt"
	"math"

	"github.com/theryanclausen/go_crash_course/03_packages/strutil"
)

func main(){
	num := 2.5
	fmt.Println(math.Floor(num))
	fmt.Println(math.Sqrt(num))
	fmt.Println(strutil.Reverse("nayR"))
}