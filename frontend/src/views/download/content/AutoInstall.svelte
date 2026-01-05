<script lang="ts">
    import { onDestroy, onMount } from "svelte";
    import MySelectCard from "../../../component/card/MySelectCard.svelte";
    import { HttpGet } from "../../../../wailsjs/go/launcher/Network";
    import MyLoadingPickaxe from "../../../component/card/MyLoadingPickaxe.svelte";
    import {
        slide_opacity_transcale,
        slide_opacity,
        slide_up,
    } from "../../../store/functions";
    import {
        game_log,
        latest_release,
        latest_snapshot,
        mc_list_fool,
        mc_list_ok,
        mc_list_old_alpha,
        mc_list_old_beta,
        mc_list_release,
        mc_list_snapshot,
    } from "../../../store/mc";
    import Grass from "../../../assets/images/Blocks/Grass.png";
    import CommandBlock from "../../../assets/images/Blocks/CommandBlock.png";
    import CobbleStone from "../../../assets/images/Blocks/CobbleStone.png";
    import MyCardButton from "../../../component/button/MyCardButton.svelte";
    import GoldBlock from "../../../assets/images/Blocks/GoldBlock.png"
    import { download_list } from "../../../store/changeBody";

    export let slide = null;
    export let after_leave = null;
    let loading_text = "";
    let loading_state = false;
    let temp_list = [
        {
            list: [],
            version: 0,
            height: 0,
        },
        {
            list: [],
            version: 1,
            height: 0,
        },
        {
            list: [],
            version: 2,
            height: 0,
        },
        {
            list: [],
            version: 3,
            height: 0,
        },
        {
            list: [],
            version: 4,
            height: 0,
        },
        {
            list: [],
            version: 5,
            height: 0,
        }
    ];
    function isAprilFirst(dateString: string) {
        try {
            const date = new Date(dateString);
            if (isNaN(date.getTime())) 
                return false;
            const month = date.getMonth();
            const day = date.getDate();
            return month === 3 && day === 1;
        } catch (error) {
            return false;
        }
    }
    onMount(async () => {
        if (!$mc_list_ok) {
            loading_text = "正在获取版本列表";
            loading_state = false;
            let meta = await HttpGet(
                "https://piston-meta.mojang.com/mc/game/version_manifest_v2.json",
                "",
            );
            if (meta == "") {
                loading_text =
                    "网络环境不佳，请稍后重试，或使用 VPN 以改善网络环境";
                loading_state = true;
                return;
            }
            let json = JSON.parse(meta);
            let ver = json.versions!;
            for (let i = 0; i < ver.length; i++) {
                let obj = ver[i];
                if (obj.id! == json.latest!.release!) latest_release.set(obj);
                if (obj.id! == json.latest!.snapshot!) latest_snapshot.set(obj);
                if (isAprilFirst(obj.releaseTime!))
                    mc_list_fool.update((value) => [...value, obj])
                else if (obj.type! == "release")
                    mc_list_release.update((value) => [...value, obj]);
                else if (obj.type! == "snapshot")
                    mc_list_snapshot.update((value) => [...value, obj]);
                else if (obj.type! == "old_beta")
                    mc_list_old_beta.update((value) => [...value, obj]);
                else if (obj.type! == "old_alpha")
                    mc_list_old_alpha.update((value) => [...value, obj]);
            }
            mc_list_ok.set(true);
        }
        temp_list[0].height = $mc_list_release.length * 50 + 20;
        temp_list[1].height = $mc_list_snapshot.length * 50 + 20;
        temp_list[2].height = $mc_list_old_beta.length * 50 + 20;
        temp_list[3].height = $mc_list_old_alpha.length * 50 + 20;
        temp_list[4].height = $mc_list_fool.length * 50 + 20;
    });
    let isTransitioning = true;

    function control_leave() {
        isTransitioning = true;
    }

    let f = false;
    const unsubscribe_mc_list_ok = mc_list_ok.subscribe((value) => {
        if (!f) {
            f = true;
            isTransitioning = true;
        } else {
            isTransitioning = !isTransitioning;
        }
    });
    onDestroy(unsubscribe_mc_list_ok);
    // 以下均是来自 回到顶部 按钮的函数
    let scrollStep = 10;
    let scrollTop = 0;

    function onDivScroll(e: Event) {
        let target = e.target as HTMLElement;
        scrollTop = target.scrollTop;
    }

    function onTopDiv() {
        let target = document.getElementsByClassName(
            "component-auto_install",
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

    function onComp(event: CustomEvent) {
        let { height, isExpand, title } = event.detail;
        scrollStep += (isExpand ? height : -height) * 0.05;
        if (title == "正式版") {
            temp_list[0].list = isExpand ? $mc_list_release : [];
        } else if (title == "快照版") {
            temp_list[1].list = isExpand ? $mc_list_snapshot : [];
        } else if (title == "远古 Beta 版") {
            temp_list[2].list = isExpand ? $mc_list_old_beta : [];
        } else if (title == "远古 Alpha 版") {
            temp_list[3].list = isExpand ? $mc_list_old_alpha : [];
        } else if (title == "愚人节版") {
            temp_list[4].list = isExpand ? $mc_list_fool : [];
        }
    }
</script>

<div
    class="component-auto_install"
    in:slide
    out:slide
    on:outroend={after_leave}
    on:scroll={onDivScroll}
>
    {#if !$mc_list_ok && isTransitioning}
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
            <MySelectCard title="最新版本">
                <div class="version-all">
                    <MyCardButton
                        image={Grass}
                        title={$latest_release["id"]}
                        desc="最新正式版，发布于：{$latest_release[
                            'releaseTime'
                        ]
                            .substring(
                                0,
                                $latest_release['releaseTime'].length - 9,
                            )
                            .replaceAll('T', ' ')
                            .replaceAll('-', '/')}"
                        click={() => {}}
                        buttons={[
                            {
                                title: "下载服务端",
                                icon: `
                                        <svg
                                                xmlns="http://www.w3.org/2000/svg"
                                                viewBox="0 0 24 24"
                                                fill="none"
                                                stroke-width="1.5"
                                                stroke-linecap="round">
                                            <g>
                                                <path
                                                        d="M2 12c0-4.714 0-7.071 1.464-8.536C4.93 2 7.286 2 12 2s7.071 0 8.535 1.464C22 4.93 22 7.286 22 12s0 7.071-1.465 8.535C19.072 22 16.714 22 12 22s-7.071 0-8.536-1.465C2 19.072 2 16.714 2 12Z"/>
                                                <path
                                                        d="M2 13h3.16c.905 0 1.358 0 1.756.183s.692.527 1.281 1.214l.606.706c.589.687.883 1.031 1.281 1.214s.85.183 1.756.183h.32c.905 0 1.358 0 1.756-.183s.692-.527 1.281-1.214l.606-.706c.589-.687.883-1.031 1.281-1.214S17.934 13 18.84 13H22M8 7h8m-6 3.5h4"/>
                                            </g>
                                        </svg>
                                    `,
                                click: () => {},
                            },
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
                    <MyCardButton
                        image={CommandBlock}
                        title={$latest_snapshot["id"]}
                        desc="最新快照版，发布于：{$latest_snapshot[
                            'releaseTime'
                        ]
                            .substring(
                                0,
                                $latest_snapshot['releaseTime'].length - 9,
                            )
                            .replaceAll('T', ' ')
                            .replaceAll('-', '/')}"
                        click={() => {}}
                        buttons={[
                            {
                                title: "下载服务端",
                                icon: `
                                        <svg
                                                xmlns="http://www.w3.org/2000/svg"
                                                viewBox="0 0 24 24"
                                                fill="none"
                                                stroke-width="1.5"
                                                stroke-linecap="round">
                                            <g>
                                                <path
                                                        d="M2 12c0-4.714 0-7.071 1.464-8.536C4.93 2 7.286 2 12 2s7.071 0 8.535 1.464C22 4.93 22 7.286 22 12s0 7.071-1.465 8.535C19.072 22 16.714 22 12 22s-7.071 0-8.536-1.465C2 19.072 2 16.714 2 12Z"/>
                                                <path
                                                        d="M2 13h3.16c.905 0 1.358 0 1.756.183s.692.527 1.281 1.214l.606.706c.589.687.883 1.031 1.281 1.214s.85.183 1.756.183h.32c.905 0 1.358 0 1.756-.183s.692-.527 1.281-1.214l.606-.706c.589-.687.883-1.031 1.281-1.214S17.934 13 18.84 13H22M8 7h8m-6 3.5h4"/>
                                            </g>
                                        </svg>
                                    `,
                                click: () => {},
                            },
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
                </div>
            </MySelectCard>
            {#each temp_list as temp}
                <MySelectCard
                    title={temp.version === 0
                        ? "正式版"
                        : temp.version === 1
                          ? "快照版"
                          : temp.version === 2
                            ? "远古 Beta 版"
                            : temp.version === 3
                              ? "远古 Alpha 版"
                              : temp.version === 4
                                ? "愚人节版"
                                : temp.version === 5
                                  ? "Nova 特供" : ""}
                    on:comp={onComp}
                    maxHeight={temp.height}
                    isExpand={true}
                    canExpand={true}
                    in_style={temp.version === 5 ? 'margin-bottom: 15px;' : ''}
                >
                    <div class="version-all">
                        {#each temp.list as list}
                            <MyCardButton
                                image={temp.version === 0
                                    ? Grass
                                    : temp.version === 1
                                      ? CommandBlock
                                      : temp.version === 2
                                        ? CobbleStone
                                        : temp.version === 3
                                          ? CobbleStone
                                          : temp.version === 4
                                            ? GoldBlock : ""}
                                title={list["id"]}
                                desc={list["releaseTime"]
                                    .substring(
                                        0,
                                        list["releaseTime"].length - 9,
                                    )
                                    .replaceAll("T", " ")
                                    .replaceAll("-", "/")}
                                click={() => {}}
                                buttons={[
                                    {
                                        title: "下载服务端",
                                        icon: `
                                             <svg
                                                     xmlns="http://www.w3.org/2000/svg"
                                                     viewBox="0 0 24 24"
                                                     fill="none"
                                                     stroke-width="1.5"
                                                     stroke-linecap="round">
                                                 <g>
                                                     <path
                                                             d="M2 12c0-4.714 0-7.071 1.464-8.536C4.93 2 7.286 2 12 2s7.071 0 8.535 1.464C22 4.93 22 7.286 22 12s0 7.071-1.465 8.535C19.072 22 16.714 22 12 22s-7.071 0-8.536-1.465C2 19.072 2 16.714 2 12Z"/>
                                                     <path
                                                             d="M2 13h3.16c.905 0 1.358 0 1.756.183s.692.527 1.281 1.214l.606.706c.589.687.883 1.031 1.281 1.214s.85.183 1.756.183h.32c.905 0 1.358 0 1.756-.183s.692-.527 1.281-1.214l.606-.706c.589-.687.883-1.031 1.281-1.214S17.934 13 18.84 13H22M8 7h8m-6 3.5h4"/>
                                                 </g>
                                             </svg>
                                         `,
                                        click: () => {},
                                    },
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

    .component-auto_install {
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        overflow-y: auto;
    }

    .component-auto_install > div {
        width: 100%;
        height: 100%;
    }

    .component-auto_install > div:last-child {
        height: calc(100% - 20px);
    }
</style>
