package lsof

import (
	"os"
	"path/filepath"
	"strconv"
)

// LsofUsePid 获取pid的打开文件
func LsofUsePid(pid int) []string {
	openFiles := make([]string, 0)
	fdPath := filepath.Join("/proc", strconv.Itoa(pid), "fd")
	fFiles, _ := os.ReadDir(fdPath)
	for _, f := range fFiles {
		fdPath, err := os.Readlink(filepath.Join(fdPath, f.Name()))
		if err != nil {
			continue
		}
		openFiles = append(openFiles, fdPath)
	}
	return openFiles
}

// OpenFilePids 获取打开文件的pids
func OpenFilePids() map[string][]string {
	files, _ := os.ReadDir("/proc")
	openFiles := make(map[string][]string)
	for _, f := range files {
		fName := f.Name()
		m, _ := filepath.Match("[0-9]*", fName)
		if f.IsDir() && m {
			fPid := fName
			fdPath := filepath.Join("/proc", f.Name(), "fd")
			fFiles, _ := os.ReadDir(fdPath)
			for _, f := range fFiles {
				fdPath, err := os.Readlink(filepath.Join(fdPath, f.Name()))
				if err != nil {
					continue
				}
				if v, ok := openFiles[fdPath]; ok {
					openFiles[fdPath] = append(v, fPid)
				} else {
					openFiles[fdPath] = []string{fPid}
				}
			}
		}
	}
	return openFiles
}

// FileIsOpen 判断文件是否打开
func FileIsOpen(fullPath string) bool {
	files, _ := os.ReadDir("/proc")
	for _, f := range files {
		fName := f.Name()
		m, _ := filepath.Match("[0-9]*", fName)
		if f.IsDir() && m {
			fdPath := filepath.Join("/proc", fName, "fd")
			fFiles, _ := os.ReadDir(fdPath)
			for _, f1 := range fFiles {
				fdPath, err := os.Readlink(filepath.Join(fdPath, f1.Name()))
				if err != nil {
					continue
				}
				if fdPath == fullPath {
					return true
				}
			}
		}
	}
	return false
}
