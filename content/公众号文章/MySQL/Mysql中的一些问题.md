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

### 2. 一张自增表里面总共有 7 条数据，删除了最后 2 条数据，重启 mysql 数据库，又插入了一条数据，此时 id 是几？

这个问题需要考虑mysql存储引擎的类型，不同的存储引擎对于自增主键的存储情况不一样。

 一般Mysql的默认存储引擎是InnoDB, 它在存储数据时会把自增主键的最大ID记录到内存中，所以重启数据库挥着对表OPTIMIZE操作

1.InnoDB：因为InnoDB表只把自增主键的最大ID记录到内存中，所以重启数据库或者对表OPTIMIZE操作，都会使最大ID丢失。所以记录的ID是6

2.MylSAM：因为MylSAM表会把自增主键的最大ID记录到数据文件里面，重启MYSQL后，自增主键的最大ID也不会丢失。所以记录的ID是8

