package excel

import (
	"github.com/xuri/excelize/v2"
)

// GetBaseStyle 基础样式
func GetBaseStyle() *excelize.Style {
	return &excelize.Style{
		Font: &excelize.Font{
			Size:   9,
			Color:  "#000000",
			Family: "微软雅黑",
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	}
}

// GetThStyle 获取表头样式
func GetThStyle() *excelize.Style {
	style := GetBaseStyle()
	style.Font.Color = "#C00000"
	style.Fill = excelize.Fill{
		Type:    "pattern",
		Pattern: 1,
		Color:   []string{"#FDE9D9"},
	}
	return style
}

// GetTdStyle 获取内容样式
func GetTdStyle() *excelize.Style {
	return GetBaseStyle()
}

//GetSpecialColStyle 获取特殊的样式
func GetSpecialColStyle() *excelize.Style {
	style := GetBaseStyle()
	style.Fill = excelize.Fill{
		Type:    "pattern",
		Pattern: 1,
		Color:   []string{"#d71345"},
	}
	return style
}
