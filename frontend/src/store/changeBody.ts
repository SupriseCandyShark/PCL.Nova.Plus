import { writable } from "svelte/store";
export const CurseForgeBaseURL = "https://api.curseforge.com/v1";
export const CurseForgeSearchURL =
    "https://api.curseforge.com/v1/mods/search?gameId=432&page_size=50&sortOrder=desc";
export const ModrinthBaseURL = "https://api.modrinth.com/v2";
export const ModrinthSearchURL =
    "https://api.modrinth.com/v2/search?limit=50&index=relevance";
export const Steve =
    "iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAYAAACqaXHeAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsQAAA7EAZUrDhsAAAHMSURBVHhe7dsxS0JRHAVwr099WqQRYVIgzUHQ1tSsEA3RF6iGICForKmmigja6hPUFC1FNIRfpCAsoYJCUMSe+p41nPHvfW" +
    "BDw/n/lnMWBc/jwlXQzE5nehELP0DpwzFRNFnb76LJul0fTRaN2t8/HnPQZD3rp/t9fyQtHQBJSwdA0tIBkLTMTN5+D/grNxFDk6Vd+zMIQp7RV8NDk/lByD0DSUsHQNLSAZC0dAAkrdB7wOHWKposEU+gyVLJNJqs1ayj9eHYn1Ht4w1NdnJ1hybTI4CkpQMgaekASFr0A5iLvU3r" +
    "PSDlumiylmf/Pu449tfX6jU0WX4iiyarVB7RBqNHAElLB0DS0gGQtOgHMNcH29Z7wPLuKZpsqXiEJluff0WTuSPDaLLzhzE02c39Dprscr+EJtMjgKSlAyBp6QBIWvQDmLONRes9IJubQpPF3Qya7LZcRhvMSnEBTVZ9eUKTPb9/osn0CCBp6QBIWjoAkhb9AOZ4rWC9BySM/Q8DlW" +
    "oD7X/kJofQZI1mB02mRwBJSwdA0tIBkLToBzClwpz1HhAE9nvA+Kj994AwXqeNJku5SbTBtLxvNJkeASQtHQBJSwdA0iIfIBL5AYycVHLWJ+hEAAAAAElFTkSuQmCC";
export const Alex =
    "iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAYAAACqaXHeAAABg0lEQVR4nO2asUoDQRRFZ5zRjSBYWqbVTqzstEvvJ9hYWVv4G4KkyWf4CZbaqZWdhY2iKLohQf/gPpZFDrj3tG9n9nB58DKbyU/TyU9SfHzLctoY6Xq0PqLv/sH6lY46/w4HQAvQOABagMYB0AI09c/nfLD+a/4p6+" +
    "t6995+g+8AB0AL0DgAWoDGAdACNJWe85uHZ3r/gPn1Ra/1g+8AB0AL0DgAWoDGAdACNDWa02kt2KHnnG+aRtbf3171+/eOZTnfzGR98B3gAGgBGgdAC9A4AFqApo72T+UDW7tHsr59siPrtxM951PR3xMOLs9l/WF6L+vP+u3uAAdAC9A4AFqAxgHQAjT55e5K3hPMZVVuULK+ZhjN" +
    "+bQM/nfIRZYXi4Ws15JlffAd4ABoARoHQAvQOABagKbWWvUTJTjPB3N8+Rh81w8oY31TMJrzbdvK+uA7wAHQAjQOgBagcQC0AE3wIyDF5/XgvF/GXXS67x8RXD9wBzgAWoDGAdACNA6AFqD5BVHZQDbja48HAAAAAElFTkSuQmCC";
export const Zuri =
    "iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAYAAACqaXHeAAABZ0lEQVR4nO2bPUoDURSFM/ok+JOIELG0tdQijWBjpWDjKqziHtyC6xDcgY0gQjYgWA3YSAJi4g9BRXuL85BL/NB3vvbkDV8OF+4ww1TtVH02ArRaSzJvzqXI5RuTt3eZj8dPoevPhE7/A1wALUDjAmgBGhdAC9Bkl/" +
    "Th9obM2wtNma8s6jzHw/NE5qMXnZ9f3ci8+AlwAbQAjQugBWhcAC1AU/X2NuXzgOgep8ndRxQ/AS6AFqBxAbQAjQugBWjSX9/zOXL/r/gJcAG0AI0LoAVoXAAtQJN9L3Bydi3z3fVZme90uz8z+sZlvy/zi/pD5sf7WzIvfgJcAC1A4wJoARoXQAvQVKdHB/K9wHA4+C2XqdDprMq8" +
    "+AlwAbQAjQugBWhcAC1Ak6a951/vH0Pn59eWQ+dv6zuZFz8BLoAWoHEBtACNC6AFaFI9GMkfRL8HiO7xHNHvCYqfABdAC9C4AFqAxgXQAjRfq1A1NTxw6xsAAAAASUVORK5CYII=";
