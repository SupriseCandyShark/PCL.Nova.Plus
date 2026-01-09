<script lang="ts">
    import NavBar from "./views/NavBar.svelte";
    import Body from "./views/Body.svelte";
    import {
        dark_mode,
        dont_click,
        launcher_version,
        theme_mode,
        download_list,
        unlock_theme,
        current_view,
    } from "./store/changeBody";
    import { onMount } from "svelte";
    import {
        GetConfigIniPath,
        ReadConfig,
        GetOtherIniPath,
        WriteConfig,
    } from "../wailsjs/go/launcher/ReaderWriter";
    import {
        GetBackgroundImage,
        Operation,
        Version,
    } from "../wailsjs/go/launcher/MainMethod";
    import { GetUniqueAddress, CryptoUnlock } from "../wailsjs/go/main/App";
    import {
        DarkAndThemeToMain,
        OpenCustomURL,
        slide_opacity_transcale,
    } from "./store/functions";
    import MyMessageBox from "./component/card/MyMessageBox.svelte";
    import MyNormalHint from "./component/card/MyNormalHint.svelte";
    import MyInputBox from "./component/card/MyInputBox.svelte";
    import { game_log } from "./store/mc";
    import { messagebox, MSG_INFO, MSG_WARNING } from "./store/messagebox";
    import { Quit } from "../wailsjs/runtime/runtime";
    function ConvertDarkToRGB(dark: boolean, theme: number) {
        const d = DarkAndThemeToMain(dark, theme);
        return {
            dark1: d.substring(
                d.indexOf("bottom, ") + 8,
                d.indexOf("), rgb(") + 1,
            ),
            dark2: d.substring(d.indexOf("), rgb(") + 3, d.length - 1),
        };
    }
    $: ({ dark1, dark2 } = ConvertDarkToRGB($dark_mode, $theme_mode));
    function getRotate(dont: number) {
        return {
            rotateX: dont == 1 || dont == 3 ? "180deg" : "0",
            rotateY: dont == 1 || dont == 2 ? "180deg" : "0",
        };
    }
    $: ({ rotateX, rotateY } = getRotate($dont_click));
    let backImage = "";
    onMount(async () => {
        let ver = await Version();
        launcher_version.set(ver);
        let d = await ReadConfig(await GetConfigIniPath(), "Misc", "DarkMode");
        dark_mode.set(
            d == "1"
                ? true
                : d == "0"
                  ? false
                  : window.matchMedia("(prefers-color-scheme: light)").matches,
        );
        const jt = await GetUniqueAddress();
        if (jt.indexOf("ERR_") === -1) {
            const l = await ReadConfig(await GetOtherIniPath(), "Unlock", jt);
            if (await CryptoUnlock(l)) {
                unlock_theme.set(true);
                let t = Number(
                    await ReadConfig(
                        await GetOtherIniPath(),
                        "Unlock",
                        "ThemeMode",
                    ),
                );
                theme_mode.set(Number.isNaN(t) || t < 1 || t > 15 ? 1 : t);
            }
        }
        let back = await GetBackgroundImage(-1);
        if (back.length != 0) {
            backImage = `url('data:image/${back[1]};base64,${back[0]}')`;
        } else {
            backImage = ``;
        }
        if (
            (await ReadConfig(await GetOtherIniPath(), "EULA", "Agree")) !== "1"
        ) {
            if (
                (await messagebox(
                    "最终用户协议",
                    "虽然 Nova 是纯开源软件！但是你依旧需要遵守一些规定！请参考我们的最终用户协议！",
                    MSG_INFO,
                    ["阅读并同意", "不同意"],
                )) === 0
            ) {
                OpenCustomURL(
                    "https://github.com/SupriseCandyShark/PCL.Nova.Plus/blob/main/docs/GUIDELINES.md",
                );
                await WriteConfig(
                    await GetOtherIniPath(),
                    "EULA",
                    "Agree",
                    "1",
                );
            } else {
                Quit();
            }
        }
        if ((await Operation()) === 2) {
            if (
                (await ReadConfig(
                    await GetOtherIniPath(),
                    "Arm64",
                    "CheckWarning",
                )) !== "1"
            ) {
                let c = await messagebox(
                    "OpenGL 兼容层提示",
                    "检测到您目前处于 Windows arm64 环境！如果您的处理器是 高通 系列处理器，您可能需要下载 OpenGL 兼容层才能正常使用游戏。",
                    MSG_WARNING,
                    ["我已知晓，不再提示", "我已知晓，并进入网站下载"],
                );
                if (c === 0) {
                    await WriteConfig(
                        await GetOtherIniPath(),
                        "Arm64",
                        "CheckWarning",
                        "1",
                    );
                }
                OpenCustomURL("https://apps.microsoft.com/detail/9nqpsl29bfff");
            }
        }
    });
