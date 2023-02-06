package pathlib

// Path 路径
type Path struct {
	Path string
}

// NewPath 转化路径
func NewPath(p string) *Path {
	return &Path{Path: p}
}

// String 使得打印输出为字符串
func (h *Path) String() string {
	return h.Path
}
