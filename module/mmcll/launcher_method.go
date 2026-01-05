package mmcll

import (
	"archive/zip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"

	"github.com/go-vgo/robotgo"
	"github.com/shirou/gopsutil/mem"
)

func ConvNameToPath(name string) string {
	var suffix string
	if strings.Contains(name, "@") {
		i := strings.Index(name, "@")
		suffix = name[i+1:]
		name = name[0:i]
	} else {
		suffix = "jar"
	}
	s := strings.Split(name, ":")
	if len(s) == 3 {
		result := strings.ReplaceAll(s[0], ".", "/")
		return filepath.Join(result, s[1], s[2], fmt.Sprintf("%s-%s.%s", s[1], s[2], suffix))
	} else if len(s) == 4 {
		result := strings.ReplaceAll(s[0], ".", "/")
		return filepath.Join(result, s[1], s[2], fmt.Sprintf("%s-%s-%s.%s", s[1], s[2], s[3], suffix))
	}
	return ""
}

// GetMCRealPath 用来获取MC的文件绝对路径。（可以通过后缀或者sha1来判断）
func GetMCRealPath(versionPath, suffix string) (string, error) {
	fpr, err := os.Stat(versionPath)
	if err != nil || !fpr.IsDir() {
		return "", err
	}
	dir, err := os.Open(versionPath)
	if err != nil {
		return "", err
	}
	defer dir.Close()
	files, err := dir.Readdir(-1)
	if err != nil {
		return "", err
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		name := file.Name()
		fp := filepath.Join(versionPath, name)
		if strings.Contains(name, suffix) {
			if suffix == ".json" && filepath.Ext(name) == ".json" {
				content, err := GetFile(fp)
				if err != nil {
					continue
				}
				var m map[string]any
				if err := json.Unmarshal([]byte(content), &m); err != nil {
					continue
				}
				if res := Safe(m, []any{}, "libraries").([]any); len(res) == 0 {
					continue
				}
				if res := Safe(m, "abcdefghijklmnopqrstuvwxyz", "mainClass"); res == "abcdefghijklmnopqrstuvwxyz" {
					continue
				}
				if res := Safe(m, "abcdefghijklmnopqrstuvwxyz", "id"); res == "abcdefghijklmnopqrstuvwxyz" {
					continue
				}
				return fp, nil
			} else {
				return fp, nil
			}
		} else if !strings.Contains(suffix, ".") {
			sha1, err := GetSha1(fp)
			if err != nil {
				continue
			}
			if sha1 == suffix {
				return fp, nil
			}
		}
	}
	return "", NewMMCLLError(-80, "Cannot Find Any Real Path in your dir, please try again!")
}

// GetMCInheritsFrom 获取当前 versionPath 下的 JSON，并且附带 suffix 键值对的 原版键值
func GetMCInheritsFrom(versionPath, suffix string) (string, error) {
	realPath, err := GetMCRealPath(versionPath, ".json")
	if err != nil {
		return "", err
	}
	realFile, err := GetFile(realPath)
	if err != nil {
		return "", err
	}
	var root map[string]any
	if err = json.Unmarshal([]byte(realFile), &root); err != nil {
		return "", err
	}
	inherits := Safe(root, "", suffix)
	if inherits == "" {
		return versionPath, nil
	}
	parent := filepath.Dir(versionPath)
	if parent == versionPath {
		return "", NewMMCLLError(-400, "Parent Path is already in Home Dir!")
	}
	dir, err := os.Open(parent)
	if err != nil {
		return "", err
	}
	defer dir.Close()
	files, err := dir.Readdir(-1)
	if err != nil {
		return "", err
	}
	for _, file := range files {
		if !file.IsDir() {
			continue
		}
		fp := filepath.Join(parent, file.Name())
		versionJson, err := GetMCRealPath(fp, ".json")
		if err != nil {
			continue
		}
		versionContent, err := GetFile(versionJson)
		if err != nil {
			continue
		}
		vanillaVersion, err := GetVanillaVersion(versionContent)
		if err != nil {
			continue
		}
		if vanillaVersion == inherits {
			return fp, nil
		}
	}
	return "", NewMMCLLError(-82, "Cannot Find Any InheritsFrom Path in your dir, please try again!")
}

// GetVanillaVersion 通过JSON获取到原版键值
func GetVanillaVersion(j string) (string, error) {
	var root map[string]any
	if err := json.Unmarshal([]byte(j), &root); err != nil {
		return "", err
	}
	// 判断 PCL2 核心
	clientVersion := Safe(root, "", "clientVersion").(string)
	if clientVersion != "" {
		return clientVersion, nil
	}
	// 判断 HMCL 核心
	patches := Safe(root, []any{}, "patches").([]any)
	if len(patches) != 0 {
		for _, patch := range patches {
			id := Safe(patch, "", "id").(string)
			if id == "game" {
				mcid := Safe(patch, "", "version").(string)
				if mcid != "" {
					return mcid, nil
				}
			}
		}
	}
	// TODO: 从 releaseTime 中判断其的原版键值
	// TODO: 从 NeoForge 的额外 game 参数中提取出原版键值
	// 直接从 id 号中取出（最保险但也最容易被其他启动器修改
	id := Safe(root, "", "id").(string)
	if id != "" {
		return id, nil
	}
	return "", NewMMCLLError(-500, "Cannot Find Any Vanilla Version in JSON Path!")
}

