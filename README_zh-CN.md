<p align="center">
    <a href="https://openkf.github.io/website" target="_blank">
        <img src="assets/logo-gif/openkf-logo.gif" width="60%" height="30%"/>
    </a>
</p>
<h3 align="center" style="border-bottom: none">
    ⭐️  OpenKF（开放知识流）是一个在线智能客服系统。 ⭐️ <br>
<h3>

<p align=center>
<a href="https://goreportcard.com/report/github.com/OpenIMSDK/OpenKF"><img src="https://goreportcard.com/badge/github.com/OpenIMSDK/OpenKF" alt="A+"></a>
<a href="https://github.com/OpenIMSDK/OpenKF/issues?q=is%3Aissue+is%3Aopen+sort%3Aupdated-desc+label%3A%22good+first+issue%22"><img src="https://img.shields.io/github/issues/OpenIMSDK/OpenKF/good%20first%20issue?logo=%22github%22" alt="good first"></a>
<a href="https://github.com/OpenIMSDK/OpenKF"><img src="https://img.shields.io/github/stars/OpenIMSDK/OpenKF.svg?style=flat&logo=github&colorB=deeppink&label=stars"></a>
<a href="https://join.slack.com/t/openimsdk/shared_invite/zt-1tmoj26uf-_FDy3dowVHBiGvLk9e5Xkg"><img src="https://img.shields.io/badge/Slack-100%2B-blueviolet?logo=slack&amp;logoColor=white"></a>
<a href="https://github.com/OpenIMSDK/OpenKF/blob/main/LICENSE"><img src="https://img.shields.io/badge/license-Apache--2.0-green"></a>
<a href="https://golang.org/"><img src="https://img.shields.io/badge/Language-Go-blue.svg"></a>
</p>

</p>

<p align="center">
    <a href="./README.md"><b>英文</b></a> •
    <a href="./README_zh-CN.md"><b>中文</b></a>
</p>

</p>
<br>


## 🧩 特性介绍

1. [OpenKF](https://github.com/OpenIMSDK/OpenKF) 是一个基于 [OpenIM](https://github.com/OpenIMSDK) 的开源客服系统。
2. 支持 LLM（本地知识库）客服。
3. 支持多渠道客服，并且易于与第三方系统集成。
4. 易于部署和二次开发。

## 📺 系统预览

![登录页面](./assets/images/login_page.png)

## 🛫 快速开始 

> **注意**：你可以快速开始使用 OpenKF。

### 📦 安装

```bash
git clone https://github.com/OpenIMSDK/OpenKF openkf && export openkf=$(pwd)/openkf && cd $openkf && make
```

### 🚀 运行

> **注意**： 我们需要先运行后端服务器

```
bashCopy code
make build
```

> 打开另一个终端并运行以下命令

```
bashCopy code# make dev
cd web
npm run dev
```

### 📖 贡献者快速入门

善用 Makefile，它可以确保你的项目的质量。

```
bashCopy codeUsage: make <TARGETS> ...

Targets:
  all                          Build all the necessary targets. 🏗️
  build                        Build binaries by default. 🛠️
  go.build                     Build the binary file of the specified platform. 👨‍💻
  build-multiarch              Build binaries for multiple platforms. 🌍
  tidy                         tidy go.mod 📦
  style                        Code style -> fmt,vet,lint 🎨
  fmt                          Run go fmt against code. ✨
  vet                          Run go vet against code. 🔍
  generate                     Run go generate against code and docs. ✅
  lint                         Run go lint against code. 🔎
  test                         Run unit test ✔️
  cover                        Run unit test with coverage. 🧪
  docker-build                 Build docker image with the manager. 🐳
  docker-push                  Push docker image with the manager. 🔝
  docker-buildx-push           Push docker image with the manager using buildx. 🚢
  copyright-verify             Validate boilerplate headers for assign files. 📄
  copyright-add                Add the boilerplate headers for all files. 📝
  swagger                      Generate swagger document. 📚
  serve-swagger                Serve swagger spec and docs. 🌐
  clean                        Clean all builds. 🧹
  help                         Show this help info. ℹ️
```

> **注意**： 我们强烈推荐你在提交代码之前运行 `make all`。🚀

```
bashCopy code
make all
```

## 🕋 架构图

![Architecture](./assets/images/architecture.png)

**MVC 架构设计：**

![MVC](./assets/images/mvc.png)

## 🤖 文件目录描述

目录标准化设计结构：

```
bashCopy code.
├── assets
│   └── images
├── build
├── deploy
├── docs
├── kf_plugins # 本地知识库和LLM
│   ├── chat
│   ├── config
│   ├── data
│   ├── logs
│   ├── model
│   └── utils
├── scripts
│   ├── githooks
│   └── LICENSE
├── server # OpenKF 后端
│   ├── cmd
│   ├── data
│   ├── docs
│   ├── examples
│   ├── internal
│   ├── logs
│   ├── pkg
│   ├── test
│   └── tools
└── web # OpenKF 前端
    ├── public
    ├── scripts
    └── src
```

## 🗓️ 社区会议

我们欢迎任何人参与我们的社区，我们提供礼物和奖励，我们欢迎你每个周四晚上加入我们。

我们的会议在 [OpenIM Slack](https://join.slack.com/t/openimsdk/shared_invite/zt-1tmoj26uf-_FDy3dowVHBiGvLk9e5Xkg) 🎯 `openkf` 管道中，然后你可以搜索 openkf 管道加入
