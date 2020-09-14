package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {

	//views.NotUploadedCity()
	r := IsNumeric("20181112")
	fmt.Println(r)

}

func IsNumeric(val interface{}) bool {
	switch val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		//return true
	case float32, float64, complex64, complex128:
		return true
	case string:
		str := val.(string)
		if str == "" {
			return false
		}
		// Trim any whitespace
		str = strings.Trim(str, " \\t\\n\\r\\v\\f")
		if str[0] == '-' || str[0] == '+' {
			if len(str) == 1 {
				return false
			}
			str = str[1:]
		}
		// hex
		if len(str) > 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X') {
			for _, h := range str[2:] {
				if !((h >= '0' && h <= '9') || (h >= 'a' && h <= 'f') || (h >= 'A' && h <= 'F')) {
					return false
				}
			}
			return true
		}
		// 0-9,Point,Scientific
		p, s, l := 0, 0, len(str)
		for i, v := range str {
			if v == '.' { // Point
				if p > 0 || s > 0 || i+1 == l {
					return false
				}
				p = i
			} else if v == 'e' || v == 'E' { // Scientific
				if i == 0 || s > 0 || i+1 == l {
					return false
				}
				s = i
			} else if v < '0' || v > '9' {
				return false
			}
		}
		return true
	}

	return false
}

func testmap(m map[string]string) {
	m["a"] = "a"

	//return m
}
func CreateExcel() {
	var path string
	for {
		fmt.Println("输入文件路径")
		fmt.Scanf("%s", &path)
		if path != "" {
			break
		}
	}
	//var path string

	var save_path string

	PathList := Read(path)
	//fmt.Println(PathList)
	path_info := PathList[0]
	save_path = PathList[1]
	templateFile, err := excelize.OpenFile(path + "\\" + save_path)
	f, err := excelize.OpenFile(path + "\\" + path_info)
	if err != nil {
		fmt.Println(err)
	}
	rows := f.GetRows("Sheet1")

	for _, row := range rows {
		if len(row) == 0 {
			continue
		}
		file_name := row[0]
		file_path := path + "\\" + file_name
		//fmt.Println(file_path)
		templateFile.SaveAs(file_path)
	}
	fmt.Println("复制完成")
	time.Sleep(time.Second * 3)
}

func Read(path string) []string {

	f, err := os.Open(path + "\\path.txt")
	if err != nil {
		fmt.Println("read file fail", err)
		//return ""
	}
	defer f.Close()
	decoder := mahonia.NewDecoder("gbk")
	fd, err := ioutil.ReadAll(decoder.NewReader(f))
	if err != nil {
		fmt.Println("read to fd fail", err)
		//return ""
	}
	//fmt.Println(string(fd))
	return strings.Split(string(fd), "\r\n")
}
