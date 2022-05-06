package excel

import (
	"context"
	"fmt"
	"github.com/xuri/excelize/v2"
	"strings"
)

// 常量定义
const (
	firstRowHeight = 20
	rowHeight      = 20
)

// ConvertInterface ConvertInterface
type ConvertInterface interface {
	Apply(string) string
}

type cut2point struct {
}

// Apply Apply
func (*cut2point) Apply(val string) string {
	idx := strings.Index(val, ".")
	if idx > 0 {
		if idx+3 > len(val) {
			return val
		}
		return string([]byte(val)[0 : idx+3])
	}
	return val
}

// Save2Excel 保存到excel文件
func Save2Excel(list []*ExportData, f *excelize.File, workSheet string) error {

	if f.GetSheetIndex(workSheet) == -1 {
		f.NewSheet(workSheet)
	}

	//获取表头
	cols := GetExcelColsSlice()
	colsName := make([]string, 0, len(cols))
	for _, col := range cols {
		colsName = append(colsName, col.Title)
	}

	//设置表头
	err := f.SetSheetRow(workSheet, "A1", &colsName)
	if err != nil {
		fmt.Println("set table header failed. err:" + err.Error())
		return err
	}

	//设置首行的高度
	err = f.SetRowHeight(workSheet, 1, firstRowHeight)
	if err != nil {
		fmt.Println("set table header row height failed. err:" + err.Error())
		return err
	}

	//设置表格样式
	err = SetExcelStyle(context.TODO(), f, workSheet, len(list)+1)
	if err != nil {
		return err
	}

	//写入数据
	lint := 0

	//行数据写入
	for index, item := range list {
		lint = index + 2
		err = f.SetSheetRow(workSheet, fmt.Sprintf("A%d", lint), &[]interface{}{
			item.ProductID,
			item.ProductName,
			item.SearchKeyword,
			item.CommentTime,
			item.CommentKeyword,
			item.CommentText,
			item.IsHasKeyword,
			item.CommentImg,
		})
		if err != nil {
			fmt.Println("SetSheetRow failed. err:" + err.Error())
			return err
		}
		err = f.SetRowHeight(workSheet, lint, rowHeight)
		if err != nil {
			fmt.Println("set table row height failed. err:" + err.Error())
			return err
		}

	}

	return nil
}