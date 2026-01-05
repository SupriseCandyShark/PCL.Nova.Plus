<script lang="ts">
    import {
        OpenCustomURL,
        slide_up,
        slide_opacity,
        slide_opacity_transcale,
    } from "../../../../store/functions";
    import MySelectCard from "../../../../component/card/MySelectCard.svelte";
    import MyNormalSpan from "../../../../component/input/MyNormalSpan.svelte";
    import MyNormalButton from "../../../../component/button/MyNormalButton.svelte";
    import { fabric_list, game_log } from "../../../../store/mc";
    import { download_list } from "../../../../store/changeBody";
    import { onDestroy, onMount } from "svelte";
    import MyLoadingPickaxe from "../../../../component/card/MyLoadingPickaxe.svelte";
    import { HttpGet } from "../../../../../wailsjs/go/launcher/Network";
    import MyCardButton from "../../../../component/button/MyCardButton.svelte";
    import Fabric from "../../../../assets/images/Blocks/Fabric.png";
    let loading_text = "";
    let loading_state = false;
    export let slide = null;
    export let after_leave = null;
    onMount(async () => {
        if ($fabric_list.length <= 1) {
            loading_text = "正在获取 Fabric 列表";
            loading_state = false;
            let meta = await HttpGet(
                "https://meta.fabricmc.net/v2/versions/installer",
                "",
            );
            if (meta == "") {
                loading_text =
                    "网络环境不佳，请稍后重试，或使用 VPN 以改善网络环境";
                loading_state = true;
                return;
            }
            let json = JSON.parse(meta);
            for (let i = 0; i < json.length; i++) {
                fabric_list.update((value) => [
                    ...value,
                    {
                        url: json[i].url,
                        version: json[i].version,
                        stable: json[i].stable,
                    },
                ]);
            }
        }
    });
    let isTransitioning = true;
    function control_leave() {
        isTransitioning = true;
    }
    let f = false;
    const unsubscribe_fabric_list = fabric_list.subscribe((value) => {
        if (!f) {
            f = true;
            isTransitioning = true;
        } else {
            isTransitioning = !isTransitioning;
        }
    });
    onDestroy(unsubscribe_fabric_list);
    let scrollStep = 100;
    let scrollTop = 0;
    function onDivScroll(e: Event) {
        let target = e.target as HTMLElement;
        scrollTop = target.scrollTop;
    }
    function onTopDiv() {
        let target = document.getElementsByClassName(
            "component-fabric",
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
    class="component-fabric"
    in:slide
    out:slide
    on:outroend={after_leave}
    on:scroll={onDivScroll}
>
    <div style="width: 100%; height: max-content">
        <MySelectCard title="Fabric 简介" maxHeight={150}>
            <div class="version-all">
                <MyNormalSpan
                    >Fabric Loader 是新版 Minecraft 下的轻量化 Mod
                    加载器，你需要先安装它才能安装各种 Fabric 模组</MyNormalSpan
                ><br />
                <MyNormalSpan
                    >本页面提供 Fabric
                    安装器下载，在下载后你需要手动打开安装器进行安装。</MyNormalSpan
                ><br />
                <MyNormalButton
                    style_in="width: 150px; height: 40px; border: 1px solid skyblue"
                    on:click={() => {
                        OpenCustomURL("https://www.fabricmc.net/");
                    }}>打开官网</MyNormalButton
                >
            </div>
        </MySelectCard>
    </div>
    {#if $fabric_list.length <= 1 && isTransitioning}
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
            <MySelectCard title="版本列表 ({$fabric_list.length})">
                <div class="version-all">
                    {#each $fabric_list as list}
                        <MyCardButton
                            image={Fabric}
                            title={list.version}
                            desc={list.stable ? "正式版" : "测试版"}
                            click={() => {}}
                            buttons={[]}
                        />
                    {/each}
                </div>
            </MySelectCard>
        </div>
        {#if scrollTop >= 1000}
            <button
                id="topButton"
                title="回到顶部"
                on:click={onTopDiv}
                in:slide_opacity_transcale
                out:slide_opacity_transcale
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
    .component-fabric {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        margin-bottom: -15px;
        overflow-y: auto;
    }
    .component-fabric > div {
        width: 100%;
        height: calc(100% - 203px);
    }
    .component-fabric > div:last-child {
        height: calc(100% - 223px);
    }
</style>
