package file

// NotExists 文件不存在错误
type NotExists string

// Error 实现报错的接口
func (f NotExists) Error() string {
	return string(f)
}
