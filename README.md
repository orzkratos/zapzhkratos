[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/orzkratos/zapzhkratos/release.yml?branch=main&label=BUILD)](https://github.com/orzkratos/zapzhkratos/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/orzkratos/zapzhkratos)](https://pkg.go.dev/github.com/orzkratos/zapzhkratos)
[![Coverage Status](https://img.shields.io/coveralls/github/orzkratos/zapzhkratos/main.svg)](https://coveralls.io/github/orzkratos/zapzhkratos?branch=main)
[![Supported Go Versions](https://img.shields.io/badge/Go-1.25+-lightgrey.svg)](https://go.dev/)
[![GitHub Release](https://img.shields.io/github/release/orzkratos/zapzhkratos.svg)](https://github.com/orzkratos/zapzhkratos/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/orzkratos/zapzhkratos)](https://goreportcard.com/report/github.com/orzkratos/zapzhkratos)

# zapzhkratos

Zap logging integration with Kratos microservice framework using Chinese function names.

---

<!-- TEMPLATE (EN) BEGIN: LANGUAGE NAVIGATION -->
## CHINESE README

[中文说明](README.zh.md)
<!-- TEMPLATE (EN) END: LANGUAGE NAVIGATION -->

## Main Features

- 🎯 Chinese Function Names - Use intuitive Chinese names (`Get奎沱日志`, `New奎沱主簿`, etc.)
- 📊 Structured Logging - Use Uber Zap's fast structured logging
- ⚡ High Performance - Benefit from Zap's zero-allocation design
- 🔄 Kratos Compatible - Seamless integration with Kratos framework
- 🌍 Module Tracking - Auto add module info to logs
- 📋 Flexible Config - Custom module field naming and options

## Installation

```bash
go get github.com/orzkratos/zapzhkratos
```

## Quick Start

```go
package main

import (
    "github.com/orzkratos/zapzhkratos"
    "github.com/yyle88/zaplog"
)

func main() {
    // Create instance with Chinese functions
    v匝普日志 := zapzhkratos.New匝普日志(
        zaplog.LOGGER,
        zapzhkratos.New日志配置(),
    )

    // Get Kratos log.Logger
    logger := v匝普日志.Get奎沱日志("my-app")

    // Get Kratos log.Helper (主簿)
    slog := v匝普日志.Get奎沱主簿("my-module")
    slog.Info("app started")

    // Get module-aware Zap
    zapLog := v匝普日志.Sub模块匝普()
    zapLog.LOG.Info("with module context")
}
```

## Complete Examples

See [zapzhkratos-demos](https://github.com/orzkratos/zapzhkratos-demos) to view complete integration in working Kratos projects:

- **[demo1kratos](https://github.com/orzkratos/zapzhkratos-demos/tree/main/demo1kratos)** - Basic integration with HTTP and gRPC
- **[demo2kratos](https://github.com/orzkratos/zapzhkratos-demos/tree/main/demo2kratos)** - Advanced usage with Wire DI

The demos show:
- Integration in main.go and Wire setup
- Usage across biz/service/data tiers
- HTTP and gRPC setup with zapzhkratos
- Log Helper usage in business logic

## API Reference

### Types

**T日志配置** - Configuration options (wraps `zapkratos.Options`)

```go
// Create config with default settings
cfg := zapzhkratos.New日志配置()

// Customize module field name
cfg.With模块位置键名("module")
```

**T匝普日志** - Main struct (wraps `zapkratos.ZapKratos`)

```go
// Create instance
zk := zapzhkratos.New匝普日志(zaplog.LOGGER, cfg)
```

### Methods

#### Get奎沱日志 / New奎沱日志
Creates `log.Logger` with given caption:

```go
logger := zk.Get奎沱日志("my-service")
// Same as New奎沱日志
logger = zk.New奎沱日志("my-service")
```

#### Get奎沱主簿 / New奎沱主簿
Creates `log.Helper` ("主簿" means chief clerk):

```go
slog := zk.Get奎沱主簿("module-name")
slog.Info("message")
slog.Infow("key", "value")
```

#### Get奎沱秘书 / New奎沱秘书
Creates `log.Helper` ("秘书" means secretary, alternative name):

```go
slog := zk.Get奎沱秘书("module-name")
```

#### Get基本匝普
Returns underlying Zap instance:

```go
zap := zk.Get基本匝普()
zap.LOG.Info("message")
```

#### Sub模块匝普
Creates child Zap with module context:

```go
zapLog := zk.Sub模块匝普()
// Auto adds module field with filename
```

## Name Translations

- **zapzhkratos**:
  - `zap` - Zap logging
  - `zh` - Chinese (中文)
  - `kratos` - Kratos framework

- **日志** - Logger interface
- **主簿** - log.Helper (chief clerk)
- **秘书** - log.Helper (secretary, alternative)
- **模块** - Module
- **配置** - Configuration

## Chinese Programming

This package uses Chinese function names, making it intuitive to Chinese developers while maintaining the same core functions as `zapkratos`.

The naming "主簿" (chief clerk) was chosen as `log.Helper` is difficult to translate - "侍者" (attendant), "史官" (historian), "助理" (assistant) didn't fit the context. "秘书" is also provided as an alternative.

**Note**: This is a test package with Chinese names. If you don't like this approach, use [zapkratos](https://github.com/orzkratos/zapkratos) instead.

## Benefits of Chinese Names

Chinese function names bring semantic precision and self-documentation abilities. The characters ensure alignment with unambiguous meanings. Words like "秘书" achieve the same expressiveness with less space.

## Dependencies

- `github.com/go-kratos/kratos/v2` - Kratos microservice framework
- `github.com/orzkratos/zapkratos` - Core zapkratos package
- `github.com/yyle88/zaplog` - Zap management package

## Related Projects

**Frameworks:**
- [Kratos](https://github.com/go-kratos/kratos) - Go microservice framework
- [Zap](https://github.com/uber-go/zap) - Uber's structured logging

**zapkratos Ecosystem:**
- [zapkratos](https://github.com/orzkratos/zapkratos) - Core integration package
- [zapkratos-demos](https://github.com/orzkratos/zapkratos-demos) - Demo projects
  - [demo1kratos](https://github.com/orzkratos/zapkratos-demos/tree/main/demo1kratos) - Basic integration
  - [demo2kratos](https://github.com/orzkratos/zapkratos-demos/tree/main/demo2kratos) - Advanced usage

**zapzhkratos Ecosystem (Chinese):**
- [zapzhkratos](https://github.com/orzkratos/zapzhkratos) - This project
- [zapzhkratos-demos](https://github.com/orzkratos/zapzhkratos-demos) - Chinese version demos
  - [demo1kratos](https://github.com/orzkratos/zapzhkratos-demos/tree/main/demo1kratos) - 基础集成
  - [demo2kratos](https://github.com/orzkratos/zapzhkratos-demos/tree/main/demo2kratos) - 高级用法

<!-- TEMPLATE (EN) BEGIN: STANDARD PROJECT FOOTER -->
<!-- VERSION 2025-11-20 04:26:32.402216 +0000 UTC -->

## 📄 License

MIT License - see [LICENSE](LICENSE).

---

## 💬 Contact & Feedback

Contributions are welcome! Report bugs, suggest features, and contribute code:

- 🐛 **Mistake reports?** Open an issue on GitHub with reproduction steps
- 💡 **Fresh ideas?** Create an issue to discuss
- 📖 **Documentation confusing?** Report it so we can improve
- 🚀 **Need new features?** Share the use cases to help us understand requirements
- ⚡ **Performance issue?** Help us optimize through reporting slow operations
- 🔧 **Configuration problem?** Ask questions about complex setups
- 📢 **Follow project progress?** Watch the repo to get new releases and features
- 🌟 **Success stories?** Share how this package improved the workflow
- 💬 **Feedback?** We welcome suggestions and comments

---

## 🔧 Development

New code contributions, follow this process:

1. **Fork**: Fork the repo on GitHub (using the webpage UI).
2. **Clone**: Clone the forked project (`git clone https://github.com/yourname/repo-name.git`).
3. **Navigate**: Navigate to the cloned project (`cd repo-name`)
4. **Branch**: Create a feature branch (`git checkout -b feature/xxx`).
5. **Code**: Implement the changes with comprehensive tests
6. **Testing**: (Golang project) Ensure tests pass (`go test ./...`) and follow Go code style conventions
7. **Documentation**: Update documentation to support client-facing changes and use significant commit messages
8. **Stage**: Stage changes (`git add .`)
9. **Commit**: Commit changes (`git commit -m "Add feature xxx"`) ensuring backward compatible code
10. **Push**: Push to the branch (`git push origin feature/xxx`).
11. **PR**: Open a merge request on GitHub (on the GitHub webpage) with detailed description.

Please ensure tests pass and include relevant documentation updates.

---

## 🌟 Support

Welcome to contribute to this project via submitting merge requests and reporting issues.

**Project Support:**

- ⭐ **Give GitHub stars** if this project helps you
- 🤝 **Share with teammates** and (golang) programming friends
- 📝 **Write tech blogs** about development tools and workflows - we provide content writing support
- 🌟 **Join the ecosystem** - committed to supporting open source and the (golang) development scene

**Have Fun Coding with this package!** 🎉🎉🎉

<!-- TEMPLATE (EN) END: STANDARD PROJECT FOOTER -->

---

## GitHub Stars

[![Stargazers](https://starchart.cc/orzkratos/zapzhkratos.svg?variant=adaptive)](https://starchart.cc/orzkratos/zapzhkratos)
