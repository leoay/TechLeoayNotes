---
title: "常用快捷命令"
date: 2022-01-09T10:33:41+09:00
description: "常用快捷命令"
draft: false
hideToc: false
enableToc: true
enableTocContent: false
author: leoay
authorEmoji: 🎅
pinned: true
tags:
- shell
- bash
- terminal
series:
- terminal
categories:
- shell
- bash
- terminal
image: images/feature2/color-palette.png
---


## 进程管理
### 根据进程名杀死进程

##### MacOS:

```bash
ps -ef | grep pycharm | awk '{print $2}' | xargs kill -9
```
# ps -ef | grep pycharm | awk '{print $2}' | xargs 执行之后，会将所有含有pycharm都进程号列出来