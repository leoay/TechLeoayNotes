---
title: "Go进阶训练营学习笔记☍Go语言实践error2"
date: 2022-01-02T14:28:31+08:00
draft: false
hideToc: false
enableToc: true
enableTocContent: false
author: leoay
authorEmoji: 🎅
pinned: false
description = "Go进阶训练营学习笔记 Go语言实践error"
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

### Error Vs Exception

#### Error

Go语言中的error本质上就是一个普通的接口，我们可以从源码`builtin.go`中看到

```go
//https://github.com/golang/go/blob/master/src/builtin/builtin.go

// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type error interface {
	Error() string
}
```

在开发过程中我们常常使用`errors.New()` 函数返回一个`error`对象。

```go
//https://github.com/golang/go/blob/master/src/errors/errors.go
// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}
```

Go语言的基础库中，有大量自定义的 error，我们以bufio举例。
```go
//https://github.com/golang/go/blob/master/src/bufio/bufio.go
var (
	ErrInvalidUnreadByte = errors.New("bufio: invalid use of UnreadByte")
	ErrInvalidUnreadRune = errors.New("bufio: invalid use of UnreadRune")
	ErrBufferFull        = errors.New("bufio: buffer full")
	ErrNegativeCount     = errors.New("bufio: negative count")
)
```

除此以外，go语言源码中还有其他的包中有类似的自定义error。

PS： 通过上面的自定义`error`，我们可以观察到官方在自定义`error`时，前面会加上包名 `bufio`，这一点我们也可以用于我们自己的项目开发中。

而errors.New()返回的是内部 `errorString` 对象的指针
```go
// New returns an error that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
func New(text string) error {
	return &errorString{text}
}
```

下面我们有一段代码进一步说明error包的使用：

```go
package main

import (
	"errors"
	"fmt"
)

type errorString struct {
	s string
}

func (e errorString) Error() string {
	return e.s
}

func NewError(text string) error {
	return errorString{text}
}

var ErrType = NewError("EOF")

func main() {
	if ErrType == NewError("EOF") {
		fmt.Println("Error:", ErrType)
	}
}
```

#### 各个语言中Error或者Exception的演进历史
* C
	单反回值，一般通过传递指针作为入参，返回值为 int 表示成功还是失败。
	```c
	ngx_int_t ngx_create_path(ngx_file_t *file, ngx_path_t *path);
	```

* C++
	引入exception, 但是无法知道被调用方会抛出异常

* Javad
	引入了checked exception, 方法的所有者必须申明，调用者必须处理。在启动时抛出大量的异常是司空见惯的事情，并在它们的调用堆栈中尽职地记录下来。Java异常不再是异常，而是变得司空见惯了。
	它们从良性到灾难性都有使用，异常的严重性由函数调用者来区分。
	```Java
	catch(e Exception) {
		//ignore
	}
	```

* Go
	而Go的处理异常逻辑是不引入 `exception`， 支持多参数返回，所以你很容易在函数签名中带上实现了`error interface`的对象，交由调用者来判定。

	如果一个函数返回了(value, error)， 你不能对这个`value`做任何假设，必须先判定error。唯一可以忽略error的情况是，你连value也不关心。

	Go中有panic机制，如果你认为和其他语言的exception一样，那你就错了。当我们抛出异常的时候，相当于你把exception扔给了调用者来处理。

	比如，你在C++中，把string转为int, 如果转换失败了




##### 感谢朋友们的阅读，欢迎大家关注我的微信公众号 leoay技术圈，扫描下方二维码查看详情。
![](/images/whoami/leoaytechgzh.jpg)