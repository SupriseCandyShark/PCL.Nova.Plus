import { BrowserOpenURL } from "../../wailsjs/runtime";
import { quadInOut } from "svelte/easing";
import {
    GetConfigIniPath,
    ReadConfig,
} from "../../wailsjs/go/launcher/ReaderWriter";
import { GetUniqueAddress } from "../../wailsjs/go/main/App";
import {
    current_homepage,
    homepage_list,
    homepage_url,
    select_homepage,
} from "./changeBody";
import { GetAllHomePage, ReadFile } from "../../wailsjs/go/launcher/MainMethod";
import { OpenExplorer } from "../../wailsjs/go/launcher/ReaderWriter";
import { HttpGet } from "../../wailsjs/go/launcher/Network";
import { messagebox, MSG_ERROR, showHint } from "./messagebox";

export function DarkAndThemeToConst(
    darkMode: boolean,
    themeMode: number,
): string {
    let co = [
        [
            "linear-gradient(to left, rgb(10, 114, 189), rgb(11, 129, 213), rgb(10, 114, 189))", //-6, +8, -7, -7, +10, -8
            "linear-gradient(to left, rgb(40, 147, 132), rgb(45, 166, 148), rgb(40, 147, 132))",
            "linear-gradient(to left, rgb(58, 156, 78), rgb(66, 180, 73), rgb(58, 156, 78))",
            "linear-gradient(to left, rgb(117, 155, 42), rgb(131, 176, 47), rgb(117, 155, 42))",
            "linear-gradient(to left, rgb(144, 129, 74), rgb(161, 146, 85), rgb(144, 129, 74))",
            "url('/src/assets/images/Themes/BlackDark.png')",
            "url('/src/assets/images/Themes/FoolRainbowDark.png')",
            "url('/src/assets/images/Themes/PinkDark.png')",
            "url('/src/assets/images/Themes/PurpleDark.png')",
            "url('/src/assets/images/Themes/LuckyRainbowDark.png')",
            "url('/src/assets/images/Themes/GoldDark.png')",
            "url('/src/assets/images/Themes/OrangeDark.png')",
            "url('/src/assets/images/Themes/MojangRedDark.png')",
            "url('/src/assets/images/Themes/HackBlueDark.png')",
        ],
        [
            "linear-gradient(to left, rgb(16, 106, 196), rgb(18, 119, 221), rgb(16, 106, 196))",
            "linear-gradient(to left, rgb(46, 139, 139), rgb(52, 156, 156), rgb(46, 139, 139))",
            "linear-gradient(to left, rgb(64, 148, 85), rgb(73, 170, 81), rgb(64, 148, 85))",
            "linear-gradient(to left, rgb(123, 147, 49), rgb(138, 166, 55), rgb(123, 147, 49))",
            "linear-gradient(to left, rgb(150, 121, 81), rgb(168, 136, 93), rgb(150, 121, 81))",
            "url('/src/assets/images/Themes/Black.png')",
            "url('/src/assets/images/Themes/FoolRainbow.png')",
            "url('/src/assets/images/Themes/Pink.png')",
            "url('/src/assets/images/Themes/Purple.png')",
            "url('/src/assets/images/Themes/LuckyRainbow.png')",
            "url('/src/assets/images/Themes/Gold.png')",
            "url('/src/assets/images/Themes/Orange.png')",
            "url('/src/assets/images/Themes/MojangRed.png')",
            "url('/src/assets/images/Themes/HackBlue.png')",
        ],
    ];
    return (
        co[darkMode ? 0 : 1][themeMode - 1] ??
        "linear-gradient(to left, rgb(16, 106, 196), rgb(18, 119, 221), rgb(16, 106, 196))"
    );
}

