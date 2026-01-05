<script lang="ts">
    import { onDestroy, onMount } from "svelte";
    export let slide = null;
    export let after_leave = null;
    import { download_list } from "../../../../store/changeBody";
    import { forge_list, forge_support, game_log } from "../../../../store/mc";
    import { HttpGet } from "../../../../../wailsjs/go/launcher/Network";
    import MyLoadingPickaxe from "../../../../component/card/MyLoadingPickaxe.svelte";
    import { quadInOut } from "svelte/easing";
    import MySelectCard from "../../../../component/card/MySelectCard.svelte";
    import MyNormalSpan from "../../../../component/input/MyNormalSpan.svelte";
    import MyNormalButton from "../../../../component/button/MyNormalButton.svelte";
    import {
        OpenCustomURL,
        SortForgeVersion,
    } from "../../../../store/functions";
    import MyCardButton from "../../../../component/button/MyCardButton.svelte";
    import Anvil from "../../../../assets/images/Blocks/Anvil.png";
    let loading_text = "";
    let loading_state = false;
    let temp_forge_list = {};
    let temp_forge_support = [];
    onMount(async () => {
        if ($forge_support.length <= 1) {
            loading_text = "正在获取 Forge 列表";
            loading_state = false;
            let meta = await HttpGet(
                "https://maven.minecraftforge.net/net/minecraftforge/forge/maven-metadata.xml",
                "",
            );
            if (meta == "") {
                loading_text =
                    "网络环境不佳，请稍后重试，或使用 VPN 以改善网络环境";
                loading_state = true;
                return;
            }
            let parser = new DOMParser();
            let xml = parser.parseFromString(meta, "text/xml");
            let versions = xml.getElementsByTagName("version");
            let temp_list = {};
            let temp_support = [];
            for (let i = 0; i < versions.length; i++) {
                let version = versions[i].textContent;
                let mcv = version.substring(0, version.indexOf("-"));
                if (!temp_support.includes(mcv)) {
                    temp_support.push(mcv);
                }
                if (!temp_list[mcv]) {
                    temp_list[mcv] = [];
                }
                temp_list[mcv].push(version);
            }
            forge_list.set(temp_list);
            SortForgeVersion(temp_support);
            forge_support.set(temp_support);
        }
        Object.keys($forge_list).forEach((name) => {
            temp_forge_support.push({
                version: name,
                height: $forge_list[name].length * 50 + 20,
                count: $forge_list[name].length,
            });
            if (!temp_forge_list[name]) {
                temp_forge_list[name] = [];
            }
        });
        SortForgeVersion(temp_forge_support, true);
    });
    // 当任意卡片下拉的事件
    function onComp(event: CustomEvent) {
        let { height, isExpand, title } = event.detail;
        scrollStep += (isExpand ? height : -height) * 0.05;
        title = title.split("(")[0].trim();
        if (isExpand) {
            temp_forge_list[title] = $forge_list[title];
        } else {
            temp_forge_list[title] = [];
        }
    }
    let isTransitioning = true;
    function control_leave() {
        isTransitioning = true;
    }
    let f = false;
    const unsubscribe_forge_support = forge_support.subscribe((value) => {
        if (!f) {
            f = true;
            isTransitioning = true;
        } else {
            isTransitioning = !isTransitioning;
        }
    });
    function slide_opacity(node: HTMLElement) {
        return {
            duration: 200,
            easing: quadInOut,
            css: (t: number) => {
                return `opacity: ${t};`;
            },
        };
    }
    function slide_up(node: HTMLElement) {
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
    onDestroy(unsubscribe_forge_support);
    function slide_button_opacity(node: HTMLElement) {
        return {
            duration: 200,
            easing: quadInOut,
            css: (t: number) => {
                return `
                    transform: scale(${t});
                `;
            },
        };
    }
    let scrollStep = 166;
    let scrollTop = 0;
    function onDivScroll(e: Event) {
        let target = e.target as HTMLElement;
        scrollTop = target.scrollTop;
    }
    function onTopDiv() {
        let target = document.getElementsByClassName(
            "component-forge",
        )[0] as HTMLElement;
        let s = setInterval(function () {
            target.scrollTop -= scrollStep;
            scrollTop -= scrollStep;
            if (scrollTop <= scrollStep) {
                scrollTop = 0;
                target.scrollTop = 0;
                clearInterval(s);
            }
        }, 10);
    }
</script>

<div
    class="component-forge"
    in:slide
    out:slide
    on:outroend={after_leave}
    on:scroll={onDivScroll}
>
    <div style="width: 100%; height: max-content">
        <MySelectCard title="Forge 简介">
            <div class="version-all">
                <MyNormalSpan
                    >Forge 是一个 Mod 加载器，你需要先安装 Forge 才能安装各种
                    Forge 模组。</MyNormalSpan
                ><br />
                <MyNormalButton
                    style_in="width: 150px; height: 40px; border: 1px solid skyblue"
                    on:click={() => {
                        OpenCustomURL("https://files.minecraftforge.net/");
                    }}>打开官网</MyNormalButton
                >
            </div>
        </MySelectCard>
    </div>
    {#if temp_forge_support.length <= 1 && isTransitioning}
        <div
            style="display: flex; align-items: center; justify-content: center;"
            in:slide_opacity
            out:slide_opacity
            on:outroend={control_leave}
        >
            <MyLoadingPickaxe {loading_text} state={loading_state} />
        </div>
    {:else if isTransitioning}
        <div in:slide_up out:slide_up on:outroend={control_leave} style="height: auto;">
            {#each temp_forge_support as support}
                <MySelectCard
                    title={support.version + " (" + support.count + ")"}
                    isExpand={true}
                    on:comp={onComp}
                    maxHeight={support.height}
                    canExpand={true}
                >
                    <div class="version-all">
                        {#each temp_forge_list[support.version] as f}
                            <MyCardButton
                                image={Anvil}
                                title={f}
                                desc="发布于 Unknown"
                                click={() => {}}
                                buttons={[
                                    {
                                        title: "更新日志",
                                        icon: `
                                            <svg
                                                    role="img"
                                                    xmlns="http://www.w3.org/2000/svg"
                                                    viewBox="0 0 24 24"
                                                    stroke-width="1.5"
                                                    stroke-linecap="round"
                                                    stroke-linejoin="round"
                                                    fill="none" class="version-buttons-icon">
                                                <path stroke-width="2" d="M12,11 L12,17"/>
                                                <line stroke-width="2" x1="12" y1="8" x2="12" y2="8"/>
                                                <circle cx="12" cy="12" r="10"/>
                                            </svg>
                                        `,
                                        click: () => {},
                                    },
                                ]}
                            />
                        {/each}
                    </div>
                </MySelectCard>
            {/each}
        </div>
        {#if scrollTop >= 1000}
            <button
                id="topButton"
                title="回到顶部"
                on:click={onTopDiv}
                in:slide_button_opacity
                out:slide_button_opacity
                style:bottom={$game_log !== "" && $download_list.length !== 0
                    ? "160px"
                    : $game_log !== "" || $download_list.length !== 0
                      ? "90px"
                      : "20px"}
            >
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 24 24"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="3"
                >
                    <path d="M5 3l15 0M12 21l0 -13.5M12 7l7 7M12 7l-7 7" />
                </svg>
            </button>
        {/if}
    {/if}
</div>

<style>
    #topButton {
        position: fixed;
        right: 20px;
        width: 50px;
        height: 50px;
        border-radius: 50%;
        background-color: #00aaffcf;
        border: 0;
        transition: all 0.2s;
    }
    #topButton:hover {
        background-color: #0077ffcf;
    }
    #topButton svg {
        stroke: white;
        width: 25px;
        height: 25px;
        vertical-align: middle;
    }
    .component-forge {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        overflow-y: auto;
    }
    .component-forge > div {
        width: 100%;
        height: calc(100% - 166px);
    }
    .component-forge > div:last-child {
        height: calc(100% - 186px);
    }
</style>
