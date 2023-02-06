package pathlib

import "fmt"

type Err struct {
	Path *Path
}

// NotExistsErr 不存在错误
type NotExistsErr Err

func (e NotExistsErr) Error() string {
	return fmt.Sprintf("该路径不存在：%s", e.Path.Path)
}

// ExistsErr 已存在错误
type ExistsErr Err

func (e ExistsErr) Error() string {
	return fmt.Sprintf("该文件、文件夹已存在：%s", e.Path.Path)
}

// NotFileErr 非文件错误
type NotFileErr Err

func (e NotFileErr) Error() string {
	return fmt.Sprintf("该路径不是文件：%s", e.Path.Path)
}

// NotDirErr 非文件夹错误
type NotDirErr Err

func (e NotDirErr) Error() string {
	return fmt.Sprintf("该路径不是文件夹：%s", e.Path.Path)
}

// UnrecognizedType 无法识别类型错误
type UnrecognizedType Err

func (e UnrecognizedType) Error() string {
	return fmt.Sprintf("无法识别文件类型：%s", e.Path.Path)
}