export function DarkAndThemeToMain(
    darkMode: boolean,
    themeMode: number,
): string {
    let co = [
        [
            "linear-gradient(to left bottom, rgb(4, 22, 27), rgb(6, 17, 35))",
            "linear-gradient(to left bottom, rgb(4, 22, 27), rgb(6, 17, 35))",
            "linear-gradient(to left bottom, rgb(4, 22, 27), rgb(6, 17, 35))",
            "linear-gradient(to left bottom, rgb(4, 22, 27), rgb(6, 17, 35))",
            "linear-gradient(to left bottom, rgb(4, 22, 27), rgb(6, 17, 35))",
            "linear-gradient(to left bottom, rgb(4, 22, 27), rgb(6, 17, 35))",
            "linear-gradient(to left bottom, rgb(4, 22, 27), rgb(6, 17, 35))",
            "linear-gradient(to left bottom, rgb(4, 22, 27), rgb(6, 17, 35))",
            "linear-gradient(to left bottom, rgb(4, 22, 27), rgb(6, 17, 35))",
            "linear-gradient(to left bottom, rgb(4, 22, 27), rgb(6, 17, 35))",
            "linear-gradient(to left bottom, rgb(4, 22, 27), rgb(6, 17, 35))",
            "linear-gradient(to left bottom, rgb(4, 22, 27), rgb(6, 17, 35))",
            "linear-gradient(to left bottom, rgb(4, 22, 27), rgb(6, 17, 35))",
            "linear-gradient(to left bottom, rgb(4, 22, 27), rgb(6, 17, 35))",
        ],
        [
            "linear-gradient(to left bottom, rgb(150, 212, 235), rgb(179, 196, 241))",
            "linear-gradient(to left bottom, rgb(214, 237, 232), rgb(216, 233, 239))",
            "linear-gradient(to left bottom, rgb(223, 238, 222), rgb(220, 237, 227))",
            "linear-gradient(to left bottom, rgb(236, 237, 213), rgb(229, 239, 217))",
            "linear-gradient(to left bottom, rgb(238, 230, 226), rgb(235, 233, 223))",
            "linear-gradient(to left bottom, rgb(230, 232, 234), rgb(231, 231, 234))",
            "linear-gradient(to left bottom, rgb(255, 250, 254), rgb(255, 126, 143))",
            "linear-gradient(to left bottom, rgb(245, 219, 231), rgb(246, 223, 224))",
            "linear-gradient(to left bottom, rgb(234, 224, 242), rgb(239, 219, 240))",
            "linear-gradient(to left bottom, rgb(235, 223, 241), rgb(240, 220, 238))",
            "linear-gradient(to left bottom, rgb(243, 233, 215), rgb(238, 239, 207))",
            "linear-gradient(to left bottom, rgb(246, 226, 219), rgb(244, 232, 212))",
            "linear-gradient(to left bottom, rgb(246, 222, 224), rgb(245, 227, 219))",
            "linear-gradient(to left bottom, rgb(221, 229, 243), rgb(227, 226, 245))",
        ],
    ];
    return (
        co[darkMode ? 0 : 1][themeMode - 1] ??
        "linear-gradient(to left bottom, rgb(150, 212, 235), rgb(179, 196, 241))"
    );
}

function judgeVersion(left: string, right: string): boolean {
    let splitLeft = left.split(".");
    let splitRight = right.split(".");
    for (let i = 0; i < 3; i++) {
        let sLeft = parseInt(splitLeft[i]) || 0;
        let sRight = parseInt(splitRight[i]) || 0;
        if (sLeft > sRight) {
            return false;
        } else if (sLeft < sRight) {
            return true;
        }
    }
    return false;
}

// 按照 Forge 版本进行排序（boo是指：当boo为true时，则调用arr[i].version，否则直接排序arr。）
// 下列方法只适用于：arr要么是个里面包含对象的键值数组，要么直接就是个字符串数组
export function SortForgeVersion(arr: any[], boo: boolean = false) {
    for (let i = 0; i < arr.length; i++) {
        for (let j = 0; j < arr.length - i - 1; j++) {
            if (
                judgeVersion(
                    boo ? arr[j].version : arr[j],
                    boo ? arr[j + 1].version : arr[j + 1],
                )
            ) {
                let a = arr[j];
                arr[j] = arr[j + 1];
                arr[j + 1] = a;
            }
        }
    }
}

export function OpenCustomURL(url: string) {
    BrowserOpenURL(url);
}

export function copyToClipBoard(message: string) {
    let aux = document.createElement("input");
    aux.setAttribute("value", message);
    document.body.appendChild(aux);
    aux.select();
    document.execCommand("copy");
    document.body.removeChild(aux);
}

export function slide_up(node: HTMLElement) {
    return {
        duration: 200,
        easing: quadInOut,
        css: (t: number, n: number) => {
            return `
                transform: translateY(${-50 * n}%);
                opacity: ${t};
            `;
        },
    };
}

