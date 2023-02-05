/*
	文件操作
	包含文件的集合操作和文件的便捷操作，达到一个包处理文件的所有事情

	注意：
		针对已存在的，不存在的在路径里面
		错误使用方式 Err("文件不存在或不是文件：" + inPath)
		路径操作对 win 都不友好，最好将路径的斜杠都转换为 /
*/

package file

import (
	"io"
	"log"
	"os"
)

// Info 文件详情
func Info(p string) (os.FileInfo, error) {
	return os.Stat(p)
}

// IsFile 判断是不是文件，必须是存在的文件
func IsFile(p string) bool {
	info, err := Info(p)
	if err != nil {
		return false
	}

	return !info.IsDir()
}

// IsDir 判断是不是目录，必须是存在的目录
func IsDir(p string) bool {
	info, err := Info(p)
	if err != nil {
		return false
	}

	return info.IsDir()
}

// Exists 判断文件是否存在
func Exists(p string) bool {
	_, err := Info(p)
	if err != nil {
		return false
	}

	return true
}

// Open 打开文件
func Open(p string) (*os.File, error) {
	return os.Open(p)
}

// Copy 复制文件
func Copy(inPath string, outPath string) error {
	if IsFile(inPath) {
		// 读取源文件
		inFile, err := Open(inPath)
		if err != nil {
			return err
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {

			}
		}(inFile)

		// 创建新文件
		outFile, err := os.Create(outPath)
		if err != nil {
			return err
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {

			}
		}(outFile)

		// 拷贝源文件
		_, err = io.Copy(outFile, inFile)
		if err != nil {
			return err
		}

		// flush 到硬盘
		return outFile.Sync()
	} else {
		return NotExists("文件不存在或不是文件：" + inPath)
	}
}

// WriteByte 以 byte 形式写入文件
func WriteByte(f *os.File, b []byte) error {
	_, err := f.Write(b)
	return err
}

// WriteStr 以 string 形式写入文件
func WriteStr(f *os.File, s string) error {
	_, err := f.WriteString(s)
	return err
}

// Read 读取文件：一次性读取
func Read(p string) ([]byte, error) {
	if IsFile(p) {
		file, err := os.Open(p)
		if err != nil {
			return nil, err
		}

		buf := make([]byte, 1024)
		for {
			_, err := file.Read(buf)
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln(err)
				return nil, err
			}
		}
		return buf, nil
	} else {
		return nil, NotExists("文件不存在或不是文件：" + p)
	}
}
