package main

// 通过反射读取结构体的属性 调用方法
import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func (s Student) Print() {
	fmt.Println("调用了Print方法")
}

func (s Student) GetName() string {
	return s.Name
}

func (s Student) SetName(name string) {
	s.Name = name
}

func (s Student) SetStudent(name string, age int) {
	s.Name = name
	s.Age = age
}

func main() {
	stu := Student{Name: "张三", Age: 18}

	testStructReflect(stu)

}

func testStructReflect(s interface{}) {

	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)
	fmt.Println(t.Name(), t.Kind())
	fmt.Println(v)
	// 获取字段个数
	fmt.Println("字段个数：", v.NumField())

	// 获取字段对应的值
	for i := 0; i < v.NumField(); i++ {
		fmt.Println(v.Field(i))

	}
	// 获取方法个数
	fmt.Println("方法个数：", v.NumMethod())
	// 获取方法 方法首字母要大写，顺序按照ASCII码顺序
	// 调用 Print()
	v.Method(1).Call(nil)
	// 调用 GetName()
	name := v.Method(0).Call(nil)[0].String()
	fmt.Println(name)

	// 调用 SetStudent()
	var params []reflect.Value
	params = append(params, reflect.ValueOf("王五"))
	params = append(params, reflect.ValueOf(20))
	v.Method(3).Call(params)

}
