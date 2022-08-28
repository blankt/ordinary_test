package main

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

type Student struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func main() {
	students := []Student{
		{1, "tlf", "xxx"},
		{2, "pxy", "xxx"},
		{3, "xxx", "xxx"},
		{4, "yyy", "xxx"},
		{5, "zzz", "xxx"},
		{6, "qqq", "xxx"},
		{7, "www", "xxx"},
		{8, "eee", "xxx"},
	}
	ExportExcel(students)
}

func ExportExcel(students []Student) {
	f := excelize.NewFile()
	streamWriter, err := f.NewStreamWriter("sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	header := []interface{}{}
	for _, cell := range []string{
		"ID", "姓名", "创建日期",
	} {
		header = append(header, cell)
	}
	if err := streamWriter.SetRow("A1", header); err != nil {
		fmt.Println(err)
		return
	}
	for rowId := 2; rowId < len(students)+2; rowId++ {
		student := students[rowId-2]
		data := make([]interface{}, 0)
		data = append(data, student.ID)
		data = append(data, student.Name)
		data = append(data, student.Address)
		cell, _ := excelize.CoordinatesToCellName(1, rowId)
		if err := streamWriter.SetRow(cell, data); err != nil {
			fmt.Println(err)
			return
		}
	}
	if err := streamWriter.Flush(); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SaveAs("学生信息表.xlsx"); err != nil {
		fmt.Println(err)
	}
}
