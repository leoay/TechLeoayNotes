---
title: "MySQL主从同步"
date: 2022-01-03T14:28:31+08:00
draft: false
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

#### 什么是MySQL的主从同步？

当master(主)库的数据发生变化的时候，变化会实时地同步到slave(从)库。

#### 为什么要做主从同步？

1. 扩展数据库的负载能力
2. 容错，高可用(在做读写分离时，需要先实现主从同步)
3. 数据备份

#### Mysql的主从同步是怎么实现的？(原理)

图示：



MySQL 主从同步是通过复制实现的，复制是MySQL数据库提供的一种高可用高性能的解决方案，一般用来建立大型的应用。
总体来说，复制的工作原理分为以下3个步骤：
1. 主服务器(master) 把数据更改记录到二进制日志(bin log)中
2. 从服务器(slave)把主服务器的二进制复制到自己的中继日志(relay log)中。
3. 从服务器重做中继日志中的日志，把更改应用到自己的数据库中，以达到数据的最终一致性。

复制的工作原理并不复杂，其实就是一个完全备份加上二进制日志备份的还原。不同的是这个二进制日志的还原操作基本上是实时在进行中。
这里特别需要注意的是，复制不是完全实时地进行同步，而是异步实时。

#### MySQL主从同步的最佳实践

1. 准备两个MySQL数据库，一个作为主数据库，一个作为从数据库
2. 配置主数据库
3. 配置从数据库
4. 查看主数据库状态
5. 查看从数据库状态