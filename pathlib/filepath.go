// Package pathlib filepath 包的部分覆写，返回值也是一致的
package pathlib

import (
	"io/fs"
	"path/filepath"
)

// IsAbs 判断是否是绝对路径
func (h *Handler) IsAbs() bool {
	return filepath.IsAbs(h.Path)
}

// Split 分割成路径和文件名：斜杠结尾则文件名为空
func (h *Handler) Split() (dir, file string) {
	return filepath.Split(h.Path)
}

// Abs 返回绝对路径
func (h *Handler) Abs() (string, error) {
	return filepath.Abs(h.Path)
}

// Walk 迭代文件夹：含有 FileInfo 信息
func (h *Handler) Walk(fn filepath.WalkFunc) error {
	return filepath.Walk(h.Path, fn)
}

// WalkDir 迭代文件：含有 DirEntry 信息
func (h *Handler) WalkDir(fn fs.WalkDirFunc) error {
	return filepath.WalkDir(h.Path, fn)
}

// Ext 获取文件后缀：无所谓存不存在，文件返回空
func (h *Handler) Ext() string {
	return filepath.Ext(h.Path)
}
