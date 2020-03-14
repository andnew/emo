package ch17

// 值接收器和指针接收器的区别

type Employee1 struct {
	Name string
	Age  int
}

// 值接收器的方法
func (e Employee1) ChangeName(newName string) {
	e.Name = newName
}

// 指针接收器
func (e *Employee1) ChangeAge(newAge int) {
	e.Age = newAge
}
