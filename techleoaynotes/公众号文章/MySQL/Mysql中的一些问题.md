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


### 3. 如何获取当前数据库版本？
1.cmd mysql -V(V大写)
2.mysql命令行里：select version();

#### 4. 什么是事务？



### 4. 什么是ACID？

ACID 主要是针对数据库中的事务而言的。

原子性(Atomic)：事务中各项操作，要么全做要么全不做，任何一项操作的失败都会导致整个事务的失败；

一致性(Consistent)：事务结束后系统状态是一致的；

隔离性(Isolated)：并发执行的事务彼此无法看到对方的中间状态；

持久性(Durable)：事务完成后所做的改动都会被持久化，即使发生灾难性的失败。通过日志和同步备份可以在故障发生后重建数据。