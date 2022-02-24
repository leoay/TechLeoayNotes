---
title: "Go进阶训练营学习笔记☍Go语言实践error2"
date: 2022-01-14T14:28:31+08:00
draft: false
hideToc: false
enableToc: true
enableTocContent: false
author: leoay
authorEmoji: 🎅
pinned: true
tags:
- Error
- Go进阶训练营
series:
- Go进阶训练营
categories:
- Go进阶训练营
- Golang
- Error
image: 
---

![](https://pic4.zhimg.com/v2-683be6cff5288cd457d0241e4b760c6c)

#### Error Vs Exception

##### Error

Go语言中的error就是一个普通的接口，这一点我们可以直接从源码中看到

```go
http://golang.org/pkg/builtin/#error

type error interface {
    Error() string
}
```

在开发过程中我们常常使用errors.New() 函数返回一个error对象。

```go
http://golang.org/src/pkg/errors/errors.go

type errorString struct {
    s string
}


http://golang.org/src/pkg/errors/errors.go

func (e *errorString) Error() string {
    return e.s
}

```









#### Error Type

AAAAAAAAAAAA
AAA
AAAAAAAAAAAAAAAA
AAA
AAAAAAAAAAAAAAAA
AAA
AAAAAAAAAAAAAAAA
AAA
AAAAAAAAAAAAAAAA
AAA
AAAAAAAAAAAAAAAA
AAA
AAAAAAAAAAAAAAAA
AAA
AAAA





#### Handling Error

AAAAAAAAAAAA


#### Go 1.13 errors



#### Go 2 Error Inspection





#### 参考文档

