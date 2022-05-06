package tool

// @Title  相关的model结构体
// @Description  请填写文件描述（需要改）
// @Author  clx  2022/5/6 13:35
// @Update  clx  2022/5/6 13:35

// ImgListItem 京东图片结构体子项
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

// GetProductPageImageCommentListRsp 图片接口返回结构体
type GetProductPageImageCommentListRsp struct {
	ImgComments struct {
		ImgCommentCount int            `json:"imgCommentCount"`
		ImgList         []*ImgListItem `json:"imgList"`
	} `json:"imgComments"`
	ReferenceID int64 `json:"referenceId"`
}

// ProductDetail 详情
type ProductDetail struct {
	ProductID string `json:"product_id"`
	Title     string `json:"title"`
}
