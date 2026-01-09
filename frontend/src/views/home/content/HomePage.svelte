<script lang="ts">
    import MyNormalSpan from "../../../component/input/MyNormalSpan.svelte";
    import MySelectCard from "../../../component/card/MySelectCard.svelte";
    import MyNormalButton from "../../../component/button/MyNormalButton.svelte";
    import MyTextButton from "../../../component/button/MyTextButton.svelte";
    import { GetCurrentExeDir } from "../../../../wailsjs/go/launcher/ReaderWriter.js";
    import { onMount } from "svelte";
    import Anvil from "../../../assets/images/Blocks/Anvil.png";
    import CobbleStone from "../../../assets/images/Blocks/CobbleStone.png";
    import CommandBlock from "../../../assets/images/Blocks/CommandBlock.png";
    import Egg from "../../../assets/images/Blocks/Egg.png";
    import Fabric from "../../../assets/images/Blocks/Fabric.png";
    import GoldBlock from "../../../assets/images/Blocks/GoldBlock.png";
    import Grass from "../../../assets/images/Blocks/Grass.png";
    import GrassPath from "../../../assets/images/Blocks/GrassPath.png";
    import NeoForge from "../../../assets/images/Blocks/NeoForge.png";
    import RedstoneBlock from "../../../assets/images/Blocks/RedstoneBlock.png";
    import RedstoneLampOn from "../../../assets/images/Blocks/RedstoneLampOn.png";
    import RedstoneLampOff from "../../../assets/images/Blocks/RedstoneLampOff.png";
    import MyLoadingPickaxe from "../../../component/card/MyLoadingPickaxe.svelte";
    import {
        homepage_cache,
        homepage_value,
    } from "../../../store/changeBody.js";
    import { messagebox, showHint } from "../../../store/messagebox.js";
    import { OpenCustomURL } from "../../../store/functions";
    interface controlStruct {
        type?: string;
        name?: string;
        title?: string;
        style?: string;
        children?: Array<any>;
        isExpand?: string;
        canExpand?: string;
        content?: string;
        width?: string;
        height?: string;
        viewBox?: string;
        src?: string;
        alt?: string;
        path?: string;
        color?: string;
        event?: string;
        id?: string;
    }
    export let control: controlStruct = {};
    let exePath = "";
    onMount(async () => {
        exePath = await GetCurrentExeDir();
    });
    async function onButtonClick(ev: string) {
        if (ev == undefined || ev == "") return;
        let event_split = ev.split("|");
        let name = event_split[0];
        if (name === "messagebox") {
            await messagebox(
                event_split[1],
                event_split[2],
                Number(event_split[3]),
            );
        } else if (name === "openurl") {
            OpenCustomURL(event_split[1]);
        } else if (name === "refresh") {
            showHint("正在刷新主页中~");
            homepage_cache.set("");
        }
    }
    function addValue(id: string, children: Array<any>): string {
        if (children.length != 1 || children[0].name != "text") return "";
        homepage_value.update((value) => {
            value[id] = children[0].content;
            return value;
        });
        return "";
    }
    function replaceValue(raw: string): string {
        if (raw == undefined || raw == "") return "";
        Object.entries($homepage_value).forEach((k) => {
            raw = raw.replaceAll(`\${${k[0]}}`, k[1].toString());
        });
        return raw;
    }
</script>

{#if control.name === "MySpan"}
    <MyNormalSpan title={control.title} style_in={replaceValue(control.style)}>
        {#if control.children}
            {#each control.children as child}
                <svelte:self control={child} />
            {/each}
        {/if}
    </MyNormalSpan>
{:else if control.name === "MyCard"}
    <MySelectCard
        title={control.title}
        isExpand={control.isExpand === "True"}
        canExpand={control.canExpand === "True"}
    >
        <div class="version-all">
            {#if control.children}
                {#each control.children as child}
                    <svelte:self control={child} />
                {/each}
            {/if}
        </div>
    </MySelectCard>
{:else if control.name === "text"}
    {control.content
        ? replaceValue(control.content.replaceAll("{path}", exePath))
        : ""}
{:else if control.name === "MyDiv"}
    <div
        title={control.title}
        style="margin-top: 5px; {replaceValue(control.style)}"
    >
        {#if control.children}
            {#each control.children as child}
                <svelte:self control={child} />
            {/each}
        {/if}
    </div>
{:else if control.name === "MyButton"}
    {#if control.type === "label"}
        <MyTextButton
            title={control.title}
            hover_style_in={replaceValue(control["hov-style"])}
            active_style_in={replaceValue(control["active-style"])}
            style_in={replaceValue(control.style)}
            on:click={() => {
                onButtonClick(control.event);
            }}
        >
            {#if control.children}
                {#each control.children as child}
                    <svelte:self control={child} />
                {/each}
            {/if}
        </MyTextButton>
    {:else}
        <MyNormalButton
            title={control.title}
            style_in="width: {control.width}; height: {control.height}; {control.color
                ? `border: 1px solid ${control.color};`
                : ''}"
            on:click={() => onButtonClick(control.event)}
        >
            {#if control.children}
                {#each control.children as child}
                    <svelte:self control={child} />
                {/each}
            {/if}
        </MyNormalButton>
    {/if}
{:else if control.name === "MySvg"}
    <svg
        xmlns="http://www.w3.org/2000/svg"
        role="img"
        viewBox={control.viewBox}
        style={control.style}
    >
        <path d={control.path} />
    </svg>
{:else if control.name === "MyImg"}
    <img
        src={[
            CommandBlock,
            CobbleStone,
            GoldBlock,
            Grass,
            GrassPath,
            Anvil,
            RedstoneBlock,
            RedstoneLampOn,
            RedstoneLampOff,
            Egg,
            Fabric,
            NeoForge,
        ][
            [
                "CommandBlock",
                "Cobblestone",
                "GoldBlock",
                "Grass",
                "GrassPath",
                "Anvil",
                "RedstoneBlock",
                "RedstoneLampOn",
                "RedstoneLampOff",
                "Egg",
                "Fabric",
                "NeoForge",
            ].findIndex((item) => item === control.src)
        ]}
        alt={control.alt}
        style={control.style}
    />
{:else if control.name === "MyValue"}
    {addValue(control.id ? control.id : "", control.children)}
{:else if control.name === "MyLoadingPickaxe"}
    <div
        style="width: 100%; height: auto; display: flex; flex-direction: column; align-items: center;"
    >
        <MyLoadingPickaxe
            style_in={`width: ${control.width ?? "auto"}; height: ${control.height ?? "auto"}`}
            loading_text={control.title}
            state={control.state === "true" ? true : false}
        ></MyLoadingPickaxe>
    </div>
{/if}

<style>
</style>
