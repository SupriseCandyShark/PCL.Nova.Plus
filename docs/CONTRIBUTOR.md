# 贡献者名单

## 原作者

1. [MoYuan-CN](https://github.com/MoYuan-CN)

## 原贡献者

1. [EncVar](https://github.com/EncVar)
2. [tangge233](https://github.com/tangge233)
3. [shenjack](https://github.com/shenjackyuanjie)

## 现作者

1. [Xphost008](https://github.com/xphost008)
2. [FireDragon0659](https://github.com/firedragon0659)

## 现贡献者

1. [Hill233](https://github.com/hill23333)
2. Xysgh: 提供了 Linux 环境的测试！

# 如何发起 Pull Request 请求？

1. 首先，请不要动 Privacy.go 里面的内容，你可以对应用程序做对应的修改，随后提交你的 Pull Request 即可！
2. 如果你需要调试，并且构建平台是 Windows，请在本地新建一个 `dev.bat` 文件，随后在里面使用 wails 的 `wails dev -ldflags "-X mmcll.LauncherName=PCL.Nova.Plus"` 进行赋值
3. 因为 .gitignore 里面已经有了 `build.bat` 和 `dev.bat` 内容，因此你可以很安全的构建。
4. 如果你是 macOS，也请新建一个 `build.sh`、`dev.sh` 进行构建
5. 具体 Golang 缺失的参数是：
   1. mmcll.LauncherName ==> 启动器名称（仅启动游戏时需要）
   2. mmcll.LauncherVersion ==> 启动器版本（界面上会显示）
   3. mmcll.LauncherUserAgent ==> 启动器请求UserAgent（请求网站默认附带的 UserAgent）
   4. launcher.XApiKey ==> CurseForge 的 APIKey（如果你没有，你只是无法获取 CurseForge 网站内容，但是你依旧可以正常使用启动器其他内容！）
   5. launcher.ClientId ==> Microsoft 的 ClientId（如果你没有，你只是无法登录 Microsoft 正版，但是你依旧可以正常使用启动器其他内容！）
   6. launcher.PrivacyKey ==> 构建密钥（原作者专属，你们可以无需赋值并在意）

- 大概有以上 6 个值需要赋值，因此