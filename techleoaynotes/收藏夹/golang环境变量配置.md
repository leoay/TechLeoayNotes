---
title: "Golang环境变量配置"
date: 2022-05-18T10:33:41+09:00
description: "Golang环境变量配置"
draft: false
hideToc: false
enableToc: true
enableTocContent: false
author: leoay
authorEmoji: 🎅
pinned: false
tags:
- 
series:
- 
categories:
- 
image: /images/face/link.jpg
---

1. 基础环境变量
```bash
export GOROOT=/usr/local/go
export GOPATH=/root/go  #路径自定义，多GOPATH以 ： 分割
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```
2. 开启Mod


