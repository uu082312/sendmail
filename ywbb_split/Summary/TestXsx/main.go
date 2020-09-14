package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("D:\\project6\\huobanyun\\a.xlsx")
	if err != nil {
		fmt.Println(err)
	}
	err = f.UnmergeCell("Sheet1", "A1", "D100")
	if err != nil {
		fmt.Println(err)
	}
	f.SaveAs("./a.xlsx")
}
