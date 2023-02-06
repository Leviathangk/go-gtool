package pathlib

import (
	"io/fs"
	"runtime"
)

const (
	Unknown = iota
	Linux
	Windows
)

// SysType 系统类型
var SysType int

// Handler 记录路径相关的参数
// FileMode 在 0777 时，windows 打印出来是 511，实际上是 os.ModePerm
type Handler struct {
	Path     string
	FileMode fs.FileMode
}

type HandlerSetter func(h *Handler)

func init() {
	SysType = GetSysType()
}

// GetSysType 获取系统类型
func GetSysType() int {
	switch runtime.GOOS {
	case "linux":
		return Linux
	case "windows":
		return Windows
	default:
		return Unknown
	}
}

// NewHandler 新建一个处理器
func NewHandler(p string) *Handler {
	return &Handler{
		Path:     p,
		FileMode: 0777,
	}
}

// Path 转化路径
func Path(p string, options ...HandlerSetter) *Handler {
	handler := NewHandler(p)

	// 根据输入修改设置
	for _, opt := range options {
		opt(handler)
	}

	return handler
}

// SetFileMode 设置文件的权限：默认 0777
func SetFileMode(mode fs.FileMode) HandlerSetter {
	return func(h *Handler) {
		h.FileMode = mode
	}
}

// CopyHandler 复制 Handler 到另一个路径上面
func CopyHandler(p string, h *Handler) *Handler {
	newHandler := *h
	newHandler.Path = p

	return &newHandler
}

// String 使得打印输出为字符串
func (h *Handler) String() string {
	return h.Path
}
