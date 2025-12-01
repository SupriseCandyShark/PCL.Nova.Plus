
# 自定义主页规范与支持

### 首先，`NovaPlus` 的自定义主页支持两种格式，第一种是纯 `XML` 形式，第二种是 `HTML` 形式。

- 首先来介绍一下 `XML` 形式的主页！
  - 在 `XML` 主页里，你仅能使用我自制的控件，包括但不限于 `MyCard`、`MyButton` 等。
  - 你无法使用除了我给你定义的控件以外的任何控件与任何事件。
  - 与 `PCL` 原版的 `XAML` 不同，这种形式的主页仅能支持我自己写的控件与事件，很大程度上限制了各位的想象力。。
  - 与此同时，使用 `XML` 自定义的主页无法影响到程序本身的 `DOM`！
  - 是的，无论你如何做，使用 `HTML` 主页还是使用 `XML` 主页，均无法修改程序本身的 `DOM`！
  - `XML` 形式的主页支持读取本地文件和联网更新~
  - 在主页界面点击上方【生成教学文件】按钮即可自动生成一个 `XML` 主页的教程！里面涵盖了我目前写完的所有 `XML` 主页控件！
- 接下来来介绍一下 `HTML` 形式的主页啦~
  - 说实话，这种形式的主页我是被迫支持的，因为 `HTML` 所能带来的视觉冲击实在是太强了，很容易就能做出精美的主页！但与此同时，也可能对用户的电脑造成危害。
  - 那么，`HTML` 主页我是如何支持的呢？很简单！就是 `iframe`！这样的话，请求不会采取一些极端的手段。
  - 使用 `iframe` 有很大一个好处就是：解决了跨域问题！由于 `iframe` 的跨域导致了你无法在你本身的 `HTML` 主页里修改我程序原本的 `DOM`，与此同时，如果你需要申请什么权限，也将是你自己主动申请的，与程序作者无关！
  - `HTML` 形式的主页只能联网更新，无法读取本地。你可以使用 `NodeJS` 的 `anywhere` 库随时打开一个本地的 `http` 服务器，随后制作主页~
  - `HTML` 形式的主页，`iframe` 的 `allow` 属性为 `none`，学过前端的你应该知道这是什么意思。

## 好了，说完了自定义主页！接下来我将说一下二者的区别与执行标准：

1. 我是如何判断你的联网更新主页是 `XML` 还是 `HTML` 的呢？很简单，我判断标头：当标头是 `<!DOCTYPE html>` 的时候，这就是一个 `HTML` 主页。紧接着我再判断请求结果是否是 `<?xml version="1.0" encoding="UTF-8"?>` 的，如果是，则默认判断其为 `XML` 主页！否则这个网址就是无效请求。
2. 你当然也可以在联网更新的网址里输入 `www.baidu.com`，这样请求出来的就会是百度的首页！
3. 这可能会导致在运行到 `HTML` 请求时执行两次，一次是通过网站请求判断标头，一次是 `iframe` 的执行。
4. 在请求任意一个网站资源的主页时，附带的 `UserAgent` 均为 `PCL.Nova.Plus/<version>`，各位可以判断请求的 `UserAgent` 是否是前面所述，从而判断出应该显示 `HTML` 还是主页代码。

## 说完了执行标准，接下来说说使用方法

1. 很明显，对于 `HTML` 主页，我这边必须采取自建服务器，你可以使用 `NodeJS` 的 `anywhere` 库随时打开一个本地的 `http` 服务器，随后制作主页~
2. 如果是 `XML` 文件的话，你可以将你的 `XML` 文件放入到 `{exe}/PCL.Nova/HomePage/<YourName>.xml` 路径下，`Nova` 会自动检测当前目录下的所有 `.xml` 文件！
3. 最后，如果你制作好了你的主页，你可以选择一个网站并放上去，之后在联网更新选项填入【精确到 `【YourXML】.xml`】的网站，你也可以使用上述的 `NodeJS` 搭一个本地服务器自行测试！

## 说完了使用方法，接下来说说看我的预设主页收录原则

1. 首先，收录的所有自定义主页，必须是 `XML` 主页，你无法让我收录任何一个 `HTML` 主页！
2. 如果你想收录任何一个 `XML` 主页的话，你可以自行打开一个 `discussion`，随后在里面写入你的主页预设网址，只要在国内大陆正常访问，并且无不良信息，即可收录！
3. 预设主页无法收录服务器宣传主页！！