// MergeMCJson 合并两个JSON文件（用于给Forge、Fabric等含有inheritsFrom键的JSON文件与原版JSON文件合并成一个JSON文件。）
// rawContent: 填入带有inheritsFrom的JSON文件
// insContent: 填入原版的JSON文件
// result: 返回两个JSON文件被合并后的产物（只是合并，即使键值不存在也不会抛出异常）
func MergeMCJson(raw, ins string) map[string]any {
	var insContent map[string]any
	var rawContent map[string]any
	if err := json.Unmarshal([]byte(raw), &rawContent); err != nil {
		return make(map[string]any)
	}
	if err := json.Unmarshal([]byte(ins), &insContent); err != nil {
		return make(map[string]any)
	}
	if len(rawContent) == 0 || len(insContent) == 0 {
		return make(map[string]any)
	}
	if raw == ins {
		return rawContent
	}
	// 重载 mainClass
	mc := Safe(rawContent, "", "mainClass").(string)
	if mc != "" {
		insContent["mainClass"] = mc
	}
	// 重载 id
	id := Safe(rawContent, "", "id").(string)
	if id != "" {
		insContent["id"] = id
	}
	// 补充 libraries
	rawLib := Safe(rawContent, []any{}, "libraries").([]any)
	insLib := Safe(insContent, []any{}, "libraries").([]any)
	if len(rawLib) > 0 {
		for _, lib := range rawLib {
			insLib = append([]any{lib}, insLib...)
		}
		insContent["libraries"] = insLib
	}
	// 判断 arguments
	saf := Safe(insContent, map[string]any{}, "arguments").(map[string]any)
	if len(saf) == 0 {
		insContent["arguments"] = make(map[string]any)
	}
	// 补充 jvm
	rawJvm := Safe(rawContent, []any{}, "arguments", "jvm").([]any)
	insJvm := Safe(insContent, []any{}, "arguments", "jvm").([]any)
	if len(rawJvm) > 0 {
		for _, jvm := range rawJvm {
			insJvm = append(insJvm, jvm)
		}
		insContent["arguments"].(map[string]any)["jvm"] = insJvm
	}
	// 补充 game
	rawGame := Safe(rawContent, []any{}, "arguments", "game").([]any)
	insGame := Safe(insContent, []any{}, "arguments", "game").([]any)
	if len(rawGame) > 0 {
		for _, game := range rawGame {
			insGame = append(insGame, game)
		}
		insContent["arguments"].(map[string]any)["game"] = insGame
	}
	// 判断 1.12 minecraftArguments
	ma := Safe(rawContent, "", "minecraftArguments")
	if ma != "" {
		insContent["minecraftArguments"] = ma
	}
	return insContent
}

// JudgeRule 判断 libraries 的 rules 键值，以判断该 library 是否添加
func JudgeRule(rules []any) bool {
	boo := true
	for _, rule := range rules {
		if action := Safe(rule, "", "action"); action == "allow" {
			if aos := Safe(rule, "", "os", "name"); aos == "" {
				boo = true
			} else if aos != GetMcOs() {
				boo = false
			} else if aos == GetMcOs() {
				return true
			}
		} else if action == "disallow" {
			if aos := Safe(rule, "", "os", "name"); aos == "" {
				boo = false
			} else if aos != GetMcOs() {
				boo = true
			} else if aos == GetMcOs() {
				return false
			}
		}
	}
	return boo
}

// Unzip 解压缩
func Unzip(zipPath, extractPath string) error {
	archive, err := zip.OpenReader(zipPath)
	if err != nil {
		return NewMMCLLError(-1451, "Cannot Open Zip File!")
	}
	defer archive.Close()
	for _, f := range archive.File {
		fp := filepath.Join(extractPath, f.Name)
		if f.FileInfo().IsDir() {
			_ = os.MkdirAll(fp, os.ModePerm)
			continue
		}
		dstFile, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			continue
		}
		fileIn, err := f.Open()
		if err != nil {
			continue
		}
		if _, err := io.Copy(dstFile, fileIn); err != nil {
			continue
		}
		fileIn.Close()
		dstFile.Close()
	}
	return nil
}

