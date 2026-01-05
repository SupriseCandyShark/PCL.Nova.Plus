# 如果你有任何的问题，请优先进入 Github 链接提出你的 issue！

# 如果是游戏崩溃问题，请提供你的 `版本 JSON`、`启动参数`、`平台（Windows、macOS、Linux）` 在 Github Issue 上！

# 0.0.18.2

### fix

1. fix(Download): 修复了下载界面的卡片未能正确在页面底部加上margin-bottom: 10px。（加个height: auto即可）
2. fix(Launch): 修复了在启动游戏时，Nova 未能正确的替换 Forge 老版本的键值：${primary_jar_name} 键。
3. fix(RightBottomButton): 调整了右下角的按钮位置，现在应该会显示在 messagebox 的遮罩层下方了！不会在弹出了！

### feat

1. feat(Classification): 下载界面新增愚人节版本分类！以及Nova特供版本（未完成）！
2. feat(FontSize): 稍微修复了一下字体大小（在总的css里面加上了* {font-size: 16px}，仅此而已。。）
3. feat(QualcommTips): 在任何 Windows arm64 版本的 Nova 上，现在默认会弹出是否是高通骁龙处理器的提示了！
    1. 除了高通骁龙处理器是 arm64 以外，还有国产麒麟芯片也是 arm64 架构，但是该架构在 HMCL 上未作提示。。
    2. 现在会弹出一个框框，确认是否安装了 OpenGL 兼容层提示！
4. feat(EULA): 在刚打开启动器时，会弹出一个框，提示最终用户协议！您需要点击 同意 才能正常使用启动器！

# 0.0.18.1

### fix

