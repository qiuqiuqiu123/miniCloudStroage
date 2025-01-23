package iserver

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

/*
*

	总的根目录由rootPath指定,用户的根目录为rootPath/{user_id}/
*/
type SimpleFileServer struct {
	rootPath string
}

func NewSimpleFileServer(path string) *SimpleFileServer {
	return &SimpleFileServer{
		rootPath: path,
	}
}

func (s *SimpleFileServer) Upload(data []byte, name string, path string) {
	targetDir := filepath.Join(s.rootPath, path)   // 目标目录
	filename := name                               // 文件名
	filePath := filepath.Join(targetDir, filename) // 拼接完整路径

	fmt.Println(filePath)

	// 1. 创建目标目录（如果不存在）
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		err := os.MkdirAll(targetDir, os.ModePerm) // 递归创建目录
		if err != nil {
			fmt.Println("Failed to create directory:", err)
		}
		fmt.Println("Directory created:", targetDir)
	}

	// 2. 创建或打开目标文件
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}
	defer file.Close() // 确保文件关闭

	// 3. 写入文件内容
	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Failed to write file:", err)
		return
	}

	// 4. 打印成功信息
	fmt.Println("File saved successfully at:", filePath)
}

func (s *SimpleFileServer) Download(path string) ([]byte, error) {
	// 1. 创建或打开目标文件
	targetDir := filepath.Join(s.rootPath, path)
	file, err := os.Open(targetDir)
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return nil, err
	}
	defer file.Close() // 确保文件关闭

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Failed to read file:", err)
		return nil, err
	}
	return data, nil
}

func (s *SimpleFileServer) ListDirs(path string) []string {
	curPath := s.rootPath
	if path != "" {
		curPath = filepath.Join(curPath, path)
	}

	dirs, err := os.ReadDir(curPath)
	if err != nil {
		fmt.Println("Failed to list dirs:", err)
		return make([]string, 0)
	}

	dirNames := make([]string, 0)
	for _, dir := range dirs {
		if dir.IsDir() {
			dirNames = append(dirNames, dir.Name())
		}
	}

	return dirNames
}

func (s *SimpleFileServer) AddDir(path string) error {
	curPath := filepath.Join(s.rootPath, path)

	fmt.Println("cur path :", curPath)
	// 如果文件夹不存在，则进行创建
	if _, err := os.Stat(curPath); os.IsNotExist(err) {
		err := os.Mkdir(curPath, fs.ModeDir)
		if err != nil {
			fmt.Println("create dir fail")
			return err
		}
	}

	// 在就不进行处理
	return nil
}

func (s *SimpleFileServer) DelDir(path string) error {
	curPath := filepath.Join(s.rootPath, path)

	fmt.Println("cur path :", curPath)
	// 如果文件夹不在,不进行处理
	if _, err := os.Stat(curPath); os.IsNotExist(err) {
		return nil
	}

	err := os.Remove(curPath)
	if err != nil {
		fmt.Println("del dir fail")
		return err
	}
	return nil
}