// UnzipNative 解压 Natives，从 {rootPath}/libraries中，找到 finalJson 里的 libraries 对应的键值，随后解压到 {versionPath}/{版本号}-natives下
// isCheckLibraries 的意思是：你是否需要检查 Libraries 的正确性？如果不需要检查，可以设置成 false。
func UnzipNative(finalJson, rootPath, versionPath string, isCheckLibraries bool) error {
	nativeDir := filepath.Join(versionPath, fmt.Sprintf("%s-%s-native", filepath.Base(versionPath), LauncherName))
	if err := os.MkdirAll(nativeDir, os.ModePerm); err != nil {
		return NewMMCLLError(-233, "Cannot Mkdir in "+versionPath+"!")
	}
	var f map[string]any
	if err := json.Unmarshal([]byte(finalJson), &f); err != nil {
		return NewMMCLLError(-234, "Cannot Parse JSON!")
	}
	libs := Safe(f, []any{}, "libraries").([]any)
	for _, lib := range libs {
		name := Safe(lib, "", "name").(string)
		if name == "" {
			continue
		}
		natives := Safe(lib, map[string]any{}, "natives").(map[string]any)
		if len(natives) == 0 {
			continue
		}
		if rules := Safe(lib, []any{}, "rules").([]any); len(rules) > 0 {
			if !JudgeRule(rules) {
				continue
			}
		}
		CurOsNative := Safe(natives, "", GetMcOs()).(string)
		if CurOsNative == "" {
			continue
		}
		if runtime.GOARCH == "386" || runtime.GOARCH == "arm" {
			CurOsNative = strings.ReplaceAll(CurOsNative, "${arch}", "32")
		} else if runtime.GOARCH == "amd64" || runtime.GOARCH == "arm64" {
			CurOsNative = strings.ReplaceAll(CurOsNative, "${arch}", "64")
		}
		path := filepath.Join(rootPath, "libraries", ConvNameToPath(name+":"+CurOsNative))
		if isCheckLibraries {
			sha1 := Safe(lib, "", "downloads", "classifiers", CurOsNative, "sha1").(string)
			if sha1 != "" {
				if sha, err := GetSha1(path); err != nil {
					return NewMMCLLError(-200, "Read Natives sha1 has occurred Error!, Please download it again! Error: "+err.Error())
				} else if sha != sha1 {
					return NewMMCLLError(-350, "Native Sha1 Not match your Libraries, Please download it again!")
				}
			}
		}
		if err := Unzip(path, nativeDir); err != nil {
			return NewMMCLLError(-201, "Cannot Unzip Native, Please download it again! Error: "+err.Error())
		}
	}
	return nil
}

// JudgeArguments 获取到 jvm 或者 game 的参数
func JudgeArguments(realJson map[string]any, arg string) []string {
	arguments := Safe(realJson, []any{}, "arguments", arg).([]any)
	var res []string
	for _, argument := range arguments {
		if i, ok := argument.(string); ok {
			res = append(res, strings.ReplaceAll(i, " ", ""))
		}
	}
	return res
}

// ExtractNumber 提取字符串中的所有字母（排除数字）
func ExtractNumber(str string) string {
	var result string
	for _, char := range str {
		if char < '0' || char > '9' {
			result += string(char)
		}
	}
	return result
}

// LibsIndexOf 用于查找所有 Lib 中可能出现的重复字段并且返回索引
func LibsIndexOf(libs []string, lib string) int {
	for i, l := range libs {
		left1 := ExtractNumber(l)
		left2 := ExtractNumber(lib)
		if left1 == left2 {
			return i
		}
	}
	return -1
}

