package main

import (
	"fmt"
	"github.com/tidwall/gjson"
	"myhttp"
	"myrpc"
	"runtime"
)

const (
	i=1<<iota
	j=2<<iota
	k
	l
)

//结构体
type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}

func Factorial(n uint64)(result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

// fib returns a function that returns
// successive Fibonacci numbers.
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

//最大值
func max(num1, num2 int) int {
	/* 声明局部变量 */
	var result int

	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func isPrime(value int) bool {
	if value <= 3 {
		return value >= 2
	}
	if value%2 == 0 || value%3 == 0 {
		return false
	}
	for i := 5; i*i <= value; i += 6 {
		if value%i == 0 || value%(i+2) == 0 {
			return false
		}
	}
	return true
}



func main() {

	//go遍历
	fmt.Println("i=",i)
	fmt.Println("j=",j)
	fmt.Println("k=",k)
	fmt.Println("l=",l)
	fmt.Println("\nHello, world or 你好，世界 or καλημ ́ρα κóσμ or こんにちはせかい\n")
	fmt.Println(max(1,2))

	//指针 int默认值 0
	var  ptr *int
	fmt.Printf("ptr 的值为 : %x\n", ptr  )

	// 创建一个新的结构体
	fmt.Println(Books{"Go 语言", "shishijie.cc", "Go 语言", 6495407})
	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Go 语言", author: "shishijie.cc", subject: "Go 语言", bookId: 6495407})
	// 忽略的字段为 0 或 空
	fmt.Println(Books{title: "Go 语言", author: "shishijie.cc"})

	//数组的遍历
	nums := []int{1,2,3,4}
	for i,num := range nums {
		fmt.Printf("索引是%d,长度是%d\n",i, num)
	}

	//判断数字是否是质数
	fmt.Println(isPrime(7))

	//阶乘
	var i int = 15
	fmt.Println("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))

	//json的运用 gjson使用
	const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`
	value := gjson.Get(json, "name.last")
	fmt.Println(value.String())

	//defer关键字，默认 最后执行，并且是先进后出（栈模式）
	for	i:=0;i<5;i++{
		defer fmt.Printf("%d	",	i)
	}

	//查看cpu核数
	fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))

	//闭包 求斐波那契数列
	var f = fib()
	fmt.Println(f(),f(),f(),f(),f(),f())


	myrpc.Rpcgo()

	//myhttpfunc
	myhttp.HandleFunc()
	//myhttp 与func只能开启一个
	myhttp.Handles()
}

