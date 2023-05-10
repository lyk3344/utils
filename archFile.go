package utils

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	ezip "github.com/alexmullins/zip"
)

// 在sourceDir目录下方生成zip文件，压缩sourceDir目录的所有文件或目录
func ZipMultiFile(sourceDir, zipFileName string) error {

	var (
		oldZipFilePath string
		newZipFilePath string
	)

	err := os.Chdir(sourceDir)
	if err != nil {
		return err
	}

	//获取源文件列表
	f, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		return err
	}

	fzip, _ := os.Create(zipFileName)
	//defer fzip.Close()

	w := zip.NewWriter(fzip)
	//defer w.Close()


	for _, file := range f {
		//fmt.Println(file.Name())
		if file.Name() == "call_result" {
			continue
		}
		filepath.Walk(file.Name(), func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			header, err := zip.FileInfoHeader(info)
			if err != nil {
				return err
			}

			header.Name = path
			if info.IsDir() {
				header.Name += "/"
				//header.Name += "\\"
			} else {
				header.Method = zip.Deflate
			}

			writer, err := w.CreateHeader(header)
			if err != nil {
				return err
			}

			if ! info.IsDir() {
				file, err := os.Open(path)
				if err != nil {
					return err
				}
				//defer file.Close()
				_, err = io.Copy(writer, file)
				file.Close()
			}
			return err
		})
	}


	w.Close()
	fzip.Close()

	//获取上一级目录路径
	parentDir := getParentDirectory(sourceDir)

	//切换到上一级目录
	err = os.Chdir(parentDir)
	if err != nil {
		return err
	}

	//将zip文件移出到上一级目录中
	if strings.Contains(parentDir, "\\") {
		oldZipFilePath = sourceDir + "\\" + zipFileName
		newZipFilePath = parentDir + "\\" + zipFileName
	} else if strings.Contains(parentDir, "/") {
		oldZipFilePath = sourceDir + "/" + zipFileName
		newZipFilePath = parentDir + "/" + zipFileName
	}
	err = os.Rename(oldZipFilePath, newZipFilePath)
	if err != nil {
		return err
	}

	//删除数据目录

	return err
}


// src could be a single file or a directory
//@@param: src:
//@@param: dst: zip file name
func Zip(src string, dst string) error {
	zipfile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		header.Name = path
		if info.IsDir() {
			header.Name += "/"
			//header.Name += "\\"
		} else {
			header.Method = zip.Deflate
		}

		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if ! info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
		}
		return err
	})

	err = os.RemoveAll(src)
	if err != nil {
		return err
	}

	return err
}

func Unzip(zipFile string, destDir string) error {
	zipReader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer zipReader.Close()

	for _, f := range zipReader.File {
		fpath := filepath.Join(destDir, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
		} else {
			if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
				return err
			}

			inFile, err := f.Open()
			if err != nil {
				return err
			}
			defer inFile.Close()

			outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// src could be a single file or a directory
//@@param: src 待压缩的文件名
//@@param: dst 压缩后的文件名
func EncryptZip(src, dst, passwd string) error {
	zipfile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := ezip.NewWriter(zipfile)
	defer archive.Close()

	filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		header, err := ezip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = strings.TrimPrefix(path, filepath.Dir(src)+"/")
		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}
		// 设置密码
		header.SetPassword(passwd)
		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
		}
		return err
	})
	return err
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func getParentDirectory(dirctory string) string {
	if strings.Contains(dirctory, "\\") {
		//windows写法
		return substr(dirctory, 0, strings.LastIndex(dirctory, "\\"))
	} else if strings.Contains(dirctory, "/") {
		//linux写法
		return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
	}
	return dirctory
}

// 在sourceDir目录下方生成zip文件，压缩sourceDir目录下指定的文件或目录
// fileList: 要压缩的文件列表
func ZipSpecifyMultiFile(sourceDir, zipFileName string, fileList []string) error {

	err := os.Chdir(sourceDir)
	if err != nil {
		return err
	}

	fzip, _ := os.Create(zipFileName)
	//defer fzip.Close()

	w := zip.NewWriter(fzip)
	//defer w.Close()

	for i := 0; i < len(fileList); i++ {
		filepath.Walk(fileList[i], func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			header, err := zip.FileInfoHeader(info)
			if err != nil {
				return err
			}

			header.Name = path
			if info.IsDir() {
				header.Name += "/"
				//header.Name += "\\"
			} else {
				header.Method = zip.Deflate
			}

			writer, err := w.CreateHeader(header)
			if err != nil {
				return err
			}

			if ! info.IsDir() {
				file, err := os.Open(path)
				if err != nil {
					return err
				}
				//defer file.Close()
				_, err = io.Copy(writer, file)
				file.Close()
			}
			return err
		})

		//删除原文件
		err = os.RemoveAll(fileList[i])
		if err != nil {
			return err
		}
	}


	w.Close()
	fzip.Close()


	return err
}