// GetMCLibs 通过 realJson 获取到 MC 所有的 Libraries，并将其拼接到 -cp 参数的后面。与此同时，还会拼接 versionPath 下的主 jar 文件！
// result: 返回一个切片，你需要自己处理分号还是冒号的问题，在 Windows 上以 分号 作为分隔符，在类 Unix 系统上以 冒号 作为分隔符。
// isCheckLibraries 的意思是：你是否需要检查 Libraries 的正确性？如果不需要检查，可以设置成 false。
func GetMCLibs(realJson map[string]any, rootPath, versionPath string, isCheckLibraries bool) ([]string, error) {
	var result []string
	libs := Safe(realJson, []any{}, "libraries").([]any)
	// Fuck you Optifine!!
	var optifines []string
	for _, lib := range libs {
		name := Safe(lib, "", "name").(string)
		if name == "" {
			continue
		}
		natives := Safe(lib, map[string]any{}, "natives").(map[string]any)
		if len(natives) > 0 {
			continue
		}
		rules := Safe(lib, []any{}, "rules").([]any)
		if !JudgeRule(rules) {
			continue
		}
		path := filepath.Join(rootPath, "libraries", ConvNameToPath(name))
		// 新增加了 isCheckLibraries【是否校验 Libraries！】现在可以手动判断是否需要校验了！
		if isCheckLibraries {
			sha1 := Safe(lib, "", "downloads", "artifact", "sha1")
			if sha1 != "" {
				if sha, err := GetSha1(path); err != nil {
					// 最近有朋友反馈：手动修改了 glfw-patcher 之后，由于 sha1 对不上，导致这里直接返回错误，
					// 由于现已更新 是否校验 sha1，因此，此处可以直接返回 NewMMCLLError 了！
					// 这里不特判 asm，将直接 continue 掉。
					// Fuck you asm!!
					if name == "org.ow2.asm:asm:9.6" {
						continue
					}
					return nil, NewMMCLLError(-400, "Read Natives sha1 has occurred Error!, Please download it again! Error: "+err.Error())
				} else if sha != sha1 {
					if name == "org.ow2.asm:asm:9.6" {
						continue
					}
					return nil, NewMMCLLError(-450, "Native Sha1 Not match your Libraries, Please download it again!")
				}
			}
		}
		if index := LibsIndexOf(result, path); index > 0 {
			result[index] = path
			continue
		}
		if strings.Contains(name, "optifine") {
			optifines = append(optifines, name)
			continue
		}
		result = append(result, path)
	}
	for _, optifine := range optifines {
		result = append(result, filepath.Join(rootPath, "libraries", ConvNameToPath(optifine)))
	}
	// 拼接 主 JAR
	// 卧槽这逻辑好 jb 难写。。如果有有缘之人可以来帮我修改一下，谢谢啦亲亲~以下我先把逻辑写清楚
	// 首先判断 JSON 里面是否有 jar 键值对
	// 由于 Fabric 以及 Forge 的影响，现在版本 json 里没有 jar 这个键值对，因此这一段掐掉。与此同时下方替换成直接获取 inheritsFrom。原有逻辑被注释（
	inhJar, err := GetMCInheritsFrom(versionPath, "inheritsFrom")
	if err != nil {
		return nil, NewMMCLLError(-1450, "Cannot Get MC InheritsFrom Key!")
	}
	//inhJar, err := GetMCInheritsFrom(versionPath, "jar")
	//if err != nil {
	//	return nil, err
	//}
	//// 如果没有 jar 键值对，则判断里面是否有 inheritsFrom 键值对
	//if inhJar == versionPath {
	//	// 获取 inheritsFrom 键值对
	//	if inhJar, err = GetMCInheritsFrom(versionPath, "inheritsFrom"); err != nil {
	//		return nil, err
	//	}
	//	// 如果有 inheritsFrom 键值对，但是没有 jar 键值对，则直接返回 result 切片，因为此时多半是无需拼接 主 jar 的（
	//	if inhJar != versionPath {
	//		return result, nil
	//	}
	//}
	// 如果有 jar 键值对，或者 没有 jar 键值对以及 inheritsFrom 键值对，则获取原版，随后尝试拼接主 jar
	// 这里先获取 sha1 值，如果 sha1 获取不了，则返回错误（
	sha1 := Safe(realJson, "", "downloads", "client", "sha1").(string)
	if sha1 != "" {
		// 从 前面获取到的 路径，通过 sha1 获取到真实路径。随后拼接 主 jar，然后返回切片！
		if inhPath, err := GetMCRealPath(inhJar, sha1); err == nil {
			result = append(result, inhPath)
			return result, nil
		} else {
			return nil, NewMMCLLError(-208, "Cannot find Main Jar file, Please download it again! Error: "+err.Error())
		}
	}
	return nil, NewMMCLLError(-208, "Cannot find Main Jar file, Please download it again!")
}

// LaunchAccount 用来登录账号所需要的所有键，仅能new一次，后续无法修改。
// 清尽量使用下方的 New 进行初始化，不要直接初始化这个列表（否则会出事！！
type LaunchAccount struct {
	name        string // name 账户名字
	uuid        string // uuid 账号UUID
	accessToken string // accessToken 正版账号登录密钥
	accType     string // accType 账号登录类型
	base        string // base 账号登录第三方数据（仅限外置登录）
	url         string // url 账号登录第三方网址（仅限外置登录）
	alijPath    string // AuthLib-Injector 路径
	online      int8   // online 账号类型（1：离线、2：正版、3：外置）
}

func (l LaunchAccount) Name() string {
	return l.name
}

func (l LaunchAccount) Uuid() string {
	return l.uuid
}

func (l LaunchAccount) AccessToken() string {
	return l.accessToken
}

func (l LaunchAccount) AccType() string {
	return l.accType
}

func (l LaunchAccount) Base() string {
	return l.base
}

func (l LaunchAccount) Url() string {
	return l.url
}

func (l LaunchAccount) AlijPath() string {
	return l.alijPath
}

func (l LaunchAccount) Online() int8 {
	return l.online
}