export function slide_left(node: HTMLElement, param: { x: number }) {
    const x = param.x || 144;
    return {
        duration: 150,
        easing: quadInOut,
        css: (t: number, n: number) => {
            return `
                transform: translateX(${-x * n}px);
                opacity: ${t};
            `;
        },
    };
}

export function slide_opacity_transcale(node: HTMLElement) {
    return {
        duration: 200,
        easing: quadInOut,
        css: (t: number) => {
            return `
                opacity: ${t};
                transform: scale(${t});
            `;
        },
    };
}
export function slide_opacity(node: HTMLElement) {
    return {
        duration: 200,
        easing: quadInOut,
        css: (t: number) => {
            return `opacity: ${t};`;
        },
    };
}

export function parseXmlToJson(xmlString: string): any[] {
    const parser = new DOMParser();
    const xmlDoc = parser.parseFromString(xmlString, "text/xml");

    // 处理解析错误
    const parseError = xmlDoc.querySelector("parsererror");
    if (parseError) {
        throw new Error(parseError.textContent);
    }

    return convertNode(xmlDoc.documentElement);
}

function convertNode(node: Node): any[] {
    const result: any[] = [];

    // 只处理元素节点和文本节点
    if (node.nodeType === Node.ELEMENT_NODE) {
        const element = node as Element;
        const obj: any = { name: element.tagName };

        // 处理属性
        for (const attr of Array.from(element.attributes)) {
            if (attr.name !== "name") {
                // 跳过 name 属性
                obj[attr.name] = attr.value;
            }
        }

        // 处理子节点
        const children: any[] = [];
        for (const childNode of Array.from(element.childNodes)) {
            if (childNode.nodeType === Node.ELEMENT_NODE) {
                children.push(...convertNode(childNode));
            } else if (childNode.nodeType === Node.TEXT_NODE) {
                const textContent = childNode.textContent?.trim();
                if (textContent) {
                    children.push({
                        name: "text",
                        content: textContent,
                    });
                }
            }
        }

        // 双标签必须有 children 字段
        if (element.childNodes.length > 0) {
            obj.children = children;
        }

        result.push(obj);
    }

    return result;
}

export async function getHomePage(): Promise<string> {
    let ind = Number(
        await ReadConfig(await GetConfigIniPath(), "Misc", "SelectHomePage"),
    );
    if (ind && ind > 0 && ind <= 3) {
        select_homepage.set(ind);
        switch (ind) {
            case 1:
                //TODO: 预设主页
                break;
            case 2:
                let ld = await GetAllHomePage();
                if (!ld.status) {
                    await messagebox(
                        "错误的主页",
                        "加载到错误的主页！错误信息：" + ld.message,
                        MSG_ERROR,
                    );
                    return;
                }
                homepage_list.set([...ld.data]);
                let i = Number(
                    await ReadConfig(
                        await GetConfigIniPath(),
                        "Misc",
                        "HomePageValue",
                    ),
                );
                if (i >= 0 && i < ld.data.length) {
                    current_homepage.set(i);
                    showHint("主页已设置完毕喵~");
                    return await ReadFile(ld.data[i].path);
                }
                break;
            case 3:
                let up = await ReadConfig(
                    await GetConfigIniPath(),
                    "Misc",
                    "HomePageURL",
                );
                homepage_url.set(up);
                let hp = await HttpGet(up, "");
                if (hp && hp != "") {
                    showHint("主页已设置完毕喵~");
                    return hp;
                }
                break;
        }
    }
    return "none";
}

export async function openExplorer(path: string) {
    if (!(await OpenExplorer(path))) {
        await messagebox(
            "无法打开 文件管理器",
            "请确保你的系统已经内置 文件管理器。如果你使用的 Linux 发行版与常见的不同，请尝试安装 xdg-open 以打开文件管理器~",
            MSG_ERROR,
        );
    }
}

// export function syncDo<T>(p: Promise<T>): T {
//     let res: T;
//     p.then((r) => (res = r));
//     while (!res) {}
//     return res;
// }

export function sleep(time: number): Promise<void> {
    return new Promise((resolve) => {
        setTimeout(resolve, time);
    });
}

export function GetLittleSkinUrl(): string[] {
    return [
        "https://open.littleskin.cn/oauth/link",
        "https://open.littleskin.cn/oauth/device_code",
        "https://open.littleskin.cn/oauth/token",
        "https://littleskin.cn/api/yggdrasil",
    ];
}

