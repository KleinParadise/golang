
package main
import "fmt"
//闭包
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}
//adder()函数返回一个参数为int并且返回为int的匿名函数
func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}
//pos,neg等于adder()return的func(x,int) int函数。通过该main函数实现对adder()函数的sum变量闭包操作。

