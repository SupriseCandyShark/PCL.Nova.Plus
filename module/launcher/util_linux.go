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
	"strings"
	"time"
)

func GetCurrentDir() string {
	exePath, err := os.Executable()
	return mmcll.If(err != nil, "", filepath.Dir(exePath)).(string)
}

func GetPCL2Identify() string {
	return "DNS"
}

func GetUniqueAddress() string {
	out, err := mmcll.GetFile("/sys/class/dmi/id/product_uuid")
	if err != nil {
		return "ERR_LINUX"
	}
	re := regexp.MustCompile("[0-9a-f]{8}-([0-9a-f]{4}-){3}[0-9a-f]{12}")
	return strings.ToUpper(fmt.Sprintf("%x", md5.Sum([]byte(re.FindString(strings.ToLower(string(out)))))))
}
func GetOtherDir() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	return currentUser.HomeDir, nil
}
func PingCMD(ip string, timeout time.Duration) *exec.Cmd {
	return CMD(GetCurrentDir(), "ping", "-c", "1", "-W", fmt.Sprintf("%.0f", timeout.Seconds()), ip)
}

func CMD(dir string, name string, args ...string) *exec.Cmd {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	return cmd
}

func openExp(fpath string) error {
	// 尝试检测桌面环境并打开文件管理器
	desktopEnv := os.Getenv("XDG_CURRENT_DESKTOP")
	desktopEnv = strings.ToLower(desktopEnv)

	// 根据桌面环境选择相应的文件管理器
	switch {
	case strings.Contains(desktopEnv, "gnome"):
		return exec.Command("nautilus", fpath).Start()
	case strings.Contains(desktopEnv, "kde"):
		return exec.Command("dolphin", fpath).Start()
	case strings.Contains(desktopEnv, "xfce"):
		return exec.Command("thunar", fpath).Start()
	case strings.Contains(desktopEnv, "mate"):
		return exec.Command("caja", fpath).Start()
	case strings.Contains(desktopEnv, "lxde"):
		return exec.Command("pcmanfm", fpath).Start()
	case strings.Contains(desktopEnv, "cinnamon"):
		return exec.Command("nemo", fpath).Start()
	case strings.Contains(desktopEnv, "pantheon"):
		return exec.Command("io.elementary.files", fpath).Start()
	case strings.Contains(desktopEnv, "budgie"):
		return exec.Command("nautilus", fpath).Start()
	default:
		// 尝试使用xdg-open作为后备方案
		if err := exec.Command("xdg-open", fpath).Start(); err != nil {
			return fmt.Errorf("无法识别桌面环境或缺少文件管理器: %s", desktopEnv)
		}
		return nil
	}
}
