# 各位好呀！这里是教你如何在NovaPlus中自制魔塔！

## 基础知识

1. 需要具备 Lua 的编程基础
2. 需要了解 魔塔 的基本概念和规则
3. 需要熟悉魔塔的关卡设计和怪物属性

## 开始制作魔塔

1. 首先，新建一个魔塔项目，在 `./.PCL.Nova/MagicTower` 目录下，新建一个名为 `yourname.lua` 的文件。
    - macos 在 `~/Library/Application Support/.PCL.Nova/MagicTower` 下新建
2. 随后，尽情开始编写魔塔代码吧！这里会给各位介绍一下示例塔的制作过程。

## Lua 中的所有函数

所有函数和方法参见下表所示，所有函数类型和函数名在表中都有展示！

| 函数名 | 参数列表 | 返回值 | 功能 |
| ---- | ----- | ----- | ---- |
| init | `difficulty/int` | 无 | 初始化函数，需要写在外面当作 function init(difficulty) ... end 可以做几个 `local` 函数来初始化小怪的逻辑，difficulty 是难度，1表示简单，2表示普通，3表示困难，可以根据这个来初始化你的数据。  |
| set_chess | `width/int`, `height/int` | 无 | 设置棋盘大小（3<width<=20，3<height<=20），如果不提前设置则无法进行游戏 |
| set_level | `level/int` | 无 | 设置最大层数（2<level<=100），如果不提前设置则无法进行游戏 |
| set_lang | `type/string` | 无 | 默认zh_cn |
| new_player | `x/int`, `y/int`, `health/int`, `attack/int`, `defense/int`, `money/int` | `player` | 新建一个玩家，后续将使用这个玩家操控！ |
| add_player | `level/int`, `x/int`, `y/int`, `player/player` | `bool` | 添加一个玩家！ |
| add_wall | `level/int`, `x/int`, `y/int` | `bool` | 添加墙，返回是否成功 |
| new_monster | `name/string`, `health/int`, `attack/int`, `defense/int`, `money/int` | `monster` | 创建一个新的怪物，返回怪物对象 |
| add_monster | `level/int`, `x/int`, `y/int`, `monster/monster` | `bool` | 添加怪物，第三个参数填入上面 new 出来的属性，返回是否成功 |
| add_fake_wall  | `level/int`, `x/int`, `y/int` | `bool` | 添加假墙壁，走过去会自动变成地板。 |
| add_fake_floor | `level/int`, `x/int`, `y/int` | `bool` | 添加假地板，走过去会自动变成墙壁 |
| new_business | `name/string`, `sale/item`, `money/int` | `business` | 新增商人，sale是一个物品对象，晚点会说。当count是大于100的任何数时，是无限兑换。|
| add_business | `level/int`, `x/int`, `y/int`, `business/business` | `bool` | 添加商人，返回是否成功。 |
| add_lava | `level/int`, `x/int`, `y/int` | `bool` | 添加熔岩，必须当玩家背包里有 `ice` 属性时，方可通过！ |
| new_item | `type/string`, `health/int`, `attack/int`, `defense/int`, `money/int` | `item` | 新建物品，返回物品对象，type有8种类型，分别是：生命I（h1），生命II（h2），攻击I（a1），攻击II（a2），防御I（d1），防御II（d2），金钱I（m1），金钱II（m2）填入括号里的即可！填入其他的会报错！然后加多少值可以填在后面。 |
| new_magic_item | `name/string`, `callback(func(chess))` | `item` | 新增一个特殊物品，这个特殊物品会在玩家点击使用后，自动调用 `callback` 遍历整个棋盘，随后每遍历一个格子，就会输出这个格子的元数据，以及一些内置函数。可以使用 `run_process` 运行所有的内置函数！ |
| new_event | `name/string`, `callback(func(chess))` | `event` | 新增一个特殊事件，这个事件会在 `name` 事件名时，遍历整个棋盘。同上所示。type有4种类型，分别是：玩家走路（walk），玩家击杀怪物（kill），游戏时间（time），点击屏幕（click） |
| add_item | `level/int`, `x/int`, `y/int`, `item/item` | `bool` | 添加一个物品，返回是否成功 |
| new_messenger | `name/string`, `message/string`, `count/int` | `messenger` | 新建一个 `messenger`，count 大于 100 时将是无限对话。 |
| add_messenger | `level/int`, `x/int`, `y/int`, `messenger/messenger` | `bool` | 添加一个传话着，message以 `\|` 进行分割。每一个都是一句话。 |
| add_door | `level/int`, `x/int`, `y/int`, `type/string` | `bool` | 定义一扇门（type有4种取值，red，yellow，blue，green），其中绿钥匙原版未出现。 |
| add_key | `level/int`, `x/int`, `y/int`, `type/string` | `bool` | 定义一把钥匙，同上。 |
| new_shop | `name/string`, `initial/int` | `shop` | 定义一个商店，initial的意思是初始金币，后续会随着购买次数增长而指数型增长（每买一次价格翻一倍） |
| add_shop | `level/int`, `x/int`, `y/int`, `shop/shop` | `bool` | 新增一个商店。 |
| play_sound | `name/string` | `bool` | 播放一段音乐。 |

