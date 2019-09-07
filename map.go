
import "fmt"
//创建Vertex结构体
type Vertex struct {
	Lat, Long float64
}
//声明一个key为string类型,value为Vertex类型的map
var m map[string]Vertex

func main() {
  //通过make函数生成map
	m = make(map[string]Vertex)
  //对map m进行赋值
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
  //删除map中的元素
	delete(m,"Bell Labs")
  //查找map是否有key为Bell Labs的元素
	_,ok := m["Bell Labs"]
	if ok {
		fmt.Println("is in map",m)
	}else {
		fmt.Println("not in map",m)
	}

	fmt.Println(m)
}
