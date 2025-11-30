<!--
    仅用于在 ExtensionXXX 中使用到的控件~用于显示一个类似于 CardButton 那样的按钮~
-->
<script lang="ts">
    import { createEventDispatcher } from "svelte";
    import MyNormalSpan from "../input/MyNormalSpan.svelte";
    import MyCodeSpan from "../input/MyCodeSpan.svelte";

    export let image_url = "";
    export let mod_name = "";
    export let mod_id = "";
    export let libraries = [];
    export let description = "";
    export let mod_type = "";
    export let download_count = "";
    export let update_date = "";
    export let mod_source = "";
    const dispatch = createEventDispatcher();
    function onButtonClick() {
        dispatch("click");
    }
</script>

<div on:click={onButtonClick} on:keydown|preventDefault class="comp">
    <img src={image_url} alt="Logo" class="comp-image" />
    <div class="comp-desc">
        <div class="comp-header">
            <MyNormalSpan style_in="font-size: 16px; vertical-align: middle;"
                >{mod_name}
            </MyNormalSpan><span class="comp-description">| {mod_id}</span>
        </div>
        <div class="comp-center">
            {#each libraries as lib, index}
                <MyCodeSpan
                    style_in={'user-select: none;' + (index === 0 ? 'margin-left: 1px;' : 'margin-left: 5px;')}>{lib}</MyCodeSpan>
            {/each}
            <span class="comp-description">{description}</span>
        </div>
        <div class="comp-footer">
            <div class="comp-description">
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 16 14"
                    fill="currentColor"
                    style="width: 16px; height: 14px; vertical-align: middle; margin-right: 5px;"
                >
                    <path
                        d="M15.7633 6.12296L12.7153 0.877039C12.3997 0.33342 11.8178 0 11.1867 0L4.81514 0C4.184 0 3.60212 0.335232 3.28655 0.877039L0.236676 6.12296C-0.0788919 6.66658 -0.0788919 7.33523 0.236676 7.87704L3.28655 13.123C3.60212 13.6666 4.184 14 4.81514 14L11.1849 14C11.816 14 12.3979 13.6648 12.7134 13.123L15.7633 7.87704C16.0789 7.33523 16.0789 6.66658 15.7633 6.12296ZM14.0377 7.87704L11.8507 11.6389C11.5351 12.1825 10.9532 12.5159 10.3221 12.5159L5.67793 12.5159C5.0468 12.5159 4.46491 12.1807 4.14935 11.6389L1.96409 7.87704C1.64852 7.33342 1.64852 6.66477 1.96409 6.12296L4.14935 2.36293C4.46491 1.81931 5.0468 1.48589 5.67793 1.48589L10.3221 1.48589C10.9532 1.48589 11.5351 1.82112 11.8507 2.36293L14.0359 6.12296C14.3515 6.66658 14.3515 7.33523 14.0377 7.87704ZM8.13681 3.54077C8.0912 3.53896 8.0456 3.53715 8 3.53715C5.81292 3.53715 4.09645 5.53585 4.60172 7.7955C4.88993 9.08025 5.90595 10.0896 7.19922 10.3759C7.51844 10.4465 7.83218 10.4737 8.13681 10.461C9.9992 10.3904 11.4858 8.86824 11.4858 7C11.4877 5.13357 9.9992 3.61144 8.13681 3.54077ZM9.38813 8.41703C9.05068 8.74321 8.60377 8.94072 8.13681 8.97334L8.02736 8.97877C7.47649 8.98602 6.95845 8.7722 6.57356 8.38079C6.21786 8.01657 6.01357 7.52912 6.00992 7.02174C6.00445 6.48537 6.21239 5.98162 6.59363 5.6029C6.96939 5.22961 7.46919 5.02304 8.00182 5.02304C8.04743 5.02304 8.09303 5.02485 8.13863 5.02848C8.62566 5.06109 9.07804 5.26767 9.42461 5.6174C9.78214 5.97981 9.98643 6.46725 9.99373 6.97463C9.99738 7.52187 9.78396 8.0365 9.38813 8.41703Z"
                    />
                </svg>{mod_type}
            </div>
            <div class="comp-description">
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 13 13"
                    fill="currentColor"
                    style="width: 13px; height: 13px; vertical-align: middle; margin-right: 5px;"
                >
                    <path
                        d="M12.5726 7.86335L11.7593 7.86335C11.5525 7.86335 11.3595 8.04237 11.3595 8.26271L11.3595 10.9619C11.3595 11.1684 11.1803 11.3612 10.9597 11.3612L2.01273 11.3612C1.80594 11.3612 1.61294 11.1822 1.61294 10.9619L1.61294 8.26271C1.61294 8.05614 1.43372 7.86335 1.21315 7.86335L0.399788 7.86335C0.193001 7.86335 0 8.04237 0 8.26271L0 11.9258C0 12.518 0.482503 13 1.07529 13L11.9247 13C12.5175 13 13 12.518 13 11.9258L13 8.27648C13 8.06991 12.8208 7.87712 12.6002 7.87712L12.5726 7.86335ZM6.18982 9.6536C6.34146 9.80509 6.58961 9.80509 6.75504 9.6536L10.4083 6.00424C10.5599 5.85275 10.5599 5.60487 10.4083 5.43962L9.84305 4.875C9.69141 4.72352 9.44327 4.72352 9.27784 4.875L7.7614 6.38983C7.60976 6.54131 7.30647 6.43114 7.30647 6.2108L7.30647 0.399364C7.2789 0.192797 7.05832 0 6.86532 0L6.05196 0C5.84517 0 5.65217 0.179025 5.65217 0.399364L5.65217 6.14195C5.65217 6.37606 5.36267 6.48623 5.19724 6.32097L3.68081 4.80614C3.52916 4.65466 3.28102 4.65466 3.11559 4.80614L2.5228 5.42585C2.37116 5.57733 2.37116 5.82521 2.5228 5.99047L6.18982 9.6536Z"
                    />
                </svg>{download_count}
            </div>
            <div class="comp-description">
                <svg
                    viewBox="0 0 1024 1024"
                    fill="currentColor"
                    style="width: 15px; height: 15px; vertical-align: middle; margin-right: 5px;"
                >
                    <path
                        d="M810 403l-39 39c-10 10-28 10-39 0l-162-162v412c0 15-12 28-28 28h-56c-15 0-28-12-28-28v-412l-162 162c-10 10-28 10-39 0l-39-39c-10-10-10-28 0-39l278-278c10-10 28-10 39 0l278 278c10 10 10 28-0 39zM568 946h-112v-112h112v112z"
                    />
                </svg>{update_date}
            </div>
            <div class="comp-description">
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    stroke="currentColor"
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    viewBox="0 0 24 24"
                    fill="none"
                    style="width: 13px; height: 13px; vertical-align: middle; margin-right: 5px;"
                >
                    <circle cx="12" cy="12" r="10" />
                    <path
                        d="M2 12h20M12 2a15.3 15.3 0 0 1 4 10 15.3 15.3 0 0 1-4 10 15.3 15.3 0 0 1-4-10 15.3 15.3 0 0 1 4-10"
                    />
                </svg>{mod_source}
            </div>
        </div>
    </div>
</div>

<style>
    .comp {
        width: 100%;
        height: 80px;
        transition: all 0.2s;
        border-radius: 10px;
        display: flex;
        align-items: center;
        cursor: pointer;
    }
    .comp:hover {
        background-color: rgba(100, 100, 156, 0.2);
    }
    .comp:active {
        transform: scale(98%);
        background-color: rgba(100, 100, 156, 0.3);
    }
    .comp-image {
        width: 70px;
        height: 70px;
        margin: 0 5px 0 20px;
    }
    .comp-desc {
        width: calc(100% - 120px);
        display: flex;
        justify-content: space-between;
        flex-direction: column;
        height: 80px;
        margin: 0 20px 0 5px;
    }
    .comp-description {
        font-size: 12px;
        color: gray;
        vertical-align: middle;
    }
    .comp-footer {
        display: flex;
        align-items: center;
        justify-content: space-between;
        margin-bottom: 5px;
    }
    .comp-header {
        margin-top: 2px;
    }
    .comp-center {
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
    }
    .comp-center span {
        margin-left: 5px;
    }
    .comp-center span:first-child {
        margin-left: 1px;
    }
</style>
