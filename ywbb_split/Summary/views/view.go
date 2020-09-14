package views

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"log"
	"strconv"
	"strings"
	"time"
	_ "ywbb_split/Summary/logconfig"
	"ywbb_split/Summary/model"
)

type Table struct {
	col1  string
	col2  string
	start int
	end   int
}

func MapWeekRow(startrow int) map[int]int {
	var WrMap = make(map[int]int)
	var WeekIndex = 1
	for i := startrow; i < startrow+5; i++ {
		WrMap[WeekIndex] = i
		WeekIndex++
	}
	return WrMap
}

// 返回区域和对应城市
func GetAreaCity(c string) string {
	// 南区
	var citys_south = []string{"无锡", "上海", "南京", "苏州", "广州", "深圳", "福厦", "武汉"}
	// 北区
	var citys_north = []string{"宁波", "杭州", "济青", "京津", "西安", "郑州", "大连", "川渝"}
	var m = make(map[string][]string)
	m["南区"] = citys_south
	m["北区"] = citys_north
	for area, citys := range m {

		for _, city := range citys {
			if strings.TrimSpace(c) == strings.TrimSpace(city) {
				return area + "-" + strings.TrimSpace(city)
			}
		}
	}

	return "未知地区-" + c

}

func NotUploadedCity() map[string]string {
	now_date := time.Now().Format("2006-01-02")
	var no_up_map = make(map[string]string)
	city_objs := model.QueryCity()
	for _, city := range city_objs {
		if strings.TrimSpace(now_date) != strings.TrimSpace(city.Date) {
			no_up_map[city.Name] = ""
		}

	}

	//fmt.Println(no_up_map)
	return no_up_map

}

func UpdateDataYj(month string, f *excelize.File, sheet_name, area string, m int) {

	defer func() {
		errs := recover()
		if errs != nil {
			log.Println(errs)

			panic("UpdateDataYj")

		}

	}()

	var cols = []string{"C", "D", "F", "G", "I", "J", "K", "M", "N", "P", "Q", "S", "T", "U", "Y", "Z", "AA", "AB"}
	// 获取数据对象列表
	var DataObjList = model.QueryMultiRowDemo(month, area)

	// 写入查出的多条数据
	for _, rowData := range DataObjList {
		// 查询已上传城市
		//AreaNameList := model.QueryCity(month, rowData.Week, area)
		//NotUploadedCity(month, rowData.Week, area, AreaNameList, c_map)

		//fmt.Println(month, rowData.Week, rowData.C, rowData.D, rowData.F, rowData.G, rowData.I, rowData.J,
		//	rowData.L, rowData.M, rowData.O, rowData.P, rowData.R, rowData.S, rowData.W, rowData.X, rowData.Y, rowData.Z)
		var DataList = []string{rowData.C, rowData.D, rowData.F, rowData.G, rowData.I, rowData.J, rowData.K,
			rowData.M, rowData.N, rowData.P, rowData.Q, rowData.S, rowData.T, rowData.U, rowData.Y, rowData.Z, rowData.AA, rowData.AB}

		var startrow int
		switch m {
		case 7:
			startrow = 8

		case 8:
			startrow = 14
		case 9:
			startrow = 20
		case 10:
			startrow = 27
		case 11:
			startrow = 33
		case 12:
			startrow = 39

		}
		//map[1:8 2:9 3:10 4:11 5:12]
		for week, row := range MapWeekRow(startrow) {
			//fmt.Println(week, row)
			// 数据库里的周
			w, err := strconv.Atoi(strings.Replace(rowData.Week, "周", "", 1))
			if err != nil {
				log.Println(err)
				panic("业绩管理转换int")
			}
			if week == w {
				for index, col := range cols {
					v, err := strconv.Atoi(DataList[index])
					if err != nil {
						log.Println(err)
						panic("业绩管理转换int")
					}
					f.SetCellInt(sheet_name, fmt.Sprintf("%s%d", col, row), v)
				}
			}

		}

	}

	// 返回 未上传的城市map
}

