package Gfile

import (
	"io/ioutil"
	"log"
	"os"
)

func ReadFileToString(path string) string {
	ret, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println("[!] ReadFileToString Error: ", err)
		return ""
	}
	return string(ret)
}
func WriteString(path, data string) bool {
	err := ioutil.WriteFile(path, []byte(data), 0777)
	if err != nil {
		log.Println("[!] WriteString Error: ", err)
		return false
	}
	return true
}

func CheckExist(path string) bool {
	_, oserr := os.Stat(path)
	if os.IsNotExist(oserr) {
		return false
	}
	return true
}
func UnZip(src string, dst string) {
	if CheckExist(dst) == false {
		os.MkdirAll(dst, 0777)
	}
	read, err := zip.OpenReader(src)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, file := range read.Reader.File {
		p := path.Join(dst, file.Name)
		if file.FileInfo().IsDir() {
			fmt.Println("[C] ", p)
			_ = os.MkdirAll(p, file.FileInfo().Mode())
			continue
		}
		r, err := file.Open()
		if err != nil {
			continue
		}
		defer r.Close()
		outFile, err := os.Create(p)
		if err != nil {
			continue
		}
		fmt.Println("[W] ", p)
		defer outFile.Close()
		_, _ = io.Copy(outFile, r)
	}
}

func GetFileList(dir string) interface{} {
	dirList := []map[string]interface{}{}
	dir_list, err := ioutil.ReadDir(dir)
	if err != nil {
		return dirList
	}
	for i, v := range dir_list {
		dir := map[string]interface{}{
			"Index":      i,
			"Name":       v.Name(),
			"Size":       v.Size(),
			"Mode":       v.Mode().String(),
			"ModifyTime": v.ModTime().Format("2006/1/2 15:04:05"),
			"IsDir":      v.IsDir(),
		}
		dirList = append(dirList, dir)
	}
	return dirList
}
