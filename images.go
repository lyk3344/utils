package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

//获取图片文件后缀
func GetImgExt(file string) (ext string, err error) {
	var headerByte []byte
	headerByte = make([]byte, 8)
	fd, err := os.Open(file)
	if err != nil {
		return "", err
	}

	defer fd.Close()
	_, err = fd.Read(headerByte)
	if err != nil {
		return "", err
	}
	xStr := fmt.Sprintf("%x", headerByte)
	switch {
	case xStr == "89504e470d0a1a0a":
		ext = ".png"
	case xStr == "0000010001002020":
		ext = ".ico"
	case xStr == "0000020001002020":
		ext = ".cur"
	case xStr[:12] == "474946383961" || xStr[:12] == "474946383761":
		ext = ".gif"
	case xStr[:10] == "0000020000" || xStr[:10] == "0000100000":
		ext = ".tga"
	case xStr[:8] == "464f524d":
		ext = ".iff"
	case xStr[:8] == "52494646":
		ext = ".ani"
	case xStr[:4] == "4d4d" || xStr[:4] == "4949":
		ext = ".tiff"
	case xStr[:4] == "424d":
		ext = ".bmp"
	case xStr[:4] == "ffd8":
		ext = ".jpg"
	case xStr[:2] == "0a":
		ext = ".pcx"
	default:
		ext = ""
	}
	return ext, nil
}


//通过Url下载图片并将图片body用base64编码
func DownloadImgTransToBase64(imgPath string) (imgstring string){
	if imgPath == "" {
		return imgPath
	}

	//去掉可能出现的双引号
	imgPath = strings.Replace(imgPath, "\"", "", -1)

	//fmt.Println("##########图片地址： ", imgPath)
	//通过http请求获取图片的流文件
	resp, err := http.Get(imgPath)
	if err != nil {
		panic(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	//保存成文件
	//out, _ := os.Create(name)
	//io.Copy(out, bytes.NewReader(body))
	//return
	//base64压缩
	imgstring = base64.StdEncoding.EncodeToString(body[:])
	return imgstring

	/*
		//还原测试
		//写入临时文件
		ioutil.WriteFile("a.jpg.txt", []byte(sourcestring), 0667)
		//读取临时文件
		cc, _ := ioutil.ReadFile("a.jpg.txt")
		//解压
		dist, _ := base64.StdEncoding.DecodeString(string(cc))
		//写入新文件
		f, _ := os.OpenFile("b.jpg", os.O_RDWR|os.O_CREATE, os.ModePerm)
		defer f.Close()
		f.Write(dist)

	*/
}

func DownloadImg(urlPath, savePath string) (localImgPath string) {
	//fmt.Println("##########图片地址： ", urlPath)

	//创建多级目录
	err := os.MkdirAll(savePath, 0766)
	if err != nil {
		panic(err)
	}

	//通过http请求获取图片的流文件
	resp, err := http.Get(urlPath)
	if err != nil {
		//return "未能获取头像信息"
		return ""
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	//localImgPath := config.SysConf.LocalImagePath

	//根据图片url获取其文件名
	//如果url中包含/132，则在生成的文件名中去掉/132
	if strings.Contains(urlPath, "/132") {
		urlPath = strings.Replace(urlPath, "/132", "", 1)
	}

	//返回地址中的最后一个文件名,例如：
	fileName := path.Base(urlPath)

	//如果文件名中包含了 ? ， 则去掉？和？后面的内容
	//按？分割字符串
	sep := "?"
	fileArr := strings.Split(fileName, sep)
	fileName = fileArr[0]

	//保存成文件
	file, err := os.Create(savePath + fileName)
	if err != nil {
		panic(err)
	}

	//获得文件的writer对象
	//writer := bufio.NewWriter(file)

	io.Copy(file, bytes.NewReader(body))

	file.Close()

	//重命名前文件路径
	fp := savePath + fileName

	if find := strings.Contains(fileName, "."); !find {
		//获取图片格式：
		fileType, err := GetImgExt(fp)
		if err != nil {
			panic(err)
		}

		//新图片名称
		fileName = GetRandomString(40) + fileType

		//fmt.Println("自定义图片名： " ,fileName)
		//新文件路径
		nfp := savePath + fileName
		err = os.Rename(fp, nfp)
		if err != nil {
			panic(err)
		}
	}
	return  "image/" + fileName
}

//读取本地图片并将图片body用base64编码
//该路径为全路径，包含图片名称
func LocalImgTransToBase64(imgPath string) (imgstring string){
	if imgPath == "" {
		return imgPath
	}

	//去掉可能出现的双引号
	imgPath = strings.Replace(imgPath, "\"", "", -1)

	//读取本地图片
	imgBody, _ := ioutil.ReadFile(imgPath)
	//base64编码
	imgstring = base64.StdEncoding.EncodeToString(imgBody[:])
	return imgstring
}
