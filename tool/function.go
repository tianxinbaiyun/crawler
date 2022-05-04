package tool

import (
	"encoding/json"
	"fmt"
	"github.com/tianxinbaiyun/crawler/config"
	"io/ioutil"
	"net/http"
	"strings"
)

// @Title  请填写文件名称（需修改）
// @Description  请填写文件描述（需修改）
// @Author  clx  2022/5/4 9:36 上午
// @Update  clx  2022/5/4 9:36 上午

// GetSearchUrl 获取搜索地址
func GetSearchUrl(keyword string, page int) (s string) {
	s = config.C.Crawler.SearchUrl
	if keyword != "" {
		s = strings.ReplaceAll(s, "{{keyword}}", keyword)

	} else {
		s = strings.ReplaceAll(s, "&keyword={{keyword}}", "")
	}

	if page > 0 {
		s = strings.ReplaceAll(s, "{{page}}", fmt.Sprintf("%d", page))
	} else {
		s = strings.ReplaceAll(s, "&page={{page}}", "")
	}

	return
}

//GetProductPageImageCommentList 获取京东评论列表
func GetProductPageImageCommentList(productID string, page, pageSize int) (rsp *GetProductPageImageCommentListRsp, err error) {

	url := "https://club.jd.com/discussion/getProductPageImageCommentList.action?isShadowSku=0"

	if productID != "" {
		url = fmt.Sprintf("%s&productId=%s", url, productID)
	}
	if page > 0 {
		url = fmt.Sprintf("%s&page=%d", url, page)
	}
	if pageSize > 0 {
		url = fmt.Sprintf("%s&pageSize=%d", url, pageSize)
	}
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("authority", "club.jd.com")
	req.Header.Add("accept", "*/*")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Add("cookie", "__jdv=76161171|direct|-|none|-|1651628159902; __jdu=1651628159902882856289; areaId=19; PCSYCityID=CN_440000_440300_0; shshshfpa=c5cf4e2b-ec5a-9305-7b65-5d56637f9d7a-1651628162; shshshfpb=dkymtl45vL25neqoOlmP6RQ; __jdc=122270672; shshshfp=a8d5f7226be831ee42175b035de435b2; ip_cityCode=1607; ipLoc-djd=19-1607-4773-62121; jwotest_product=99; user-key=df0c4835-1931-4b7d-bd9a-bb7eb5e685e9; __jda=122270672.1651628159902882856289.1651628160.1651640464.1651644444.5; wlfstk_smdl=90g5mpmw7tpoowoqcrsxp0gtf5ng0zra; thor=166B0217F5BD758A102A863263DA550C95A8CE525929534878474DFBF9C17B1B2CE51475261841D915208ECB2085BB0A5D3CBC149197282F36AF61549133EEDCC719CD9C211C9DFBA6B8DD59B658E72CE7E55049A147B80CDA6805AB840D59F9BEF7484A3EB2CE41E90816F704A55B6762AE36CA7461781D097BCA53DCE56420; pinId=Mh5NbTnoOSUY_Wsft5KKsQ; pin=%E5%A4%A9%E6%AD%86%E7%99%BD%E4%BA%91; unick=%E5%A4%A9%E6%AD%86%E7%99%BD%E4%BA%91; ceshi3.com=000; _tp=4JDUexuHnTB942ZV8wEtRW4vB8iHnL9pnIGnhNwq1h8r3rXK61svcL2yNKeHZSeh; _pst=%E5%A4%A9%E6%AD%86%E7%99%BD%E4%BA%91; token=5737d200b8b2ba6d2e2360574a3cf3c6,3,917580; __tk=s4YBI12zsalulcxOsrYkX1TzsgnvJDoOd4PuetvpJqlvJDxNs3YvXErpsriDJcbEsorglroE,3,917580; shshshsID=f5301d9f3d2ff6393e825eb01ed76906_2_1651644585323; __jdb=122270672.10.1651628159902882856289|5.1651644444; 3AB9D23F7A4B3C9B=3MPX4KF54SRNU7YEJBOLBYV55R3LVP3KOFCQVEAJ5HZYYRDTJHAZLDZBZVRIOAUZVGTM3DYV7EWB7T7KL7M4NSBMA4; JSESSIONID=5FE89382E3A442A0E54907D57B496D3D.s1; ipLoc-djd=1-72-2799-0; JSESSIONID=59BEEFD4AFFB06CC53CA8DF15894EF9B.s1")
	req.Header.Add("referer", "https://item.jd.com/")
	req.Header.Add("sec-ch-ua", "\" Not A;Brand\";v=\"99\", \"Chromium\";v=\"100\", \"Google Chrome\";v=\"100\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"macOS\"")
	req.Header.Add("sec-fetch-dest", "script")
	req.Header.Add("sec-fetch-mode", "no-cors")
	req.Header.Add("sec-fetch-site", "same-site")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(string(body))
	rsp = &GetProductPageImageCommentListRsp{}
	err = json.Unmarshal(body, rsp)
	if err != nil {
		return
	}
	return
}

