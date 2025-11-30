<script lang="ts">
    import MySelectCard from "../../../component/card/MySelectCard.svelte";
    import { OpenCustomURL } from "../../../store/functions";
    import MyNormalButton from "../../../component/button/MyNormalButton.svelte";
    import MyNormalSpan from "../../../component/input/MyNormalSpan.svelte";
    // 导入图片
    import Xphost from "../../../assets/images/Heads/xphost.jpg";
    import FireDragon0659 from "../../../assets/images/Heads/firedragon0659.jpg";
    import CandyShark from "../../../assets/images/Heads/CandyShark.png";
    import NovaPlus from "../../../assets/images/NovaPlus.png";
    import LtCatt from "../../../assets/images/Heads/ltcatt.jpg";
    import MoYuan from "../../../assets/images/Heads/moyuan.jpg";
    import Hill233 from "../../../assets/images/Heads/hill233.jpg";
    import MyCardButton from "../../../component/button/MyCardButton.svelte";
    import { dark_mode, launcher_version } from "../../../store/changeBody";
    import Afdian from "../../../assets/images/afdian.ico";
    import MyCodeSpan from "../../../component/input/MyCodeSpan.svelte";

    export let slide = null;
    export let after_leave = null;
    const dev = [
        {
            avatar: Xphost,
            title: "xphost008",
            desc: "Plain Craft Launcher II: Nova Plus 的现作者之一",
            buttons: [
                {
                    title: "Github 主页",
                    click: () => {
                        OpenCustomURL("https://github.com/xphost008");
                    },
                    icon: `<svg
                                    xmlns="http://www.w3.org/2000/svg"
                                    viewBox="0 0 24 24"
                                    stroke-linecap="round"
                                    stroke-linejoin="round"
                                    stroke-width="1.5"
                                    style="width: 30px; height: 30px"
                                    fill="${$dark_mode ? "#ffffff" : "#000000"}">
                                    <path d="M6.517 17.113c.395.578 1.592 1.81 3.225 2.12M9.864 22C8.836 21.83 2 19.606 2 12.093C2 5.063 8.002 2 12 2c4 0 10 3.063 10 10.093c0 7.513-6.836 9.738-7.864 9.907c0 0-.21-3.417-.087-4.003c.122-.586-.294-1.528-.294-1.528c.971-.364 2.45-.884 2.945-2.282c.385-1.084.627-2.658-.45-4.138c0 0 .282-2.39-.25-2.484c-.533-.092-2.1.947-2.1.947c-.457-.13-1.476-.377-1.898-.333c-.423-.044-1.445.203-1.902.333c0 0-1.568-1.04-2.1-.947s-.25 2.484-.25 2.484c-1.077 1.48-.835 3.054-.45 4.138c.496 1.398 1.974 1.918 2.945 2.282c0 0-.416.942-.294 1.528S9.864 22 9.864 22" color="currentColor"/>
                        </svg>`,
                },
                {
                    title: "爱发电",
                    click: () => {
                        OpenCustomURL("https://afdian.com/a/xphost");
                    },
                    iconType: "img",
                    icon: Afdian,
                },
            ],
        },
        {
            avatar: FireDragon0659,
            title: "FireDragon0659",
            desc: "Plain Craft Launcher II: Nova Plus 的现作者之一",
            buttons: [
                {
                    title: "Github 主页",
                    click: () => {
                        OpenCustomURL("https://github.com/FireDragon0659");
                    },
                    icon: `<svg
                                    xmlns="http://www.w3.org/2000/svg"
                                    viewBox="0 0 24 24"
                                    stroke-linecap="round"
                                    stroke-linejoin="round"
                                    stroke-width="1.5"
                                    style="width: 30px; height: 30px"
                                    fill="${$dark_mode ? "#ffffff" : "#000000"}">
                                    <path d="M6.517 17.113c.395.578 1.592 1.81 3.225 2.12M9.864 22C8.836 21.83 2 19.606 2 12.093C2 5.063 8.002 2 12 2c4 0 10 3.063 10 10.093c0 7.513-6.836 9.738-7.864 9.907c0 0-.21-3.417-.087-4.003c.122-.586-.294-1.528-.294-1.528c.971-.364 2.45-.884 2.945-2.282c.385-1.084.627-2.658-.45-4.138c0 0 .282-2.39-.25-2.484c-.533-.092-2.1.947-2.1.947c-.457-.13-1.476-.377-1.898-.333c-.423-.044-1.445.203-1.902.333c0 0-1.568-1.04-2.1-.947s-.25 2.484-.25 2.484c-1.077 1.48-.835 3.054-.45 4.138c.496 1.398 1.974 1.918 2.945 2.282c0 0-.416.942-.294 1.528S9.864 22 9.864 22" color="currentColor"/>
                        </svg>`,
                },
            ],
        },
        {
            avatar: CandyShark,
            title: "Candy Shark",
            desc: "惊喜糖鲨开发团队",
            buttons: [
                {
                    title: "Github 主页",
                    click: () => {
                        OpenCustomURL("https://github.com/3TProject");
                    },
                    icon: `<svg
                                    xmlns="http://www.w3.org/2000/svg"
                                    viewBox="0 0 24 24"
                                    stroke-linecap="round"
                                    stroke-linejoin="round"
                                    stroke-width="1.5"
                                    style="width: 30px; height: 30px"
                                    fill="${$dark_mode ? "#ffffff" : "#000000"}">
                                    <path d="M6.517 17.113c.395.578 1.592 1.81 3.225 2.12M9.864 22C8.836 21.83 2 19.606 2 12.093C2 5.063 8.002 2 12 2c4 0 10 3.063 10 10.093c0 7.513-6.836 9.738-7.864 9.907c0 0-.21-3.417-.087-4.003c.122-.586-.294-1.528-.294-1.528c.971-.364 2.45-.884 2.945-2.282c.385-1.084.627-2.658-.45-4.138c0 0 .282-2.39-.25-2.484c-.533-.092-2.1.947-2.1.947c-.457-.13-1.476-.377-1.898-.333c-.423-.044-1.445.203-1.902.333c0 0-1.568-1.04-2.1-.947s-.25 2.484-.25 2.484c-1.077 1.48-.835 3.054-.45 4.138c.496 1.398 1.974 1.918 2.945 2.282c0 0-.416.942-.294 1.528S9.864 22 9.864 22" color="currentColor"/>
                        </svg>`,
                },
            ],
        },
        {
            avatar: NovaPlus,
            title: "Plain Craft Launcher II: Nova Plus",
            desc: `当前版本：${$launcher_version}`,
            buttons: [
                {
                    title: "Github 主页",
                    click: () => {
                        OpenCustomURL(
                            "https://github.com/3TProject/PCL.Nova.Plus",
                        );
                    },
                    icon: `<svg
                                    xmlns="http://www.w3.org/2000/svg"
                                    viewBox="0 0 24 24"
                                    stroke-linecap="round"
                                    stroke-linejoin="round"
                                    stroke-width="1.5"
                                    style="width: 30px; height: 30px"
                                    fill="${$dark_mode ? "#ffffff" : "#000000"}">
                                    <path d="M6.517 17.113c.395.578 1.592 1.81 3.225 2.12M9.864 22C8.836 21.83 2 19.606 2 12.093C2 5.063 8.002 2 12 2c4 0 10 3.063 10 10.093c0 7.513-6.836 9.738-7.864 9.907c0 0-.21-3.417-.087-4.003c.122-.586-.294-1.528-.294-1.528c.971-.364 2.45-.884 2.945-2.282c.385-1.084.627-2.658-.45-4.138c0 0 .282-2.39-.25-2.484c-.533-.092-2.1.947-2.1.947c-.457-.13-1.476-.377-1.898-.333c-.423-.044-1.445.203-1.902.333c0 0-1.568-1.04-2.1-.947s-.25 2.484-.25 2.484c-1.077 1.48-.835 3.054-.45 4.138c.496 1.398 1.974 1.918 2.945 2.282c0 0-.416.942-.294 1.528S9.864 22 9.864 22" color="currentColor"/>
                        </svg>`,
                },
            ],
        },
    ];
    const credit = [
        {
            avatar: LtCatt,
            title: "龙腾猫跃",
            desc: "Plain Craft Launcher II 的原作者",
            buttons: [
                {
                    title: "Github 主页",
                    click: () => {
                        OpenCustomURL("https://github.com/LTCatt");
                    },
                    icon: `<svg
                                xmlns="http://www.w3.org/2000/svg"
                                viewBox="0 0 24 24"
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                stroke-width="1.5"
                                style="width: 30px; height: 30px"
                                fill="${$dark_mode ? "#ffffff" : "#000000"}">
                                <path d="M6.517 17.113c.395.578 1.592 1.81 3.225 2.12M9.864 22C8.836 21.83 2 19.606 2 12.093C2 5.063 8.002 2 12 2c4 0 10 3.063 10 10.093c0 7.513-6.836 9.738-7.864 9.907c0 0-.21-3.417-.087-4.003c.122-.586-.294-1.528-.294-1.528c.971-.364 2.45-.884 2.945-2.282c.385-1.084.627-2.658-.45-4.138c0 0 .282-2.39-.25-2.484c-.533-.092-2.1.947-2.1.947c-.457-.13-1.476-.377-1.898-.333c-.423-.044-1.445.203-1.902.333c0 0-1.568-1.04-2.1-.947s-.25 2.484-.25 2.484c-1.077 1.48-.835 3.054-.45 4.138c.496 1.398 1.974 1.918 2.945 2.282c0 0-.416.942-.294 1.528S9.864 22 9.864 22" color="currentColor"/>
                    </svg>`,
                },
                {
                    title: "爱发电",
                    click: () => {
                        OpenCustomURL("https://afdian.com/a/LTCat");
                    },
                    iconType: "img",
                    icon: Afdian,
                },
            ],
        },
        {
            avatar: MoYuan,
            title: "墨渊",
            desc: "Plain Craft Launcher II: Nova 的原作者",
            buttons: [
                {
                    title: "Github 主页",
                    click: () => {
                        OpenCustomURL("https://github.com/MoYuan-CN");
                    },
                    icon: `<svg
                                xmlns="http://www.w3.org/2000/svg"
                                viewBox="0 0 24 24"
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                stroke-width="1.5"
                                style="width: 30px; height: 30px"
                                fill="${$dark_mode ? "#ffffff" : "#000000"}">
                                <path d="M6.517 17.113c.395.578 1.592 1.81 3.225 2.12M9.864 22C8.836 21.83 2 19.606 2 12.093C2 5.063 8.002 2 12 2c4 0 10 3.063 10 10.093c0 7.513-6.836 9.738-7.864 9.907c0 0-.21-3.417-.087-4.003c.122-.586-.294-1.528-.294-1.528c.971-.364 2.45-.884 2.945-2.282c.385-1.084.627-2.658-.45-4.138c0 0 .282-2.39-.25-2.484c-.533-.092-2.1.947-2.1.947c-.457-.13-1.476-.377-1.898-.333c-.423-.044-1.445.203-1.902.333c0 0-1.568-1.04-2.1-.947s-.25 2.484-.25 2.484c-1.077 1.48-.835 3.054-.45 4.138c.496 1.398 1.974 1.918 2.945 2.282c0 0-.416.942-.294 1.528S9.864 22 9.864 22" color="currentColor"/>
                    </svg>`,
                },
            ],
        },
        {
            avatar: Hill233,
            title: "Hill233",
            desc: "Plain Craft Launcher II: Neo 的原作者",
            buttons: [
                {
                    title: "Github 主页",
                    click: () => {
                        OpenCustomURL("https://github.com/Hill23333");
                    },
                    icon: `<svg
                                xmlns="http://www.w3.org/2000/svg"
                                viewBox="0 0 24 24"
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                stroke-width="1.5"
                                style="width: 30px; height: 30px"
                                fill="${$dark_mode ? "#ffffff" : "#000000"}">
                                <path d="M6.517 17.113c.395.578 1.592 1.81 3.225 2.12M9.864 22C8.836 21.83 2 19.606 2 12.093C2 5.063 8.002 2 12 2c4 0 10 3.063 10 10.093c0 7.513-6.836 9.738-7.864 9.907c0 0-.21-3.417-.087-4.003c.122-.586-.294-1.528-.294-1.528c.971-.364 2.45-.884 2.945-2.282c.385-1.084.627-2.658-.45-4.138c0 0 .282-2.39-.25-2.484c-.533-.092-2.1.947-2.1.947c-.457-.13-1.476-.377-1.898-.333c-.423-.044-1.445.203-1.902.333c0 0-1.568-1.04-2.1-.947s-.25 2.484-.25 2.484c-1.077 1.48-.835 3.054-.45 4.138c.496 1.398 1.974 1.918 2.945 2.282c0 0-.416.942-.294 1.528S9.864 22 9.864 22" color="currentColor"/>
                    </svg>`,
                },
            ],
        },
    ];
