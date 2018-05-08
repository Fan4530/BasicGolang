package BasicGolang

import (
	"fmt"
	"strings"
	"math/big"
	"time"
	"math/rand"
	//为什么？？？
	//"lib"
)
type people struct {
	name string
	age int
}
//还未总结
func main() {

	//var s string
	//fmt.Scanln(&s)
	//fmt.Printf(s)
	//
	//reader := bufio.NewReader(os.Stdin)// from the standard input, from the console
	//fmt.Println("Enter text: ")
	//// strconv
	//// string converter -> float 64
	//str, _ := reader.ReadString('\n')
	//f, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println("value of number:", f)
	//}
	// strings package
	l1 := "hello world"
	l2 := "HELLO WORLD"
	fmt.Println(strings.EqualFold(l1, l2))
	fmt.Println(strings.Contains(l1, "ell"))

	// math
	var b1, b2, b3, bigSum big.Float
	b1.SetFloat64(23.5)
	b1.SetFloat64(65.1)
	b1.SetFloat64(65.3)
	// %g: 根据数值给的数值自动选择e还是f
	// bigSum, 因为是个object，
	bigSum.Add(&b1, &b2).Add(&bigSum, &b3)
	fmt.Printf("Bigsum = %.10g\n", &bigSum)
	fmt.Printf("%f", &bigSum)
	p := people{"zhanfan", 5}
	//问题： 为什么结构体不用&而big sum要用&，bigsum可以理解因为他是object，struct就不是了么
	// %v default type
	fmt.Printf("\n%v\n", p)

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("\ngo launch at: %s\n", t)
	fmt.Printf("\ngo launch at: %s\n", time.Now())
	//make and new(no initialize, be careful with map)

	// randomly, generate the random according to the tiem of milli
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(6) + 1
	fmt.Println("day", dow)

	//return type可以在 signiture里面定义，也可以在函数里面定义。
	fmt.Println(fullName1())
	fmt.Println(fullName2())

}
func fullName1() (nameTest string) {
	nameTest = "zhangfan"
	return
}
func fullName2() (string) {
	nameTest := "zhangfan"
	return nameTest
}