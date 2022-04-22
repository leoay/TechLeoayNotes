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

1. 常见的垃圾回收算法
【1】引用计数 PHP、Swift、Python
【2】标记-清除 Golang（三色标记清除法）
【3】 分代收集 JAVA