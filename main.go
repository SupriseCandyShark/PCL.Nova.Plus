/*
 *                        _oo0oo_
 *                       o8888888o
 *                       88" . "88
 *                       (| -_- |)
 *                       0\  =  /0
 *                     ___/`---'\___
 *                   .' \\|     |// '.
 *                  / \\|||  :  |||// \
 *                 / _||||| -:- |||||- \
 *                |   | \\\  - /// |   |
 *                | \_|  ''\---/''  |_/ |
 *                \  .-\__  '-'  ___/-. /
 *              ___'. .'  /--.--\  `. .'___
 *           ."" '<  `.___\_<|>_/___.' >' "".
 *          | | :  `- \`.;`\ _ /`;.`/ - ` : | |
 *          \  \ `_.   \_ __\ /__ _/   .-` /  /
 *      =====`-.____`.___ \_____/___.-`___.-'=====
 *                        `=---='
 *
 *
 *      ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *
 *            佛祖保佑       永不宕机     永无BUG
 */
package main

import (
	"NovaPlus/module/launcher"
	"NovaPlus/module/mmcll"
	"bufio"
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	Ctx context.Context
}

var app = &App{}
var readerWriter = &launcher.ReaderWriter{}
var mainMethod = &launcher.MainMethod{}
var accountMethod = &launcher.AccountMethod{}
var network = &launcher.Network{}
var launchMethod = &launcher.LaunchMethod{}
var downloadMethod = &launcher.DownloadMethod{}

var logLevel = "0"

// startup is called when the app starts. The context is saved
// so we can call the runtime methods

func startup(ctx context.Context) {
	app.Ctx = ctx
	readerWriter.Ctx = ctx
	mainMethod.Ctx = ctx
	accountMethod.Ctx = ctx
	network.Ctx = ctx
	launchMethod.Ctx = ctx
	downloadMethod.Ctx = ctx
}

func (a *App) StartDownload() {
	for progress := 10; progress <= 100; progress += 10 {
		time.Sleep(time.Second)
		runtime.EventsEmit(a.Ctx, "download_progress", progress)
	}
	runtime.EventsEmit(a.Ctx, "download_success")
}

// GetMachineCode 在Windows上尝试生成PCL2原版的识别码~
func (a *App) GetPCL2Identify() string {
	return strings.ToUpper(launcher.GetPCL2Identify())
}
func (a *App) GetUniqueAddress() string {
	return strings.ToUpper(launcher.GetUniqueAddress())
}
func genRandomInt(seed int64) int {
	return rand.New(rand.NewSource(seed)).Intn(100)
}

// 此处使用了 Privacy.go 里面的 PV_KEY，用来加密解密你的数据！
func (a *App) CryptoUnlock(wc string) bool {
	article := strings.Split(launcher.PV_KEY, "@@")
	key := article[0]
	iv := article[1]
	wm, err := launcher.AESDecrypt(key, iv, wc)
	if err != nil {
		return false
	}
	ua := a.GetUniqueAddress()
	wt := "<moduleKey>" + ua + "</moduleKey><version>" + strconv.Itoa(genRandomInt(int64(launcher.GetHash(ua)))) + "</version>"
	return wm == wt
}
func BankerRound(x float64) float64 {
	floor := math.Floor(x)
	ceil := math.Ceil(x)
	if floor == ceil {
		return x
	}
	if x-floor < ceil-x {
		return floor
	}
	return ceil
}

// GetTodayLucky 生成今日人品，有3个报错类型。-1：未能找到注册表中的机器码，你可能在虚拟机中运行Nova、-2：未能找到注册表中的PCL-Identify键值，你可能需要从原版PCL里主动生成一次识别码才能。-3：不支持的系统。
func (a *App) GetTodayLucky(code string) int {
	switch code {
	case "ERR_1":
		return -1
	case "ERR_2":
		return -2
	case "DNS":
		return -3
	default:
		// 以下为标准的PCL生成今日人品的代码【参数code为PCL原版识别码】~
		now := time.Now()
		str1 := fmt.Sprintf("asdfgbn%d12#3$45%dIUY", now.YearDay(), now.Year())
		str2 := fmt.Sprintf("QWERTY%s0*8&6%dkjhg", "2E97-2E0C-4038-BE81", now.Day())
		hash1 := launcher.GetHash(str1)
		hash2 := launcher.GetHash(str2)
		combined := math.Abs((float64(hash1)/3.0 + float64(hash2/3.0)) / 527.0)
		num := int(BankerRound(math.Mod(combined, 1001.0)))
		var lucky int
		if num >= 970 {
			lucky = 100
		} else {
			lucky = int(BankerRound(float64(num) / 969.0 * 99.0))
		}
		return lucky
	}
}

//go:embed all:frontend/dist
var assets embed.FS