// NewLaunchAccountOffline 新建一个离线登录模块
func NewLaunchAccountOffline(name, uuid string) LaunchAccount {
	return LaunchAccount{
		name:        name,
		uuid:        uuid,
		accessToken: uuid,
		accType:     "Legacy",
		base:        "",
		url:         "",
		alijPath:    "",
		online:      1,
	}
}

// NewLaunchAccountMicrosoft 新建一个微软登录模块
func NewLaunchAccountMicrosoft(name, uuid, accessToken string) LaunchAccount {
	return LaunchAccount{
		name:        name,
		uuid:        uuid,
		accessToken: accessToken,
		accType:     "msa",
		base:        "",
		url:         "",
		alijPath:    "",
		online:      2,
	}
}

// NewLaunchAccountThirdParty 新建一个外置登录模块
func NewLaunchAccountThirdParty(name, uuid, accessToken, base, url, alijPath string) LaunchAccount {
	return LaunchAccount{
		name:        name,
		uuid:        uuid,
		accessToken: accessToken,
		accType:     "msa",
		base:        base,
		url:         url,
		alijPath:    alijPath,
		online:      3,
	}
}

// LaunchOption 启动设置类（新建后无法修改account、javaPath、rootPath、versionPath几个必填项。）
type LaunchOption struct {
	account          LaunchAccount
	javaPath         string
	rootPath         string
	versionPath      string
	gamePath         string
	isCheckLibraries bool
	windowHeight     uint32
	windowWidth      uint32
	minMemory        uint32
	maxMemory        uint32
	customInfo       string
	additionalJvm    string
	additionalGame   string
}

// NewLaunchOption 新建一个启动设置类。（以下非必填的可以直接链式调用设置初始值）
func NewLaunchOption(account LaunchAccount, javaPath, rootPath, versionPath, gamePath string) *LaunchOption {
	return &LaunchOption{
		account:          account,
		javaPath:         javaPath,
		rootPath:         rootPath,
		versionPath:      versionPath,
		gamePath:         gamePath,
		isCheckLibraries: true,
		windowHeight:     480,
		windowWidth:      854,
		minMemory:        256,
		maxMemory:        1024,
		customInfo:       LauncherName + "-" + LauncherVersion,
		additionalJvm:    "",
		additionalGame:   "",
	}
}
func (opt *LaunchOption) SetIsCheckLibraries(isCheckLibraries bool) {
	opt.isCheckLibraries = isCheckLibraries
}
func (opt *LaunchOption) SetWindowWidth(windowWidth uint32) *LaunchOption {
	opt.windowWidth = windowWidth
	return opt
}
func (opt *LaunchOption) SetWindowHeight(windowHeight uint32) *LaunchOption {
	opt.windowHeight = windowHeight
	return opt
}
func (opt *LaunchOption) SetMinMemory(minMemory uint32) *LaunchOption {
	opt.minMemory = minMemory
	return opt
}
func (opt *LaunchOption) SetMaxMemory(maxMemory uint32) *LaunchOption {
	opt.maxMemory = maxMemory
	return opt
}
func (opt *LaunchOption) SetCustomInfo(customInfo string) *LaunchOption {
	opt.customInfo = customInfo
	return opt
}
func (opt *LaunchOption) SetAdditionalJvm(additionalJvm string) *LaunchOption {
	opt.additionalJvm = additionalJvm
	return opt
}
func (opt *LaunchOption) SetAdditionalGame(additionalGame string) *LaunchOption {
	opt.additionalGame = additionalGame
	return opt
}
func (opt *LaunchOption) Account() LaunchAccount {
	return opt.account
}

func (opt *LaunchOption) JavaPath() string {
	return opt.javaPath
}

func (opt *LaunchOption) RootPath() string {
	return opt.rootPath
}

func (opt *LaunchOption) VersionPath() string {
	return opt.versionPath
}

func (opt *LaunchOption) GamePath() string {
	return opt.gamePath
}

func (opt *LaunchOption) IsCheckLibraries() bool {
	return opt.isCheckLibraries
}

func (opt *LaunchOption) WindowHeight() uint32 {
	return opt.windowHeight
}

func (opt *LaunchOption) WindowWidth() uint32 {
	return opt.windowWidth
}

func (opt *LaunchOption) MinMemory() uint32 {
	return opt.minMemory
}

func (opt *LaunchOption) MaxMemory() uint32 {
	return opt.maxMemory
}

func (opt *LaunchOption) CustomInfo() string {
	return opt.customInfo
}

func (opt *LaunchOption) AdditionalJvm() string {
	return opt.additionalJvm
}

func (opt *LaunchOption) AdditionalGame() string {
	return opt.additionalGame
}