ok！上述就是目前 魔塔游戏 里所有的交互函数啦！

## 资源文件

我们可以找到【`Nova` 配置文件的 `MagicTower` 下面的 与`lua`文件同名的文件夹里，将图片资源保存到里面。其中，里面的目录大致如下：】

lua同名文件\
├─ assets\
│   ├─ monster\
│   │   ├─ \[name\].png\
│   ├─ items\
│   │   ├─ \[name\].png\
│   ├─ builtin\
│   │   ├─ red_key.png\
│   │   ├─ blue_key.png\
│   │   ├─ yellow_key.png\
│   │   ├─ green_key.png\
│   │   ├─ red_door.png\
│   │   ├─ blue_door.png\
│   │   ├─ yellow_door.png\
│   │   ├─ green_door.png\
│   │   ├─ wall.png\
│   │   ├─ floor.png\
│   │   ├─ lava.png\
│   │   ├─ player_front.png\
│   │   ├─ player_left.png\
│   │   ├─ player_right.png\
│   │   ├─ player_back.png\
│   │   ├─ bg.png\
│   │   ├─ health1.png\
│   │   ├─ health2.png\
│   │   ├─ attack1.png\
│   │   ├─ attack2.png\
│   │   ├─ defense1.png\
│   │   ├─ defense2.png\
│   │   ├─ money1.png\
│   │   ├─ money2.png\
│   └─ other\
├─ sounds\
│   ├─ kill\
│   │   ├─ kill1.ogg\
│   │   ├─ kill2.ogg\
│   │   ├─ ...\
│   │   ├─ killn.ogg\
│   ├─ messenger\
│   │   ├─ messenger1.ogg\
│   │   ├─ messenger2.ogg\
│   │   ├─ ...\
│   │   ├─ messengern.ogg\
│   ├─ gain\
│   │   ├─ gain1.ogg\
│   │   ├─ gain2.ogg\
│   │   ├─ ...\
│   │   ├─ gainn.ogg\
│   ├─ door\
│   │   ├─ door1.ogg\
│   │   ├─ door2.ogg\
│   │   ├─ ...\
│   │   ├─ door4.ogg\
│   └─ bgm\
│   -   └─ \[name\].mp3\
└─ lang\
\-   └─ \[name\].json

## 答疑解惑

其中，各位可能会有点疑惑：new里面的name是什么意思呢？
其实很简单嗷！这个是资源文件的意思！我们可以看到在每 new 一个怪物、物品时，都并不会使用任何图片资源文件，所以如果我们想自定义图片资源，就得用到name！

在 `lua` 文件同名的文件夹下，新建一个 `assets`，随后便可将自己的图片资源放进去啦！

这里提供一个[示例游戏](https://wwdy.lanzoub.com/b0sy6w7zi)，密码:60nw，作者 `xphost` 自己写的魔塔原版，各位可以自行下载，随后将其解压到 `PCL.Nova/MagicTower` 目录下，如果是 Linux 或者 macOS，你们则需要打开隐藏文件夹显示才能看见配置文件！