func UpdateExcel(area, fileName string) {

	defer func() {
		errs := recover()
		// 有错 发送邮件
		if errs != nil {
			log.Println(errs)
			err := Send("发送失败", "jc@joyowo.com", "Asd159357", "发送失败", []string{"andapy@163.com"}, []string{}, []string{})
			log.Println(err)
			panic("UpdateExcel")

		}

	}()

	var sheet_name = "业绩管理"

	// 在这个文件上做修改
	//var path = "C:\\template\\周业绩管理表_今元人才20200713.xlsx"
	//var path = "C:\\Users\\andap\\Desktop\\周业绩管理表_今元人才20200720.xlsx"
	//var path = "C:\\Users\\andap\\Desktop\\周业绩统计模板\\城市_周业绩管理表_今元人才20200720.xlsx"
	var path = "./周业绩管理表模板.xlsx"
	//var path = "C:\\Users\\andap\\Desktop\\周业绩统计模板\\周业绩管理表_今元人才2016.xlsx"
	//var path = "C:\\Users\\andap\\Desktop\\周业绩统计模板\\北区_周业绩管理表_今元人才2010.xlsx"
	//var path = "C:\\Users\\andap\\Desktop\\周业绩统计模板\\城市_周业绩管理表_今元人才2019.xlsx"
	//var path = "C:\\Users\\andap\\Desktop\\周业绩管理表.xlsx"

	f, err := excelize.OpenFile(path)
	if err != nil {
		log.Println(err)
	}
	f.UpdateLinkedValue()
	//var month = 7
	//m := time.Now().Month()
	t := time.Now().Format("2006-01-02 15:04:05")
	mm := strings.Split(t, "-")[1]
	month, err := strconv.Atoi(mm)
	if err != nil {
		log.Println(err)
	}
	//month = 12
	//var c_map = make(map[string][]string)

	for m := 7; m < month+1; m++ {
		MonthStr := fmt.Sprintf("%d月", m)
		UpdateDataYj(MonthStr, f, sheet_name, area, m)
	}
	//sheet_name = "客户商机池管理"
	sheet_name = "客户商机池管理"

	for m := 7; m < month+1; m++ {

		MonthStr := fmt.Sprintf("%d月", m)
		UpDateCustom(MonthStr, f, sheet_name, area, m)

	}

	sheet_name = "新签客户成交"

	for m := 7; m < month+1; m++ {

		MonthStr := fmt.Sprintf("%d月", m)
		UpDateNewCustom(MonthStr, f, sheet_name, area, m)
	}

	sheet_name = "业绩年完成率管理"
	QueryObjList := model.QueryComplete(area)
	var num int
	for _, obj := range QueryObjList {
		if area == strings.TrimSpace(obj.Region) {

			v, err := strconv.Atoi(obj.Num)
			if err != nil {
				fmt.Println(err)
				panic("完成率转换int")
			}
			f.SetCellInt(sheet_name, "B2", v)
			//num := 1
		}

		v, err := strconv.Atoi(obj.Num)
		if err != nil {
			fmt.Println(err)
			panic("完成率转换int")
		}
		num += v

	}
	if area == "南区+北区" {
		f.SetCellInt(sheet_name, "B2", num)
	}

	//f.SaveAs(fmt.Sprintf("C:\\Users\\andap\\Desktop\\财务数据切分\\%s.xlsx", area))
	f.SaveAs(fileName)

	//return c_map
}

// 完成率
func CompletionRate() {

}