// launchGame 正式启动游戏的类
type launchGame struct {
	account          LaunchAccount
	javaPath         string
	rootPath         string
	versionPath      string
	gamePath         string
	isCheckLibraries bool
	windowHeight     uint32
	windowWidth      uint32
	minMemory        uint32
	maxMemory        uint32
	customInfo       string
	additionalJvm    string
	additionalGame   string
	callback         func([]string)
}

// NewLaunchStart 初始化启动类
func newLaunchStart(option LaunchOption, callback func([]string)) launchGame {
	ls := launchGame{
		account:          option.account,
		javaPath:         option.javaPath,
		rootPath:         option.rootPath,
		versionPath:      option.versionPath,
		gamePath:         option.gamePath,
		isCheckLibraries: option.isCheckLibraries,
		windowHeight:     option.windowHeight,
		windowWidth:      option.windowWidth,
		minMemory:        option.minMemory,
		maxMemory:        option.maxMemory,
		customInfo:       option.customInfo,
		additionalJvm:    option.additionalJvm,
		additionalGame:   option.additionalGame,
		callback:         callback,
	}
	return ls
}

// checkError 检查参数是否有误
func (lg launchGame) checkError() error {
	if lg.account.Online() == 0 {
		if b, _ := regexp.MatchString("^[a-zA-Z0-9]{3,16}$", lg.account.Name()); !b {
			return NewMMCLLError(1, "user name is invalid")
		}
		if b, _ := regexp.MatchString("^[a-f0-9]{32}$", lg.account.Uuid()); !b {
			return NewMMCLLError(2, "user uuid is invalid")
		}
	} else if lg.account.Online() == 1 {
		um := NewUrlMethod("https://api.minecraftservices.com/minecraft/profile")
		ih, err := um.Get(lg.account.accessToken)
		if err != nil {
			return NewMMCLLError(95, "Get user info failed, Maybe your Access Token is Expired")
		}
		var m map[string]any
		if err = json.Unmarshal([]byte(ih), &m); err != nil {
			return NewMMCLLError(96, "Unmarshal user info failed, Maybe your Access Token is Expired")
		}
		name := Safe(m, "", "name").(string)
		if name == "" {
			return NewMMCLLError(97, "user name is invalid, Maybe your Access Token is Expired")
		}
		uuid := Safe(m, "", "id").(string)
		if uuid == "" {
			return NewMMCLLError(98, "user uuid is invalid, Maybe your Access Token is Expired")
		}
		if name != lg.account.Name() || uuid != lg.account.Uuid() {
			return NewMMCLLError(99, "user name or uuid not equals your access token, if you try to login with another account, please logout and try again!")
		}
		//TODO: 微软账户判断
	} else if lg.account.Online() == 2 {
		fmt.Println("mmm")
		ap, err := os.Stat(lg.account.AlijPath())
		if err != nil || !ap.IsDir() {
			return NewMMCLLError(50, "AuthLib-Injector File Is Not Found!")
		}
		t := fmt.Sprintf("%s/authserver/validate", lg.account.Url())
		p := fmt.Sprintf(`{"accesstoken":"%s"}`, lg.account.AccessToken())
		po := NewUrlMethod(t)
		_, err = po.Post(p, true)
		if err != nil {
			return NewMMCLLError(51, "ThirdParty Login Is Expire!")
		}
	}
	cInfo, err := os.Stat(lg.javaPath)
	if os.IsNotExist(err) {
		return NewMMCLLError(3, "javaPath does not exist")
	} else if cInfo.IsDir() {
		return NewMMCLLError(3, "javaPath is a directory")
	}
	cInfo, err = os.Stat(lg.rootPath)
	if os.IsNotExist(err) {
		return NewMMCLLError(4, "rootPath does not exist")
	} else if !cInfo.IsDir() {
		return NewMMCLLError(4, "rootPath is not a directory")
	}
	cInfo, err = os.Stat(lg.versionPath)
	if os.IsNotExist(err) {
		return NewMMCLLError(5, "versionPath does not exist")
	} else if !cInfo.IsDir() {
		return NewMMCLLError(5, "versionPath is not a directory")
	}
	cInfo, err = os.Stat(lg.gamePath)
	if os.IsNotExist(err) {
		return NewMMCLLError(6, "gamePath does not exist")
	} else if !cInfo.IsDir() {
		return NewMMCLLError(6, "gamePath is not a directory")
	}
	sx, sy := robotgo.GetScreenSize()
	if lg.windowWidth < 854 || lg.windowWidth > uint32(sx) {
		return NewMMCLLError(7, "window width out of range")
	}
	if lg.windowHeight < 480 || lg.windowHeight > uint32(sy) {
		return NewMMCLLError(8, "window height out of range")
	}
	if lg.minMemory < 256 || lg.minMemory > 1024 {
		return NewMMCLLError(9, "minMemory out of range")
	}
	v, _ := mem.VirtualMemory()
	sysMem := v.Total / 1024 / 1024
	if lg.maxMemory < 1024 || lg.maxMemory > uint32(sysMem) {
		return NewMMCLLError(10, "maxMemory out of range")
	}
	if lg.customInfo == "" {
		return NewMMCLLError(11, "customInfo is empty")
	}
	return nil
}

