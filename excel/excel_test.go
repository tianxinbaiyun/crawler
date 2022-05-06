package excel

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"testing"
	"time"
)

// @Title  请填写文件名称（需要改）
// @Description  请填写文件描述（需要改）
// @Author  clx  2022/5/6 14:29
// @Update  clx  2022/5/6 14:29

func TestSave2Excel(t *testing.T) {
	list := make([]*ExportData, 0)
	list = append(list, &ExportData{
		ProductID:      "123",
		ProductName:    "123",
		SearchKeyword:  "123",
		CommentTime:    "123",
		CommentKeyword: "123",
		CommentText:    "123",
		IsHasKeyword:   false,
		CommentImg:     "123",
	})
	f := excelize.NewFile()
	err := Save2Excel(list, f, "test")
	if err != nil {
		fmt.Println(err)
		return
	}
	f.DeleteSheet("Sheet1")

	_ = f.SaveAs(fmt.Sprintf("test_%s.xlsx", time.Now().Format("20060102150405")))
}
