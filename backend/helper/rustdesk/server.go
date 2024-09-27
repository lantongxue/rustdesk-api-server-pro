package rustdesk

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/shirou/gopsutil/v3/process"
)

var (
	serverBinDir = GetRustdeskServerBinDir()

	hbbrPidFile = filepath.Join(serverBinDir, "hbbr.pid")
	hbbsPidFile = filepath.Join(serverBinDir, "hbbs.pid")
)

func GetRustdeskServerBinDir() string {
	pwd, _ := os.Getwd()
	return path.Join(pwd, "rustdesk-server")
}

func GetRustdeskServerBin() (hbbr, hbbs string) {
	dir := GetRustdeskServerBinDir()
	switch runtime.GOOS {
	case "windows":
		hbbr = path.Join(dir, "hbbr.exe")
		hbbs = path.Join(dir, "hbbs.exe")
	case "linux":
		hbbr = path.Join(dir, "hbbr")
		hbbs = path.Join(dir, "hbbs")
	default:
		// now, rustdesk-server only support windows and linux.
	}

	return hbbr, hbbs
}

func StartServer() (bool, error) {
	hbbr, hbbs := GetRustdeskServerBin()

	pHbbr := exec.Command(hbbr)
	pHbbr.Dir = serverBinDir
	err := pHbbr.Start()
	if err != nil {
		fmt.Println("hbbr start error:", err.Error())
		return false, err
	}
	err = os.WriteFile(hbbrPidFile, []byte(strconv.Itoa(pHbbr.Process.Pid)), os.ModePerm)
	if err != nil {
		fmt.Println("write hbbr pid file error:", err.Error())
		return false, err
	}

	pHbbs := exec.Command(hbbs)
	pHbbs.Dir = serverBinDir
	err = pHbbs.Start()
	if err != nil {
		fmt.Println("hbbs start error:", err.Error())
		return false, err
	}
	err = os.WriteFile(hbbsPidFile, []byte(strconv.Itoa(pHbbs.Process.Pid)), os.ModePerm)
	if err != nil {
		fmt.Println("write hbbs pid file error:", err.Error())
		return false, err
	}
	return true, nil
}

func read_server_pid() (hbbrPid, hbbsPid int) {
	hbbrPidBytes, err := os.ReadFile(hbbrPidFile)
	if err != nil {
		return -1, -1
	}
	hbbrPid, err = strconv.Atoi(string(hbbrPidBytes))
	if err != nil {
		return -1, -1
	}

	hbbsPidBytes, err := os.ReadFile(hbbsPidFile)
	if err != nil {
		return -1, -1
	}

	hbbsPid, err = strconv.Atoi(string(hbbsPidBytes))
	if err != nil {
		return -1, -1
	}

	return hbbrPid, hbbsPid
}

func StopServer() bool {

	hbbrPid, hbbsPid := read_server_pid()

	hbbr, _ := process.NewProcess(int32(hbbrPid))
	hbbs, _ := process.NewProcess(int32(hbbsPid))

	hbbr.Kill()
	hbbs.Kill()

	return true
}

func Status() (hbbrIsRunning, hbbsIsRunning bool) {
	hbbrPid, hbbsPid := read_server_pid()

	hbbr, _ := process.NewProcess(int32(hbbrPid))
	hbbs, _ := process.NewProcess(int32(hbbsPid))

	hbbrIsRunning, _ = hbbr.IsRunning()
	hbbsIsRunning, _ = hbbs.IsRunning()

	return hbbrIsRunning, hbbsIsRunning
}

func Keys() (public, private string) {
	publicKeysBytes, err := os.ReadFile(path.Join(serverBinDir, "id_ed25519.pub"))
	if err != nil {
		return "", ""
	}

	privateKeysBytes, err := os.ReadFile(path.Join(serverBinDir, "id_ed25519"))
	if err != nil {
		return "", ""
	}

	public = string(publicKeysBytes)

	private = string(privateKeysBytes)

	return public, private
}