1. fix(OnlineText): 修复了【联机界面】的文本错误 ([#14](https://github.com/SupriseCandyShark/PCL.Nova.Plus/issues/14))
2. fix(TodayLucky): 修复了在该电脑并未打开过任何一次 `PCL` 时，可能会遇到写入注册表时出现错误
3. fix(About): 将 3TProject 换成了 SupriseCandyShark！
4. fix(Mod): 修复了 Mod 显示标签的问题！（简单来说就是把文字缩小了两个像素）([#6](https://github.com/SupriseCandyShark/PCL.Nova.Plus/issues/6))
5. fix(AccountAvatar): 修复了【账号】头像双层的问题！现在头像是纯方形的，并且现在应该可以正常显示账号头像是双层了！([#8](https://github.com/SupriseCandyShark/PCL.Nova.Plus/issues/8))

# 0.0.18

### feat

1. feat(???): 在 ??????? 有一个 ??????? ，点了之后会 ??????? 随后就 ???????
    - 这是一个彩蛋功能喵~放心吧！你们不会很轻易的就找到这个在哪的！
2. feat(TitleBar): 现在标题栏在 macOS 上已经能显示红绿灯啦~

### fix

1. fix(Launch): 修复了许许多多的启动问题，现在应该可以启动全版本 MC 了！！【经过 macOS 测试与 Linux 测试，已经全部可以启动了（真的）】
    - 目前启动游戏不支持 `Prism`、`MultiMC`、`XMCL` 等启动器下载的 MC，请使用 `PCL`、`HMCL`、`BakaXL` 等规范意义上的启动器下载 MC，随后启动！！
    - 如果再有启动不了的错误，麻烦请你把游戏 `crash-report`、`启动器内导出启动参数`、`游戏版本 JSON` 一并打包成 `zip` 并发送到作者的 Github issue！
    - 如果是拼接启动参数中出现的报错，可以尝试直接给我 `游戏版本 JSON`、`启动器错误截图` 二者打包成 `zip` 并发送到作者的 Github issue！
2. fix(macOS): 修复了 macOS 上的配置文件写出，现在应该会默认写出到 ~/Library/Application Support/PCL.Nova下了。

### refactor:

1. refactor(ThirdParty): 已正式支持第三方登录啦！同时也支持获取第三方登录的头像啦~~已经可以正常登录 MC 并且开玩了喵！
2. refactor(Downloader): 初步实现了一下**下载单个文件**的 Golang 逻辑，现在应该已经能下载单个文件了，已支持切块下载！
    - 目前该功能仅用于下载 `Authlib-Injector` 功能。还未能成功获取到 `CurseForge`、`Modrinth` 等网站上的模组获取（也许 0.0.19 可以呢喵）
3. refactor(CoreSettings): 核心设置现在已经可以被点击了，初步新增了【导出启动参数】！在导出启动参数时会默认为 Access Token 打码。当然你也可以选择不打，但是这样会有很大的安全隐患。。
    - 目前核心设置仅有以下几个功能：
        1. 显示版本名称
        2. 导出启动参数
4. refactor(GameCrash): 初步实现了一下报错处理！虽然没有完善的处理，但是你可以在游戏发生报错时，打包启动器作者所需要的一些事情并发送给作者！放心，作者绝对不会将你的任何隐私信息打包的！
5. refactor(Log): 初步写了一点启动器日志服务，但是日志目前还有很多地方没有考虑到。。为前端代码和后端代码都生成了一个函数，很方便的可以写出日志！你也可以在【设置 -> 启动器】里设置是否写出日志。
6. refactor(GameLog): 游戏日志终于回归啦！现在开始！游戏日志会在你启动了 MC 之后，在窗口的右下角会多一个按钮，鼠标悬停可以看到是【游戏日志】！已经正式完善并发布啦！！
7. refactor(MicrosoftLogin): 现在开始，微软登陆已经可以正常回显进度了，不会只显示登陆成功或失败了！
8. refactor(Proxy): 在启动器 -> 设置 -> 启动器里，新增了下载代理功能，现在应该可以正常使用代理了喵~（代理仅适用于下载任意文件，对于一般的网络请求将默认不使用代理）
9. 上述所有更新日志，同样会同步到 `MMCLL` 里的喵~各位不用担心哈！

### style:

1. style(CodeStyle): 简化了代码风格，增加了一点点后端代码的注释，新增了几个 only windows 的函数，处理了一下今日人品的获取问题（今日人品真 tmd 难嗷嗷嗷）！虽然最终还是没能与 PCL2 一致，但从源码来看，我真的是努力啦！！
2. style(TitleBar): 修正了标题栏，现在只有一个 Badge 了【 `Nova+` 】
3. style(Privacy): 修正了一下 Privacy.go 内容，现在应该会自动从 github actions 里面生了，下次下载源码不用再新建源文件了。【使用了 go build -ldflags 属性】
4. style(Version): 现在 Github Actions 获取文件版本号，将会直接从 主目录 -> VERSION 文件中读取，下次版本号只需要改动 VERSION 下的内容即可啦！

### chore:

1. refactor(Game): 为 NovaPlus 新增了一点点小游戏喵~扫雷、2048 强势回归！！与此同时还加入了【魔塔】供各位游玩嗷~
    - 魔塔游戏为 NovaPlus 作者 xphost 倾情制作！不仅如此还有完善的文档可供各位自定义！请看[NovaPlus 魔塔自定义](./MagicTower.md)文档！
    - 该游戏复刻于 [xphost 的魔塔游戏](https://github.com/xphost008/magic-tower)！作者强势重启项目！使用前端框架狠狠的搓了一顿~
2. refactor(BackgroundImage): 初步实现了一下加载 gif 图片的功能，现在应该没啥 bug

### build

1. build(AES): 新增了 AES 函数加密功能，相关密钥、偏移等存放在 privacy.go 里！
2. build(CI): 彻底重构了 CI/CD 文件，目前不依赖于 wails-app-action，是依赖自己的控制语句！

# 0.0.17

### refactor:

1. refactor(Microsoft): 新增了微软登录！现在可以使用微软登录啦~
    - 点击微软登录之后，启动器开始进行登录操作，在此期间会弹出一个无法关闭的信息框。请进入到网页鉴权，随后窗口会自动成功并且消失！

### feat:

1. feat(HomePage): 在主页里面新增了一个 MyValue 标签，具体如何使用请点击【生成主页模板】或者查阅[主页规范](HOMEPAGE_STANDARD)~
2. feat(Controls): 新增了 1 个新的控件：MyCustomBox，可以显示一个自定义的信息框，可以在页面里随时引用并自主采取是否关闭的策略！

### fix:

1. 修复了一个之前写坏的功能（账号保存功能，这个功能之前写坏了💦💦）

# 0.0.16.2

### feat:

1. feat(IsCheckLibraries): 新增了 是否校验 Libraries 的按钮！当打勾时，将会校验 Libraries 中的 sha1 是否匹配。如果匹配，则添加，反之则跳过（
    - 如果不勾选的话，则不会校验 libraries 的 sha1。。
    - 在某些不需要校验 sha1 的情况下（例如你手动为某个 MC Libraries 打了个 patcher 补丁时），这非常有用！
    - 默认勾选。。可以在设置里手动关闭。
2. feat(MyCheckButton): 更新了一个复选框组件！总算是更新了（不容易啊。。

# 0.0.16.1

### fix:

1. fix(ModSearch): 修复了在搜索模组时不识别空格导致的意外错误
2. fix(ModSearch): 修复了搜索无结果时，不会产生报错的 bug（一直停留在加载界面）
3. fix(JavaSearch): 修复了在类 Unix 系统上，无法找 Java 的问题（解决方案是，不设置 Java 的判断！）
4. fix(VersionPath): 由于在 0.0.15.2-0.0.16 的过程中出现了毁灭性更改，因此启动游戏、选择核心部分写炸了（按钮无法跳转），现已修复！

# 0.0.16

### refactor:

1. refactor(HomePage): 制作了主页功能！各位可以在【设置 -> 主页】里面，输出一个自定义主页的教程！你也可以联网更新自己的自定义主页网址~
    - 自定义主页的教程与规范在[这里](HOMEPAGE_STANDARD)
    - 目前还没有支持 HTML 主页（呜呜呜，也许下个版本支持了呢？）
    - HTML 主页目前我还在纠结之中，是否应该支持，如果支持的话，会有什么安全隐患。。接下来就看用户反馈了吧~
2. refactor(Mod): 新增 Modrinth 的 Mod 搜索界面喵，现在已经可以正常搜索了喵！~

### exception:

1. 在 Config.go 文件里新增了 ExceptionHandler 的类型，用于在反馈给前端时提交一些重要信息~

### style:

1. style(MessageBox): 调整了信息框的 `rotate`，现在应该比 `PCL2` 更可爱了~
2. style(DontClick): 为【千万别点】新设计了两种，分别是单独将 `rotateX` 和 `rotateY` 转换成 `180deg`，仅此而已（
3. style(MyNormalSpan): 将 `MyNormalLabel` 修改成了 `MyNormalSpan`，并且将内部元素也改成了 `span` 标签，现在应该更加合理了！
4. style(Controls): 将几乎所有的控件的点击事件，全部设计成了 dispatch！现在应该可以在使用该控件位置的地方使用【on:事件名】来调用了！

### fix:

1. fix(afdian): 修复了爱发电图标失效的 bug
2. fix(IPv6): 彻底修复了 IPv6 检测~
3. fix(MCVersion): MC 文件夹管理可能会导致一些问题
4. fix(AccountSelect): 启动游戏时可能出现的 Cannot convert account index to int 报错~

# 0.0.15.2

### fix:

1. fix(cmd): 稍微优化了一下执行终端的逻辑，现在开始，所有执行终端的输出，都将会静默执行，不会再显示终端的窗口了。（仅 Windows 的 cmd 的终端）
2. fix(Java): 修复了 Java 在前端展示信息时，误把路径写成版本号的 bug（
3. fix(IPv6): 稍微修复了一丢丢的 IPv6 获取逻辑（现在**应该**不会再导致窗口卡死了）

# 0.0.15.1

### fix:

1. fix(Java): 修复了在选择 Java 的时候，在没有选择任何 Java，第一次进入设置界面并生成 JavaJson.json 配置文件时，会弹出一个本不应该出现的报错（

### style:

1. style(messagebox): 略微调整了一下信息框的 padding

# 0.0.15

### refactor:

1. refactor(Launch): 已经可以启动游戏了！
    1. 目前，版本文件夹的选择是在【版本选择】里面，包括版本也是。点击之后的行为也与 PCL 一致。但是顶部菜单栏不会变（
    2. 版本独立设置目前暂时还没有，请见谅（
    3. 版本全局设置在【设置 -> 启动】里，这里也同样可以设置 Java
    4. Java 管理目前已被重构，重构得非常好看（确信）
2. refactor(Online): 现在可以在启动器中使用【IPv6】进行联机，
    1. 启动器支持以下 IPv6 功能：
        1. 检测本机 IPv6
        2. 尝试 ping 一次 IPv6
        3. 列举出本局域网 IPv6 是**临时**还是**永久**
        4. 输入游戏内端口，点击按钮复制 IPv6 地址与端口一同发送给朋友进行联机~
    2. 本启动器承诺绝对不会搜集您电脑本身的 IPv6 地址，如有疑问，请查阅源码。
    3. 本启动器检测 IPv6 的行为将与 LLL 启动器一致：[查看 LLL 源码](https://github.com/xphost008/lllauncher/blob/master/LittleLimboLauncher/OnlineIPv6Method.pas)

### feat:

1. feat(Controls): 更新了四个控件：
    1. MyInputBox: 用来显示输入框的
    2. MyNormalHint: 用来显示一个左下角提示。
    3. MyProgressBar: 滑动条组件
    4. MyCardButton: 在帮助页面、下载界面中显示的默认【类 PCL2 按钮】（
2. feat(GameSetting): 在 设置 -> 游戏 中新增了各种游戏设置，目前你已经可以正常的去玩了！
3. feat(about): 在 更多 -> 关于与鸣谢 中，明确指出了引用 MMCLL 的部分内容！

### style:

1. style(Left): 略微调整了主界面的左侧栏的宽度。
2. style(AccountSelect): 账号选择部分不再使用 border 来代替选中，而是在前方加上 MyRadioButton！

# 0.0.14

### fix:

1. fix(OfflineLogin): 修复了一个 bug，无法自动通过账号名称绑定 UUID 到输入框的 bug（未显示，但是能登录成功。。）
2. fix(OfflineLogin): 修复了名字可以命名超过 3-16 个字符的 bug。。
3. fix(Card's Inner Margin): 稍微调整了一下所有有关于卡片内控件的外边距，现在已经变得更加贴合卡片了！
4. fix(About): 重新修改了帮助界面，现在帮助界面的按钮全部采用组件的形式了~但、但是缺点就是右边的 Github 主页变成 icon 形式的了。。
5. fix(MessageBox): 修复了当窗口移动到特别靠右的位置时，如果此时触发了信息框，很可能拉不回来。。修复方式为：现在信息框的背景是可以拖拽窗口的了！

### feat:

1. feat(Thirdparty): 在外置登录上，添加了【应用 Littleskin 服务器】的按钮，与此同时，如果 Littleskin 将来需要采用设备验证流方式登录，这边也会同样开始适配~
    1. 目前外置登录依旧不能用，因为此时 MC 都无法启动（所以登录的话即使做好了目前也暂时不会应用上去。。

### refactor:

1. refactor(Download): 初步重构下载部分，现在仅能获取 Minecraft 版本列表和 Forge、Fabric 的手动安装列表。但是无法下载，请见谅~
    1. 默认官方源，目前还没有做适配 BMCLAPI 的部分。下个版本把 Quilt、NeoForge、Optifine、LiteLoader 四个版本适配了之后，再做适配 BMCLAPI 的吧~

# 0.0.13【以下由小万泥、xphost 合作更新~】

## 将更新日志按照标准 github 语义来书写~

### refactor:

1. refactor(frontend): 将前端架构换成了 svelte，以进一步减小打包体积。

### feat:

1. feat(Background): 翻新了一次背景样式，现在开始，如果未加入背景图片，则背景样式的渐变已经与原版 PCL、CE 一致啦~
2. feat(left): 重新制作了一次左侧栏的拉取，现在左侧栏已经和原版 PCL 类似了！切换页面时将不会收回再展开了！将会按照原版一样自动撑开~
3. feat(Thirdparty): 新增了第三方登录的初步界面，为以后添加第三方登录提供了一个接口~
4. feat(Minesweeper, P2048): 暂时移除了扫雷、2048 等游戏，以后再慢慢加回来~
5. feat(TreasureBox): 新增百宝箱：内有今日人品和千万别点~

### style:

1. style(name): 名称已更换成 PCL.Nova.Plus【不再使用原有的 PCL2.Nova.App】

# 0.0.12

1. 修改了 MyLoading 的样式，现在成功和失败时的左下角都会有一个图标了。
2. 修正了一下 MySelectCard 组件，现在已经不用 maxHeight 了~ （[PR#27](https://github.com/PCL-Community/PCL2.Nova.App/pull/27)，[@AMagicPear](https://github.com/AMagicPear)）
3. 新组件：MyCheckButton！（同时修改原来的 MyCheckButton 名称为 MyNavButton）【与 MyNavButton 不同的是，该 button 选中时中间会打勾~】
    1. 👆感觉完全可以被 MyToggleSwitch 代替（但是又转念一想，好像也可以（目前 MyCheckButton 仅在主界面有展示，在别的地方目前没有实用空间。。
4. 初步设置透明度，并且现在已经支持设置背景图片啦！只需要将背景图片放到`{exe}\PCL.Nova\BackgroundImage`文件夹下，即可随机从中抽取一张背景图片展示！！
5. 现在开始，所有**组件**均有透明度 127 了！不用担心遮挡住背景图片了！
6. 稍微调亮了一点加载框成功时的颜色（
7. 在关于与鸣谢部分新增了 PCL Community!
8. 将鸣谢改成特别鸣谢~

# 0.0.11

1. 翻新账号部分！现在终于有 CE 那味了~~（虽然现在还暂时无法登录微软，不过可以登录离线登录啦！！）
2. 调整了一下信息框的背景颜色，现在看起来应该会更柔和了一点吧~
3. 稍微调整了一下加载框的暗色模式颜色。现在中间的镐子动画应该更加明显了。
4. 为 NormalButton 新增了 disabled 属性！现在可以正常显示禁用的样式啦~（不过正常的话需要手动判断此时是否处在禁用状态。。）
5. 账号部分目前可以正常保存到外部文件了！【别轻易发给别人（】
6. 为 Nova 添加了 exe 的图标啦！

# 0.0.10

1. 更了个新的控件：MyRadioButton！
2. 现在标题栏在从亮色切换到暗色时，亮度会稍微变暗一点。
3. 现在，暗色模式会保存到外部文件了！第一次启动时，会生成`{exe}\PCL.Nova\config\PCL2.Nova.ini`文件，可以保存一些配置信息。
4. 新增关于与鸣谢啦~【添加了原作者、现作者的鸣谢人员等】

# 0.0.9

1. 修复了扫雷无论点不点击扣分模式，都会扣分的 bug
2. 将外置登录提上日程！！

# 0.0.8

1. MyToggleSwitch 控件应用暗色模式！
2. 扫雷游戏新增了【扣分模式】，此时扫雷如果踩中雷不会扣分，而是减 50 分。
3. 2048 游戏新增【作弊生成】按钮，按下即可生成 256、512、1024 三种随机数字在场上。ps：该操作不扣分也不加分，不减步数也不加步数。

# 0.0.7

1. 在主界面的 MyPrograssBar 底下的文字新增了【进度显示】
2. 更新了 MyPrograssBar 的深色模式！
3. 更新了 2048 小游戏！可以去玩了！

# 0.0.6

1. 更新了 MyPrograssBar 控件！
2. 在主界面新增了【测试进度条】框，现在可以点击【开始】，然后进行测试加载了！

# 0.0.5

1. 更换架构至`wails`，由于相较之前并没有什么新增内容，因此暂不发布二进制文件。。

# 0.0.4

1. 更新了 MyToggleSwitch（自定义拨动开关）【应用至暗色模式的调整】
2. 更新了扫雷游戏！（更多 -> 小游戏 -> 扫雷），完美仿刻于：[xphost 的 Godot 扫雷游戏](https://github.com/xphost008/Godot-Minesweeper)
3. 目前该扫雷缺陷是：仅仿刻于 0.0.3 版本，并且没有种子功能，同时游戏开局可能会直接踩雷直接输（因为开局就已经确定了所有雷的位置。。

# 0.0.3

1. 首次公开发布，用时两周
2. 目前更新自定义控件：MyDialog（自定义信息框）、MyLoading（自定义加载框）、MySelectCard（自定义下拉卡片）等一系列自定义组件
3. 更新了黑暗模式
