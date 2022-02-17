---
title: "一篇文章搞懂Golang中的interface(1)"
date: 2022-01-09T10:33:41+09:00
description: "Syntax highlighting test"
draft: false
hideToc: false
enableToc: true
enableTocContent: false
author: leoay
authorEmoji: 🎅
pinned: true
authorImageUrl: "/images/whoami/shuaian.jpg"
tags:
- golang
- interface
series:
-
categories:
- golang
image: images/feature2/color-palette.png
---

你好，我是 leoay, 一个单纯的技术人，

你好，我是 leoay，又好久不见了，好像上篇文章立的 flag 又被打破了， 如果不知道，那就前往上篇文章看看我立的 flag。 

果然， 同时更新维护两个公众号是一件不太容易的事情，熟悉我的朋友应该知道，我基本上都一值在 leoay 这个号更新，而且是日更，所以技术号有时候就会被忽略掉了，

毕竟每天要和大部分人一样要工作。

那么问题来了，什么是一个接口呢？

在Go语言中，接口是一个方法的集合，当一个类型中定义了这个接口中的所有方法时，我们也将这叫做实现了这个接口。

这个很像面向对象编程范式中提到的接口。

接口指定了一个类型中拥有的方法，也决定了这个类型怎么去实现这些方法。

举个例子，洗衣机可以被认为是一个有洗涤和干燥方法的接口，那么我们就可以把任何提供洗涤和干燥方法的类型，就叫做实现了洗衣机接口。

下面我们用代码进行更加详细地说明：

### 声明和实现一个接口

```go
package main

import (
	"fmt"
)

//interface definition
type WashingMachine interface {
	Cleaning() (string, error)
	Drying() (string, error)
}

type MyWashingMachine string

//MyWashingMachine 实现了 WashingMachine
func (mwm MyWashingMachine) Cleaning() (string, error) {
	//do clean work
	return "Clean Ok!", nil
}

func (mwm MyWashingMachine) Drying() (string, error) {
	//do dry work
	return "Drying Ok!", nil
}

func main() {
	var v WashingMachine
	name := MyWashingMachine("Leoay Home")
	v = name
	resut, err := v.Cleaning()
	fmt.Println("result: ", resut, "err is: ", err)

	resut, err = v.Drying()
	fmt.Println("result: ", resut, "err is: ", err)
}
```

在上面的代码中我们创建了一个 WashingMachine 接口类型，其中有两个方法，分别是 Cleaning ()和Drying()

然后，我们定义了一个 MyWashingMachine 类型，接着我们写了两个方法分别是 Cleaning() 和 Drying(), 并将 MyWashingMachine 作为方法的接收类型

这个时候，我们就可以说 MyWashingMachine 实现了 WashingMachine 接口。这个与 java 有很大不同，在java中我们一般使用 implements 这个关键字表示实现了一个接口。

而在Go语言中，只需要这个类型包含接口中的所有方法即可。

所以在下面的代码中，我们可以直接用 v 调用 Cleaning() 和 Drying() 这两个方法，因为 WashingMachine 已经实现了 WashingMachine 中的方法。

到了一步，我们就创建了一个接口，怎么样，是不是超级简单。

其实，如果你更深入学习 Go 语言时，你会发现接口在Go项目开发中使用的特别频繁，一不留神它就出现在你眼前，不过如果不了解的话，就会感到一头雾水。

### 接口的使用实践

通过上面的例子，我们知道了怎么创建并实现一个接口，但是并没有真正说明白怎么在实际项目中使用。

从上面代码中我们可以发现这一行代码 `name := MyWashingMachine("Leoay Home")` 

那么，这个有什么用呢？

如果我们直接使用 name 调用 Cleaning() 和 Drying() 函数时，会出现什么问题呢？

这个时候虽然也能正常输出，但是没有用到接口。

下面，我就用一个实例说明一下接口的使用。

我们将编写一个简单的程序，根据员工的个人工资计算公司的总费用。

```go
package main

import (  
    "fmt"
)

type SalaryCalculator interface {  
    CalculateSalary() int
}

type Permanent struct {  
    empId    int
    basicpay int
    pf       int
}

type Contract struct {  
    empId    int
    basicpay int
}

//正式员工的工资是基本工资和奖金的总和
func (p Permanent) CalculateSalary() int {  
    return p.basicpay + p.pf
}

//试用期员工的工资是单独的基本工资
func (c Contract) CalculateSalary() int {  
    return c.basicpay
}

//总费用是通过迭代 SalaryCalculator 切片并求和来计算的
func totalExpense(s []SalaryCalculator) {  
    expense := 0
    for _, v := range s {
        expense = expense + v.CalculateSalary()
    }
    fmt.Printf("每个月的总支出 ￥%d", expense)
}

func main() {  
    pemp1 := Permanent{
        empId:    1,
        basicpay: 5000,
        pf:       20,
    }
    pemp2 := Permanent{
        empId:    2,
        basicpay: 6000,
        pf:       30,
    }
    cemp1 := Contract{
        empId:    3,
        basicpay: 3000,
    }
    employees := []SalaryCalculator{pemp1, pemp2, cemp1}
    totalExpense(employees)
}
```

