---
title: "Golang 面试题整理及分析"
date: 2022-01-12T14:28:31+08:00
draft: false
hideToc: false
enableToc: true
description: "Golang 面试题整理及分析"
enableTocContent: false
author: leoay
authorEmoji: 🎅
pinned: true
tags:
- 
series:
- 
categories:
- 
image: images/face/面试golang.png
---

![](https://pic4.zhimg.com/v2-49ec2bdf975ead3536bbb647f12ee22c)

### 一、谈谈Golang的并发机制和它使用的CPS并发模型

CPS并发模型使用“通信的方式共享内存”， 而传统的多线程通过共享内存的方式进行通信。

 它是用于描述两个独立的并发实体通过共享的通讯channel(管道)进行通信的并发模型。

CSP 中 channel 是第一类对象，它不关注发送消息的实体，而关注与发送消息时使用的channel。

Golang 中 channel 是被单独创建并且可以在进程间传递，它的通信模式类似于 boss-worker 模式的，一个实体通过将消息发送到 channel 中，然后又监听这个 channel 的实体处理，两个实体之间是匿名的，这个就实现实体中间的解耦，其中 channel 是同步的一个消息被发送到 channel 中，最终是一定要被另外的实体消费掉的，在实现原理上其实类似一个阻塞的消息队列。

Goroutine 是 Golang 实际并发执行的实体，它底层是使用协程( coroutine )实现并发，coroutine 是一种运行在用户态的用户线程，类似于 greenthread, go底层选择并使用coroutine的出发点是因为，它具有以下特点：

* 用户空间，避免了内核态和用户态的切换导致的成本。
* 可以由语言和框架层进行调度
* 更小的栈空间，允许创建大量的实例。

 Golang 中的Goroutine的特性：

 Golang 内部有三个对象：P对象( processor )代表上下文 (或者可以认为是CPU), M（work thread）代表工作线程，G对象(goroutine).

 正常情况下，一个CPU对象开启一个工作线程对象，线程去检查并执行 Goroutine 对象。碰到 Goroutine 对象阻塞的时候，会启动一个新的工作线程，已充分利用CPU资源。

 所以有时候线程对象会比处理器对象多很多。

 我们用下图表示P、M、G

 (图)

 G (Goroutine) : 我们所说的协程，为用户级的轻量级线程，每个Goroutine 对象中的sched 保存着其上下文信息。

 M (Machine)：对OS内核级线程的封装，数量对应真实的CPU数(真正干活的对象)。

 P (Processor) ：逻辑处理器，即为G和M的调度对象，用来调度 G 和M 之间的关联关系，其数量可通过GOMAXPROCS() 设置，默认为核心数。

 在单核情况下，所有 Goroutine 运行在同一个线程(M0) 中，每一个线程维护一个上下文(P), 任何时刻，一个上下文中只有一个 Goroutine, 其它 Goroutine 在 runqueue 中 (如下图所示)。

 当正在运行的G0阻塞的时候 （可以需要IO），会再创建一个线程 （M1）, P 转到新的线程中去运行。

当M0返回时，它会尝试从其他线程中“偷”一个上下文过来，如果没有偷到，会把 Goroutine 放到 `Global runqueue` 中去，然后把自己放入线程缓存中。
上下文会定时检查 `Global runqueue`。

`Golang` 是为并发而生的语言，Go语言是为数不多的在语言层面实现并发的语言；也正是Go语言的并发特性，吸引了全球无数的开发者。

Golang的CSP并发模型，是通过Goroutine和Channel来实现的。

Goroutine 是Go语言中并发的执行单位。有点抽象，其实就是和传统概念上的“线程”类似，可以理解为“线程”。

Channel 是 Go 语言中各个并发结构体之间的通信机制。通常 Channel, 是各个 Goroutine 之间的通信的 “管道”， 有点类似于 Linux 中的管道。

通信机制channel 也很方便，传数据用 `channel <- data`, 取数据用`<- channel`。

而且不管是传还是取，肯定阻塞，直到另外的 goroutine 传或者取为止。

因此 GPM 的简要概括即为： 事件循环，线程池，工作队列。

### 二、Golang 常用的并发模型

Golang 中常用的并发模型有三种：

#### 1. 通过 channel 发送消息实现并发控制

无缓存的通道指的是通道的大小为0， 也就是说，这种类型的通道在接收前没有能力保存任何值，它要求发送 goroutine 和 接收 goroutine 同时准备好，才可以完成发送和接收操作。

从上面无缓冲的通道定义来看，发送 goroutine 和接收 goroutine 必须是同步的，同时准备后，如果没有同时准备好的话，先执行的操作会阻塞等待，直到另一个相对应的操作准备好为止。这种无缓冲的通道我们也称之为同步通道。

```go
func main() {
    ch := make(chan struct() {})
    go func() {
        fmt.Println("start working")
        time.Sleep(time.Second * 1)
        ch <- struct{}{}
    }()

    <-ch

    fmt.Println("finished")
}
```

当主 goroutine 运行到 `<-ch` 接受 channel 的值的时候，如果该 channel 中没有数据，就会一直阻塞等待，直到有值。这样就可以简单实现并发控制。

#### 2. 通过 sync 包中的WaitGroup 实现并发控制

Goroutine 是异步执行的，有时候为了防止在结束 main 函数的时候结束掉 Goroutine， 所以需要同步等待，这个时候就需要用 WaitGroup 了，在 sync 包中，提供了 WaitGroup，它会等待它收集的所有 goroutine 任务全部完成。

在 WaitGroup 里主要有三个方法：

* Add， 可以添加或减少 goroutine 的数量
* Done， 相当于Add(-1)
* Wait，执行后会堵塞主线程，直到 WaitGroup 里的值减至 0.

在主 goroutine 中 Add(delta int) 索要等待 goroutine 的数量。在每一个 goroutine 完成后 Done() 表示这一个 gorouine 已经完成，当所有的 goroutine 都完成后，在主 goroutine 中 WaitGroup 返回。

```go
func main() {
    var Wg sync.WaitGroup
    var urls = []string{
        "http://www.golang.org/",
        "http://www.google.com/",
    }
    for _, url := urls {
        wg.Add(1)
        go func(url string) {
            defer wg.Done()
            http.Get(url)
        }(url)
    }
    wg.Wait()
}
```
在 `Golang` 官网中对于 `WaitGroup` 介绍是 `A WaitGroup must not be copied after first use`, 在WaitGroup 第一次使用后，不能被拷贝。

应用示例：

```go
func main() {
    wg := sync.WaitGroup{}
    for i:= 0; i < 5; i++ {
        wg.Add()
        go func(wg sync.WaitGroup, i int) {
            fmt.Println(i)
            wg.Done()
        }(wg, i)
    }
    wg.Wait()
    fmt.Println("exit")
}
```

运行：

```go
i:1i:3i:2i:0i:4fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0xc000094018)
        /home/keke/soft/go/src/runtime/sema.go:56 +0x39
sync.(*WaitGroup).Wait(0xc000094010)
        /home/keke/soft/go/src/sync/waitgroup.go:130 +0x64
main.main()
        /home/keke/go/Test/wait.go:17 +0xab
exit status 2
```

它提示所有的 `goroutine` 都已经睡眠了，出现了死锁。这是因为 wg 给拷贝传递到了 goroutine 中，导致只有 Add 操作，其实 Done 操作是在 wg 的副本执行的。

因此 Wait 就会死锁。

这个第一个修改方式： 将匿名函数中 wg 的传入类型改为 `*sync.WaitGroup`, 这样就能引用到正确的 `WaitGroup` 了。

这个第二个修改方式：将匿名函数中的 wg 的传入参数去掉，因为Go支持闭包类型，在匿名函数中可以直接使用外面的 wg 变量。

#### 3. 在Go 1.7 以后引进的强大的Context上下文，实现并发控制

通常在一些简单场景下使用 channel 和 WaitGroup 已经足够了，但是当面临一些复杂多变的网络并发场景下 `channel` 和 `WaitGroup` 显得有些力不从心了。

比如一个网络请求 Request， 每个Request 都需要开启一个 goroutine 做一些事情， 这些 goroutine 又可能会开启其他的 goroutine, 比如数据库和RPC服务。

所以我们需要一种可以跟踪 `goroutine` 的方案，才可以达到控制他们的目的，这就是 Go 语言为我们提供的 `Context`， 称之为上下文非常贴切，它就是 goroutine 的上下文。

它是包括一个程序的运行环境、现场和快照等。每个程序要运行时，都需要知道当前程序的运行状态，通常 Go 将这些封装在一个 Context 里，再将它传给要执行的 goroutine。

context 包主要就是用来处理多个 goroutine 之间共享数据。及多个 goroutine 的管理。

context 包的核心是 struct Context, 接口声明如下：

```go
// A Context carries a deadline, cancelation signal, and request-scoped values
// across API boundaries. Its methods are safe for simultaneous use by multiple
// goroutines.
type Context interface {
    // Done returns a channel that is closed when this `Context` is canceled
    // or times out.
    // Done() 返回一个只能接受数据的channel类型，当该context关闭或者超时时间到了的时候，该channel就会有一个取消信号
    Done() <-chan struct{}

    // Err indicates why this Context was canceled, after the Done channel
    // is closed.
    // Err() 在Done() 之后，返回context 取消的原因。
    Err() error

    // Deadline returns the time when this Context will be canceled, if any.
    // Deadline() 设置该context cancel的时间点
    Deadline() (deadline time.Time, ok bool)

    // Value returns the value associated with key or nil if none.
    // Value() 方法允许 Context 对象携带request作用域的数据，该数据必须是线程安全的。
    Value(key interface{}) interface{}
}
```

Context 对象是线程安全的，你可以把一个 Context 对象传递给任意个数的 goroutine, 对它执行取消操作时，所有 goroutine 都会接受到取消信号。

一个 Context 不能拥有 Cancel 方法，同时我们也只能 Done channel 接收数据。其中的原因是一致的： 接收取消信号的函数和发送信号的函数通常不是一个。

典型的场景是：父操作作为子操作操作启动 goroutine, 子操作也就不能取消父操作。

### 三、Go 中对 nil 的 slice 和 空Slice的处理是一致的吗？

首先 Go 的 JSON 标准对 `nil slice`和 空 `slice` 的处理是不一致的。

通常错误的用户，会报数组越界的错误，因为只是声明了 slice，却没有给实例化的对象。

```go
var slice []int
slice[1] = 0
```

此时 slice 的值是 nil，这种情况可以用于需要返回 slice 的函数， 当函数出现异常的时候， 保证函数依然会有 nil 的返回值。

empty slice 是指 slice 不为 nil, 但是 slice没有值， slice 的底层的空间是空的，此时的定义如下：

```go
slice := make([]int, 0)
slice := []int{}
```

当我们查询或者处理一个空的列表的时候，这非常有用，它会告诉我们返回的是一个列表，但是列表内没有任何值。

总之， `nil slice` 和 `empty slice` 是不同的东西，需要我们加以区分的。

### 四、协程、线程、进程的区别

* 进程

进程是具有一定独立功能的程序，是系统进行资源分配和调度的一个独立单位。

每个进程都有自己的独立内存空间，不同进程通过进程间通信来通信。由于进程比较重，占据独立的内存，所以上下文进程间的切换、开销(栈、寄存器、虚拟内存、文件句柄等)比较大，但相对比较稳定安全。

* 线程

线程是进程的一个实体，线程是内核态，而且是CPU调度和分派的基本单位，它是比进程更小的能独立运行的基本单位。线程自己基本上不拥有系统资源，只拥有一点在运行中必不可少的资源(如程序计数器， 一组寄存器和栈)， 但是它可与同属一个进程的其他线程共享进程所拥有的全部资源。

线程间通信


* 协程

协程是一种用户态的轻量级线程，协程的调度完全由用户控制。协程拥有自己的寄存器上下文和栈。

协程调度切换时，将寄存器上下文和栈保存到其他地方，在切回来的时候，恢复先前保存的寄存器上下文和栈，直接操作栈则基本没有内核切换的开销，可以不加锁地访问全局变量，所以上下文的切换非常快。

### 五、Go中数据竞争问题怎么解决？

`Data Race` 问题可以使用互斥锁 `sync.Mutex`, 或者也可以通过 CAS 无锁并发解决。
其中使用同步访问共享数据或者CAS无锁并发是处理数据竞争的一种有效的方法。

golang 在1.1之后引入了竞争检测机制，可以使用 `go run -race` 或者 `go build -race` 来进行静态检测。

其在内部的实现是，开启多个协程执行同一个命令， 并且记录下每个变量的状态。

竞争检测器基于C/C++的 `ThreadSanitizer` 运行时库， 该库在 Google 内部代码基地和 Chromium 找到许多错误。这个技术在2012年九月集成到 Go 中，从那时开始，它已经在标准库中检测到42个竞争条件。现在，它已经是我们持续构建过程的一部分，当竞争条件出现时，它会继续捕捉到这些错误。

竞争检测器已经完全集成到Go工具链中，仅仅添加-race 标志到命令行就使用了检测器。

```go
$ go test -race mypkg    //测试包
$ go run -race mysrc.go  //编译和运行程序
$ go build -race mycmd   //构建程序
$ go install -race mypkg
```
要想解决数据竞争的问题，可以使用互斥锁`sync.Mutex`, 解决数据竞争(Data race)，也可以使用管道解决，使用管道的效率要比互斥锁高。


### 六、Go中值接收器和指针接收器的区别

Go 中方法能给用户自定义的类型添加新的行为。它和函数的区别在于方法有一个接收者，给一个函数添加一个接收者，那么它就变成了方法。接受者可以是值接收者，也可以是指针接收者。

在调用方法的时候，值类型既可以调用值接收者的方法，也可以调用指针接收者的方法，指针类型既可以调用指针接收者的方法，也可以调用值接收者的方法。

也就是说，不管方法的接收者是什么类型，该类型的值和指针都可以调用，不必严格符合接收者的类型。

```go
package main

import "fmt"

type Person struct {
    age int
}

func (p Person) Elegance() int {
    return p.age
}

func (p *Person) GetAge() {
    p.age += 1
}

func main() {
    // p1 是值类型
    p := Person{age: 18}

    // 值类型 调用接收者也是值类型的方法
    fmt.Println(p.howOld())
    
    // 值类型，调用接收者是指针类型的方法
    p.GetAge()

    fmt.Println(p.GetAge())

    // ---------------------------------

    // p2 是指针类型
    p2 := &Person{ age: 100}

    // 指针类型 调用接收者是值类型的方法
    fmt.Println(p2.GetAge())

    // 指针类型 调用接收者也是指针类型的防范
    p2.GetAge()
    fmt.Println(p2.GetAge())
}
```

运行
```go
18
19
100
101
```

值类型调用者： 方法会使用调用者的一个副本，类似于“传值”，使用值得引用调用方法，上例中，p1.GetAge()实际上是 (&p1).GetAge().

指针类型调用者： 指针被解引用为值，上例中，p2.GetAge() 实际上是 (*p1).GetAge(), 实际上也是“传值”，方法里的操作会影响到调用者，类似于指针传参，拷贝了一份指针

如果实现了接收者是值类型的方法，会隐含地实现了接收者是指针类型的方法。

如果方法的接收者是值类型，无论调用者是对象还是对象指针，修改都是对象的副本，不影响调用者，如果方法的接收者是指针类型，则调用者修改的是指针指向的对象本身。

通常我们使用指针作为方法的接收者的理由：

* 可以避免在每次调用方法时复制该值，在值的类型为大型结构体时，这样做会更加高效。

因而呢，我们是使用值接收者还是指针接收者，不是由该方法是否修改了调用者(也就是接收者)来决定，而是应该基于该类型的本质。

如果类型具备“原始的本质”，也就是说它的成员都是由 Go 语言里内置的原始类型，如字符串，整数值等，那就定义值接收者类型的方法。像内置的引用类型，如 slice, map, interface，channel, 这些类型比较特殊，声明他们的时候，实际上是创建了一个 header， 对于他们也是直接定义值接收者类型的方法。这样，调用函数时，是直接 copy 了这些类型的 header, 而 header 本身就是为复制设计的。

如果类型具备非原始的本质，不能被安全地复制，这种类型总是应该被共享，那就定义指针接收者的方法。比如 go 源码里的文件结构体(struct file)就不应该被复制，应该只有一份实体。

接口值得零值是指动态类型和动态值都为 nil。当仅且当这两部分的值都为 nil 的情况下，这个接口值就才会被认为接口值 == nil。


### 七、Go的对象在内存中是怎样分配的

Go中的内存分类并不像 TCMalloc 那样分成小、中、大对象，但是它的小对象里又细分了一个Tiny对象， Tiny 对象指大小在 1Byte到16Byte之间并且不包含指针的对象。

小对象和大对象只用大小划定，无其他区分。

大对象指大小大于32kb, 而小对象是在 mcache 中分配的，而大对象是直接从 mheap 分配的，从小对象的内存分配看起。

Go 的内存分配原则：

Go 在程序启动的时候，会先向操作系统申请一块内存 （注意这时还只是一段虚拟的地址空间，并不会真正地分配内存），切成小块后进行管理。

申请到的内存块被分配了三个区域，在x64上分别是 `512MB`, `16GB`， `512GB` 大小。

arena 区域就是我们所谓的堆区，Go 动态分配的内存都是在这个区域, 它把内存分割成 8kb 大小的页， 一些页组合起来称为 `mspan` 。

bitmap 区域标识 arena 区域哪些地址保存了对象，并且用 4bit 标志位表示对象是否包含指针、GC 标记信息。
bitmap 中一个 byte 大小的内存对象 arena 区域中 4个指针大小（指针大小为 8B ）的内存，所以 bitmap 区域的大小是 `512GB/(4*8B)=16GB`。

此外，我们还可以看到 bitmap 的高地址部分指向 arena 区域的低地址部分，这里 bitmap 的地址是由高地址向低地址增长的。

spans 区域存放 `mspan` （是一些 `arena` 分割的页组合起来的内存管理基本单元，后文会再讲）的指针，每个指针对应一页，所以 spans 区域的大小就是 `512GB/8KB*8B=512MB`。

除以 8KB 是计算 `arena` 区域的页数，而最后乘以 8 是计算 `spans` 区域所有指针的大小。创建 `mspan` 的时候，按页填充对应的 `spans` 区域，在回收 object 时，根据地址很容易就能找到它所属的 `mspan`。


### 八、栈的内存是怎么分配的

栈和堆只是虚拟内存上两块不同功能的内存区域。

* 栈在高地址，从高地址向低地址增长
* 堆在低地址，从低地址向高地址增长。

### 九、Golang中Context包的解析

1. Context 包的用途

在Go http包中



### 十、GC的原理

三色标记法

标记清除原理

黑白灰

STW

### 十一、Go中的内存分配