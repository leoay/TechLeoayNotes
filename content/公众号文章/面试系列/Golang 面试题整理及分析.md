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

### 谈谈Golang的并发机制和它使用的CPS并发模型

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

当M0返回时，它会尝试从其他线程中“偷”一个上下文过来，如果没有偷到，会把 Goroutine 放到 `Golbal runqueue` 中去，然后把自己放入线程缓存中。
上下文会定时检查 `Global runqueue`。

`Golang` 是为并发而生的语言，Go语言是为数不多的在语言层面实现并发的语言；也正是Go语言的并发特性，吸引了全球无数的开发者。

Golang的CSP并发模型，是通过Goroutine和Channel来实现的。

Goroutine 是Go语言中并发的执行单位。有点抽象，其实就是和传统概念上的“线程”类似，可以理解为“线程”。

Channel 是 Go 语言中各个并发结构体之间的通信机制。通常 Channel, 是各个 Goroutine 之间的通信的 “管道”， 有点类似于 Linux 中的管道。

通信机制channel 也很方便，传数据用 `channel <- data`, 取数据用`<- channel`。

而且不管是传还是取，肯定阻塞，直到另外的 goroutine 传或者取为止。

因此 GPM 的简要概括即为： 事件循环，线程池，工作队列。

###