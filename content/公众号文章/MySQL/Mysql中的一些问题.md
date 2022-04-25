---
title: "Mysql中的一些问题"
date: 2022-04-25T17:10:47+08:00
description: "Mysql中的一些问题"
hideToc: false
enableToc: true
enableTocContent: false
author: leoay
authorEmoji: 🎅
pinned: false
description: "MySQL MySQL主从同步"
tags:
- 主从同步
- MySQL
series:

categories:
- MySQL
image: /images/face/MYSQL_Color_3-405x405.png
---

### 1. 数据库的三范式是什么？

* 第一范式（1NF）：字段不可分；即原子性。
* 第二范式（2NF）：有主键，非主键字段依赖主键；即唯一性。
* 第三范式（3NF）：非主键字段不能相互依赖；即每列都与主键有直接关系，不存在传递依赖。

### 2. 