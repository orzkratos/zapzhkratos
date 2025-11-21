[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/orzkratos/zapzhkratos/release.yml?branch=main&label=BUILD)](https://github.com/orzkratos/zapzhkratos/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/orzkratos/zapzhkratos)](https://pkg.go.dev/github.com/orzkratos/zapzhkratos)
[![Coverage Status](https://img.shields.io/coveralls/github/orzkratos/zapzhkratos/main.svg)](https://coveralls.io/github/orzkratos/zapzhkratos?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/orzkratos/zapzhkratos.svg)](https://github.com/orzkratos/zapzhkratos/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/orzkratos/zapzhkratos)](https://goreportcard.com/report/github.com/orzkratos/zapzhkratos)

# zapzhkratos

使用中文函数名的 Zap 日志与 Kratos 微服务框架集成。

---

<!-- TEMPLATE (ZH) BEGIN: LANGUAGE NAVIGATION -->
## 英文文档

[ENGLISH README](README.md)
<!-- TEMPLATE (ZH) END: LANGUAGE NAVIGATION -->

## 核心特性

- 🎯 中文函数名 - 使用直观的中文函数名 (`Get奎沱日志`、`New奎沱主簿` 等)
- 📊 结构化日志 - 使用 Uber Zap 快速的结构化日志能力
- ⚡ 高性能 - 受益于 Zap 的零内存分配设计
- 🔄 Kratos 兼容 - 与 Kratos 框架无缝集成
- 🌍 模块追踪 - 自动添加模块信息到日志
- 📋 灵活配置 - 自定义模块字段命名和选项

## 安装

```bash
go get github.com/orzkratos/zapzhkratos
```

## 快速开始

```go
package main

import (
    "github.com/orzkratos/zapzhkratos"
    "github.com/yyle88/zaplog"
)

func main() {
    // 使用中文函数创建实例
    v匝普日志 := zapzhkratos.New匝普日志(
        zaplog.LOGGER,
        zapzhkratos.New日志配置(),
    )

    // 获取 Kratos log.Logger
    logger := v匝普日志.Get奎沱日志("my-app")

    // 获取 Kratos log.Helper（主簿）
    slog := v匝普日志.Get奎沱主簿("my-module")
    slog.Info("app started")

    // 获取带模块信息的 Zap
    zapLog := v匝普日志.Sub模块匝普()
    zapLog.LOG.Info("with module context")
}
```

## 完整示例

查看 [zapzhkratos-demos](https://github.com/orzkratos/zapzhkratos-demos) 了解在实际 Kratos 项目中的完整集成：

- **[demo1kratos](https://github.com/orzkratos/zapzhkratos-demos/tree/main/demo1kratos)** - HTTP 和 gRPC 基础集成
- **[demo2kratos](https://github.com/orzkratos/zapzhkratos-demos/tree/main/demo2kratos)** - Wire 依赖注入高级用法

演示项目展示：
- main.go 和 Wire 配置中的集成方式
- biz/service/data 各层的使用方法
- HTTP 和 gRPC 服务中的 zapzhkratos 配置
- Log Helper 在业务逻辑中的使用

## API 参考

### 类型

**T日志配置** - 配置选项（包装 `zapkratos.Options`）

```go
// 创建默认配置
cfg := zapzhkratos.New日志配置()

// 自定义模块字段名
cfg.With模块位置键名("module")
```

**T匝普日志** - 主结构（包装 `zapkratos.ZapKratos`）

```go
// 创建实例
zk := zapzhkratos.New匝普日志(zaplog.LOGGER, cfg)
```

### 方法

#### Get奎沱日志 / New奎沱日志
创建带说明的 `log.Logger`：

```go
logger := zk.Get奎沱日志("my-service")
// 与 New奎沱日志 相同
logger = zk.New奎沱日志("my-service")
```

#### Get奎沱主簿 / New奎沱主簿
创建 `log.Helper`（"主簿" 意为主管文书）：

```go
slog := zk.Get奎沱主簿("module-name")
slog.Info("message")
slog.Infow("key", "value")
```

#### Get奎沱秘书 / New奎沱秘书
创建 `log.Helper`（"秘书" 为备选名称）：

```go
slog := zk.Get奎沱秘书("module-name")
```

#### Get基本匝普
返回底层 Zap 实例：

```go
zap := zk.Get基本匝普()
zap.LOG.Info("message")
```

#### Sub模块匝普
创建带模块上下文的子 Zap：

```go
zapLog := zk.Sub模块匝普()
// 自动添加文件名模块字段
```

## 命名翻译

- **zapzhkratos**:
  - `zap` - Zap 日志
  - `zh` - 中文
  - `kratos` - Kratos 框架

- **日志** - Logger 接口
- **主簿** - log.Helper（主管文书）
- **秘书** - log.Helper（秘书，备选）
- **模块** - Module
- **配置** - Configuration

## 中文编程

本包使用中文函数名，让中文开发者使用更直观，同时保持与 `zapkratos` 相同的核心功能。

命名"主簿"（主管文书）的原因是 `log.Helper` 很难翻译 - "侍者"、"史官"、"助理"等备选方案都不太合适。"秘书"也作为备选选项提供。

**注意**: 这是一个使用中文命名的测试包。如果不喜欢这种方式，请使用 [zapkratos](https://github.com/orzkratos/zapkratos)。

## 中文命名的优势

中文函数名带来语义精确性和自文档化能力。汉字确保对齐和明确含义。像"秘书"这样的词用更少空间达到同等表现力。

## 依赖项

- `github.com/go-kratos/kratos/v2` - Kratos 微服务框架
- `github.com/orzkratos/zapkratos` - 核心 zapkratos 包
- `github.com/yyle88/zaplog` - Zap 管理包

## 关联项目

**框架：**
- [Kratos](https://github.com/go-kratos/kratos) - Go 微服务框架
- [Zap](https://github.com/uber-go/zap) - Uber 的结构化日志

**zapkratos 生态：**
- [zapkratos](https://github.com/orzkratos/zapkratos) - 核心集成包
- [zapkratos-demos](https://github.com/orzkratos/zapkratos-demos) - 演示项目
  - [demo1kratos](https://github.com/orzkratos/zapkratos-demos/tree/main/demo1kratos) - 基础集成
  - [demo2kratos](https://github.com/orzkratos/zapkratos-demos/tree/main/demo2kratos) - 高级用法

**zapzhkratos 生态（中文版）：**
- [zapzhkratos](https://github.com/orzkratos/zapzhkratos) - 本项目
- [zapzhkratos-demos](https://github.com/orzkratos/zapzhkratos-demos) - 中文版演示项目
  - [demo1kratos](https://github.com/orzkratos/zapzhkratos-demos/tree/main/demo1kratos) - 基础集成
  - [demo2kratos](https://github.com/orzkratos/zapzhkratos-demos/tree/main/demo2kratos) - 高级用法

<!-- TEMPLATE (ZH) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-20 04:26:32.402216 +0000 UTC -->

## 📄 许可证类型

MIT 许可证 - 详见 [LICENSE](LICENSE)。

---

## 💬 联系与反馈

非常欢迎贡献代码！报告 BUG、建议功能、贡献代码：

- 🐛 **问题报告？** 在 GitHub 上提交问题并附上重现步骤
- 💡 **新颖思路？** 创建 issue 讨论
- 📖 **文档疑惑？** 报告问题，帮助我们改进文档
- 🚀 **需要功能？** 分享使用场景，帮助理解需求
- ⚡ **性能瓶颈？** 报告慢操作，帮助我们优化性能
- 🔧 **配置困扰？** 询问复杂设置的相关问题
- 📢 **关注进展？** 关注仓库以获取新版本和功能
- 🌟 **成功案例？** 分享这个包如何改善工作流程
- 💬 **反馈意见？** 欢迎提出建议和意见

---

## 🔧 代码贡献

新代码贡献，请遵循此流程：

1. **Fork**：在 GitHub 上 Fork 仓库（使用网页界面）
2. **克隆**：克隆 Fork 的项目（`git clone https://github.com/yourname/repo-name.git`）
3. **导航**：进入克隆的项目（`cd repo-name`）
4. **分支**：创建功能分支（`git checkout -b feature/xxx`）
5. **编码**：实现您的更改并编写全面的测试
6. **测试**：（Golang 项目）确保测试通过（`go test ./...`）并遵循 Go 代码风格约定
7. **文档**：为面向用户的更改更新文档，并使用有意义的提交消息
8. **暂存**：暂存更改（`git add .`）
9. **提交**：提交更改（`git commit -m "Add feature xxx"`）确保向后兼容的代码
10. **推送**：推送到分支（`git push origin feature/xxx`）
11. **PR**：在 GitHub 上打开 Merge Request（在 GitHub 网页上）并提供详细描述

请确保测试通过并包含相关的文档更新。

---

## 🌟 项目支持

非常欢迎通过提交 Merge Request 和报告问题来为此项目做出贡献。

**项目支持：**

- ⭐ **给予星标**如果项目对您有帮助
- 🤝 **分享项目**给团队成员和（golang）编程朋友
- 📝 **撰写博客**关于开发工具和工作流程 - 我们提供写作支持
- 🌟 **加入生态** - 致力于支持开源和（golang）开发场景

**祝你用这个包编程愉快！** 🎉🎉🎉

<!-- TEMPLATE (ZH) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/orzkratos/zapzhkratos.svg?variant=adaptive)](https://starchart.cc/orzkratos/zapzhkratos)
