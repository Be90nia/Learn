package main

import "fmt"

type person struct {
	name, sex string
	age       int
}

type student struct {
	person
	mail string
	qq   int
}

type persons struct {
	prev      *persons
	name, sex string
	age       int
	next      *persons
}

func main() {
	//初始化结构体
	var people person
	people.name = "name"
	people.sex = "male"
	people.age = 1
	fmt.Println(people)

	//指针声明结构体
	var user *person //这样声明指针结构体是不是初始化结构体,因此为nil
	fmt.Printf("%#v\n", user)

	user1 := new(person)
	fmt.Printf("%#v\n", user1)  //使用new进行声明结构体,new会初始化结构体
	fmt.Printf("%#v\n", *user1) //指针指向的数据值
	fmt.Printf("%#v\n", &user1) //引用指向的指针值

	//结构体指针初始化必须要通过new 或者 &的方式进行初始化
	var user2 *person = &person{"", "", 1}
	fmt.Println(user2)

	//一次性初始化
	//变量名 = 结构体名{成员值1,成员值2}
	//变量名 = 结构体名{成员名:成员值1,成员名:成员值2}

	//嵌套结构体

	user3 := new(person)
	user3.name = "begonia"
	user3.sex = "male"
	user3.age = 1
	stu := student{*user3, "begoniali@vip.qq.com", 12345678645}
	fmt.Printf("%#v\n", stu)

	//结构体指针
	var stu1, stu2, stu3 persons
	stu1 = persons{nil, "张三", "男", 32, &stu2}
	stu2 = persons{&stu1, "李四", "男", 31, &stu3}
	stu3 = persons{&stu2, "王五", "男", 33, nil}
	fmt.Printf("%#v\n", stu2)

	TreeNodeTest1()
}
