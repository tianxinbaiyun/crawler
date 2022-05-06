package excel

import (
	"context"
	"fmt"
	"github.com/xuri/excelize/v2"
	"os"
	"strings"
)

// 常量定义
const (
	firstRowHeight = 20
	rowHeight      = 20
)

var (
	BasePath, _ = os.Getwd()
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

	// 设置特殊颜色
	specialStyle := GetSpecialColStyle()
	specialStyleID, err := f.NewStyle(specialStyle)
	if err != nil {
		return err
	}

	urlStyle := GetURLColStyle()
	urlStyleID, err := f.NewStyle(urlStyle)
	if err != nil {
		return err
	}

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
			item.CommentRemoteImg,
		})
		if err != nil {
			fmt.Println("SetSheetRow failed. err:" + err.Error())
			return err
		}
		//设置超链接
		if item.ProductID != "" {
			url := fmt.Sprintf("https://item.jd.com/%s.html", item.ProductID)
			err = f.SetCellHyperLink(workSheet, fmt.Sprintf("A%d", lint), url, "External")
			if err != nil {
				fmt.Println("set table row height failed. err:" + err.Error())
				return err
			}
			err = f.SetCellStyle(workSheet, fmt.Sprintf("A%d", lint), fmt.Sprintf("A%d", lint), urlStyleID)
			if err != nil {
				return err
			}
		}
		//设置超链接
		if item.CommentImg != "" {
			imgPath := item.CommentImg
			if strings.HasPrefix(imgPath, ".") {
				imgPath = fmt.Sprintf("%s%s", BasePath, strings.Trim(imgPath, "."))
			}
			imgPath = strings.ReplaceAll(imgPath, `/`, `\`)
			err = f.SetCellHyperLink(workSheet, fmt.Sprintf("H%d", lint), imgPath, "External")
			if err != nil {
				fmt.Println("set table row height failed. err:" + err.Error())
				return err
			}
			err = f.SetCellStyle(workSheet, fmt.Sprintf("H%d", lint), fmt.Sprintf("H%d", lint), urlStyleID)
			if err != nil {
				return err
			}
		}
		//设置超链接
		if item.CommentRemoteImg != "" {
			err = f.SetCellHyperLink(workSheet, fmt.Sprintf("I%d", lint), item.CommentRemoteImg, "External")
			if err != nil {
				fmt.Println("set table row height failed. err:" + err.Error())
				return err
			}
			err = f.SetCellStyle(workSheet, fmt.Sprintf("I%d", lint), fmt.Sprintf("I%d", lint), urlStyleID)
			if err != nil {
				return err
			}
		}
		err = f.SetRowHeight(workSheet, lint, rowHeight)
		if err != nil {
			fmt.Println("set table row height failed. err:" + err.Error())
			return err
		}

		//颜色标记
		if item.IsHasKeyword {
			err = f.SetCellStyle(workSheet, fmt.Sprintf("A%d", lint), fmt.Sprintf("I%d", lint), specialStyleID)
			if err != nil {
				return err
			}
		}

	}

	return nil
}
