---
title: "殷浩详解DDD系列 第一讲 - Domain Primitive"
date: 2022-07-19T14:28:31+08:00
draft: false
hideToc: false
enableToc: true
description: "殷浩详解DDD系列 第一讲 - Domain Primitive"
enableTocContent: false
author: 殷浩
authorEmoji: 🎅
pinned: true
tags:
- 
series:
- 
categories:
- 
image: images/face/面试golang.png
---

> 版权说明：本文系转载，原文链接[殷浩详解DDD系列 第一讲 - Domain Primitive](https://developer.aliyun.com/article/713097?spm=a2c6h.12873639.article-detail.80.9c2527c9PpOitu)

## 写在最前面
对于一个架构师来说，在软件开发中如何降低系统复杂度是一个永恒的挑战，无论是94年GoF的Design Patterns，99年的Martin Fowler的Refactoring，02年的P of EAA，还是03年的Enterprise Integration Patterns，都是通过一系列的设计模式或范例来降低一些常见的复杂度。

但是问题在于，这些书的理念是通过技术手段解决技术问题，但并没有从根本上解决业务的问题。

所以03年Eric Evans的Domain Driven Design一书，以及后续Vaughn Vernon的Implementing DDD，Uncle Bob的Clean Architecture等书，真正的从业务的角度出发，为全世界绝大部分做纯业务的开发提供了一整套的架构思路。

由于DDD不是一套框架，而是一种架构思想，所以在代码层面缺乏了足够的约束，导致DDD在实际应用中上手门槛很高，甚至可以说绝大部分人都对DDD的理解有所偏差。

举个例子，Martin Fowler在他个人博客里描述的一个Anti-pattern，Anemic Domain Model （贫血域模型）在实际应用当中层出不穷，而一些仍然火热的ORM工具比如Hibernate，Entity Framework实际上助长了贫血模型的扩散。

同样的，传统的基于数据库技术以及MVC的四层应用架构（UI、Business、Data Access、Database），在一定程度上和DDD的一些概念混淆，导致绝大部分人在实际应用当中仅仅用到了DDD的建模的思想，而其对于整个架构体系的思想无法落地。

我第一次接触DDD应该是2012年，当时除了大型互联网公司，基本上商业应用都还处于单机的时代，服务化的架构还局限于单机+LB用MVC提供Rest接口供外部调用，或者用SOAP或WebServices做RPC调用，但其实更多局限于对外部依赖的协议。

让我关注到DDD思想的是一个叫Anti-Corruption Layer（防腐层）的概念，特别是其在解决外部依赖频繁变更的情况下，如何将核心业务逻辑和外部依赖隔离的机制。到了2014年，SOA开始大行其道，微服务的概念开始冒头，而如何将一个Monolith应用合理的拆分为多个微服务成为了各大论坛的热门话题，而DDD里面的Bounded Context（限界上下文）的思想为微服务拆分提供了一套合理的框架。

而在今天，在一个所有的东西都能被称之为“服务”的时代（XAAS），DDD的思想让我们能冷静下来，去思考到底哪些东西可以被服务化拆分，哪些逻辑需要聚合，才能带来最小的维护成本，而不是简单的去追求开发效率。

所以今天，我开始这个关于DDD的一系列文章，希望能继续在总结前人的基础上发扬光大DDD的思想，但是通过一套我认为合理的代码结构、框架和约束，来降低DDD的实践门槛，提升代码质量、可测试性、安全性、健壮性。

未来会覆盖的内容包括：

* 最佳架构实践：六边形应用架构 / Clean架构的核心思想和落地方案
* 持续发现和交付：Event Storming > Context Map > Design Heuristics > Modelling
* 降低架构腐败速度：通过Anti-Corruption Layer集成第三方库的模块化方案
* 标准组件的规范和边界：Entity, Aggregate, Repository, Domain Service, Application Service, Event, DTO Assembler等
* 基于Use Case重定义应用服务的边界
* 基于DDD的微服务化改造及颗粒度控制
* CQRS架构的改造和挑战
* 基于事件驱动的架构的挑战

今天先给大家带来一篇最基础，但极其有价值的Domain Primitive的概念.

## 第一讲 - Domain Primitive

就好像在学任何语言时首先需要了解的是基础数据类型一样，在全面了解DDD之前，首先给大家介绍一个最基础的概念: Domain Primitive（DP）。

Primitive的定义是：

> 不从任何其他事物发展而来
> 初级的形成或生长的早期阶段

就好像Integer、String是所有编程语言的Primitive一样，在DDD里，DP可以说是一切模型、方法、架构的基础，而就像Integer、String一样，DP又是无所不在的。所以，第一讲会对DP做一个全面的介绍和分析，但我们先不去讲概念，而是从案例入手，看看为什么DP是一个强大的概念。

