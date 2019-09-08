package main

import "fmt"

type Observer interface {
	update()
}

type ConcreteObserverA struct {
}
// 此方法表示类型 T 实现了接口 I，但我们无需显式声明此事。
func (observerA ConcreteObserverA) update(){
	fmt.Println("ConcreteObserverA update enter")
}

type ConcreteObserverB struct {
}
// 此方法表示类型 T 实现了接口 I，但我们无需显式声明此事。
func (observerB ConcreteObserverB) update(){
	fmt.Println("ConcreteObserverB update enter")
}

type Subject struct {
	vectorArray []Observer
}

func NewSubject() *Subject {
    return &Subject{}
}

func (subject *Subject) attach(po Observer)  {
	_ = append(subject.vectorArray, po)
	return
}

func main(){
	subject := NewSubject()

	a1 := new(ConcreteObserverA)

	subject.attach(a1)

	for observer := range subject.vectorArray{
		fmt.Println("----------- AAAAAAAA")
		fmt.Println(observer)
	}

	fmt.Println(subject.vectorArray)

}




