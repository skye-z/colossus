package common

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
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

func ZipTarGz(dir, name, outName string) error {
	// 创建输出文件
	out, err := os.Create(dir + "/" + outName)
	if err != nil {
		return err
	}
	defer out.Close()
	// 创建 gzip 写入器
	gw := gzip.NewWriter(out)
	defer gw.Close()
	// 创建 tar 写入器
	tw := tar.NewWriter(gw)
	defer tw.Close()
	// 遍历目录并存入压缩包
	return filepath.Walk(dir+"/"+name, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 跳过根目录
		if path == dir {
			return nil
		}
		// 创建信息标头
		hdr, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return err
		}
		// 设置相对根目录的标头名称
		hdr.Name = path[len(dir)+1:]
		// 标头写入压缩包
		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}
		// 常规文件直接写入
		if info.Mode().IsRegular() {
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()
			if _, err := io.Copy(tw, f); err != nil {
				return err
			}
		}
		return nil
	})
}
