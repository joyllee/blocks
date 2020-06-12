package utils

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

func WriteFile(file_path string, imageBase64 []byte) (result bool, file string, err error) {
	uuid := uuid.New().String()
	dir1 := uuid[0:3]
	dir2 := uuid[3:6]
	dir := filepath.Join(file_path, dir1, dir2)
	if ok := IsFileExist(dir); !ok {
		err = os.MkdirAll(dir, 0766)
		if err != nil {
			logrus.Error("err===", err)
			return false, "", err
		}
	}
	file = filepath.Join(dir, uuid + ".jpg")
	out, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return false, "", err
	}
	defer out.Close()
	err = ioutil.WriteFile(file, imageBase64, 0666)
	if err != nil {
		return false, "", err
	}
	return true, file, nil
}

//判断文件是否存在
func IsFileExist(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
