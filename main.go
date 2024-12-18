package main

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

func main() {
	fmt.Println("Hello Daksh!")
	fmt.Println(rand.Int())
	fmt.Printf("Sq.rt :%f\n", math.Sqrt(9))

	var i int = 29
	var j int
	k := 15
	j = 9
	fmt.Println(i, j, k)

	var (
		Dad   string = "Konda Reddy"
		Mom   string = "Sireesha"
		Son   string = "Daksh"
		Total int    = 3
	)
	fmt.Println(Dad, Mom, Son, Total)

	var c string
	var a int
	var b float32
	var d bool
	fmt.Printf("%d %f %s %t\n", a, b, c, d)

	nr := 12
	dr := 0
	var div, rem, err = division(nr, dr)
	if err != nil {
		fmt.Println(err.Error())

	} else if rem == 0 {
		fmt.Printf("reslt of div:%d \n", div)
	} else {
		fmt.Printf("reslt of div:%d and rem:%v", div, rem)
	}

	///////////----ARRAY------//////////
	var iArr [3]int = [3]int{1, 2, 3}
	inArr := [3]int{1, 2, 3}
	intArr := [...]int{1, 2, 3}
	fmt.Println(iArr, inArr, intArr)

	/////////-----SLICE------////////
	var iSlice []int = []int{4, 5, 6}
	iSlice = append(iSlice, 7)
	fmt.Println(iSlice)
	fmt.Printf("len %v and cap %v\n", len(iSlice), cap(iSlice))

	var jSlice []int = []int{8, 9}
	iSlice = append(iSlice, jSlice...)
	fmt.Println(iSlice)

	var kSlice []int32 = make([]int32, 3, 8) //(type,len,capacity)
	fmt.Printf("len %v and cap %v\n", len(kSlice), cap(kSlice))

	var n int = 1000000
	var testSlice1 = []int{}
	var testSlice2 = make([]int, 0, n)

	fmt.Printf("withot allocation: %v\n", Timeloop(testSlice1, n))
	fmt.Printf("with allocation: %v\n", Timeloop(testSlice2, n))

	/////////---MAPS-----/////////
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	var Boys = map[string]uint8{"Konda": 29, "Daksh": 15}
	fmt.Println(Boys["Konda"])
	fmt.Println(Boys["Siri"])
	var DOB, ok = Boys["Siri"]
	if ok {
		fmt.Printf("DOB:%v\n", DOB)
	} else {
		fmt.Println("Invalid Boy Name")
	}

	var fam = map[string]uint8{"Konda": 30, "Daksh": 2, "Siri": 29}
	for name, age := range fam {
		fmt.Printf("Name:%s Age:%v\n", name, age)

	}

	var myStr = "rèsumè"
	var index = myStr[1]
	fmt.Printf("%v,%T\n", index, index)
	fmt.Printf("len of str:%v\n", len(myStr))
	var myStrr = []rune("rèsumè")
	var indexed = myStrr[1]
	fmt.Printf("%v,%T\n", indexed, indexed)
	fmt.Printf("len of str:%v\n", len(myStrr))
	for i, v := range myStr {
		fmt.Println(i, v)

	}

	var strSlice = []string{"m", "y", " ", "D", "A", "K", "S", "H", "!"}
	/*var sn = ""
	for dd := range strSlice {
		sn += strSlice[dd]
	}
	fmt.Printf("Mine :%v\n", sn)*/

	var strBuilder strings.Builder
	for dd := range strSlice {
		strBuilder.WriteString(strSlice[dd])
	}
	var ddd = strBuilder.String()
	fmt.Printf("Mine :%v\n", ddd)

}

func Timeloop(slc []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slc) <= n {
		slc = append(slc, 1)
	}
	return time.Since(t0)

}
func division(nr int, dr int) (int, int, error) {

	var err error
	if dr == 0 {
		err = errors.New("can't divided by Zero")
		return 0, 0, err
	}

	div := nr / dr
	rem := nr % dr
	return div, rem, err
}