从上面的代码中我们可以看到 我们在 SalaryCalculator 接口中声明了一个 CalculateSalary() 方法

在公司里有两种雇员，正式员工和试用期员工，分别用 Permanent 和 Contract 两个结构体表示，正式员工的工资包含基本工资和奖金，试用期员工的工资只有基本工资，但是我们希望只用一个方法计算员工的工资，所以我们就分别用 Permanent 和 Contract 实现了 SalaryCalculator 接口，这样无论员工是哪种类型，都有可以用 CalculateSalary 方法计算薪水了。

然后我们定义了一个总的计算薪水支出的方法 totalExpense， 这个方法将 SalaryCalculator 切片作为参数，然后通过这个切片将所有的员工信息传到方法中去，然后在内部调用 CalculateSalary 方法计算每个员工的薪水并求和

执行上面的代码我们可以最后的输出结果:

每个月的总支出 ￥14050

这样做的最大优点是 totalExpense 可以扩展到任何新员工类型而无需更改任何代码。

假设公司增加了 Freelancer 一种工资结构不同的新型员工。

Freelancer 可以只在 slice 参数中传递给 totalExpense 而无需对 totalExpense 函数进行任何一行代码更改 

这个方法会做它应该做的事情，Freelancer 也会实现 SalaryCalculator 接口

下面我们就修改这个程序增加一种新的雇员 Freelancer， 其薪资是收入效率和总工作时间的乘积

```go
package main

import (  
    "fmt"
)

type SalaryCalculator interface {  
    CalculateSalary() int
}

type Permanent struct {  
    empId    int
    basicpay int
    pf       int
}

type Contract struct {  
    empId    int
    basicpay int
}

type Freelancer struct {  
    empId       int
    ratePerHour int
    totalHours  int
}

//正式员工的工资是基本工资和奖金的总和
func (p Permanent) CalculateSalary() int {  
    return p.basicpay + p.pf
}

//试用期员工的工资是单独的基本工资
func (c Contract) CalculateSalary() int {  
    return c.basicpay
}

//自由职业者的薪资
func (f Freelancer) CalculateSalary() int {  
    return f.ratePerHour * f.totalHours
}

func totalExpense(s []SalaryCalculator) {  
    expense := 0
    for _, v := range s {
        expense = expense + v.CalculateSalary()
    }
    fmt.Printf("每月的总支出 ￥%d", expense)
}

func main() {  
    pemp1 := Permanent{
        empId:    1,
        basicpay: 5000,
        pf:       20,
    }
    pemp2 := Permanent{
        empId:    2,
        basicpay: 6000,
        pf:       30,
    }
    cemp1 := Contract{
        empId:    3,
        basicpay: 3000,
    }
    freelancer1 := Freelancer{
        empId:       4,
        ratePerHour: 70,
        totalHours:  120,
    }
    freelancer2 := Freelancer{
        empId:       5,
        ratePerHour: 100,
        totalHours:  100,
    }
    employees := []SalaryCalculator{pemp1, pemp2, cemp1, freelancer1, freelancer2}
    totalExpense(employees)
}
```

我们添加了一个 Freelancer 结构体。并声明了一个用Freelancer实现的CalculateSalary方法。

新的 totalExpense 方法中不需要更改其他代码， 因为 Freelancer 结构体也实现了该 SalaryCalculator 接口。

然后我们在main方法中添加了几个 Freelancer 类型的员工。执行程序后打印，

每月的总支出 ￥32450  

### 接口内部表示

可以认为接口在内部由 tuple(type, value) 中表示的。type 是接口的底层具体类型， value 保存具体类型的值。

为了更好地理解，我们写一段代码展示：

```go
package main

import (  
    "fmt"
)

type Worker interface {  
    Work()
}

type Person struct {  
    name string
}

func (p Person) Work() {  
    fmt.Println(p.name, "is working")
}

func describe(w Worker) {  
    fmt.Printf("Interface type %T value %v\n", w, w)
}

func main() {  
    p := Person{
        name: "Naveen",
    }
    var w Worker = p
    describe(w)
    w.Work()
}
```

从上面的代码我们可以看到，Worker 接口有一个方法 Work(), 而 Person 结构体类型实现了该接口。

在main函数中我们定义了一个 Person 类型的 p, 并将他赋值给 Worker 类型的变量 w, 那么现在 w 的类型就变成了 Person, 而且其包含一个变量 name 值为 Naveen。

而 describe 函数则打印了 Worker 接口的具体类型和值，结果输出：

```shell
Interface type main.Person value {Naveen}
```

接下来我们更深入地了解一些怎么获取接口底层的值。


### 空接口

一个没有方法的接口就是空接口， 用 interface{} 表示， 因为空接口中没有方法，所以所有的类型都实现了一个空接口。

