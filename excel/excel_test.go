package excel

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/xuri/excelize/v2"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"strings"
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

func TestGetBaseStyle(t *testing.T) {
	b, err := ToUTF8("gbk", []byte("���Ч����Ч���ú�ʵ���ʺϼ�ͥʹ��"))
	if err != nil {
		panic(err)
	}
	t.Log(string(b))
	t.Log(FromUTF8("gbk", []byte("���Ч����Ч���ú�ʵ���ʺϼ�ͥʹ��")))
}

// ToUTF8 convert from CJK encoding to UTF-8
func ToUTF8(from string, s []byte) ([]byte, error) {
	var reader *transform.Reader
	switch strings.ToLower(from) {
	case "gbk", "cp936", "windows-936":
		reader = transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	case "gb18030":
		reader = transform.NewReader(bytes.NewReader(s), simplifiedchinese.GB18030.NewDecoder())
	case "gb2312":
		reader = transform.NewReader(bytes.NewReader(s), simplifiedchinese.HZGB2312.NewDecoder())
	case "big5", "big-5", "cp950":
		reader = transform.NewReader(bytes.NewReader(s), traditionalchinese.Big5.NewDecoder())
	case "euc-kr", "euckr", "cp949":
		reader = transform.NewReader(bytes.NewReader(s), korean.EUCKR.NewDecoder())
	case "euc-jp", "eucjp":
		reader = transform.NewReader(bytes.NewReader(s), japanese.EUCJP.NewDecoder())
	case "shift-jis", "iso-2022-jp", "cp932", "windows-31j":
		reader = transform.NewReader(bytes.NewReader(s), japanese.ShiftJIS.NewDecoder())
	case "iso-8859-1":
		return charmap.ISO8859_1.NewDecoder().Bytes(s)
	case "cp1252", "windows-1252":
		//return fromWindows1252(s), nil
		//return s, nil
	// case "iso-2022-jp", "cp932", "windows-31j":
	// 	reader = transform.NewReader(bytes.NewReader(s), japanese.ISO2022JP.NewDecoder())
	default:
		return s, errors.New("Unsupported encoding " + from)
	}
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// FromUTF8 convert from UTF-8 encoding to CJK encoding
func FromUTF8(to string, s []byte) ([]byte, error) {
	var reader *transform.Reader
	switch strings.ToLower(to) {
	case "gbk", "cp936", "windows-936":
		reader = transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	case "gb18030":
		reader = transform.NewReader(bytes.NewReader(s), simplifiedchinese.GB18030.NewEncoder())
	case "gb2312":
		reader = transform.NewReader(bytes.NewReader(s), simplifiedchinese.HZGB2312.NewEncoder())
	case "big5", "big-5", "cp950":
		reader = transform.NewReader(bytes.NewReader(s), traditionalchinese.Big5.NewEncoder())
	case "euc-kr", "euckr", "cp949":
		reader = transform.NewReader(bytes.NewReader(s), korean.EUCKR.NewEncoder())
	case "euc-jp", "eucjp":
		reader = transform.NewReader(bytes.NewReader(s), japanese.EUCJP.NewEncoder())
	case "shift-jis", "iso-2022-jp", "cp932", "windows-31j":
		reader = transform.NewReader(bytes.NewReader(s), japanese.ShiftJIS.NewEncoder())
	// case "iso-2022-jp", "cp932", "windows-31j":
	// 	reader = transform.NewReader(bytes.NewReader(s), japanese.ISO2022JP.NewEncoder())
	default:
		return s, errors.New("Unsupported encoding " + to)
	}
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}
