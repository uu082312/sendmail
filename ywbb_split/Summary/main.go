package main

import (
	"fmt"
	"github.com/robfig/cron"
	"log"
	"os"
	"strings"
	"time"
	_ "ywbb_split/Summary/logconfig"
	"ywbb_split/Summary/views"
)

func main() {
	c := cron.New()

	_, err := os.Stat("./周业绩管理表模板.xlsx")
	if err != nil {
		fmt.Println("模板文件不存在")
	}

	//run()
	//fmt.Println("发送完成")

	defer func() {
		c.Stop()
		err := recover()
		if err != nil {
			log.Println(err)
		}

	}()

	c.AddFunc("10 1 2 * * 0", run)
	c.AddFunc("10 1 2 * * 7", run)
	//c.AddFunc("10 15 16 * * 3", run)
	//c.AddFunc("10 33 10 * * 1", run)
	c.Start()
	select {}

}

func run() {
	// 文件模板！！！！

	save_path := "./SendEmail"
	if _, err := os.Stat(save_path); err != nil {
		os.MkdirAll(save_path, os.ModePerm)
	}

	// lips5748@joyowo.com
	var EmailAddr = make(map[string]string)

	EmailAddr["南区"] = "xxx.com"
	EmailAddr["北区"] = "xxx.com"
	EmailAddr["南区+北区"] = "xxx.com"

	//EmailAddr["南区"] = "xxxx"
	//EmailAddr["北区"] = "xxxx"
	//EmailAddr["南区+北区"] = "xxx"

	// 查询所有城市修改时间
	c_map := views.NotUploadedCity()

	date := time.Now().Format("20060102")

	var NorthAndSouthFilePaths []string
	// 北区全部附件
	var NorthPaths []string
	// 南区全部附件
	var SouthPaths []string

	var citys = []string{"无锡", "上海", "南京", "苏州", "广州", "深圳", "福厦", "武汉", "宁波", "杭州", "济青", "京津", "西安", "郑州", "大连", "川渝"}

	for _, city := range citys {

		CityPath := fmt.Sprintf("%s\\%s_周业绩管理表_今元人才%s.xlsx", save_path, city, date)
		AreaCity := strings.Split(views.GetAreaCity(city), "-")
		if AreaCity[0] == "南区" {
			SouthPaths = append(SouthPaths, CityPath)
		} else if AreaCity[0] == "北区" {
			NorthPaths = append(NorthPaths, CityPath)
		}
		views.UpdateExcel(city, CityPath)
	}

	var areas = []string{"南区", "北区", "南区+北区"}
	for _, area := range areas {

		var fileName string
		var content string
		if area == "南区+北区" {
			fileName = fmt.Sprintf("%s\\周业绩管理表_今元人才%s.xlsx", save_path, date)
			content = fmt.Sprintf("周业绩管理表_今元人才%s", date)

		} else {
			fileName = fmt.Sprintf("%s\\%s_周业绩管理表_今元人才%s.xlsx", save_path, area, date)
			content = fmt.Sprintf("%s_周业绩管理表_今元人才%s", area, date)
			if area == "南区" {
				SouthPaths = append(SouthPaths, fileName)
			} else if area == "北区" {
				NorthPaths = append(NorthPaths, fileName)
			}
		}

		NorthAndSouthFilePaths = append(NorthAndSouthFilePaths, fileName)

		// 修改数据 保存文件
		views.UpdateExcel(area, fileName)

		var AreaAndCitys = make(map[string][]string)
		// 未上传城市列表
		var southCitys []string
		var northCitys []string

		// 统计未上传城市

		// 未上传城市
		var text string
		text = "<b>未上传城市：</b><br>"

		//var citysList []string
		// 未上传   k 城市名  v 上次上传时间
		for k, v := range c_map {
			AreaAndCity := views.GetAreaCity(k)
			//fmt.Println(AreaAndCity)
			AreaAndCityList := strings.Split(AreaAndCity, "-")

			switch AreaAndCityList[0] {
			case "南区":
				southCitys = append(southCitys, AreaAndCityList[1]+v+"<br>")
			case "北区":
				northCitys = append(northCitys, AreaAndCityList[1]+v+"<br>")
			}

		}
		AreaAndCitys["南区"] = southCitys
		AreaAndCitys["北区"] = northCitys
		var city_str string
		// a 地区 c 城市列表
		if area == "南区+北区" {
			for _, cs := range AreaAndCitys {
				for _, c := range cs {
					city_str += c
				}
			}
		} else {
			for _, c := range AreaAndCitys[area] {
				city_str += c
			}
		}

		if city_str == "" {
			text = ""
		}
		text += city_str

		if _, err := os.Stat(fileName); err != nil {
			fmt.Println(err)
			err := views.Send(area+"没有附件", "xxxx.com", "xxx", "没有附件", []string{"xxxx"}, []string{}, []string{""})
			log.Println(err)
			continue
		}

		var enclosure []string
		if area == "南区+北区" {
			enclosure = NorthAndSouthFilePaths
		} else if area == "南区" {
			enclosure = SouthPaths

		} else if area == "北区" {
			enclosure = NorthPaths
		}
		// 添加抄送人
		if area == "南区+北区" {
			// "chenm@joyoget.com"
			err := views.Send(content+"<br>"+text, "xxx", "xxx", content, []string{EmailAddr[area]}, []string{"xxxx.com"}, enclosure)
			//err := views.Send(content + "<br>" + text, "xxx", "xxx", content,  []string{EmailAddr[area]}, []string{}, enclosure)
			if err != nil {
				//fmt.Println(err)
				log.Println(err)
			}
		} else {
			// "chenm@joyoget.com"
			err := views.Send(content+"<br>"+text, "xxxx", "xxx", content, []string{EmailAddr[area]}, []string{"xxx", "xxx"}, enclosure)
			//err := views.Send(content + "<br>" + text, "xxx", "xxxx", content,  []string{EmailAddr[area]}, []string{}, enclosure)
			if err != nil {
				//fmt.Println(err)
				log.Println(err)
			}

		}

		err := views.Send(content+"<br>"+text, "xxx", "xxx", content, []string{"andapy@163.com"}, []string{}, enclosure)
		if err != nil {
			//fmt.Println(err)
			log.Println(err)
		}

	}

}
