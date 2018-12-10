// go 语言的每一个文件都是一个包，以小写字母开头的函数是私有的
// 每个文件都应该声明自己所处的包，一般包名与其所处的文件夹同名
// main 函数必须在 main 包里
package main

import (
	"log"
	"fmt"
	// 导入网络包；导入而不使用包中的方法，默认情况下 go 不允许导入不使用的包
	_ "github.com/goinaction/code/chapter2/sample/matchers"
)

// 每个包可以包含任意多的 init 函数，这些函数都会在程序执行开始的时候被调用
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// 与其他语言相反，go 语言中的变量描述在变量名与类型名之后
type user struct {
	name  string
	email string
}

// 与类型关联的“函数”被称为方法，自由的“函数”被称为 函数
func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n",
		u.name,
		u.email)
}

// changeEmail implements a method with a pointer receiver.
func (u *user) changeEmail(email string) {
	u.email = email
}

// 
func main() {
	// go 是静态语言； := 使用初始化值定义变量的类型与值
	bill := user{"Bill", "bill@email.com"}
	bill.notify()

	// Pointers of type user can also be used to call methods
	// declared with a value receiver.
	lisa := &user{"Lisa", "lisa@email.com"}
	lisa.notify()

	// Values of type user can be used to call methods
	// declared with a pointer receiver.
	bill.changeEmail("bill@newdomain.com")
	bill.notify()

	// Pointers of type user can be used to call methods
	// declared with a pointer receiver.
	lisa.changeEmail("lisa@newdomain.com")
	lisa.notify()
}
