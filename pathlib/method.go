// Package pathlib 自定义的方法
package pathlib

import (
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
)

// MkdirAll 创建路径：包含父路径，一般 0777
func (h *Path) MkdirAll(mode os.FileMode) error {
	return os.MkdirAll(h.Path, mode)
}

// FindFunc 配合 FindFiles 使用
type FindFunc func(path *Path, info fs.FileInfo, err error) error

// FindFiles 查找指定文件，返回路径处理器
func (h *Path) FindFiles(pattern string, f FindFunc) (err error) {
	var re *regexp.Regexp

	re, err = regexp.Compile(pattern)
	if err != nil {
		return err
	}

	// 遍历查找
	err = h.Walk(func(path string, info fs.FileInfo, err error) error {
		if re.MatchString(path) {
			err = f(NewPath(path), info, err)
			if err != nil {
				return err
			}
		}
		return err
	})

	return err
}

// Info 文件详情
func (h *Path) Info() (os.FileInfo, error) {
	return os.Stat(h.Path)
}

// IsFile 判断是否是文件：存在的
func (h *Path) IsFile() bool {
	file, err := h.Info()
	if err != nil {
		return false
	}

	return !file.IsDir()
}

// IsDir 判断是否是文件夹：存在的
func (h *Path) IsDir() bool {
	file, err := h.Info()
	if err != nil {
		return false
	}

	return file.IsDir()
}

// IsMatch 正则匹配路径
func (h *Path) IsMatch(pattern string) bool {
	matchString, err := regexp.MatchString(pattern, h.Path)
	if err != nil {
		return false
	}
	return matchString
}

// Parent 获取路径的父路径
func (h *Path) Parent() *Path {
	return NewPath(filepath.Dir(h.Path))
}

// Name 获取名字：不分文件文件夹，斜杠结尾将会去除斜杠
func (h *Path) Name() string {
	_, name := filepath.Split(h.Path)
	return name
}

// Exists 判断是否存在
func (h *Path) Exists() bool {
	_, err := h.Info()
	if err != nil {
		return false
	}

	return true
}

// Join 合并路径：不修改原来的，类似创建副本
func (h *Path) Join(paths ...string) *Path {
	newPath := NewPath(h.Path)

	for _, p := range paths {
		newPath.Path = filepath.Join(newPath.Path, p)
	}

	return newPath
}

// ShowDir 返回文件夹列表：大文件建议 walk、iter
func (h *Path) ShowDir() (allPaths []*Path, err error) {
	var dir *os.File

	if h.Exists() {
		if h.IsDir() {
			var names []string

			dir, err = os.Open(h.Path)
			if err != nil {
				return
			}

			names, err = dir.Readdirnames(0) // <=0 返回所有
			if err != nil {
				return
			}

			for _, f := range names {
				allPaths = append(allPaths, NewPath(filepath.Join(h.Path, f)))
			}

			defer func(dir *os.File) {
				err = dir.Close()
				if err != nil {
				}
			}(dir)

			return
		}
		return nil, NotDirErr{Path: h}

	}

	return nil, NotExistsErr{Path: h}
}

// Rename 重命名：真的只针对名字，输入新名字即可（含后缀）
// name 新的名字
// override 是否存在即覆盖，为 false 时，重复将会报 ExistsErr err
func (h *Path) Rename(name string, override bool) (err error) {
	if h.Exists() {
		newPath := h.Parent().Join(name)
		if !override && newPath.Exists() {
			return ExistsErr{Path: h}
		} else {
			return os.Rename(h.Path, newPath.Path)
		}
	}
	return NotExistsErr{Path: h}
}

// Move 移动：包含路径及名字
// toPath：全路径，含有名字
// override 是否存在即覆盖，为 false 时，重复将会报 ExistsErr err
func (h *Path) Move(toPath string, override bool) error {
	if !override && NewPath(toPath).Exists() {
		return ExistsErr{Path: h}
	}
	return os.Rename(h.Path, toPath)
}

// Delete 删除：是文件夹则会整个文件夹及内部文件都被删除
func (h *Path) Delete() error {
	if h.IsFile() {
		return os.Remove(h.Path)
	} else {
		return os.RemoveAll(h.Path)
	}
}

// IterFunc Iter 的接收函数
type IterFunc func(path *Path, err error) error

// iter Iter 的迭代函数
func (h *Path) iter(f IterFunc) (err error) {
	var file *os.File
	var names []string

	// 打开文件夹
	file, err = os.Open(h.Path)
	if err != nil {
		return err
	}

	for {
		// 每次读一个文件
		names, err = file.Readdirnames(1)
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return err
		}

		// 构造新 Path
		newPath := NewPath(h.Join(names[0]).Path)

		// 读完立即处理
		err = f(newPath, err)
		if err != nil {
			return err
		}

		// 判断是否是文件夹继续处理
		if newPath.IsDir() {
			err = newPath.iter(f)
			if err != nil && !errors.Is(err, io.EOF) {
				return err
			}
		}
	}

	return nil
}

// Iter 深度迭代
// 和 walk 功能一致
// walk 是一次性读取文件夹的所有再去遍历
// iter 是逐个读取，超大文件会好一些
func (h *Path) Iter(f IterFunc) (err error) {
	if h.IsDir() {
		return h.iter(f)
	}
	return f(h, nil)
}
