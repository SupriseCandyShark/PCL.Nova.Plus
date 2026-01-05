// 以下是操作 MC 版本数组的类
import { writable } from "svelte/store";

// 以下是下载部分
export const mc_list_ok = writable(false);
export const latest_release = writable({});
export const latest_snapshot = writable({});
export const mc_list_release = writable([]);
export const mc_list_snapshot = writable([]);
export const mc_list_old_beta = writable([]);
export const mc_list_old_alpha = writable([]);
export const mc_list_fool = writable([]);
export const forge_list = writable({});
export const forge_support = writable([]);
export const fabric_list = writable([]);
export const quilt_list = writable([]);
export const neoforge_list = writable([]);
export const optifine_list = writable([]);
export const liteloader_list = writable([]);

// 以下是核心部分
export const select_mc = writable([]);
export const current_mc_index = writable(0);
export const select_mc_version = writable([]);
export const current_mc_version_index = writable(-1);
export const current_mc_version_path = writable("");

// 以下是 Java 管理
export const select_java = writable([]);
export const current_java_index = writable(-1);

// 以下是资源部分
export const search_mod_cache = writable([]);

// 以下是游戏启动
export const game_log = writable("");
