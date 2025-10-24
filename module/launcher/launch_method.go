package launcher

import (
	"NovaPlus/module/mmcll"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	runtime2 "runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type MCConfig struct {
	Path string `json:"path"`
	Name string `json:"name"`
}
type MCConfigs struct {
	Mc []MCConfig `json:"mc"`
}
type JavaConfig struct {
	Path    string `json:"path"`
	Version string `json:"version"`
	Arch    string `json:"arch"`
	Vendor  string `json:"vendor"`
}
type JavaConfigs struct {
	Java []JavaConfig `json:"java"`
}
type LaunchMethod struct {
	Ctx context.Context
}

func (lm *LaunchMethod) GetMCVersionConfig() ExceptionHandler[MCConfigs] {
	exception := func(err error) ExceptionHandler[MCConfigs] {
		return NewExceptionHandler(500, false, "File System Error: "+err.Error(), MCConfigs{Mc: []MCConfig{}})
	}
	cur := filepath.Join(GetCurrentDir(), "PCL.Nova", "config", "MCJson.json")
	err := EnsureConfigFile(cur)
	if err != nil {
		return exception(err)
	}
	if content, _ := mmcll.GetFile(cur); content == "" {
		err := mmcll.SetFile(cur, "{\n\t\"mc\": []\n}")
		if err != nil {
			return exception(err)
		}
		return NewExceptionHandler(201, true, "OK!", MCConfigs{Mc: []MCConfig{}})
	} else {
		var arr MCConfigs
		err = json.Unmarshal([]byte(content), &arr)
		if err != nil {
			return exception(err)
		}
		return NewExceptionHandler(200, true, "OK!", arr)
	}
}

func (lm *LaunchMethod) SetMCVersionConfig(mc MCConfigs) ExceptionHandler[any] {
	exception := func(err error) ExceptionHandler[any] {
		return NewExceptionHandler[any](500, false, "File System Error: "+err.Error(), nil)
	}
	cur := filepath.Join(GetCurrentDir(), "PCL.Nova", "config", "MCJson.json")
	err := EnsureConfigFile(cur)
	if err != nil {
		return exception(err)
	}
	res, err := json.MarshalIndent(mc, "", "\t")
	if err != nil {
		return exception(err)
	}
	err = mmcll.SetFile(cur, string(res))
	if err != nil {
		return exception(err)
	}
	return NewExceptionHandler[any](201, true, "Created!", nil)
}

func (lm *LaunchMethod) GetJavaConfig() ExceptionHandler[JavaConfigs] {
	exception := func(err error) ExceptionHandler[JavaConfigs] {
		return NewExceptionHandler(500, false, "File System Error: "+err.Error(), JavaConfigs{Java: []JavaConfig{}})
	}
	cur := filepath.Join(GetCurrentDir(), "PCL.Nova", "config", "JavaJson.json")
	err := EnsureConfigFile(cur)
	if err != nil {
		return exception(err)
	}
	if content, _ := mmcll.GetFile(cur); content == "" {
		err := mmcll.SetFile(cur, "{\n\t\"java\": []\n}")
		if err != nil {
			return exception(err)
		}
		return NewExceptionHandler(201, true, "Created!", JavaConfigs{Java: []JavaConfig{}})
	} else {
		var arr JavaConfigs
		err = json.Unmarshal([]byte(content), &arr)
		if err != nil {
			return exception(err)
		}
		return NewExceptionHandler(200, true, "OK!", arr)
	}
}
func (lm *LaunchMethod) SetJavaConfig(java JavaConfigs) ExceptionHandler[any] {
	exception := func(err error) ExceptionHandler[any] {
		return NewExceptionHandler[any](500, false, "File System Error: "+err.Error(), nil)
	}
	cur := filepath.Join(GetCurrentDir(), "PCL.Nova", "config", "JavaJson.json")
	err := EnsureConfigFile(cur)
	if err != nil {
		return exception(err)
	}
	res, err := json.MarshalIndent(java, "", "\t")
	if err != nil {
		return exception(err)
	}
	err = mmcll.SetFile(cur, string(res))
	if err != nil {
		return exception(err)
	}
	return NewExceptionHandler[any](201, true, "Created!", nil)
}

