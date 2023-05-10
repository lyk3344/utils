package utils

import (
	"os"
	"io"
	"path/filepath"
)

//扫描指定目录下的所有文件，获取文件列表（带路径）
//会扫描子目录
func GetFileListByDir(filePath string) (fileList []string){
	var files []string
	err := filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error{
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return files
}

//复制文件到新目录下
func CopyFile(srcFile, destFile string) {
	file1, err := os.Open(srcFile)
	if err != nil {
		panic(err)
	}
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	defer file1.Close()
	defer file2.Close()

	//拷贝数据
	bs := make([]byte, 1024, 1024)
	n := -1 //读取的数据量
	total := 0
	for {
		n, err = file1.Read(bs)
		if err == io.EOF || n == 0 {
			break //拷贝完毕
		} else if err != nil {
			panic(err)
		}
		total += n
		file2.Write(bs[:n])
	}
}
