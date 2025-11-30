package launcher

import (
	"NovaPlus/module/mmcll"
	"crypto/md5"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/shirou/gopsutil/mem"
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/registry"
)

func GetCurrentDir() string {
	exePath, err := os.Executable()
	return mmcll.If(err != nil, "", filepath.Dir(exePath)).(string)
}

// getTimestampIdentify 获取 PCL 身份标识（根据时间戳+当前可用虚拟内存随机生成）（私有函数～）
func getTimestampIdentify() string {
	kernel32 := windows.NewLazySystemDLL("kernel32.dll")
	proc := kernel32.NewProc("GetTickCount")
	ret, _, _ := proc.Call()
	vm, _ := mem.VirtualMemory()
	result := strconv.FormatUint(uint64(ret)+2147483651, 10) + strconv.FormatUint(vm.Available, 10)
	return result
}

func GetOtherDir() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(currentUser.HomeDir, "AppData", "Roaming"), nil
}

func GetPCL2Identify() string {
	key, _ := registry.OpenKey(registry.LOCAL_MACHINE, "SYSTEM\\HardwareConfig", registry.READ)
	defer key.Close()
	str, _, err := key.GetStringValue("LastConfig")
	if err != nil {
		return "ERR_1"
	}
	str = strings.Trim(strings.ToUpper(str), "{}")
	key2, err := registry.OpenKey(registry.CURRENT_USER, "Software\\PCL", registry.WRITE|registry.READ|registry.CREATE_SUB_KEY|registry.SET_VALUE)
	if err != nil {
		key2, _, _ = registry.CreateKey(registry.CURRENT_USER, "Software\\PCL", registry.WRITE|registry.READ|registry.CREATE_SUB_KEY|registry.SET_VALUE)
	}
	defer key2.Close()
	str2, _, err := key2.GetStringValue("Identify")
	if err != nil {
		str2 = getTimestampIdentify()
		if err := key2.SetStringValue("Identify", str2); err != nil {
			println(err)
			return "ERR_2"
		}
	}
	str3 := StrFill(strconv.FormatUint(GetHash(str+str2), 16), "7", 16)
	return Mid(str3, 5, 4) + "-" + Mid(str3, 13, 4) + "-" + Mid(str3, 1, 4) + "-" + Mid(str3, 9, 4)
}

func GetUniqueAddress() string {
	cmd := CMD(GetCurrentDir(), "wmic", "csproduct", "get", "uuid")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "ERR_WINDOWS"
	}
	re := regexp.MustCompile("[0-9a-f]{8}-([0-9a-f]{4}-){3}[0-9a-f]{12}")
	return strings.ToUpper(fmt.Sprintf("%x", md5.Sum([]byte(re.FindString(strings.ToLower(string(out)))))))
}

func PingCMD(ip string, timeout time.Duration) *exec.Cmd {
	return CMD(GetCurrentDir(), "ping", "-n", "1", "-w", fmt.Sprintf("%d", timeout.Milliseconds()), ip)
}
func CMD(dir string, name string, args ...string) *exec.Cmd {
	cmd := exec.Command(name, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}
	cmd.Dir = dir
	return cmd
}
func openExp(fpath string) error {
	return exec.Command("explorer", fpath).Start()
}
