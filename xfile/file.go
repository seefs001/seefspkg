package xfile

import (
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
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
	if exist := CheckFileExist(src); exist == false {
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