// 客户商机
func UpDateCustom(month string, f *excelize.File, sheet_name, area string, m int) {

	defer func() {
		errs := recover()
		if errs != nil {
			log.Println(errs)

			panic("UpDateCustom")

		}

	}()

	var cols = []string{"C", "D", "F", "G", "I", "J", "L", "M", "O", "P"}
	var DataObjList = model.QueryCustom(month, area)

	// 保存第一行的数据 C F I
	var c int
	var f_col int
	var i int
	// 循环每一行数据结果
	for num, rowData := range DataObjList {

		var DataList = []string{rowData.C, rowData.D, rowData.F, rowData.G, rowData.I, rowData.J,
			rowData.L, rowData.M, rowData.O, rowData.P}
		//fmt.Println(rowData.Week, DataList)

		var startrow int
		switch m {
		case 7:
			startrow = 7

		case 8:
			startrow = 13
		case 9:
			startrow = 19
		case 10:
			startrow = 26
		case 11:
			startrow = 32
		case 12:
			startrow = 38

		}
		//map[1:8 2:9 3:10 4:11 5:12]
		for week, row := range MapWeekRow(startrow) {
			// 数据库里的周
			w, err := strconv.Atoi(strings.Replace(rowData.Week, "周", "", 1))
			if err != nil {
				log.Println(err)
				panic("客户商机获取周")
			}
			if week == w {
				// 写入每一列数据
				for index, col := range cols {

					v, err := strconv.Atoi(DataList[index])
					if err != nil {
						log.Println(err)
						panic("客户商机单元格转换成int")
					}
					f.SetCellInt(sheet_name, fmt.Sprintf("%s%d", col, row), v)

				}
			}

		}

		if m == 8 || m == 10 || m == 11 {
			if num == 1 {
				v, err := strconv.Atoi(DataList[0])
				if err != nil {
					log.Println(err)
					panic("客户商机转换成int")
				}
				c = v

				vf, err := strconv.Atoi(DataList[2])
				if err != nil {
					log.Println(err)
					panic("客户商机转换成int")
				}
				f_col = vf

				vi, err := strconv.Atoi(DataList[4])
				if err != nil {
					log.Println(err)
					panic("客户商机转换成int")
				}
				i = vi

			}

		} else {

			if num == 0 {
				v, err := strconv.Atoi(DataList[0])
				if err != nil {
					log.Println(err)
					panic("客户商机转换成int")
				}
				c = v

				vf, err := strconv.Atoi(DataList[2])
				if err != nil {
					log.Println(err)
					panic("客户商机转换成int")
				}
				f_col = vf

				vi, err := strconv.Atoi(DataList[4])
				if err != nil {
					log.Println(err)
					panic("客户商机转换成int")
				}
				i = vi

			}

		}

		if num == 4 {

			f.SetCellInt(sheet_name, fmt.Sprintf("C%d", startrow+1), c)
			f.SetCellInt(sheet_name, fmt.Sprintf("F%d", startrow+1), f_col)
			f.SetCellInt(sheet_name, fmt.Sprintf("I%d", startrow+1), i)
		}

	}

}

//  新客户成交
func UpDateNewCustom(month string, f *excelize.File, sheet_name, area string, m int) {
	defer func() {
		errs := recover()
		if errs != nil {

			log.Println(errs)

			panic("UpDateNewCustom")
		}

	}()
	var cols = []string{"C", "D", "E", "F", "G"}
	// 获取数据对象列表
	var DataObjList = model.QueryNewCustom(month, area)
	// 写入查出的多条数据
	for _, rowData := range DataObjList {
		//fmt.Println(month, rowData.Week, rowData.C, rowData.D, rowData.E, rowData.F, rowData.G)
		var DataList = []string{rowData.C, rowData.D, rowData.E, rowData.F, rowData.G}

		var startrow int
		switch m {
		case 7:
			startrow = 6

		case 8:
			startrow = 12
		case 9:
			startrow = 18
		case 10:
			startrow = 25
		case 11:
			startrow = 31
		case 12:
			startrow = 37

		}
		//map[1:8 2:9 3:10 4:11 5:12]
		for week, row := range MapWeekRow(startrow) {
			//fmt.Println(week, row)
			// 数据库里的周
			w, err := strconv.Atoi(strings.Replace(rowData.Week, "周", "", 1))
			if err != nil {
				log.Println(err)
				panic("新客户成交转换int")
			}
			if week == w {
				for index, col := range cols {
					v, err := strconv.Atoi(DataList[index])
					if err != nil {
						fmt.Println(err)
						panic("新客户成交转换int")
					}
					f.SetCellInt(sheet_name, fmt.Sprintf("%s%d", col, row), v)
				}
			}

		}

	}

}
