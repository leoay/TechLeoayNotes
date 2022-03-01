---
title: "一篇文章搞懂Golang中的interface(2)"
date: 2022-01-09T10:33:41+09:00
description: "Syntax highlighting test"
draft: false
hideToc: false
enableToc: true
enableTocContent: false
author: leoay
authorEmoji: 🎅
pinned: false
tags:
- golang
- interface
series:
-
categories:
- golang
image: images/feature2/color-palette.png
---

Hi，大家好，我是 leoay，又好长一段儿时间不见了。

今天继续上次没写完的Go语言中的接口，上次写了《一篇文章彻底搞懂Go语言中的接口(1)》， 看标题就知道应该还有一篇， 不过这篇也拖得太久了， 再过几天就相隔3个月了。

不知道为什么我一开始写技术文，就觉得这是一件比较严谨的事情，所以很多地方写得比较谨慎一些，以至于每写完一篇就觉得又要很长时间不写了，所以每次都是理所当然地拖更了。

我很长时间都以为我拖更真实的原因就是这样，是因为我是一个严谨的人，所以不随随便便写技术文，而对于非技术号则更新得很是勤快，已经日更280多天了。

我也不知道为什么我每次写技术文的时候一定要带上自己的非技术号，可能这就是“爱”吧，而对于我拖更技术文的原因，我最近两天才算想明白了，一个字 —— 懒。

所以，想明白了之后，我立即开始写技术文章了，我也把这个号的简介改成了“学技术嘛，嘿嘿，就是玩儿!”。

其实，用玩的心态学技术，更容易让我们的大脑处于放松的状态，此时我们也更容易进入心流状态，学习东西效率也是最高的。

而以后我在写文章时，也会更多地考虑如何让技术变得好玩儿一些，我觉得这样的文章也容易被大多数人接受，如果只是堆一些官方理论文字和一堆代码，我自己想想都觉得都读不下去。

好了，废话先不多说了，下面开始进入正题。

今天继续说Go语言中的interface，如果还没有阅读过第一篇的请点击链接前往阅读。

# 指针接收器VS值接收器实现接口

在前面的文章中其实我们主要使用值接收器来实现接口，比如下面这样：

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

主要观察下面这部分代码，Work函数前面的```(p Person)``` 就是函数的接收器，我们看到传进来的是 Person 类型的值

```
func (p Person) Work() {  
    fmt.Println(p.name, "is working")
}
```

实现接口除了这样的方式之外，还有另一种就是用指针作为函数的接收器，我们看下面的代码：

```
package main

import "fmt"

type Describer interface {  
    Describe()
}
type Person struct {  
    name string
    age  int
}

func (p Person) Describe() { //使用值接收器实现 
    fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {  
    state   string
    country string
}

func (a *Address) Describe() { //使用指针接收器实现
    fmt.Printf("State %s Country %s", a.state, a.country)
}

func main() {  
    var d1 Describer
    p1 := Person{"Sam", 25}
    d1 = p1
    d1.Describe()
    p2 := Person{"James", 32}
    d1 = &p2
    d1.Describe()

    var d2 Describer
    a := Address{"Washington", "USA"}

    /* 不能直接用Address类型的值调用 Describe 函数，因为是 *Address类型的指针实现了 Describe 函数
    */
    //d2 = a
    d2 = &a //
    d2.Describe()
}
```

我们重点关注一下下面这段代码：
```
func (a *Address) Describe() { //使用指针接收器实现
    fmt.Printf("State %s Country %s", a.state, a.country)
}
```

这就是用指针实现了 Describer 接口中的 Describe函数，也就是实现了 Describer 这个接口。

因此，我们既可以用值作为函数的接收器实现接口，也可以用指针作为函数的接收器实现接口，但是需要注意的是，对于使用指针作为函数接收器实现时，我们在使用时，也只能用指针调用。

因此对于上面代码中的*Address实现的 Describe 函数，我们也只能用 *Address 类型的值 &a 调用 Describe。

那么这时候就有朋友要问了，这两种方式有什么区别呢？

请看下面的代码：

```
package main

import "fmt"

type Describer interface {
	Change()
}
type Person struct {
	name string
	age  int
}

func (p Person) Change() { //使用值接收器实现
	p.name = "leoay"
}

type Address struct {
	state   string
	country string
}

func (a *Address) Change() { //使用指针接收器实现
	a.state = "BeiJing"
	a.country = "China"
}

func main() {
	var d1 Describer
	p1 := Person{"Sam", 25}
	d1 = p1
	d1.Change()
	p2 := Person{"James", 32}
	d1 = &p2
	d1.Change()

	fmt.Println(p2)

	var d2 Describer
	a := Address{"Washington", "USA"}
	d2 = &a
	d2.Change()
	fmt.Println(d2)
}
```

运行之后输出：
```
{James 32}
&{BeiJing China}
```

其实该用哪种实现方式，还要看我们的目的，如果我们想要改变接收的数据的值，比如代码中要改变 a 也就是 Address类型的数据中参数的值，就用指针作为函数接收器，相反则用值作为函数接收器。


# 实现多个接口
一个类型可以实现多个接口，我们以下面的代码为例：

package main

import (  
    "fmt"
)

type SalaryCalculator interface {  
    DisplaySalary()
}

