package excel

import (
	"context"
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
)

//DataCheck 数据检测错误输出
type DataCheck struct {
	Row int    `json:"row"`
	Err string `json:"err"`
}

//Excel 表格定义
type Excel struct {
	Cols             []*ExcelCol          //列定义
	ColsMap          map[string]*ExcelCol //列定义
	ThStyle          *excelize.Style      //表头样式
	TdStyle          *excelize.Style      //表数据样式
	DefaultRowHeight float64              //默认行高
	DefaultColWidth  float64              //默认列宽
	AutoFilter       bool                 //表头是否增加自动过滤器
}

//ExcelCol 字段定义
type ExcelCol struct {
	Title          string                   // 标题
	Field          string                   // 字段名，对应结构中的json串
	RequiredCol    bool                     // 是否为必需字段，为上传时使用
	RequiredValue  bool                     // 是否为必需有值的字段
	Style          *excelize.Style          // 表格样式
	DvRange        *excelize.DataValidation // 数据有效性校验规则
	Seq            int                      // 序号,为导出时使用
	ConvertAdapter []ConvertInterface       // 内容执行过滤器
}

var gExcel *Excel

func init() {
	gExcel = initExcelConfig()
}

func initExcelConfig() *Excel {
	cols := make([]*ExcelCol, 0)
	cols = []*ExcelCol{
		{Title: "产品ID", Field: "product_id", RequiredCol: true},
		{Title: "产品名称", Field: "product_name", RequiredCol: true},
		{Title: "搜索关键词", Field: "search_keyword", RequiredCol: true},
		{Title: "评论时间", Field: "comment_time", RequiredCol: true},
		{Title: "评论关键词", Field: "comment_keyword", RequiredCol: true},
		{Title: "评论文字", Field: "comment_text", RequiredCol: true},
		{Title: "是否包含评论关键词", Field: "is_has_keyword", RequiredCol: true},
		{Title: "本地图片地址", Field: "comment_img", RequiredCol: true},
		{Title: "源图片地址", Field: "comment_remote_img", RequiredCol: true},
	}

	//转map
	colsMap := make(map[string]*ExcelCol)
	for idx, v := range cols {
		v.Seq = idx + 1 //字段排序标记
		colsMap[v.Title] = v
	}

	return &Excel{
		Cols:             cols,
		ColsMap:          colsMap,
		ThStyle:          GetThStyle(),
		TdStyle:          GetTdStyle(),
		DefaultRowHeight: 20,
		DefaultColWidth:  30,
		AutoFilter:       true,
	}
}

// GetExcelColsMap GetExcelColsMap
func GetExcelColsMap() map[string]*ExcelCol {
	return initExcelConfig().ColsMap
}

// GetExcelColsSlice GetExcelColsSlice
func GetExcelColsSlice() []*ExcelCol {
	return initExcelConfig().Cols
}

// SetExcelStyle 设置excel文件样式
func SetExcelStyle(ctx context.Context, f *excelize.File, workSheet string, maxRow int) error {

	if f == nil {
		fmt.Println("file pointer is nil")
		return errors.New("file pointer is nil")
	}

	maxColName, _ := excelize.ColumnNumberToName(len(gExcel.Cols))

	//行高设置
	if gExcel.DefaultRowHeight > 1 {
		f.SetSheetFormatPr(workSheet, excelize.DefaultRowHeight(gExcel.DefaultRowHeight))
	}

	//列宽设置
	if gExcel.DefaultColWidth > 1 {
		f.SetSheetFormatPr(workSheet, excelize.DefaultColWidth(gExcel.DefaultColWidth))
	}

	//添加过滤器
	if gExcel.AutoFilter {
		maxCol, _ := excelize.ColumnNumberToName(len(gExcel.Cols))
		vCell := maxCol + "1"
		err := f.AutoFilter(workSheet, "A1", vCell, "")
		if err != nil {
			fmt.Println("add auto filter failed. err:" + err.Error())
			return err
		}
	}

	//表头样式
	if gExcel.ThStyle != nil {
		styleID, err := f.NewStyle(gExcel.ThStyle)
		if err != nil {
			fmt.Println("get table header style failed. err:" + err.Error())
			return err
		}
		vCell := maxColName + "1"
		err = f.SetCellStyle(workSheet, "A1", vCell, styleID)
		if err != nil {
			fmt.Println("set table header style failed. err:" + err.Error())
			return err
		}
	}

	//表数据样式
	if gExcel.TdStyle != nil {
		styleID, err := f.NewStyle(gExcel.TdStyle)
		if err != nil {
			fmt.Println("get table body style failed. err:" + err.Error())
			return err
		}

		err = f.SetCellStyle(workSheet, "A2", fmt.Sprintf("%s%d", maxColName, maxRow), styleID)
		if err != nil {
			fmt.Println("set table td style failed. err:" + err.Error())
			return err
		}
	}

	//各列样式设置
	for _, v := range gExcel.Cols {
		//设置列格式
		if v.Style != nil {
			styleID, err := f.NewStyle(v.Style)
			if err != nil {
				fmt.Println(err)
				return err
			}
			colName, err := excelize.ColumnNumberToName(v.Seq)
			if err != nil {
				fmt.Println(err)
				return err
			}
			err = f.SetCellStyle(workSheet, colName+"2", fmt.Sprintf("%s%d", colName, maxRow), styleID)
			if err != nil {
				fmt.Println(err)
				return err
			}
		}

		//设置有效性校验规则
		if v.DvRange != nil {
			colName, err := excelize.ColumnNumberToName(v.Seq)
			if err != nil {
				fmt.Println(err)
				return err
			}
			v.DvRange.Sqref = fmt.Sprintf("%s2:%s%d", colName, colName, maxRow)
			err = f.AddDataValidation(workSheet, v.DvRange)
			if err != nil {
				fmt.Println(err)
				return err
			}
		}
	}
	return nil
}

// ExportData 导出结构体
type ExportData struct {
	ProductID        string `json:"product_id"`         // 产品id
	ProductName      string `json:"product_name"`       // 产品名称
	SearchKeyword    string `json:"search_keyword"`     // 搜索关键词
	CommentTime      string `json:"comment_time"`       // 评论时间
	CommentKeyword   string `json:"comment_keyword"`    // 评论关键词
	CommentText      string `json:"comment_text"`       // 评论文字
	IsHasKeyword     bool   `json:"is_has_keyword"`     // 是否包含评论关键词
	CommentImg       string `json:"comment_img"`        // 本地图片地址
	CommentRemoteImg string `json:"comment_remote_img"` // 原图片地址
}
