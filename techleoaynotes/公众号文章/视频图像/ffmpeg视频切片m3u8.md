---
title: "ffmpeg视频切片m3u8"
date: 2022-04-26T14:28:31+08:00
draft: false
hideToc: false
enableToc: true
enableTocContent: false
author: leoay
authorEmoji: 🎅
pinned: false
description: "ffmpeg视频切片m3u8"
tags:
- 
- 
series:

categories:
- 
image: /images/face/ffmpeg.jpg
---


https://www.cnblogs.com/YiFeng-Liu/p/14296802.html

fmpeg.exe -i C:\Users\Administrator\Videos\fuDaSchool\fuda.mp4 -c:v libx264 -c:a aac -strict -2 -f hls -hls_list_size 0 C:\Users\Administrator\Videos\fuDaSchool\fuda2.m3u8