type LeaveCalculator interface {  
    CalculateLeavesLeft() int
}

type Employee struct {  
    firstName string
    lastName string
    basicPay int
    pf int
    totalLeaves int
    leavesTaken int
}

func (e Employee) DisplaySalary() {  
    fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeavesLeft() int {  
    return e.totalLeaves - e.leavesTaken
}

func main() {  
    e := Employee {
        firstName: "Naveen",
        lastName: "Ramanathan",
        basicPay: 5000,
        pf: 200,
        totalLeaves: 30,
        leavesTaken: 5,
    }
    var s SalaryCalculator = e
    s.DisplaySalary()
    var l LeaveCalculator = e
    fmt.Println("\nLeaves left =", l.CalculateLeavesLeft())
}

上面的代码有两个接口，分别是 SalaryCalculator 和 LeaveCalculator，Employee结构体实现了 SalaryCalculator 接口中的 DisplaySalary 方法和 LeaveCalculator 接口中的 CalculateLeavesLeft 方法，所以 Employee 同时实现了 SalaryCalculator 和 LeaveCalculator 接口。


在代码的第 41 行中，我们将 e 分配给 SalaryCalculator 接口类型的变量，在第 43 行中，我们将相同的变量 e 分配给 LeaveCalculator 类型的变量。

这种骚操作是可以的，因为 e 类型的 Employee 同时实现了 SalaryCalculator 和 LeaveCalculator 接口。

最后，让我们一起看一下程序的输出：

Naveen Ramanathan has salary $5200  
Leaves left = 25  

# 嵌入接口

虽然 go 不提供继承，但可以通过嵌入其他接口来创建新接口。

通过下面的代码，我们可以看到这是如何实现的。

package main

import (  
    "fmt"
)

type SalaryCalculator interface {  
    DisplaySalary()
}

type LeaveCalculator interface {  
    CalculateLeavesLeft() int
}

type EmployeeOperations interface {  
    SalaryCalculator
    LeaveCalculator
}

type Employee struct {  
    firstName string
    lastName string
    basicPay int
    pf int
    totalLeaves int
    leavesTaken int
}

func (e Employee) DisplaySalary() {  
    fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee) CalculateLeavesLeft() int {  
    return e.totalLeaves - e.leavesTaken
}

func main() {  
    e := Employee {
        firstName: "Naveen",
        lastName: "Ramanathan",
        basicPay: 5000,
        pf: 200,
        totalLeaves: 30,
        leavesTaken: 5,
    }
    var empOp EmployeeOperations = e
    empOp.DisplaySalary()
    fmt.Println("\nLeaves left =", empOp.CalculateLeavesLeft())
}

上述程序第 15 行中的 EmployeeOperations 接口是通过嵌入 SalaryCalculator 和 LeaveCalculator 接口创建的。

如果某个类型同时实现了 SalaryCalculator 和 LeaveCalculator 接口，那么就可以说这个类型实现 EmployeeOperations 接口。

如上代码中 Employee 结构实现了 EmployeeOperations 接口，因为它分别在第 29 行和第 33 行中为 DisplaySalary 和 CalculateLeavesLeft 方法提供了定义。

在第 46 行中，类型为 Employee 的 e 变量 被分配到 EmployeeOperations 类型的 empOp。然后，empOp 就可以调用 DisplaySalary() 和 CalculateLeavesLeft() 方法。

最后，我们可以看到程序输出如下内容：

Naveen Ramanathan has salary $5200  
Leaves left = 25  

接口零值
接口的零值为 nil。nil 接口既有其基础值 nil，也有具体类型为 nil。

请看下面这部分代码：

package main

import "fmt"

type Describer interface {  
    Describe()
}

func main() {  
    var d1 Describer
    if d1 == nil {
        fmt.Printf("d1 is nil and has type %T value %v\n", d1, d1)
    }
}

上面的代码中，d1是的 nil 接口，所以程序最后输出：

d1 is nil and has type <nil> value <nil>

如果我们尝试调用nil接口的方法，程序就会报panic错误，因为 nil 接口既没有底层值，也没有具体类型。

我们可以看下面这段儿代码：

package main

type Describer interface {  
    Describe()
}

func main() {  
    var d1 Describer
    d1.Describe()
}

因为上面的程序中 d1是 nil，所以当我们运行的时候，会出现如下错误：

runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0xffffffff addr=0x0 pc=0xc8527]"

以上就是对Go语言中 interface 的讲解，这是第二篇，如果你想看第一篇，请点击下面的链接前往。



## 参考链接： 
1. https://golangbot.com/interfaces-part-2/


你好，我是leoay, 我创建了一个专门用来交流技术知识的星球，方向不限，大家有任何技术欢迎到星球提问，

只要我知道的我会毫不保留地解答，不过目前主要专注于

后端：Golang、Python以及PHP

前端：Html、CSS、React、Vue

等技术方向的分享，当然语言只是工具，要完成一个完整的系统还需要多方面的知识，

因此这个星球还会涉及微服务、系统架构、数据结构与算法、操作系统、计算机网络等方面的技术，

当然我对星球的定位是成长性质的，因此刚开始门票会比较低，后期会随着星球的成长慢慢涨价。

感兴趣的大家可以扫码加入。