export function GetAllCaves(): string[] {
    return [
        "一口一个大佬，说的好像群里全是大佬一样",
        "一觉醒来，全世界的何意味水平下降一百万倍。<br>你随口说出一句“何意味”，却不知何时身边围了一群同学，他们都张大了嘴巴，鸦雀无声，以为自己出现了幻觉。没想到这所大学里竟然会有这般顶尖人物！<br>“何意味……何意味？何意味？何意味？？！”<br>“何意味！何意味？何意味…… ”<br>而教授也呆住了，哪怕是他，也只是能勉强何意味，而他已经因此已经拿到了终身教职！至于你口中的“何意味，他只在最顶尖的何意味会议上听说过。<br>教授老泪纵横地赶忙与你握手，激动地说，“何意味！”<br>在周围人崇拜佩服的目光下，你十分得意，想到自己名声大噪后成为世界第一何意味大师的画面激动不已，忍不住轻哼起来",
        "就是要善用“利大于弊”田忌赛马式的妙比法，用你的利去比他人的弊，比弊时，压力即刻骤减；比利时，是位于欧洲西部的一个国家。",
        "乃好哇awa咱是奶兽W，窝今年十四岁惹qaq（不要嫌我年龄小www咱懂得可多了!)窝喜欢玩原神，星铁，舟舟，迷你世界。咱超厉害的~可以带乃们飞ww。咱平时可能会发怪东西qwq因为咱有重度双相，爆米花症，焦糖症，玉米症qwq虽然之前不小心砍了几个小乖乖不过咱现在很好克服!窝喜欢的兽兽有银碳，狮忆，夏夏，E 电鹿板。咱很期待和乃们成为好抛友",
        "1. 作为小县城唯一何过意味的人，我的眼神里总是透着同龄人不懂的意味，每次家里跳闸停电，我都会平静地抬手轻按空气开关：“何意味？”<br>2. 作为小县城唯一何过意味的人，我的眼神里总带着同龄人没有的意味，每次看到老人在巷口摆摊卖手作小物件，我都会蹲下来轻声问：“何意味？”<br>3. 作为小县城唯一何过意味的人，我的眼神里藏着同龄人不懂的意味，每次县里举办文艺汇演台下掌声雷动，我都会笑着拍手：“何意味？”<br>4. 作为小县城唯一何过意味的人，我的眼神里总有着同龄人没有的意味，每次在书店看到孩子捧着绘本认真阅读，我都会放轻脚步：“何意味？”",
        "你说得对，但是《PCI.Nova.Plus》是由PCL-Community自主研发的一款全新多平台启动器",
        "我的血液流淌着马铃薯",
        "//我写这一行代码的时候，只有上帝和我知道我在写什么<br>//现在，只有上帝知道了"
    ];
}

export async function GetSelfUniqueAddress(): Promise<string> {
    let identify = await GetUniqueAddress();
    if (identify == "ERR_WINDOWS") {
        messagebox(
            "获取唯一识别码失败",
            "获取唯一识别码失败了！已检测到您所使用的操作系统为 Windows，你可能是在虚拟机里面运行的 Nova，因为无法调用你的 wmic 命令！如果你的操作系统环境正常，并且仍然无法获取唯一标识符，请联系启动器作者！",
            MSG_ERROR,
        );
        return "";
    } else if (identify == "ERR_LINUX") {
        messagebox(
            "获取唯一识别码失败",
            `获取唯一识别码失败了！已检测到您所使用的操作系统为 Linux，请首先尝试使用以下命令：<br>
<span style="color: skyblue; user-select: all; -webkit-user-select: all; -moz-user-select: none; -ms-user-select: none;">sudo chmod +r /sys/class/dmi/id/product_uuid</span><br>
随后，再次尝试，如果仍然无法获取唯一标识符，请联系启动器作者！`,
            MSG_ERROR,
        );
        return "";
    } else if (identify == "ERR_MACOS") {
        messagebox(
            "获取唯一识别码失败",
            "获取唯一识别码失败了！已检测到您所使用的操作系统为 macOS，你可能是在虚拟机里面运行的 Nova，因为无法调用你的 system_profiler 命令！如果你的操作系统环境正常，并且仍然无法获取唯一标识符，请联系启动器作者！",
            MSG_ERROR,
        );
        return "";
    }
    return identify;
}
