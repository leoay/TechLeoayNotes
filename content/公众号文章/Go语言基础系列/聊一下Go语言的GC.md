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
    - 为每个对象维护一个引用计数，当引用该对象的对象销毁时，引用计数 -1，当对象引用计数为 0 时回收该对象。

* 标记-清除 Golang（三色标记清除法）
    - 从根变量开始遍历所有引用的对象，标记引用的对象，没有被标记的进行回收。

* 分代收集 JAVA
    - 按照对象生命周期长短划分不同的代空间，生命周期长的放入老年代，短的放入新生代，不同代有不同的回收算法和回收频率。

#### 2. 常用垃圾回收算法的优缺点

* 引用计数法
    优点：对象回收快，不会出现内存耗尽或达到某个阈值时才回收
    缺点：不能很好的处理循环引用，而实时维护引用计数也是有损耗的。

 * 标记-清除
    优点：解决了引用计数的缺点
    缺点：需要 STW，暂时停掉程序运行

 * 分代收集
    优点：回收性能好
    缺点：算法复杂