// GetMCAllVersion 从 versions 目录下，提取出所有的 MC 版本并返回给前端
func (lm *LaunchMethod) GetMCAllVersion(rootPath string) []string {
	rootPath = filepath.Join(rootPath, "versions")
	stat, err := os.Stat(rootPath)
	if err != nil || !stat.IsDir() {
		return []string{}
	}
	dir, err := os.Open(rootPath)
	if err != nil {
		return []string{}
	}
	defer dir.Close()
	dirs, err := dir.Readdir(-1)
	if err != nil {
		return []string{}
	}
	var result []string
	for _, d := range dirs {
		if !d.IsDir() {
			continue
		}
		result = append(result, filepath.Join(rootPath, d.Name()))
	}
	return result
}

func (lm *LaunchMethod) GetCurrentMinecraftDir() string {
	return filepath.Join(GetCurrentDir(), ".minecraft")
}

func (lm *LaunchMethod) GetJavaInfo(path string) ExceptionHandler[JavaConfig] {
	cmd := CMD(GetCurrentDir(), path, "-XshowSettings:properties", "-version")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return NewExceptionHandler(500, false, "Command failed with: "+err.Error(), JavaConfig{})
	}
	outStr := string(out)
	outStrSlice := strings.Split(outStr, "\n")
	result := JavaConfig{}
	result.Path = path
	for _, line := range outStrSlice {
		lc := strings.Split(line, "=")
		if len(lc) == 2 {
			lc0 := strings.Trim(lc[0], " \n\r")
			switch lc0 {
			case "java.vendor":
				result.Vendor = strings.Trim(lc[1], " \n\r")
			case "java.version":
				result.Version = strings.Trim(lc[1], " \n\r")
			case "os.arch":
				result.Arch = strings.Trim(lc[1], " \n\r")
			}
		}
	}
	return NewExceptionHandler(200, true, "OK!", result)
}
func (lm *LaunchMethod) LaunchGame(exportArguments bool, showAccessToken bool) ExceptionHandler[string] {
	exception := func(message string, data string) ExceptionHandler[string] {
		return NewExceptionHandler(400, false, message, data)
	}
	rw := ReaderWriter{
		Ctx: lm.Ctx,
	}
	acc := AccountMethod{
		Ctx: lm.Ctx,
	}
	configIniPath := rw.GetConfigIniPath()
	accountRaw := acc.GetAccountConfig()
	if !accountRaw.Status {
		return exception("JSON Parse Exception: "+accountRaw.Message, "Account JSON Is Error!")
	}
	accounts := accountRaw.Data.Accounts
	accIndexStr := rw.ReadConfig(rw.GetOtherIniPath(), "Account", "SelectAccount")
	accIndex, err := strconv.Atoi(accIndexStr)
	if err != nil {
		return exception("Convert Exception: "+err.Error(), "Cannot convert precon index to int")
	}
	account, ok := mmcll.SafeGet(accounts, accIndex)
	if !ok {
		return exception("Array Index Out Of Bounds", "Account array index out of bounds")
	}
	var accountStruct mmcll.LaunchAccount
	if account.Server != nil && *account.Server != "" {
		fmt.Println("aaa")
		home, err := GetOtherDir()
		if err != nil {
			return exception("Get Home Dir Exception: "+err.Error(), "Cannot get Home directory")
		}
		authlib_path := filepath.Join(home, mmcll.If(runtime2.GOOS == "windows", "PCL.Nova", ".PCL.Nova").(string), "Authlib-Injector.jar")
		av := rw.ReadConfig(rw.GetOtherIniPath(), "Account", "AuthLibInjectorVersion")
		sr := rw.ReadConfig(rw.GetConfigIniPath(), "Version", "SelectDownloadSource")
		pi, err := strconv.Atoi(sr)
		if err != nil || pi < 1 || pi > 3 {
			pi = 1
		}
		if pi == 3 {
			//TODO: 智能切换下载源~
			pi = 2
		}
		url := mmcll.NewUrlMethod(AuthlibDownloadLink[pi-1])
		gt, err := url.GetDefault()
		if err != nil {
			return exception("HTTP Get Exception: "+err.Error(), "Cannot get default AuthlibInjector version")
		}
		var gn map[string]any
		if err := json.Unmarshal(gt, &gn); err != nil {
			return exception("JSON Parse Exception: "+err.Error(), "Cannot parse AuthlibInjector version")
		}
		rav := mmcll.Safe(gn, "", "version").(string)
		downloadAuthlibInjector := func() error {
			bt := rw.ReadConfig(rw.GetConfigIniPath(), "Version", "ThreadBiggest")
			btn, err := strconv.Atoi(bt)
			if err != nil || btn < 1 || btn > 256 {
				btn = 64
			}
			u := mmcll.Safe(gn, "", "download_url").(string)
			if u == "" {
				return fmt.Errorf("cannot find download url in metadata json")
			}
			proxyType := rw.ReadConfig(rw.GetConfigIniPath(), "Download", "ProxyType")
			proxyUrl := rw.ReadConfig(rw.GetConfigIniPath(), "Download", "ProxyUrl")
			proxyPort := rw.ReadConfig(rw.GetConfigIniPath(), "Download", "ProxyPort")
			proxyUsername := rw.ReadConfig(rw.GetConfigIniPath(), "Download", "ProxyUsername")
			proxyPassword := rw.ReadConfig(rw.GetConfigIniPath(), "Download", "ProxyPassword")
			isSocks := proxyType == "3"
			isHttps := proxyType == "2"
			mm := mmcll.NewDownloadSingleWithProxy(u, authlib_path, btn, 0, mmcll.NewProxy(proxyUrl, proxyPort, proxyUsername, proxyPassword, isHttps, isSocks))
			mm.StartDownload(func(url string, retryCount int, msg string, code int, total int64, current int64) {
				fmt.Println(url, retryCount, msg, code, total, current)
			})
			rw.WriteConfig(rw.GetOtherIniPath(), "Account", "AuthLibInjectorVersion", rav)
			return nil
		}
		if rav != av {
			if err := downloadAuthlibInjector(); err != nil {
				return exception("Cannot Download AuthlibInjector! Error: "+err.Error(), "Cannot Download AuthlibInjector!")
			}
		}
		sha256, err := GetFileSHA256(authlib_path)
		if err != nil {
			return exception("Get AuthlibInjector SHA256 Exception: "+err.Error(), "Cannot get AuthlibInjector SHA256")
		}
		if mmcll.Safe(gn, "", "checksums", "sha256") != sha256 {
			if err := downloadAuthlibInjector(); err != nil {
				return exception("Cannot Download AuthlibInjector! Error: "+err.Error(), "Cannot Download AuthlibInjector!")
			}
		}
		accountStruct = mmcll.NewLaunchAccountThirdParty(account.Name, account.UUID, *account.AccessToken, "", *account.Server, authlib_path)
	} else if account.AccessToken != nil && *account.AccessToken != "" {
		accountStruct = mmcll.NewLaunchAccountMicrosoft(account.Name, account.UUID, *account.AccessToken)
	} else if account.Name != "" && account.UUID != "" {
		accountStruct = mmcll.NewLaunchAccountOffline(account.Name, account.UUID)
	} else {
		return exception("Cannot find any data!", "Account Name or UUID is Empty!")
	}
	javaRaw := lm.GetJavaConfig()
	if !javaRaw.Status {
		return exception("JSON Parse Exception: "+javaRaw.Message, "Account Name or UUID is Empty!")
	}
	javas := javaRaw.Data.Java
	//javas := lm.GetJavaConfig().Java
	javaIndexStr := rw.ReadConfig(configIniPath, "Java", "SelectJava")
	javaIndex, err := strconv.Atoi(javaIndexStr)
	if err != nil {
		return exception("Convert Exception: "+err.Error(), "Cannot convert java index to int")
	}
	java, ok := mmcll.SafeGet(javas, javaIndex)
	if !ok {
		return exception("Array Index Out Of Bounds", "Java array index out of bounds")
	}
	javaPath := java.Path
	rootIndexStr := rw.ReadConfig(configIniPath, "MC", "SelectMC")
	rootIndex, err := strconv.Atoi(rootIndexStr)
	if err != nil {
		return exception("Convert Exception: "+err.Error(), "Cannot convert root index to int")
	}
	var root MCConfig
	if rootIndex == 0 {
		root = MCConfig{
			Path: lm.GetCurrentMinecraftDir(),
		}
	} else {
		rootRaw := lm.GetMCVersionConfig()
		if !rootRaw.Status {
			return exception("JSON Parse Exception: "+rootRaw.Message, "MC Version Config JSON Is Error!")
		}
		roots := rootRaw.Data.Mc
		root, ok = mmcll.SafeGet(roots, rootIndex-1)
		if !ok {
			return exception("Array Index Out Of Bounds", "Root array index out of bounds")
		}
	}
	rootPath := root.Path
	vers := lm.GetMCAllVersion(rootPath)
	verIndexStr := rw.ReadConfig(configIniPath, "MC", "SelectVer")
	verIndex, err := strconv.Atoi(verIndexStr)
	if err != nil {
		return exception("Convert Exception: "+err.Error(), "Cannot convert root version index to int")
	}
	ver, ok := mmcll.SafeGet(vers, verIndex)
	if !ok {
		return exception("Array Index Out Of Bounds", "Version array index out of bounds")
	}
	isolation := rw.ReadConfig(configIniPath, "Document", "SelectIsolation") == "4"
	gamePath := mmcll.If(isolation, ver, rootPath).(string)
	option := mmcll.NewLaunchOption(accountStruct, javaPath, rootPath, ver, gamePath)
	heightStr := rw.ReadConfig(configIniPath, "Document", "WindowHeight")
	if height, err := strconv.Atoi(heightStr); err == nil {
		option.SetWindowHeight(uint32(height))
	}
	widthStr := rw.ReadConfig(configIniPath, "Document", "WindowWidth")
	if width, err := strconv.Atoi(widthStr); err == nil {
		option.SetWindowWidth(uint32(width))
	}
	maxMemoryStr := rw.ReadConfig(configIniPath, "Document", "MaxMemory")
	if maxMemory, err := strconv.Atoi(maxMemoryStr); err == nil {
		option.SetMaxMemory(uint32(maxMemory))
	}
	customInfo := rw.ReadConfig(configIniPath, "Version", "CustomInfo")
	if customInfo != "" {
		option.SetCustomInfo(customInfo)
	} else {
		option.SetCustomInfo("PCL.Nova.Plus")
	}
	isCheckLibraries := rw.ReadConfig(configIniPath, "Document", "IsCheckLibraries")
	if isCheckLibraries != "1" {
		option.SetIsCheckLibraries(false)
	}
	additionalJVM := rw.ReadConfig(configIniPath, "Version", "AdditionalJVM")
	if additionalJVM != "" {
		option.SetAdditionalJvm(additionalJVM)
	}
	additionalGame := rw.ReadConfig(configIniPath, "Version", "AdditionalGame")
	if additionalGame != "" {
		option.SetAdditionalGame(additionalGame)
	}
	err = mmcll.LaunchGame(*option, false, func(back []string) {
		runtime.EventsEmit(lm.Ctx, "launch_success")
		if exportArguments {
			if showAccessToken {
				if e := mmcll.SetFile(filepath.Join(GetCurrentDir(),
					"PCL.Nova",
					"args",
					mmcll.If(runtime2.GOOS == "windows",
						"launch.bat",
						"launch.sh").(string)),
					"\""+strings.Join(back, "\" \"")+"\""); e != nil {
					panic(e)
				}
			} else {
				args := "\"" + strings.Join(back, "\" \"") + "\""
				l1 := args[strings.Index(args, "\"--accessToken\"")+16:]
				l2 := strings.ReplaceAll(args, l1, "")
				l3 := l1[strings.Index(l1, " "):]
				params := l2 + "\"********************************\"" + l3
				if e := mmcll.SetFile(filepath.Join(GetCurrentDir(),
					"PCL.Nova",
					"args",
					mmcll.If(runtime2.GOOS == "windows",
						"launch.bat",
						"launch.sh").(string)),
					params); e != nil {
					panic(e)
				}
			}
		} else {
			getOutput := func(reader *bufio.Reader) {
				for {
					str, err := reader.ReadString('\n')
					if err != nil {
						if err == io.EOF {
							break
						}
						panic(err)
					}
					runtime.EventsEmit(lm.Ctx, "launch_log", str)
				}
			}
			cmd := CMD(option.RootPath(), back[0], back[1:]...)
			var wg sync.WaitGroup
			wg.Add(2)
			stdout, err := cmd.StdoutPipe()
			if err != nil {
				panic(err)
			}
			readout := bufio.NewReader(stdout)
			stderr, err := cmd.StderrPipe()
			if err != nil {
				panic(err)
			}
			readerr := bufio.NewReader(stderr)
			go func() {
				defer wg.Done()
				getOutput(readout)
			}()
			go func() {
				defer wg.Done()
				getOutput(readerr)
			}()
			e := cmd.Start()
			if e != nil {
				panic(e)
			}
			wg.Wait()
		}
	})
	if err != nil {
		return exception("Launch some Error occurred!", "You have some error! please try to view it: "+err.Error())
	}
	return NewExceptionHandler(200, true, "Ok!", "")
}