</script>

<div class="component-about" in:slide out:slide on:outroend={after_leave}>
    <MySelectCard title="开发人员">
        <div class="proc">
            {#each dev as d}
                <MyCardButton
                    image={d.avatar}
                    title={d.title}
                    desc={d.desc}
                    image_style="border-radius: 50%;"
                    button_style="width: 50px; height: 50px; top: 0;"
                    buttons={((b) => {
                        let be = [];
                        for (let i = 0; i < b.length; i++) {
                            be.push({
                                title: b[i].title,
                                click: b[i].click,
                                iconType: b[i].iconType || "",
                                icon: b[i].icon,
                            });
                        }
                        return be;
                    })(d.buttons)}
                ></MyCardButton>
            {/each}
        </div>
    </MySelectCard>
    <MySelectCard title="鸣谢人员">
        <div class="proc">
            {#each credit as d}
                <MyCardButton
                    image={d.avatar}
                    title={d.title}
                    desc={d.desc}
                    image_style="border-radius: 50%;"
                    button_style="width: 50px; height: 50px; top: 0;"
                    buttons={((b) => {
                        let be = [];
                        for (let i = 0; i < b.length; i++) {
                            be.push({
                                title: b[i].title,
                                click: b[i].click,
                                iconType: b[i].iconType || "",
                                icon: b[i].icon,
                            });
                        }
                        return be;
                    })(d.buttons)}
                ></MyCardButton>
            {/each}
        </div>
    </MySelectCard>
    <MySelectCard title="协议" isExpand={true}>
        <div class="proc">
            <div style="margin: 0 20px">
                <MyNormalSpan
                    >PCL.Nova.Plus 开源项目及其官方衍生程序均遵循 LtCat License
                    协议开源。本项目主导组织为 SupriseCandyShark 以及 PCL
                    Community，任何人都可以提 Pull Request！
                </MyNormalSpan>
                <br /><br />
                <MyNormalButton
                    style_in="width: 150px; height: 40px"
                    on:click={() =>
                        OpenCustomURL(
                            "https://github.com/3TProject/PCL.Nova.Plus",
                        )}
                    >查看源代码
                </MyNormalButton>
                <MyNormalButton
                    style_in="width: 150px; height: 40px; margin-left: 20px"
                    on:click={() =>
                        OpenCustomURL(
                            "https://github.com/3TProject/PCL.Nova.Plus/blob/main/LICENSE",
                        )}
                >
                    查看许可证
                </MyNormalButton>
            </div>
        </div>
    </MySelectCard>
    <MySelectCard title="许可与版权声明" isExpand={true}>
        <div class="proc">
            <div style="margin: 0 20px; width: calc(100% - 40px)">
                <MyNormalSpan>PCL.Nova.Plus（游戏主程序）</MyNormalSpan>
                <br />
                <span class="gray">GNU-GPL v3</span>
                <MyCodeSpan>
                    https://github.com/SupriseCandyShark/PCL.Nova.Plus
                </MyCodeSpan>
                <br />
                <MyNormalSpan>Wails（WebView程序）</MyNormalSpan>
                <br />
                <span class="gray">MIT</span>
                <MyCodeSpan>
                    https://github.com/wailsapp/wails
                </MyCodeSpan>
                <br />
                <MyNormalSpan>Golang（后端语言）</MyNormalSpan>
                <br />
                <span class="gray">BSD 3</span>
                <MyCodeSpan>
                    https://github.com/golang/go
                </MyCodeSpan>
                <br />
                <MyNormalSpan>Svelte（前端架构）</MyNormalSpan>
                <br />
                <span class="gray">MIT</span>
                <MyCodeSpan>
                    https://github.com/sveltejs/svelte
                </MyCodeSpan>
                <br />
                <MyNormalSpan>TypeScript（前端语言）</MyNormalSpan>
                <br />
                <span class="gray">Apache License 2.0</span>
                <MyCodeSpan>
                    https://github.com/microsoft/TypeScript
                </MyCodeSpan>
                <br />
                <MyNormalSpan>MMCLL（启动器类库）</MyNormalSpan>
                <br />
                <span class="gray">MIT</span>
                <MyCodeSpan>
                    https://github.com/xphost008/mmcll
                </MyCodeSpan>
                <br />
                <MyNormalSpan>skinview3d</MyNormalSpan>
                <br />
                <span class="gray">MIT</span>
                <MyCodeSpan>
                    https://github.com/bs-community/skinview3d
                </MyCodeSpan>
                <br />
            </div>
        </div>
    </MySelectCard>
</div>

<style>
    .component-about {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        overflow-y: auto;
    }

    .gray {
        color: gray;
        font-size: 14px;
    }
</style>
