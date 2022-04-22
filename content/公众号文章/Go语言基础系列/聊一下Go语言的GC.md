---
title: "聊一下Go语言的GC"
date: 2022-04-22T10:33:41+09:00
description: "聊一下Go语言的GC"
draft: false
hideToc: false
enableToc: true
enableTocContent: false
author: leoay
authorEmoji: 🎅
pinned: false
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

#### 1. 常见的垃圾回收算法
* 引用计数 PHP、Swift、Python

* 标记-清除 Golang（三色标记清除法）

* 分代收集 JAVA

#### 2. 常用垃圾回收算法的优缺点

* 引用计数法
    优点：对象回收快，不会出现内存耗尽或达到某个阈值时才回收
    缺点：不能很好的处理循环引用，而实时维护引用计数也是有损耗的。

 * 标记-清除
    优点：解决了引用计数的缺点
    缺点：需要 STW，暂时停掉程序运行


 * 分代收集
