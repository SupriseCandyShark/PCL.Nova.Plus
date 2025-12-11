<div align="center"><img src="./frontend/src/assets/images/NovaPlusRaw.png" alt="Logo" width="386"/></div>

# PCL.Nova.Plus

下一代 PCL 启动器，官方版 PCL2 的完全重构，采用 Web 技术和超高速并发的 Go 全新构建，这，就是 PCL.Nova.Plus！

目前，本项目有以下须知需要你查阅：

- 目前，本项目由 [@FireDragon0659](https://github.com/FireDragon0659) 和 [@xphost008](https://github.com/xphost008) 合作进行开发和维护，你可以随时随地查看本项目的 [更新日志](docs/CHANGELOG.md)
- 第一个可用发行版计划于 2026 年 7 月之前推出！
- 欢迎随时随地收看来自老作者 [@MoYuan-CN](https://github.com/MoYuan-CN) 的新项目 [NovaCL](https://github.com/NEXORA-Studios/NovaCL)！！
- 如果你还不认识这位新作者的话，你随时可以参阅 [作者的博客](https://xphost008.github.io)。
- 小万泥有点害羞哈~他还没有博客可以放出来。。

## 下载

如果你是 Windows 系统，那么你可以直接从 [蓝奏云](https://wwdy.lanzoub.com/b0sx0e10h) 下载！密码：`eim6`（但是一万年没更新了（小声））<br>
如果是 macOS 或 Linux，请从 [Releases](https://github.com/SupriseCandyShark/PCL.Nova.Plus/releases) 下载！<br>

## PCL.Nova.Plus 的优势

1. 目前本仓库没有使用到**任何一个第三方库**，有且仅使用到 Svelte、TypeScript。没有使用 scss、router 这些繁琐的东西，将最原始且最快速的方式展现给各位！
2. 为什么使用 Svelte 而不是使用 Vue 呢？很简单！因为要减小代码体积！原因如下：
    1. 众所周知，使用 React、Vue、Angular 这类前端库写出来的项目非常大，有些使用 router、pinia 的项目甚至代码量达到了 10MB 甚至往上。而那些使用这些库的前端网页打开起来特别慢。
    2. 即使是某些后端应用，在使用上述几种库时，都几乎会使得生成的 exe 体积非常巨大。例如 **BakaXL4**（使用 Vue+TypeScript+Tauri）写的后端应用程序。
3. 总之，本软件将尽可能的压缩自身体积，在使用 UPX 进行压缩之后，也能达到非常小，小到不足 5MB 的应用程序。

## 支持的系统

- Windows 10+ amd64/arm64
- macOS 10.13+ arm64/x86_64
- Linux amd64/arm64

## 自定义主页

- 各位可以查阅一下自定义主页的 [规范与教程](./docs/HOMEPAGE_STANDARD.md)