## 接下来在下面初步写一点自定义主页的示例代码！

1. 提一句嘴：目前 NovaPlus 的主页后缀名是<font size="60">`.nxml`</font>！不要给我写其他的，否则识别不了！

```xml

<MyCard title="有标题，但是折叠" isExpand="True" canExpand="True">
    <MySpan>有标题，但是已折叠，所有的 isExpand 和 canExpand 属性均默认为
        False，因此如果你想折叠卡片，你需要显式声明两个属性！
    </MySpan>
</MyCard>
<MyCard title="有标题，但是已展开" isExpand="False" canExpand="True">
<MySpan>有标题，但是已展开</MySpan>
</MyCard>
<MyCard title="有标题，但是无法折叠1" isExpand="True" canExpand="False">
<MySpan>有标题，但是无法折叠1</MySpan>
</MyCard>
<MyCard title="有标题，但是无法折叠2" isExpand="False" canExpand="False">
<MySpan>有标题，但是无法折叠2，当 canExpand 设置为 False 的时候，将会自动忽略 isExpand 属性，永久性无法被折叠。</MySpan>
</MyCard>
<MyCard title="" isExpand="False" canExpand="False">
<MySpan>无标题，无法折叠，会忽略 isExpand 和 canExpand 两个属性</MySpan>
</MyCard>

<MyCard title="MyDiv 测试用例" isExpand="True" canExpand="True">
<MyDiv style="display: flex; flex-direction: column;">
    <MySpan>当设计了 父组件 MyDiv 的 display: flex; flex-direction: column 之后，下面所有的 MySpan 均为垂直排列。</MySpan>
    <MySpan>颜色为<MySpan style="color: white;">纯白色</MySpan>，但是却很大<MySpan style="font-size: 50px;">50px</MySpan>的文字~
    </MySpan>
    <MySpan>请记住，无论你使不使用 MySpan 进行布局书写，只要你没单独为 MySpan 设置 color 的 style
        属性，文字均会按照【暗色模式】的样式写字体模式。
    </MySpan>
</MyDiv>
</MyCard>
<MyCard title="MyButton 测试用例 与 类 PCL Grid 布局~" isExpand="True" canExpand="True">
<MyDiv style="width: 100%; height: 60px">
    <MyButton width="100%" height="60px">MyButton 是行内样式，默认不会换行。你可以使用空的 MyDiv
        作为父样式来换行（自定义主页不支持 br 标签）
    </MyButton>
</MyDiv>
<MySpan>这是高级样式，父组件设置成这种布局之后，里面所有按钮都会扩宽等距排列。你也可以设置 justify-content 属性为
    space-around，这样会更好看。请注意，你一定要使用 width: 100%，否则可能会导致扩宽不成功（
</MySpan>
<MyDiv style="width: 100%; display: flex; align-items: center; justify-content: space-between">
    <MyButton width="30%" height="50px">等距排列按钮1</MyButton>
    <MyButton width="30%" height="50px">等距排列按钮2</MyButton>
    <MyButton width="30%" height="50px">等距排列按钮3</MyButton>
</MyDiv>
</MyCard>
<MyCard title="MyDiv 高级样式之自制仿 PCL2 的 MyHint" isExpand="True" canExpand="True">
<MyDiv style="border-radius: 4px; font-size: 14px; padding: 8px 14px; border-left: 3px solid rgb(216, 41, 41); background-color: rgb(255, 221, 223); color: rgb(216, 41, 41);">
    这也是一个模仿 PCL2 的 MyHint 的控件，但是颜色为红色~
</MyDiv>
<MyDiv style="border-radius: 4px; font-size: 14px; padding: 8px 14px; border-left: 3px solid rgb(245, 122, 0); background-color: rgb(255, 235, 215); color: rgb(245, 122, 0);">
    请记住，在这里面的所有 MyDiv 控件布局，一定要加 margin-top，不然所有控件杂糅在一起就会很难看（默认的 margin-top 是
    5px，你可以往上加很多。
</MyDiv>
<MyDiv style="border-radius: 4px; font-size: 14px; padding: 8px 14px; border-left: 3px solid rgb(17, 114, 212); background-color: rgb(217, 236, 255); color: rgb(17, 114, 212);">
    具体所有的 CSS 值，你们完全可以去 菜鸟教程 官网随意查阅嗷~
</MyDiv>
<MyDiv style="position: relative; height: 50px;">
    <MyDiv style="position: absolute; right: 0; width: 100px; height: 50px; background-color: orange; color: black;">
        这是文字处在右边~
    </MyDiv>
</MyDiv>
</MyCard>
<MyCard title="MySpan 高级样式" isExpand="True" canExpand="True">
<MyDiv style="margin: 10px 0">
    下列是一堆 MySpan 的示例！MySpan 是行内样式，是可以嵌入到任何<MySpan style="font-size: 50px;">50px</MySpan>的文字里面的！
</MyDiv>
<MySpan style="user-select: text">这串文字是可以被选择的喵~</MySpan>
<MySpan style="font-weight: bold">加粗的文字喵~</MySpan>
<MySpan style="font-style: oblique">斜体的文字喵~</MySpan>
<MySpan style="text-decoration: line-through #f00">删除的文字喵，线也是红色的喵~</MySpan>
<MySpan style="color: #00ffff">这是水蓝色的文字喵~</MySpan>
甚至还有
<MySpan style="background-color: lightgray; border: 1px solid gray; border-radius: 3px; font-size: 16px; vertical-align: middle; padding: 1px 4px; color: black">
    code
</MySpan>
代码块捏~
</MyCard>
<MyCard title="MyButton 支持的所有样式" isExpand="True" canExpand="True">
<MyDiv>这里介绍一下 MyButton 的几种目前支持的样式</MyDiv>
<MyDiv>首先是正常的 Button，这个 Button 是一个行内元素，你可以在文字之间嵌入 Button，就像以下：</MyDiv>
<MyDiv>在文字之间<MyButton>嵌入一个按钮</MyButton>捏~
</MyDiv>
<MyDiv>按钮支持的属性有：width、height、color、event、type，其中 event 属性在下一个卡片会讲解。type 总共有 2 个值。如果什么也不填的话，就是
    default。
</MyDiv>
<MyDiv>以下按钮为 设置了 color 的按钮。color 仅针对边框，如果你需要同时连文字一同设置颜色，你需要在里面再设置一层 MySpan（或者
    MyDiv 都行。。）
</MyDiv>
<MyDiv>
    <MyButton type="default" width="200px" height="100px" color="#FF0000">
        <MyDiv style="font-size: 14px;">没有人规定不能在 MyButton 里面写标签吧~</MyDiv>
        <MySpan style="color: red; font-size: 14px;">严重警告的按钮（红色）</MySpan>
    </MyButton>
</MyDiv>
<MyDiv>你可能会问，为什么按钮不支持 style，原因其实很简单，因为按钮是我自制组件的按钮，而上述 MySpan 和 MyDiv 都是 HTML
    自带的！你一旦修改了 style，很可能会破坏按钮原本的 css 样式！
</MyDiv>
<MyDiv>我们还有一种非常特殊的按钮，那就是【文本按钮】！其中，如果将 MyButton 的 type
    属性填成了【label】，就说明这个按钮是文本按钮！文本按钮有以下属性：style、hov-style、event
</MyDiv>
<MyDiv style="display: flex; flex-direction: column; align-items: center">
    <MyButton type="label" hov-style="color: skyblue;">这是一个文本按钮</MyButton>
    <MyButton type="label" style="border-radius: 10px; padding: 4px 6px;"
              hov-style="border-radius: 10px; padding: 4px 6px; background-color: rgba(0, 0, 255, 0.2); color: #0000FF">
        类似 PCL 的 MyTextButton 那样的按钮！
    </MyButton>
</MyDiv>
<MyDiv>好了，在 type 不一致时，你们不允许设置一些别的那种属性，就好比如说当 type 为 default 或者不填时，无法拥有 style 属性，而
    type 为 label 时，就有 2 个新增的属性啦！
</MyDiv>
</MyCard>
<MyCard title="图标与图片讲解" isExpand="True" canExpand="True">
<MyDiv>接下来是图标详解！我们可以通过声明一个新的元素：MySvg 去使用我们的 SVG 矢量图标~具体参见以下示例：</MyDiv>
<MyDiv>首先，MySvg 的样式与其他的可能会稍微有点不同，他自动声明了 role="img" 以及
    xmlns="http://www.w3.org/2000/svg"，你需要做的只是声明它的 stroke、icon 等一系列的属性即可！
</MyDiv>
<MyDiv>哦对了！还有一点！MySvg 是一个单标签，你无需声明它的子标签！而在这个 Svg，你需要填入一个必填项，那就是 path！这个相当于矢量图里的
    path 路径！类似通过 path d="xxx" 去绘制矢量图！
</MyDiv>
<MyDiv style="display: flex; align-items: center; justify-content: center;">
    <MySvg viewbox="0 0 24 24" style="width: 32px; height: 32px; fill: red;"
           path="M12 17q.425 0 .713-.288T13 16v-4q0-.425-.288-.712T12 11t-.712.288T11 12v4q0 .425.288.713T12 17m0-8q.425 0 .713-.288T13 8t-.288-.712T12 7t-.712.288T11 8t.288.713T12 9m0 13q-2.075 0-3.9-.788t-3.175-2.137T2.788 15.9T2 12t.788-3.9t2.137-3.175T8.1 2.788T12 2t3.9.788t3.175 2.137T21.213 8.1T22 12t-.788 3.9t-2.137 3.175t-3.175 2.138T12 22m0-2q3.35 0 5.675-2.325T20 12t-2.325-5.675T12 4T6.325 6.325T4 12t2.325 5.675T12 20m0-8"/>
</MyDiv>
<MyDiv>上述，我画了一个 圆圈，中间是一个 i 字符的一个图标~</MyDiv>
<MyDiv>与此同时，我们不仅可以用 fill 去绘制，我们还可以用 stroke 去绘制图标！例如以下：</MyDiv>
<MyDiv style="display: flex; align-items: center; justify-content: center;">
    <MySvg viewbox="0 0 16 16"
           style="width: 32px; height: 32px; stroke: red; stroke-width: 2px; stroke-linecap: round; stroke-linejoin: round; fill: none;"
           path="M3 15.5h3a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 1 .5-.5h2a.5.5 0 0 1 .5.5v4a.5.5 0 0 0 .5.5h3a.5.5 0 0 0 .5-.5V8.5h1.453a.497.497 0 0 0 .404-.836L8.354.654a.5.5 0 0 0-.708 0L.643 7.664a.497.497 0 0 0 .404.836H2.5V15a.5.5 0 0 0 .5.5z"/>
</MyDiv>
<MyDiv>由于 MySvg 不支持 hov-style 和 active-style 事件，因此你暂时无法对其进行
    鼠标悬浮和点击上去的提示（也许将来会支持呢！
</MyDiv>
<MyDiv>好，接下来说一下图片吧！图片其实非常简单！只需要输入 MyImg 标签即可！里面有三个属性，分别是：src、alt、style。。是的！这是整个
    Nova 控件库里面第四个支持 style 的控件！请看以下示例！
</MyDiv>
<MyDiv>例如，加载一个苹果~</MyDiv>
<MyDiv style="display: flex; align-items: center; justify-content: center;">
    <MyImg src="https://q8.itc.cn/images01/20250816/0626eec82e4d4c10800940ba9c3954bb.jpeg" alt="苹果"
           style="width: 100%;"/>
</MyDiv>
<MyDiv>哦对了，无论是 MyImg 还是 MySvg 都是行内元素，各位可以随时的插入到 MyButton 里面！</MyDiv>
</MyCard>
<MyCard title="MyButton 高级使用案例" isExpand="True" canExpand="True">
<MyDiv>这一张卡片，我将教会各位如何使用 MyButton 写出那种类似 PCL 的 MyListItem！在我们学习了上述 MyImg 和 MySvg 以及
    MySpan 之后，各位应该对于盒子布局有所了解了！请看下面示例！
</MyDiv>
<MyDiv style="width: 100%; height: 50px;">
    <MyButton type="label" style="display: block; width: 100%; height: 40px; border-radius: 10px;"
              hov-style="display: block; width: 100%; height: 40px; border-radius: 10px; cursor: pointer; background-color: rgba(0, 0, 255, 0.2);"
              active-style="display: block; width: 100%; height: 40px; border-radius: 10px; cursor: pointer; background-color: rgba(0, 0, 255, 0.2); scale: 0.98">
        <MyDiv style="display: flex; align-items: center; width: 100%; height: 40px;">
            <MyImg src="https://www.baidu.com/favicon.ico" alt="百度"
                   style="margin-left: 10px; height: 30px; width: 30px"/>
            <MyDiv style="flex: 1; margin-left: 10px; margin-top: 0">
                <MyDiv style="font-size: 14px; margin-top: 0;">这是 MyListItem 的主标题</MyDiv>
                <MyDiv style="font-size: 10px; color: gray; margin-top: 0;">这是 MyListItem 的副标题</MyDiv>
            </MyDiv>
        </MyDiv>
    </MyButton>
</MyDiv>
<MyDiv>我知道上述代码可能会有点小多，不过如果你但凡学过一点点的前端知识，应该能懂上面写了什么，不过不懂也没有关系~完全可以
    copy 上述的写法~
</MyDiv>
</MyCard>
<MyCard title="NovaPlus 的替换标记与内置图片~" isExpand="True" canExpand="True">
<MyDiv>在 Nova Plus 里，有着很多的内置图片与替换标记，可以参考以下示例：</MyDiv>
<MyDiv>首先，说的便是我们的 path 标记啦！和 PCL
    一致，只需要用花括号括起来，就知道了替换标记是什么了~就像这样：（{path}），可以用于执行某个事件时需要用到！
</MyDiv>
<MyDiv>目前 Nova Plus 支持以下 几种 替换标记，分别是：</MyDiv>
<MyDiv style="display: table; width: 100%;">
    <MyDiv style="display: table-header-group">
        <MyDiv style="display: table-row; height: 50px;">
            <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; font-weight: bold; vertical-align: middle;">
                替换标记
            </MyDiv>
            <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; font-weight: bold; vertical-align: middle;">
                实际内容
            </MyDiv>
        </MyDiv>
    </MyDiv>
    <MyDiv style="display: table-row-group">
        <MyDiv style="display: table-row; height: 50px;">
            <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">
                path
            </MyDiv>
            <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">
                {path}
            </MyDiv>
        </MyDiv>
    </MyDiv>
</MyDiv>
<MyDiv>啊哈！又让你们学到了一招！表格！！表格见上方呀~</MyDiv>
<MyDiv>下面是 内置图片~内置图片见下方：</MyDiv>
<MyDiv>
    <MyImg src="CommandBlock" style="width: 32px; height: 32px; margin-left: 5px;"/>
    <MyImg src="Cobblestone" style="width: 32px; height: 32px; margin-left: 5px;"/>
    <MyImg src="GoldBlock" style="width: 32px; height: 32px; margin-left: 5px;"/>
    <MyImg src="Grass" style="width: 32px; height: 32px; margin-left: 5px;"/>
    <MyImg src="GrassPath" style="width: 32px; height: 32px; margin-left: 5px;"/>
    <MyImg src="Anvil" style="width: 32px; height: 32px; margin-left: 5px;"/>
    <MyImg src="RedstoneBlock" style="width: 32px; height: 32px; margin-left: 5px;"/>
    <MyImg src="RedstoneLampOn" style="width: 32px; height: 32px; margin-left: 5px;"/>
    <MyImg src="RedstoneLampOff" style="width: 32px; height: 32px; margin-left: 5px;"/>
    <MyImg src="Egg" style="width: 32px; height: 32px; margin-left: 5px;"/>
    <MyImg src="Fabric" style="width: 32px; height: 32px; margin-left: 5px;"/>
    <MyImg src="NeoForge" style="width: 32px; height: 32px; margin-left: 5px;"/>
</MyDiv>
<MyDiv>因此，各位可以使用上述图片来随时制作自己的图片按钮~就像以下：</MyDiv>
<MyDiv style="width: 100%; height: 50px;">
    <MyButton type="label" style="display: block; width: 100%; height: 40px; border-radius: 10px;"
              hov-style="display: block; width: 100%; height: 40px; border-radius: 10px; cursor: pointer; background-color: rgba(0, 0, 255, 0.2);"
              active-style="display: block; width: 100%; height: 40px; border-radius: 10px; cursor: pointer; background-color: rgba(0, 0, 255, 0.2); scale: 0.98">
        <MyDiv style="display: flex; align-items: center; width: 100%; height: 40px;">
            <MyImg src="CommandBlock" alt="百度" style="margin-left: 10px; height: 32px; width: 32px"/>
            <MyDiv style="flex: 1; margin-left: 10px; margin-top: 0">
                <MyDiv style="font-size: 14px; margin-top: 0;">下载测试版</MyDiv>
                <MyDiv style="font-size: 10px; color: gray; margin-top: 0;">23w07a</MyDiv>
            </MyDiv>
        </MyDiv>
    </MyButton>
</MyDiv>
</MyCard>
<MyCard title="MyButton 的所有事件详解" isExpand="True" canExpand="True">
<MyDiv>
    所有的事件以【event名|event参数1|event参数2|...】的样式进行~如果某一个事件没有参数，则不需要加【|】线，末尾同样也不需要|线
</MyDiv>
<MyDiv style="display: table; width: 100%;">
    <MyDiv style="display: table-header-group">
        <MyDiv style="display: table-row; height: 50px;">
            <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; font-weight: bold; vertical-align: middle;">
                事件名
            </MyDiv>
            <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; font-weight: bold; vertical-align: middle;">
                功能
            </MyDiv>
            <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; font-weight: bold; vertical-align: middle;">
                参数1
            </MyDiv>
            <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; font-weight: bold; vertical-align: middle;">
                参数2
            </MyDiv>
            <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; font-weight: bold; vertical-align: middle;">
                参数3
            </MyDiv>
        </MyDiv>
    </MyDiv>
    <MyDiv style="display: table-row-group">
        <MyDiv style="display: table-row; height: 50px;">
            <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">
                refresh
            </MyDiv>
            <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">
                刷新主页
            </MyDiv>
            <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">\
            </MyDiv>
            <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">\
            </MyDiv>
            <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">\
            </MyDiv>
        </MyDiv>
    </MyDiv>
    <MyDiv style="display: table-row-group">
        <MyDiv style="display: table-row; height: 50px;">
            <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">
                messagebox
            </MyDiv>
            <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">
                信息框
            </MyDiv>
            <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">
                标题
            </MyDiv>
            <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">
                内容
            </MyDiv>
            <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">
                颜色【0：蓝色、1：黄色、2：红色】
            </MyDiv>
        </MyDiv>
    </MyDiv>
    <MyDiv style="display: table-row-group">
        <MyDiv style="display: table-row; height: 50px;">
            <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">
                openurl
            </MyDiv>
            <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">
                打开网页
            </MyDiv>
            <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">
                url地址
            </MyDiv>
            <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">\
            </MyDiv>
            <MyDiv style="display: table-cell; border: 1px solid gray; text-align: center; vertical-align: middle;">\
            </MyDiv>
        </MyDiv>
    </MyDiv>
</MyDiv>
<MyDiv>举个例子~</MyDiv>
<MyButton width="100px" height="50px" event="messagebox|这是标题|这是警告框的内容|1">触发警告信息框~</MyButton>
<MyButton width="100px" height="50px" event="openurl|https://github.com/xphost008">打开作者的github~</MyButton>
</MyCard>
<MyValue id="key">这是一个变量，key为【key】，value为这串文本</MyValue>
<MyValue id="mydivstyle">background-color: red;</MyValue>
<MyValue id="myspanstyle">font-size: 50px;</MyValue>
<MyCard title="变量详解" isExpand="True" canExpand="True">
<MyDiv>这是一个变量：${key}</MyDiv>
<MyDiv style="${mydivstyle}">你也可以定义 MyDiv 和 <MySpan style="${myspanstyle}">MySpan</MySpan> 的 style 作为变量。
</MyDiv>
<MyDiv>MyValue 标签的适用范围：【文本、MyDiv的style、MySpan的style】，很抱歉，没有其余的功能了……</MyDiv>
<MyDiv>style 的样式作为 Value 值时，不仅可以设置变量，还可以加一点自己的值，比如说
    <MySpan style="${myspanstyle}; color: red;">50px，但是红色</MySpan>
</MyDiv>
<MyDiv>这边建议的是：将 MyValue 设置在整个主页的最顶端。这样相当于是全局引入了~整个主页都可以使用 Value 变量~<MyDiv
        style="margin-top: 0;"/>因为主页的加载过程是主要的从上往下执行💦
</MyDiv>
</MyCard>
```

可以将上述xml复制到 `{可执行文件}/PCL.Nova/HomePage/<随便名字>.nxml`！即可在启动器里【设置 -> 主页】里设置！