package xfile

import (
	"crypto/sha512"
	"encoding/hex"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// GetFileSize get file size
func GetFileSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

// GetFileExt get file ext
func GetFileExt(fileName string) string {
	return path.Ext(fileName)
}

// CheckFileExist check file exist
func CheckFileExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

// CheckFilePermission check file permission
func CheckFilePermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

// IsNotExistMkDir is not exist then create dir
func IsNotExistMkDir(src string) error {

	err := os.Mkdir(src, os.ModePerm)
	if err != nil {
		if !os.IsExist(err) {
			return err
		}
	}
	return nil
}

// MkDir make dir
func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// OpenFile open file
func OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// IsBinary is binary
func IsBinary(content string) bool {
	for _, b := range content {
		if 0 == b {
			return true
		}
	}

	return false
}

// IsImg is image
func IsImg(extension string) bool {
	ext := strings.ToLower(extension)

	switch ext {
	case ".jpg", ".jpeg", ".bmp", ".gif", ".png", ".svg", ".ico":
		return true
	default:
		return false
	}
}

// IsDir is dir
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

// CopyFile copy file
func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer func(sourcefile *os.File) {
		err := sourcefile.Close()
		if err != nil {
			return
		}
	}(sourcefile)

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer func(destfile *os.File) {
		err := destfile.Close()
		if err != nil {
			return
		}
	}(destfile)

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}
	}

	return nil
}

// CopyDir copy dir
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

	defer func(directory *os.File) {
		err := directory.Close()
		if err != nil {
			return
		}
	}(directory)

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

// GetFileMd5 get file md5
func GetFileMd5(file *multipart.FileHeader) (string, error) {
	open, err := file.Open()
	if err != nil {
		return "", err
	}
	defer func(open multipart.File) {
		err := open.Close()
		if err != nil {
			return
		}
	}(open)

	h := sha512.New()
	if _, err := io.Copy(h, open); err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			return
		}
	}(src)

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			return
		}
	}(out)

	_, err = io.Copy(out, src)
	return err
}