export const Sunny =
    "iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAYAAACqaXHeAAABbklEQVR4nO2ZMUrEUBRF50sgBEwREYS4ALXKEmYBugRBcAm2Y+msw8YljPazC12AAUEMOELSxdbqPeQhB/n3tDf/5+TyYH4mqeu6eWFQVZUVL8ZxNHNvfdM0Zt73fWh/z2/PTDNABdACNCqAFqBRAbQATRHdoG1bMx" +
    "+GIXoLk3k2jzEu2U+ACqAFaFQALUCjAmgBGvcc4L1Pe+/z3nqPsizNfJqm0P7ZT4AKoAVoVAAtQKMCaAGaYrfbmRfUdW3m3vu+t/9f490/+wlQAbQAjQqgBWhUAC1AUzxdn5kXnKw2Zp5SMvPnu/NfS/3k9PbRzL3vAi/rCzPPfgJUAC1AowJoARoVQAvQpM/7q9gH9n9O9hOgAmgB" +
    "GhVAC9CoAFqApuhf380L2uPD0A3WD9vQ+tXlMrTee77sJ0AF0AI0KoAWoFEBtABN2t4szf8Djg72zQ2i54Qo3u/828eXmWc/ASqAFqBRAbQAjQqgBWi+Ab8TQ/22ShIlAAAAAElFTkSuQmCC";
export const Noor =
    "iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAYAAACqaXHeAAABeElEQVR4nO2bu0rEUBRFEw0KCpIECwuZxtJHYSUM2AgWfoAi2FhZ+gl+zqCNpWgvfoCPzkatVDLTKCiI1jb7ECIs8O7Vbm6ysjlwLyTJV+ZmvjPB1MS4irP3zy+ZL87XMr99amTe9f7R+jGZJoALoAVoXAAtQOMCaA" +
    "GafK1XyXPA4fqqvEA9PSnz5u0DXX90diXz5CfABdACNC6AFqBxAbQATT7Y25DngGifna0rmb82w/ZWf0h0jkh+AlwALUDjAmgBGhdAC9AU0T4/Go465WVVtjNqef2IOrh/8hPgAmgBGhdAC9C4AFqApoj22e3jS5nvLvdkvt8vWyr95vTuUeaD6weZn+z0ZZ78BLgAWoDGBdACNC6A" +
    "FqDJzw+25HuB/07yE+ACaAEaF0AL0LgAWoCm6PodXsTFzX2n9ZtLC53W+/uAABdAC9C4AFqAxgXQAjTF84v+by/L9H9/0Tmh6z4eEe3z0fMlPwEugBagcQG0AI0LoAVofgCvVkbgvjJEUQAAAABJRU5ErkJggg==";
export const Makena =
    "iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAYAAACqaXHeAAABhklEQVR4nO2azUrDQBhFE0mbSFMstURE3Na6ciPiC/Ql+qq+gTtX/mwVxFJLsWlpUNGtq/spQz2L+c72NpnTy8BMJkmH/eorEWw+3lWcDDpdmVvMVkuZF1lL5mVeyLxuNjLfkWkEeAG0AI0XQAvQeAG0AE3WyfOtDm" +
    "Ct88OTg6D7z5/WMrf2CdHPAC+AFqDxAmgBGi+AFqDJVk0jf7Dbasu8vfcp86LRz/MP9y8yt/YJ1vjzqT7PiH4GeAG0AI0XQAvQeAG0AE3Wr4xz924ZNEAZ9trgF/e3BtDnEdHPAC+AFqDxAmgBGi+AFqDJrHX06vpW5meHqcyr49GfpX4yfbyT+c2z/LwhGV+cyjz6GeAF0AI0XgAt" +
    "QOMF0AI06WR8LhfS6aL+L5etUPX0eUb0M8ALoAVovABagMYLoAVo0svRkdwH2OfumtnrIuj6wX4v6Pp66e8FJF4ALUDjBdACNF4ALUCTva3DnvetfULoOm5hrfPW/4t+BngBtACNF0AL0HgBtADNN4fxPOYStPLFAAAAAElFTkSuQmCC";