func runRawMain() {
	// Create an instance of the app structure
	// Create application with options
	err := wails.Run(&options.App{
		Title:     "PCL.Nova.Plus",
		Width:     1024,
		Height:    614,
		MinWidth:  1024,
		MinHeight: 614,
		Frameless: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		//BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		DragAndDrop: &options.DragAndDrop{
			EnableFileDrop: true,
		},
		OnStartup: startup,
		Bind: []any{
			app,
			readerWriter,
			accountMethod,
			mainMethod,
			network,
			launchMethod,
			downloadMethod,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
func runMMCLLTestMain() {
	defer func() {
		var m string
		_, _ = fmt.Scan(&m)
	}()
	if f, err := mmcll.GetFile(filepath.Join(launcher.GetCurrentDir(), "mmcll.json")); err == nil {
		var m map[string]any
		err := json.Unmarshal([]byte(f), &m)
		if err != nil {
			fmt.Println("JSON 格式有误，请重新写入文件！")
			return
		}
		name := mmcll.Safe(m, "", "name").(string)
		if name == "" {
			fmt.Println("名称没有输入，请重新输入！")
			return
		}
		uuid := mmcll.GenerateBukkitUUID(name)
		account := mmcll.NewLaunchAccountOffline(name, uuid)
		javaPath := mmcll.Safe(m, "", "java_path").(string)
		if javaPath == "" {
			fmt.Println("Java 路径没有输入，请重新输入！")
			return
		}
		rootPath := mmcll.Safe(m, "", "root_path").(string)
		if rootPath == "" {
			fmt.Println("版本根路径没有输入，请重新输入！")
			return
		}
		versPath := mmcll.Safe(m, "", "vers_path").(string)
		if versPath == "" {
			fmt.Println("版本文件路径没有输入，请重新输入！")
			return
		}
		isolation := mmcll.Safe(m, true, "isolation").(bool)
		gamePath := mmcll.If(isolation, versPath, rootPath).(string)
		option := mmcll.NewLaunchOption(account, javaPath, rootPath, versPath, gamePath)
		if height := mmcll.Safe(m, 480, "height").(int); height >= 480 {
			option.SetWindowHeight(uint32(height))
		}
		if width := mmcll.Safe(m, 854, "width").(int); width >= 854 {
			option.SetWindowWidth(uint32(width))
		}
		if memory := mmcll.Safe(m, 1024, "memory").(int); memory >= 1024 {
			option.SetMaxMemory(uint32(memory))
		}
		if info := mmcll.Safe(m, mmcll.LauncherName, "info").(string); info != mmcll.LauncherName {
			option.SetCustomInfo(info)
		}
		if addiJvm := mmcll.Safe(m, "", "addi_jvm").(string); addiJvm != "" {
			option.SetAdditionalJvm(addiJvm)
		}
		if addiGame := mmcll.Safe(m, "", "addi_game").(string); addiGame != "" {
			option.SetAdditionalGame(addiGame)
		}
		err = mmcll.LaunchGame(*option, true, func(back []string) {
			cmd := exec.Command(back[0], back[1:]...)
			stdout, _ := cmd.StdoutPipe()
			err := cmd.Start()
			if err != nil {
				fmt.Println("错误！无法运行 CMD!，错误信息：" + err.Error())
				return
			}
			scanner := bufio.NewScanner(stdout)
			for scanner.Scan() {
				fmt.Println(scanner.Text())
			}
			err = cmd.Wait()
			if err != nil {
				fmt.Println("错误！无法等待 CMD!，错误信息：" + err.Error())
				return
			}
		})
		if err != nil {
			fmt.Println("版本检查出错了！请将你的 MC 版本，Java 版本 以及 PCL2 能否正常启动的截图发送给作者！错误信息为：" + err.Error())
		}
	} else {
		fmt.Println("这里是 xphost 制作的 Golang 语言 MC 启动器 测试环境！请不要用其放入生产环境！本产品目前严禁制作整合包，仅限内部使用！如果你从外部获取了本产品，请于 24 小时内删除！")
		fmt.Println("目前仅限离线登录，无法使用正版，无法自主定义 UUID，仅能使用 Bukkit 方式生成 UUID")
		fmt.Println("请将 配置文件 保存到 mmcll.json 下")
		fmt.Println("配置格式为：")
		m := map[string]any{
			"name":      "[你的账户名称]",
			"java_path": "[你的 Java 路径（精确到 java.exe）]",
			"root_path": "[你的 游戏 根目录（精确到 .minecraft）]",
			"vers_path": "[你的 游戏 版本目录（精确到 .minecraft/versions/<版本名称>）]",
			"isolation": "[布尔值，是否版本隔离]",
			"height":    "[数字值，窗口高度]",
			"width":     "[数字值，窗口宽度]",
			"memory":    "[数字值，最大内存]",
			"info":      "[字符串，自定义信息，显示在游戏地图里 F3 后的左上角以及游戏主窗口的左下角]",
			"addi_jvm":  "[额外 JVM 参数，如果你不知道是什么，请不要设置此值]",
			"addi_game": "[额外 Game 参数，如果你想一进入游戏就全屏，可以设置 --fullScreen]",
		}
		res, _ := json.MarshalIndent(m, "", "\t")
		fmt.Println(string(res))
	}
}

type MyHandler struct {
	writer io.Writer
	level  slog.Level
}

func (h *MyHandler) Enabled(_ context.Context, level slog.Level) bool {
	minLevel := slog.LevelInfo
	if h.level != slog.LevelInfo {
		minLevel = h.level
	}
	return level >= minLevel
}
func (h *MyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}
func (h *MyHandler) WithGroup(name string) slog.Handler {
	return h
}
func (h *MyHandler) Handle(_ context.Context, record slog.Record) error {
	timeStr := record.Time.Format("2006/01/02 15:04:05")
	levelStr := record.Level.String()[0:4] // 取级别的前四个字母，如 "INFO", "WARN"
	logLine := fmt.Sprintf("%s %s %s\n", timeStr, levelStr, record.Message)
	_, err := h.writer.Write([]byte(logLine))
	return err
}
func main() {
	// 设置 slog
	p := filepath.Join(launcher.GetCurrentDir(), "PCL.Nova", "log", "Log.log")
	mmcll.SetFile(p, "")
	f, err := os.OpenFile(p, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	mw := io.MultiWriter(os.Stdout, f)
	l := slog.New(&MyHandler{
		writer: mw,
		level:  mmcll.If(logLevel == "DEBUG", slog.LevelDebug, slog.LevelInfo).(slog.Level),
	})
	slog.SetDefault(l)
	slog.Info("日志初始化完成！")
	runRawMain()
	//runMMCLLTestMain()
}