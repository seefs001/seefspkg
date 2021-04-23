package xzip

import (
	"archive/zip"
	"log"

	"io"
	"os"
	"path/filepath"
	"strings"
)

// source:http://liupzmin.com/2019/08/02/golang/working-zip-archive/

//Zip 压缩为zip格式
//zipit("/tmp/report.txt", "/tmp/report-2015.zip", "*.log")
//source为要压缩的文件或文件夹, 绝对路径和相对路径都可以
//target是目标文件
//filter是过滤正则(Golang 的 包 path.Match)
func Zip(source, target, filter string) error {
	var err error
	if isAbs := filepath.IsAbs(source); !isAbs {
		source, err = filepath.Abs(source) // 将传入路径直接转化为绝对路径
		if err != nil {
			return err
		}
	}
	//创建zip包文件
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}

	defer func() {
		if err := zipfile.Close(); err != nil {
			log.Println("*File close error: %s, file: %s", err.Error(), zipfile.Name())
		}
	}()

	//创建zip.Writer
	zw := zip.NewWriter(zipfile)

	defer func() {
		if err := zw.Close(); err != nil {
			log.Println("zipwriter close error: %s", err.Error())
		}
	}()

	info, err := os.Stat(source)
	if err != nil {
		return err
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}

	err = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		//将遍历到的路径与pattern进行匹配
		ism, err := filepath.Match(filter, info.Name())

		if err != nil {
			return err
		}
		//如果匹配就忽略
		if ism {
			return nil
		}
		//创建文件头
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		if baseDir != "" {
			header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}
		//写入文件头信息
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}
		//写入文件内容
		file, err := os.Open(path)
		if err != nil {
			return err
		}

		defer func() {
			if err := file.Close(); err != nil {
				log.Println("*File close error: %s, file: %s", err.Error(), file.Name())
			}
		}()
		_, err = io.Copy(writer, file)

		return err
	})

	if err != nil {
		return err
	}

	return nil
}

//Unzip 解压zip
// unzip("/tmp/report-2015.zip", "/tmp/reports/")
func Unzip(archive, target string) error {
	reader, err := zip.OpenReader(archive)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(target, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		unzippath := filepath.Join(target, file.Name)
		if file.FileInfo().IsDir() {
			err := os.MkdirAll(unzippath, file.Mode())
			if err != nil {
				return err
			}
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(unzippath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}