export const Kai =
    "iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAYAAACqaXHeAAABXklEQVR4nO2bvU4CURBGWUFJICExtpb0FDRqQucboPEFKHgdOxKo6Px9Aqy1oTC2dppY0YEJP2ZpKchM4O56ivuddrjLyckku5tAspzcpgWL1B4Xfhf2vFq25zAHtACNAtACNApAC9AoAC1AU3I/4d3nK0cZqWxnMv" +
    "4MOn/SrJvz6DdAAWgBGgWgBWgUgBag8Z8Dcn6f/+i/mPPjWiXo+slFy5xHvwEKQAvQKAAtQKMAtACN+xyQ93069Hwo0W+AAtACNApAC9AoAC1Ak6ym9+YPAIrVa/MCl90rcz5oHO5utUHnfWnOR71Hc/43ezDn0W+AAtACNApAC9AoAC1A4z4HfA2f/8slF05vzs159BugALQAjQLQ" +
    "AjQKQAvQlNL5T65f0Ht6CzrfbZ9lZLKd6DdAAWgBGgWgBWgUgBagSdz/DTp8371m5bIX3vu+R/QboAC0AI0C0AI0CkAL0KwBFfMxwHR5oEgAAAAASUVORK5CYII=";
export const Efe =
    "iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAYAAACqaXHeAAABpklEQVR4nO2bsUrDUBiFrcaGJoLFahA3cRILLro6uNTZVTc3fQ8Xwa26OfgAruqiD6CLUHB0E6k2ZmmCSYuu4nB+StAD3vOtJ9x8HH64SbipHO4ef44BvGoVxSZhLYB5P0tLrR/3ujCfaUQwHy9193+ACmALsFEBbA" +
    "E2KoAtwMazLsjSPsxnG3Mwf+u9wrwWhJYCJE7w+noOMFABbAE2KoAtwEYFsAXYeNb7fjEoYD49PId5tDAF86dkC+bW+/7G8iPMn/MVmDs/ASqALcBGBbAF2KgAtgAbb5Dn8IKl+iXM/QDv8xaL9Ssjt1bA988S/D3D+QlQAWwBNiqALcBGBbAF2JjfA3ZOb2G+31yH+WZrfmSp79xc" +
    "v8D8pHMH86O9NsydnwAVwBZgowLYAmxUAFuATaV9cAbPCUb+xV+5/Ardj22YOz8BKoAtwEYFsAXYqAC2ABvznGCRDmE+GUzA/P6hM5rRD9ZWmzC3/PwQn0N0fgJUAFuAjQpgC7BRAWwBNp71315Ux/u8hbWPl8V6Donf8f8Ezk+ACmALsFEBbAE2KoAtwOYLUodFeDw82gEAAAAASU" +
    "VORK5CYII=";
export const Ari =
    "iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAYAAACqaXHeAAABcElEQVR4nO2aMUoDQRhGExmSLSKIrGBjGRFsPILgDVJZeAZ7C0HxAsEzWHsDwUsIYi5gEYvAIksIaCsI388yLg+Z77X/7vLyGJhJssObk92vgeBgb0eNB5/rjZxHzE7P5Pzx+Snr+RFbvT79H+AAtACNA9ACNA5AC9" +
    "Ck8XgkL4j2+Wgfn9TbnaW6PD86J3ysGjkvfgU4AC1A4wC0AI0D0AI0aVLpc0C0D9dH07/0+UVV78v5LLg/OicUvwIcgBagcQBagMYBaAGaFF3Q9z6fS+jnc4DGAWgBGgegBWgcgBagSU27lheMzq/kfJr0UeLl4a6z1E+OL67lfLHR/1s083s5L34FOAAtQOMAtACNA9ACNMPV7aV8" +
    "TzCiOtS/2/dN+/aedX/xK8ABaAEaB6AFaByAFqBJfe/jy9dF1v3R7/65/sWvAAegBWgcgBagcQBagCa1S/19OnpPL4J+vyD6fMWvAAegBWgcgBagcQBagOYblfMz+bzIKgUAAAAASUVORK5CYII=";
// 千万别点
export const dont_click = writable(0);
// 暗色模式
export const dark_mode = writable(false);
// 启动器版本
export const launcher_version = writable("");
// 以下是 MC 的正常模式
export const theme_mode = writable(1);
export const unlock_theme = writable(false);
export const current_view = writable("home");
export const current_account = writable("Microsoft");
export const current_download = writable("Auto-Install");
export const current_online = writable("IPv6");
export const current_about = writable("Help");
export const current_setting = writable("Launch");
export const current_manual = writable("Forge");
// 为 true 则是【处在添加账号界面】，为 false 则是【处在选择添加账号类型界面】，该配置项不会保存到外部。
export const current_account_page = writable(true);
// 以下是账号选择
export const select_account = writable([]);
export const current_account_index = writable(-1);
// 以下是 IPv6 检测框
export const select_ipv6 = writable([]);
// 以下是主页设置
export const select_homepage = writable(-1);
export const current_homepage = writable(-1);
export const homepage_url = writable("");
export const homepage_list = writable([]);
// 以下是解析主页
export const homepage_cache = writable("");
export const homepage_value = writable({});
// 下载的列表
export const download_list = writable([]);
// 以下是实例设置
export const current_instance = writable("overview")