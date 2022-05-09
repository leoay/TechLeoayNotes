---
title: "Golang 面试题整理及分析"
date: 2022-01-12T14:28:31+08:00
draft: false
hideToc: false
enableToc: true
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

 我们用下图表示P、M、G。