// launch 如果参数无误则尝试启动
func (lg launchGame) launch() (err error) {
	var result []string
	result = append(result, "-XX:+UseG1GC", "-XX:-UseAdaptiveSizePolicy", "-XX:-OmitStackTraceInFastThrow", "-Dstderr.encoding=UTF-8", "-Dstdout.encoding=UTF-8", "-Dfml.ignoreInvalidMinecraftCertificates=true", "-Dfml.ignorePatchDiscrepancies=true", "-Dlog4j2.formatMsgNoLookups=true")

	if runtime.GOOS == "windows" {
		result = append(result, "-XX:HeapDumpPath=MojangTricksIntelDriversForPerformance_javaw.exe_minecraft.exe.heapdump")
		if runtime.GOARCH == "386" {
			result = append(result, "-Xss1M")
		}
		if GetWindowsVersion() {
			result = append(result, "-Dos.name=Windows 10")
			result = append(result, "-Dos.version=10.0")
		}
	} else if runtime.GOOS == "darwin" {
		result = append(result, "-XstartOnFirstThread")
	}
	jsonPath, err := GetMCRealPath(lg.versionPath, ".json")
	if err != nil {
		return NewMMCLLError(-101, err.Error())
	}
	rawJSON, err := GetFile(jsonPath)
	if err != nil {
		return NewMMCLLError(-102, err.Error())
	}
	inheritsPath, err := GetMCInheritsFrom(lg.versionPath, "inheritsFrom")
	if err != nil {
		return NewMMCLLError(-103, err.Error())
	}
	var finalJson string
	if inheritsPath != lg.versionPath {
		f, err := GetMCRealPath(inheritsPath, ".json")
		if err != nil {
			return NewMMCLLError(-104, err.Error())
		}
		f, err = GetFile(f)
		if err != nil {
			return NewMMCLLError(-105, err.Error())
		}
		finalJson = f
	} else {
		finalJson = rawJSON
	}
	if err := UnzipNative(finalJson, lg.rootPath, lg.versionPath, lg.isCheckLibraries); err != nil {
		return NewMMCLLError(-106, err.Error())
	}
	realJson := MergeMCJson(rawJSON, finalJson)
	if len(realJson) == 0 {
		return NewMMCLLError(-107, "Cannot Merge MC Forge and Fabric JSON!")
	}
	param, err := lg.putArgument(realJson, result)
	if err != nil {
		return NewMMCLLError(-108, err.Error())
	}
	param = append([]string{lg.javaPath}, param...)
	defer func() {
		if r := recover(); r != nil {
			if e, ok := r.(error); ok {
				err = e
			} else {
				err = fmt.Errorf("minecraft is CRASH!! The error message is: %v", r)
			}
		}
	}()
	lg.callback(param)
	return nil
}
func (lg launchGame) putArgument(realJson map[string]any, result []string) ([]string, error) {
	mcid := Safe(realJson, "", "id").(string)
	if mcid == "" {
		return nil, NewMMCLLError(-108, "Cannot Find MC ID key in Version JSON!")
	}
	mainClass := Safe(realJson, "", "mainClass").(string)
	if mainClass == "" {
		return nil, NewMMCLLError(-109, "Cannot Find Main Class key in Version JSON!")
	}
	assetIndex := Safe(realJson, "", "assetIndex", "id").(string)
	if assetIndex == "" {
		return nil, NewMMCLLError(-110, "Cannot Find Asset Index key in Version JSON!")
	}
	if lg.additionalJvm != "" {
		s := strings.Split(lg.additionalJvm, " ")
		result = append(result, s...)
	}
	jvmArgs := JudgeArguments(realJson, "jvm")
	if len(jvmArgs) > 0 {
		result = append(result, jvmArgs...)
	} else {
		result = append(result, "-Djava.library.path=${natives_directory}", "-cp", "${classpath}")
	}
	// 判断第三方登录
	if lg.account.Online() == 3 {
		result = append(result, fmt.Sprintf("-javaagent:%s=%s", lg.account.AlijPath(), lg.account.Url()))
		result = append(result, "-Dauthlibinjector.side=client")
		if lg.account.Base() == "" {
			l1 := NewUrlMethod(lg.account.Url())
			l2, err := l1.GetDefault()
			if err != nil {
				return nil, NewMMCLLError(-211, "Cannot Get Skin Metadata!")
			}
			b := base64.StdEncoding.EncodeToString(l2)
			result = append(result, fmt.Sprintf("-Dauthlibinjector.yggdrasil.prefetched=%s", b))
		} else {
			result = append(result, fmt.Sprintf("-Dauthlibinjector.yggdrasil.prefetched=%s", lg.account.Base()))
		}
	}
	result = append(result, fmt.Sprintf("-Xmn%dm", lg.minMemory))
	result = append(result, fmt.Sprintf("-Xmx%dm", lg.maxMemory))
	result = append(result, mainClass)
	mcArgument := Safe(realJson, "", "minecraftArguments").(string)
	if mcArgument != "" {
		s := strings.Split(mcArgument, " ")
		result = append(result, s...)
	}
	gameArgs := JudgeArguments(realJson, "game")
	if len(gameArgs) > 0 {
		result = append(result, gameArgs...)
	}
	if !strings.Contains(lg.additionalGame, "--fullScreen") {
		result = append(result, "--width", strconv.Itoa(int(lg.windowWidth)), "--height", strconv.Itoa(int(lg.windowHeight)))
	}
	if index := ArrayIndexOf(result, "optifine.OptiFineForgeTweaker"); index > 0 {
		result = append(result[:index-1], result[index+1:]...)
		result = append(result, "--tweakClass", result[index])
	}
	if index := ArrayIndexOf(result, "optifine.OptiFineTweaker"); index > 0 {
		result = append(result[:index-1], result[index+1:]...)
		result = append(result, "--tweakClass", result[index])
	}
	libs, err := GetMCLibs(realJson, lg.rootPath, lg.versionPath, lg.isCheckLibraries)
	if err != nil {
		return nil, NewMMCLLError(-212, "Cannot Get All MC Libraries!")
	}
	lv := strings.ReplaceAll(LauncherVersion, ".", "")
	lv = strings.ReplaceAll(lv, "-", "")
	lv = strings.ReplaceAll(lv, "Alpha", "")
	lv = strings.ReplaceAll(lv, "Beta", "")
	lv = strings.ReplaceAll(lv, "Pre", "")
	mainJarName := filepath.Base(libs[len(libs)-1])
	cp := strings.Join(libs, "${classpath_separator}")
	// 下列真的是火葬场了（等待有缘人的修复吧~
	for i := range result {
		result[i] = strings.ReplaceAll(result[i], "${natives_directory}", filepath.Join(lg.versionPath, fmt.Sprintf("%s-%s-native", filepath.Base(lg.versionPath), LauncherName)))
		result[i] = strings.ReplaceAll(result[i], "${launcher_name}", LauncherName)
		result[i] = strings.ReplaceAll(result[i], "${launcher_version}", lv)
		result[i] = strings.ReplaceAll(result[i], "${classpath}", cp)
		result[i] = strings.ReplaceAll(result[i], "${version_name}", mcid)
		result[i] = strings.ReplaceAll(result[i], "${library_directory}", filepath.Join(lg.rootPath, "libraries"))
		result[i] = strings.ReplaceAll(result[i], "${auth_player_name}", lg.account.Name())
		result[i] = strings.ReplaceAll(result[i], "${game_directory}", lg.gamePath)
		result[i] = strings.ReplaceAll(result[i], "${assets_root}", filepath.Join(lg.rootPath, "assets"))
		result[i] = strings.ReplaceAll(result[i], "${assets_index_name}", assetIndex)
		result[i] = strings.ReplaceAll(result[i], "${auth_uuid}", lg.account.Uuid())
		result[i] = strings.ReplaceAll(result[i], "${uuid}", lg.account.Uuid())
		result[i] = strings.ReplaceAll(result[i], "${auth_access_token}", lg.account.AccessToken())
		result[i] = strings.ReplaceAll(result[i], "${user_type}", lg.account.AccType())
		result[i] = strings.ReplaceAll(result[i], "${version_type}", lg.customInfo)
		result[i] = strings.ReplaceAll(result[i], "${auth_session}", lg.account.Uuid())
		result[i] = strings.ReplaceAll(result[i], "${game_assets}", filepath.Join(lg.rootPath, "assets", "virtual", "legacy"))
		result[i] = strings.ReplaceAll(result[i], "${user_properties}", "{}")
		result[i] = strings.ReplaceAll(result[i], "${classpath_separator}", GetSeparator())
		result[i] = strings.ReplaceAll(result[i], "${primary_jar_name}", mainJarName)
	}
	return result, nil
}

// LaunchGame
// 新增参数：isStrict，用于手动指定是否动用 MMCLL 的参数检查。
// 如果你想自己在源代码里检查的话，你完全可以将该值设为 false 以跳过自带的 MMCLL 参数检查。
func LaunchGame(option LaunchOption, isStrict bool, callback func([]string)) error {
	ls := newLaunchStart(option, callback)
	if isStrict {
		if err := ls.checkError(); err != nil {
			return err
		}
	}
	if err := ls.launch(); err != nil {
		return err
	}
	return nil
}
