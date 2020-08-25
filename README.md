# rtmp-recorder

RTMP Live Video Recorder

[![Go Report Card](https://goreportcard.com/badge/github.com/mingslife/rtmp-recorder)](https://goreportcard.com/report/mingslife/rtmp-recorder)

## 简介

> 时间就像海绵里的水，挤一挤总是有的。——鲁迅

这是一个 RTMP 直播流录制的服务程序。做这个东西的目的在于有时候想看某场直播，但是直播时间跟自己的日程冲突了，或者说有些直播可能是在深夜时分，而作为一只社畜，考虑到第二天还要上班，是不太可能跟着一直嗨到通宵达旦的，而且大部分的直播都是没有重播的，所以做了一个录播的服务，可以在我们时间跟直播时间冲突的时间，先把直播录制下来，等我们空了再去看，岂不美哉？

目前该项目仅支持 RTMP 直播流的录制，未来可能还会支持其他协议的直播流录制，本质只是改变 worker 实现而已，有兴趣的同学也可以帮忙实现了给我发 PR。这里的 worker 类似于 Drone CI 的 Runner 实现（或者说这一块根本就是一模一样）。这个项目只是临时用了一天不到的时间写的，后边如果有机会加上别的协议支持的话，我会考虑换下项目名字的。

## 原理

该项目主要是基于 [RTMPDump](http://git.ffmpeg.org/rtmpdump) 和 [FLVMeta](https://github.com/noirotm/flvmeta) 去实现，然后用 [Gin](https://github.com/gin-gonic/gin) 包装了下调度控制的 Web 服务程序。我们先是通过 rtmpdump 把 RTMP 直播流存为 flv 格式的视频，为修正 rtmpdump 录制的过程中可能出错退出导致 flv 元信息出错（最主要的表现就是播放 flv 视频时无进度条），在结束录制之后，我们再用 flvmeta 去进行视频元信息检查和修正。

## 安装

我这里用的是 Docker Compose 进行部署的，所以没什么好说的，docker-compose 一键安装即可。

```bash
docker-compose build && docker-compose up -d
```

## 启动参数

* --port 服务端口（默认：5000）
* --token 令牌（默认：12345678）

## API

如果启动参数 token 不为空的话，那么在访问以下所有 API 的时候，都必须在请求 URL 的 query 部分加上 token 参数，如 ``/api/videos?token=12345678``。

### 获取视频录制任务列表

#### Request

- Method: **GET**
- URL: ``/api/videos``

#### Response

- Body:
```json
{
  "rows": [
    {
      "id": "BEEJyW6MenrxUCnsYNqPLH",
      "status": "Up",
      "duration": "6 minutes ago",
      "name": "",
      "url": ""
    }
  ],
  "total": 1
}
```

### 获取视频录制任务详情

#### Request

- Method: **GET**
- URL: ``/api/videos/:id``

#### Response

- Body:
```json
{
  "id": "BEEJyW6MenrxUCnsYNqPLH",
  "status": "",
  "duration": "",
  "name": "zzz.flv",
  "url": "rtmp://xxx/yyy/zzz"
}
```

### 创建视频录制任务

#### Request

- Method: **POST**
- URL: ``/api/videos``
- Headers: ``Content-Type: application/json;application=UTF-8``
- Body:
```json
{
  "name": "zzz.flv",
  "url": "rtmp://xxx/yyy/zzz"
}
```

#### Response

- Body:
```json
{
  "id": "",
  "status": "",
  "duration": "",
  "name": "zzz.flv",
  "url": "rtmp://xxx/yyy/zzz"
}
```

### 删除视频录制任务

#### Request

- Method: **DELETE**
- URL: ``/api/videos/:id``

#### Response

## 相关项目

* [RTMPDump](http://rtmpdump.mplayerhq.hu/)
* [FLVMeta](https://flvmeta.com/)
* [Gin](https://gin-gonic.com/)
