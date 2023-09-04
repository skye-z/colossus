package common

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
)

func UnzipTarGz(path, name string) error {
	// 打开 .tar.gz 文件
	file, err := os.Open(path + "/" + name)
	if err != nil {
		return err
	}
	defer file.Close()
	// 创建一个 gzip 读取器
	gzr, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzr.Close()
	// 创建一个 tar 读取器
	tr := tar.NewReader(gzr)
	// 遍历 tar 文件中的每一项
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // 文件结束
		}
		if err != nil {
			return err
		}
		// 获取文件名和模式
		name := hdr.Name
		if name[0:2] == "./" {
			name = name[2:]
		}
		mode := hdr.FileInfo().Mode()
		// 如果是目录
		if hdr.Typeflag == tar.TypeDir {
			unzipPath := path + "/" + name
			// 创建目录
			err = os.MkdirAll(unzipPath, mode)
			if err != nil {
				return err
			}
			continue // 继续下一项
		}
		// 如果是普通文件
		if hdr.Typeflag == tar.TypeReg {
			// 确保目标目录存在
			err = os.MkdirAll(path, mode)
			if err != nil {
				return err
			}

			unzipPath := path + "/" + name
			// 创建文件
			fw, err := os.Create(unzipPath)
			if err != nil {
				return err
			}
			defer fw.Close()

			if _, err = io.Copy(fw, tr); err != nil { // 复制文件内容
				return err
			}
			continue // 继续下一项
		}
	}
	return nil
}