</script>

<div id="all" style="--rotate-y: {rotateY}; --rotate-x: {rotateX}">
    <NavBar />
    <!--背景颜色，在下层-->
    <div id="back" style="--dark-1: {dark1}; --dark-2: {dark2}">
        <!--背景图片，在上层-->
        <div id="main" style="--back-image: {backImage}">
            <Body />
        </div>
    </div>
    {#if $download_list.length !== 0 && $current_view !== "downloadFile"}
        <button
            class="topButton"
            title="下载文件"
            in:slide_opacity_transcale
            out:slide_opacity_transcale
            style:bottom={$game_log !== "" ? "90px" : "20px"}
            on:click={() => {
                current_view.set("downloadFile");
            }}
        >
            <svg
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 24 24"
                fill="none"
                stroke="currentColor"
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
            >
                <path
                    d="M4 17v2a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-2M7 11l5 5l5-5m-5-7v12"
                />
            </svg>
        </button>
    {/if}
    {#if $game_log !== "" && $current_view !== "log"}
        <button
            class="topButton"
            title="游戏日志"
            in:slide_opacity_transcale
            out:slide_opacity_transcale
            style:bottom="20px"
            on:click={() => {
                current_view.set("log");
            }}
        >
            <svg
                xmlns="http://www.w3.org/2000/svg"
                width="32"
                height="32"
                viewBox="0 0 14 14"
                fill="none"
                stroke="currentColor"
                stroke-linecap="round"
                stroke-linejoin="round"
            >
                <rect width="13" height="13" x=".5" y=".5" rx="1" />
                <path d="M.5 4h13M4 7l1.5 1.5L4 10m4.5-1.5h2" />
            </svg>
        </button>
    {/if}
    <MyMessageBox />
    <MyNormalHint />
    <MyInputBox />
</div>

<style>
    #all {
        position: absolute;
        display: flex;
        flex-direction: column;
        height: 100%;
        width: 100%;
        overflow: hidden;
        transform: rotateY(var(--rotate-y)) rotateX(var(--rotate-x));
        transition: transform 5s linear;
    }
    #back {
        width: 100%;
        background: linear-gradient(
            to left bottom,
            var(--dark-1),
            var(--dark-2)
        );
        height: calc(100% - 56px);
        transition: all 0.2s;
    }
    #main {
        position: relative;
        width: 100%;
        height: 100%;
        transition: all 0.2s;
        background-size: cover;
        background-repeat: no-repeat;
        background-position: 50% 50%;
        background-image: var(--back-image);
    }
    .topButton {
        position: fixed;
        right: 20px;
        width: 50px;
        height: 50px;
        border-radius: 50%;
        background-color: #00aaffcf;
        border: 0;
        transition: all 0.2s;
    }

    .topButton:hover {
        background-color: #0077ffcf;
    }

    .topButton svg {
        stroke: white;
        width: 25px;
        height: 25px;
        vertical-align: middle;
    }
</style>
