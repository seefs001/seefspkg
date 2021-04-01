package xfile

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// GetFileSize 获取文件大小
func GetFileSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

// GetFileExt 获取文件后缀
func GetFileExt(fileName string) string {
	return path.Ext(fileName)
}

// CheckFileExist 检查文件是否存在
func CheckFileExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

// CheckFilePermission 检查是否有权限
func CheckFilePermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// IsNotExistMkDir 不存在新建目录
func IsNotExistMkDir(src string) error {
	if exist := CheckFileExist(src); !exist {
		if err := MkDir(src); err != nil {
			return err
		}
	}

	return nil
}

// MkDir 新建目录
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// OpenFile 打开文件
func OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// IsBinary 是否为Binary
func IsBinary(content string) bool {
	for _, b := range content {
		if 0 == b {
			return true
		}
	}

	return false
}

// IsImg 是否为img
func IsImg(extension string) bool {
	ext := strings.ToLower(extension)

	switch ext {
	case ".jpg", ".jpeg", ".bmp", ".gif", ".png", ".svg", ".ico":
		return true
	default:
		return false
	}
}

// IsDir 是否为Dir
func IsDir(path string) bool {
	fio, err := os.Lstat(path)
	if os.IsNotExist(err) {
		return false
	}

	if err != nil {
		return false
	}

	return fio.IsDir()
}

// CopyFile CopyFile
func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}
	}

	return nil
}

// CopyDir 复制目录
func CopyDir(source string, dest string) (err error) {
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir
	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, err := os.Open(source)
	if err != nil {
		return err
	}

	defer directory.Close()

	objects, err := directory.Readdir(-1)
	if err != nil {
		return err
	}

	for _, obj := range objects {
		srcFilePath := filepath.Join(source, obj.Name())
		destFilePath := filepath.Join(dest, obj.Name())

		if obj.IsDir() {
			// create sub-directories - recursively
			err = CopyDir(srcFilePath, destFilePath)
			if err != nil {
				return err
			}
		} else {
			err = CopyFile(srcFilePath, destFilePath)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
