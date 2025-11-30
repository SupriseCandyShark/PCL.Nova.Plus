<script lang="ts">
    import {
        select_account,
        current_account_index,
        current_account_page,
    } from "../../../store/changeBody";
    import {
        GetOtherIniPath,
        ReadConfig,
        WriteConfig,
    } from "../../../../wailsjs/go/launcher/ReaderWriter";
    import {
        GetAccountConfig,
        SetAccountConfig,
    } from "../../../../wailsjs/go/launcher/AccountMethod";
    import { onMount } from "svelte";
    import MyNormalSpan from "../../../component/input/MyNormalSpan.svelte";
    import MyNormalButton from "../../../component/button/MyNormalButton.svelte";
    import { launcher } from "../../../../wailsjs/go/models";
    import { showHint } from "../../../store/messagebox";
    import MyRadioButton from "../../../component/button/MyRadioButton.svelte";

    export let opacity = null;
    export let after_leave = null;

    async function selectAccount(index: number) {
        if ($current_account_index != index) {
            current_account_index.set(index);
            await WriteConfig(
                await GetOtherIniPath(),
                "Account",
                "SelectAccount",
                index.toString(),
            );
        }
    }
    async function deleteAccount(index: number) {
        if (index < $current_account_index) {
            current_account_index.update((value) => value - 1);
        } else if (index == $current_account_index) {
            current_account_index.set(-1);
        }
        select_account.update((value) => {
            value.splice(index, 1);
            return value;
        });
        await SetAccountConfig(
            launcher.AccountList.createFrom({ accounts: $select_account }),
        );
    }
    onMount(async () => {
        let config = await GetAccountConfig();
        select_account.set([])
        select_account.update((account) => {
            for(let i = 0; i < config.data.accounts.length; i++) {
                let acc = config.data.accounts[i];
                let splitSkin = acc.head_skin.split("**")
                console.log(splitSkin)
                account.push(acc);
                account[i].innerSkin = splitSkin[0]
                if (splitSkin.length === 1) {
                    account[i].outerSkin = ''
                }else{
                    account[i].outerSkin = splitSkin[1]
                }
            }
            // Object.assign(account, config.data.accounts);
            return account;
        });
        let index = await ReadConfig(
            await GetOtherIniPath(),
            "Account",
            "SelectAccount",
        );
        if (!Number(index)) {
            await WriteConfig(
                await GetOtherIniPath(),
                "Account",
                "SelectAccount",
                "0",
            );
        }
        current_account_index.set(Number(index));
    });
</script>

