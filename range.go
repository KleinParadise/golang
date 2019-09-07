package main
import "fmt"

func main(){
  //声明长度为10,元素类型为int的切片pow
	pow := make([]int, 10)

	fmt.Println("pow = ",pow)
	//range 用于迭代切片中的元素。即将pow中的元素赋值给i
	for i := range pow {
		fmt.Println("i = ",i)
		pow[i] = 1 << uint(i) // == 2**i
	}
	fmt.Println("pow2 = ",pow)

	for k, value := range pow {
		fmt.Printf("%d\n", value)
		fmt.Printf("%d\n", k)
	}

}