```go
package main

import (  
    "fmt"
)

func describe(i interface{}) {  
    fmt.Printf("Type = %T, value = %v\n", i, i)
}

func main() {  
    s := "Hello World"
    describe(s)
    i := 55
    describe(i)
    strt := struct {
        name string
    }{
        name: "Naveen R",
    }
    describe(strt)
}
```

上面的代码中，因为函数 describe(i interface{}) 以一个空接口作为参数，所以任何类型的参数都可以被传入，这一点很像 C++ 或 Java 中的泛型。
In the program above, in line no.7, the describe(i interface{}) function takes an empty interface as an argument and hence any type can be passed.

因此，在上面的示例代码中我们可以 describe 函数中传入字符串、整型和结构体，最后结果输出：

```shell
Type = string, value = Hello World  
Type = int, value = 55  
Type = struct { name string }, value = {Naveen R}  
```

### 接口的类型断言

类型断言用于获取接口的底层值。

i.(T) 用于获取 i 具体类型为 T 的接口的底层值。

代码胜万言，下面我就用代码展示类型断言是怎么用的

```go
package main

import (  
    "fmt"
)

func assert(i interface{}) {  
    s := i.(int) //从 i 中获取 int 底层的值
    fmt.Println(s)
}
func main() {  
    var s interface{} = 56
    assert(s)
}
```
s 的实际类型是 int。我们使用 i.(int) 获取 i 底层的 int 的值

那么如果上面的代码的实际类型不是 int，会出现什么呢？看下面的代码

```go
package main

import (  
    "fmt"
)

func assert(i interface{}) {  
    s := i.(int) 
    fmt.Println(s)
}
func main() {  
    var s interface{} = "Steven Paul"
    assert(s)
}
```
上面的代码中，我们传入了一个字符串到 assert 函数中，想要从中获取一个整型的值，这段代码将会出现 panic，并打印如下信息：

“interface {} is string, not int.”

那么，现在该怎么办呢？怎么才能避免程序的崩溃呢？

其实我们可以这样解决：

因为 i.(T) 会返回一个 error 异常的，只要我们对它进行判断，就可以避免程序崩溃了，

v, ok := i.(T)  

如果 i 的实际类型是 T, 那么 v 就是i的值，ok 就是 true, 代码正常运行；

如果 i 的实际类型不是 T 的话，v 就返回空， ok就是 false, 代码也就不会崩溃了。

所以我们对上面的代码进行简单的修改，如下所示：

```go
package main

import (  
    "fmt"
)

func assert(i interface{}) {  
    v, ok := i.(int)
    fmt.Println(v, ok)
}

func main() {  
    var s interface{} = 56
    assert(s)
    var i interface{} = "Steven Paul"
    assert(i)
}
```

当 “Steven Paul” 传递给 assert 函数时，ok 将是 false ，因为 i 的具体类型不是 int，因此 v值为 0。所以该程序将输出：

```shell
56 true  
0 false  
```

### 类型开关

类型开关用于将接口的实际类型与多种情况下 case 语句中指定的类型进行比较。

类似于 switch case。唯一的区别是 case 指定的是类型而不是正常 switch 中的值。

类型开关的语法类似于类型断言。

在 i.(T) 类型断言的语法中，类型 T 应替换 type 为类型切换的关键字。

下面看一下代码中是怎么实现的：

```go
package main

import (  
    "fmt"
)

func findType(i interface{}) {  
    switch i.(type) {
    case string:
        fmt.Printf("I am a string and my value is %s\n", i.(string))
    case int:
        fmt.Printf("I am an int and my value is %d\n", i.(int))
    default:
        fmt.Printf("Unknown type\n")
    }
}
func main() {  
    findType("Naveen")
    findType(77)
    findType(89.98)
}
```

上面的代码中， switch i.(type) 指定了一个 a type switch， 每一个 case 语句将 i 的实际类型和指定类型比较。如果任何一个 case 匹配的话, 就打印出相应的语句。

最后程序输出如下：

```shell
I am a string and my value is Naveen
I am an int and my value is 77
Unknown type
```

89.98 类型 是 float64，不匹配任何情况，因此最后一行打印 Unknown type。

也可以将类型与接口进行比较。如果我们有一个类型并且该类型实现了一个接口，则可以将此类型与其实现的接口进行比较。

为了更清楚，让我们编写一个程序详细说明：

```go
package main

import "fmt"

type Describer interface {  
    Describe()
}
type Person struct {  
    name string
    age  int
}

func (p Person) Describe() {  
    fmt.Printf("%s is %d years old", p.name, p.age)
}
func findType(i interface{}) {  
    switch v := i.(type) {
    case Describer:
        v.Describe()
    default:
        fmt.Printf("unknown type\n")
    }
}

func main() {  
    findType("Naveen")
    p := Person{
        name: "Naveen R",
        age:  25,
    }
    findType(p)
}
```

在上面的程序中，Person 结构体实现了 Describer 接口。 然后我们在 findType 函数中使用 case 语句比较类型 v 和比较 Describer 接口类型。

而 p 是 Person类型，因此当我们把 p 传到 findType 中时，v 就是 Describer。

所以最后程序输出如下：

```shell
unknown type
Naveen R is 25 years old
```