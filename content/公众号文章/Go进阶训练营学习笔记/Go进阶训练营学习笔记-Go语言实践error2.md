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
//http://golang.org/pkg/builtin/#error

type error interface {
    Error() string
}
```

在开发过程中我们常常使用errors.New() 函数返回一个error对象。

```go
//https://github.com/golang/go/blob/master/src/errors/errors.go
// New returns an error that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
func New(text string) error {
	return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
```

Go语言的基础库中，有大量自定义的 error。

```go
//https://github.com/golang/go/blob/master/src/bufio/bufio.go
var (
	ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
	ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
	ErrBufferFull        = errors.New("bufio: buffer full")
	ErrNegativeCount     = errors.New("bufio: negative count")
)
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

