// package main

// import (
// "fmt"
// "golang.org/x/text/currency"
// )

// "fmt"
// "strings"

// "bytes"
// "fmt"
// "strings"
// "math"

// func foo() {

// }

// func main() {
// 1.打印
// fmt.Printf("\"a\": %v\n", "a")

// 2.变量的定义
// var name string = "wuyupei"
// var age int = 20

// var (
// 	name   string = "wuyupei"
// 	age    int    = 20
// 	isLeft bool   = false
// )

// name := "wuyupei"
// age := 20
// isLeft := true

// var name, age, isLeft = "wuyupei", 20, true

// const name string = "wuyupei"
// const age int = 20
// const isLeft = true
// fmt.Printf("name: %v\n", name)
// fmt.Printf("age: %v\n", age)
// fmt.Printf("isLeft: %v\n", isLeft)

// 3.数据类型
// var a string = "wuyupei"
// var b int = 20
// var c float32 = 1.00
// var d = &a

// var e = [...]int{1,2,3}

// fmt.Printf("%T\n", a)
// fmt.Printf("%T\n", b)
// fmt.Printf("%T\n", c)
// fmt.Printf("%T\n", d)
// fmt.Printf("%T\n", foo)
// fmt.Printf("e: %T\n", e)

// 数组
// 创建一个空数组
// var names [10]string
// for i := 0; i < 10; i++ {
// 	names[i] = "1"
// }
// fmt.Printf("names: %v\n", names)

// name := [...]string{"a"}
// fmt.Printf("name: %v\n", name)

//  for 循环
// var i int
// for i = 0; i < 10; i++ {
// 	fmt.Printf("i: %v\n", i)
// }

// var name = "hello word"
// fmt.Printf("name: %v\n", name)

// bool类型  if语句

// var age bool = true

// age := 18

// age = 2
// isMan := age >= 18
// if isMan {
// 	fmt.Printf("是一个成年人")
// } else if age == 2 {
// 	fmt.Printf("不是一个成年人")
// } else {
// 	fmt.Printf("AA")
// }

// 数字类型
// var a int
// var b int16
// var c int32
// var d int64
// fmt.Printf("a: %v %v %T\n", a, math.MaxInt, a)
// fmt.Printf("a: %v\n", b)
// fmt.Printf("a: %v\n", c)
// fmt.Printf("a: %v\n", d)

// 字符串
// var name string = "wuyupei"
// var age = "20"

// name := "wuyupei"
// age := "20"
// var res = name + age
// fmt.Printf("res: name=%sage=%s\n", name, age)
// fmt.Printf("res: name=%vage=%v\n", name, age)

// strings.Join
// var res = strings.Join([]string{name, age}, ",")
// fmt.Printf("res: %v\n", res)

// burrer.WriteString
// var buffer bytes.Buffer
// buffer.WriteString("a")
// buffer.WriteString("b")
// buffer.WriteString("c")
// fmt.Printf("value: %v\n", buffer.String())

// foo := [10]string{"10", "", "10"}
// fmt.Printf("foo: %v\n", foo)

// name := "wuyupei"
// print(name + "\r")
// text := "abcdefg hijklmnopq rstuvwxyz"
// print(text[12:] + "\n")
// print(text + "\n")
// res := strings.Split(text, "")
// fmt.Printf("res: %v\n", res)
// fmt.Printf("res: %T\n", res)
// fmt.Printf("res: %s\n", res)
// print("----------------------\n")
// print(res)
// age := 20
// print(age)

// 统计字符串中a的个数
// const str = "abadgasldkjfaaaa"
// var count int
// for i := 0; i < len(str); i++ {
// 	if str[i] == 'a' {
// 		count++
// 	}
// }
// fmt.Printf("count: %v\n", count)

// 打印9*9乘法表
// for i := 1; i < 10; i++ {
// 	for j := 1; j <= i; j++ {
// 		fmt.Printf("|%v * %v = %v|  ", i, j, i*j)
// 	}
// 	fmt.Printf("\n")
// }
// }

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/", index) // index 为向 url发送请求时，调用的函数
// 	log.Fatal(http.ListenAndServe("localhost:8000", nil))
// }
// func index(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "C语言中文网")
// }

package main

import (
	"fmt"
)

func main() {

	// 1.0
	// var person map[string]string
	// person = make(map[string]string)

	// 2.0
	// person := map[string]string{"age": "20"}

	// 3.0
	// var person = map[string]string{"age": "10"}

	// person["name"] = "wuyupei"
	// fmt.Printf("person: %v\n", person)

	// for key, value := range person {
	// 	fmt.Printf("key: %v\n", key)
	// 	fmt.Printf("value: %v\n", value)
	// }

	// 声明一个数组
	friends := [...]string{"liuxu", "wuyupei"}
	// b := friends[1:2]
	fmt.Printf("b: %T\n", "A")
	// a := append(friends[:1], friends[2:])
	// fmt.Printf("s: %v\n", a)

	var friendsTwo = [...]string{}
	var friendsThere = [...]string{"a"}
	fmt.Printf("friendsTwo: %v\n", friendsTwo)
	fmt.Printf("friendsThere: %v\n", friendsThere)
	fmt.Printf("friends: %v\n", friends)

	city := map[string]string{"name": "wuyupei", "age": "20"}
	cityTwo := make(map[string]string)
	var cityThere = map[string]string{"a": "B"}

	fmt.Printf("cityThere: %v\n", cityThere)
	fmt.Printf("cityTwo: %v\n", cityTwo)
	fmt.Printf("city: %v\n", city)
}
