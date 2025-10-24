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
	homeDir, err := GetOtherDir()
	if err != nil {
		exePath, err := os.Executable()
		return mmcll.If(err != nil, "", filepath.Dir(exePath)).(string)
	}
	return homeDir
}

func GetPCL2Identify() string {
	return "DNS"
}

func GetUniqueAddress() string {
	cmd := CMD(GetCurrentDir(), "system_profiler", "SPHardwareDataType", "|", "grep", "Hardware UUID")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "ERR_MACOS"
	}
	re := regexp.MustCompile("[0-9a-f]{8}-([0-9a-f]{4}-){3}[0-9a-f]{12}")
	return strings.ToUpper(fmt.Sprintf("%x", md5.Sum([]byte(re.FindString(strings.ToLower(string(out)))))))
}
func GetOtherDir() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(currentUser.HomeDir, "Library", "Application Support"), nil
}

func PingCMD(ip string, timeout time.Duration) *exec.Cmd {
	return CMD(GetCurrentDir(), "ping", "-c", "1", "-W", fmt.Sprintf("%d", timeout.Seconds()), ip)
}

func CMD(dir string, name string, args ...string) *exec.Cmd {
	cmd := exec.Command(name, args...)
	cmd.Dir = dir
	return cmd
}

func openExp(fpath string) error {
	return exec.Command("open", fpath).Start()
}
