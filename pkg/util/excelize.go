package util

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"ordinary_test/model"
)

func GenerateArticleExcel(articles []model.Article) (path string, err error) {
	f := excelize.NewFile()
	streamWriter, err := f.NewStreamWriter("sheet1")
	if err != nil {
		return
	}
	header := []interface{}{}
	for _, cell := range []string{
		"ID", "创建日期", "更新日期", "删除日期", "标签ID", "标签", "标题", "DESC", "内容",
		"封面图片URL", "创建人", "修改人", "状态",
	} {
		header = append(header, cell)
	}
	if err := streamWriter.SetRow("A1", header); err != nil {
		fmt.Println(err)
		return
	}
	for rowId := 2; rowId < len(articles)+2; rowId++ {
		article := articles[rowId-2]
		data := make([]interface{}, 0)
		data = append(data, article.ID)
		data = append(data, article.CreatedAt)
		data = append(data, article.UpdatedAt)
		data = append(data, article.DeletedAt)
		data = append(data, article.TagID)
		data = append(data, article.Tag)
		data = append(data, article.Title)
		data = append(data, article.Desc)
		data = append(data, article.Content)
		data = append(data, article.CoverImageUrl)
		data = append(data, article.CreatedBy)
		data = append(data, article.ModifiedBy)
		data = append(data, article.State)
		cell, _ := excelize.CoordinatesToCellName(1, rowId)
		if err := streamWriter.SetRow(cell, data); err != nil {
			fmt.Println(err)
			return
		}
	}
	if err := streamWriter.Flush(); err != nil {
		return
	}

	if err := f.SaveAs("文章表.xlsx"); err != nil {
		return
	}
	path = f.Path
	return
}