<div class="component" in:opacity out:opacity on:outroend={after_leave}>
    <div id="main-style">
        {#if $select_account.length <= 0}
            <svg
                role="img"
                xmlns="http://www.w3.org/2000/svg"
                viewBox="0 0 24 24"
                stroke-width="1.5"
                stroke-linecap="round"
                stroke-linejoin="round"
                fill="none"
                id="nologin-img"
            >
                <path
                    d="M4,20 C4,17 8,17 10,15 C11,14 8,14 8,9 C8,5.667 9.333,4 12,4 C14.667,4 16,5.667 16,9 C16,14 13,14 14,15 C16,17 20,17 20,20"
                />
            </svg>
            <br />
        {/if}
        <div id="all-account">
            <div id="account-list">
                {#each $select_account as account, index}
                    <div
                        class="a-account"
                        title="玩家 ID: {account.name}&#13;UUID: {account.uuid}"
                        on:click={() => selectAccount(index)}
                        on:keydown={(e) => {
                            e.preventDefault();
                        }}
                        style={$current_account_index === index
                            ? "cursor: default;"
                            : "cursor: pointer"}
                    >
                        <MyRadioButton
                            isChecked={index === $current_account_index}
                            style_in="margin-left: 5px"
                        />
                        <div class="a-avatar">
                            <img alt="头像内层" src="data:img/png;base64,{account.innerSkin}" class="inner-skin" style={account.outerSkin !== '' ? "width: 36px; height: 36px; top: 2px; left: 2px; filter: grayscale(0.13);" : 'width: 40px; height: 40px; top: 0; left: 0;'}/>
                            {#if account.outerSkin !== ''}
                                <img alt="头像外层" src="data:img/png;base64,{account.outerSkin}" class="outer-skin" />
                            {/if}
                        </div>
                        <div class="info" style="pointer-events: none">
                            <MyNormalSpan>{account.name}</MyNormalSpan>
                            <MyNormalSpan
                                style_in="font-size: 13px; color: gray;"
                                >{account.type === "Offline"
                                    ? "离线登录"
                                    : account.type === "Microsoft"
                                      ? "正版登录"
                                      : "第三方登录"}</MyNormalSpan
                            >
                        </div>
                        <button
                            class="a-delete cursor-pointer"
                            on:click|stopPropagation={() =>
                                deleteAccount(index)}
                        >
                            <svg
                                xmlns="http://www.w3.org/2000/svg"
                                viewBox="0 0 24 24"
                                fill="none"
                                stroke-linecap="round"
                                stroke-width="1.5"
                            >
                                <path
                                    d="M9.17 4a3.001 3.001 0 0 1 5.66 0m5.67 2h-17m15.333 2.5l-.46 6.9c-.177 2.654-.265 3.981-1.13 4.79s-2.196.81-4.856.81h-.774c-2.66 0-3.991 0-4.856-.81c-.865-.809-.954-2.136-1.13-4.79l-.46-6.9M9.5 11l.5 5m4.5-5l-.5 5"
                                />
                            </svg>
                        </button>
                    </div>
                {/each}
            </div>
            <div id="control">
                <MyNormalButton
                    style_in="width: calc(50% - 5px); height: 35px; border: 1px solid skyblue;"
                    on:click={() => current_account_page.set(false)}
                    >添加新账号</MyNormalButton
                >
                <MyNormalButton
                    style_in="width: calc(50% - 5px); height: 35px; border: 1px solid skyblue;"
                    on:click={() =>
                        showHint(
                            "目前修改选中账号暂时还没有做好😭，请敬请期待吧！",
                        )}>修改选中账号</MyNormalButton
                >
            </div>
        </div>
    </div>
</div>

<style>
    .component {
        height: 100%;
        width: 100%;
        display: flex;
        justify-content: center;
        align-items: center;
    }
    #main-style {
        display: flex;
        flex-direction: column;
        justify-content: center;
        align-items: center;
        width: calc(100% - 50px);
        height: calc(100vh - 270px);
    }
    #nologin-img {
        stroke: rgb(0, 186, 254);
        width: 80px;
        height: 80px;
        image-rendering: pixelated;
    }
    #all-account {
        width: calc(100% - 40px);
        max-height: 100%;
        display: flex;
        flex-direction: column;
        align-items: center;
        overflow: auto;
    }
    #account-list {
        display: flex;
        flex-direction: column;
        width: 100%;
        overflow-y: auto;
    }
    #control {
        width: 100%;
        height: 40px;
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-top: 10px;
    }
    /* 以下是单个账号控件的样式 */
    .a-account {
        position: relative;
        transition: all 0.2s;
        border-radius: 10px;
        height: 50px;
        width: calc(100% - 2px);
        flex-shrink: 0;
        display: flex;
        flex-direction: row;
        align-items: center;
    }
    .a-avatar {
        width: 40px;
        height: 40px;
        /* box-shadow: 0 0 5px gray; */
        margin-left: 5px;
        position: relative;
    }
    .a-avatar img {
        image-rendering: pixelated;
        position: absolute;
    }
    .a-avatar .inner-skin {
        box-shadow: 0 0 3px gray;
    }
    .a-avatar .outer-skin {
        top: 0;
        left: 0;
        width: 40px;
        height: 40px;
    }
    .a-account:hover {
        background-color: rgba(128, 128, 128, 0.5);
    }
    .a-account:active {
        transform: scale(98%);
    }
    .info {
        display: flex;
        flex-direction: column;
        margin-left: 10px;
    }
    .a-delete {
        position: absolute;
        top: 10px;
        right: 10px;
        width: 30px;
        height: 30px;
        background-color: transparent;
        border: 0;
        border-radius: 50%;
        stroke: gray;
        transition: all 0.2s;
    }
    .a-delete:hover {
        background-color: rgba(0, 0, 0, 0.2);
        stroke: red;
    }
    .a-delete svg {
        width: 20px;
        height: 20px;
        vertical-align: middle;
    }
</style>