type ImgListItem struct {
	ImageURL  string `json:"imageUrl"`
	ImageID   int    `json:"imageId"`
	MediaType int    `json:"mediaType"`
	CommentVo struct {
		ID               int64         `json:"id"`
		Topped           int           `json:"topped"`
		GUID             string        `json:"guid"`
		Content          string        `json:"content"`
		CreationTime     string        `json:"creationTime"`
		IsTop            bool          `json:"isTop"`
		ReferenceID      string        `json:"referenceId"`
		ReferenceTime    string        `json:"referenceTime"`
		ReferenceType    string        `json:"referenceType"`
		ReferenceTypeID  int           `json:"referenceTypeId"`
		FirstCategory    int           `json:"firstCategory"`
		SecondCategory   int           `json:"secondCategory"`
		ThirdCategory    int           `json:"thirdCategory"`
		ReplyCount       int           `json:"replyCount"`
		ReplyCount2      int           `json:"replyCount2"`
		Score            int           `json:"score"`
		Status           int           `json:"status"`
		UsefulVoteCount  int           `json:"usefulVoteCount"`
		UselessVoteCount int           `json:"uselessVoteCount"`
		UserImage        string        `json:"userImage"`
		UserImageURL     string        `json:"userImageUrl"`
		UserLevelID      string        `json:"userLevelId"`
		UserProvince     string        `json:"userProvince"`
		ViewCount        int           `json:"viewCount"`
		OrderID          int           `json:"orderId"`
		IsReplyGrade     bool          `json:"isReplyGrade"`
		UID              int           `json:"uid"`
		Nickname         string        `json:"nickname"`
		UserClient       int           `json:"userClient"`
		MergeOrderStatus int           `json:"mergeOrderStatus"`
		DiscussionID     int           `json:"discussionId"`
		ProductColor     string        `json:"productColor"`
		ProductSize      string        `json:"productSize"`
		Integral         int           `json:"integral"`
		UserImgFlag      int           `json:"userImgFlag"`
		AnonymousFlag    int           `json:"anonymousFlag"`
		UserLevelName    string        `json:"userLevelName"`
		PlusAvailable    int           `json:"plusAvailable"`
		ProductSales     []interface{} `json:"productSales"`
		MobileVersion    string        `json:"mobileVersion"`
		OfficialsStatus  int           `json:"officialsStatus"`
		Excellent        bool          `json:"excellent"`
		GsValueTotal     int           `json:"gsValueTotal"`
		VtFlag           int           `json:"vtFlag"`
		ExtMap           struct {
			BuyCount int `json:"buyCount"`
		} `json:"extMap"`
		Recommend      bool   `json:"recommend"`
		UserLevelColor string `json:"userLevelColor"`
		UserClientShow string `json:"userClientShow"`
		IsMobile       bool   `json:"isMobile"`
	}
}

// GetProductPageImageCommentListRsp 图片返回
type GetProductPageImageCommentListRsp struct {
	ImgComments struct {
		ImgCommentCount int            `json:"imgCommentCount"`
		ImgList         []*ImgListItem `json:"imgList"`
	} `json:"imgComments"`
	ReferenceID int64 `json:"referenceId"`
}
