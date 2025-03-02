// Code generated by goctl. DO NOT EDIT.
package types

type HomeBannerResponse struct {
	Banners []*Banner `json:"banners"`
}

type Banner struct {
	ID   int64  `json:"id"`
	Name string `json:"name"` // 名称
	URL  string `json:"url"`  // 图片地址
}

type FlashSaleResponse struct {
	StartTime int64      `json:"start_time"` // 抢购开始时间
	Products  []*Product `json:"products"`
}

type Product struct {
	ID          int64    `json:"id"`          // 商品ID
	Name        string   `json:"name"`        // 产品名称
	Images      []string `json:"images"`      // 图片
	Description string   `json:"description"` // 商品描述
	Price       float64  `json:"price"`       // 商品价格
	Stock       int64    `json:"stock"`       // 库存
	Category    string   `json:"category"`    // 分类
	Status      int64    `json:"status"`      // 状态：1-正常，2-下架
	CreateTime  int64    `json:"create_time"` // 创建时间
	UpdateTime  int64    `json:"update_time"` // 更新时间
}

type RecommendRequest struct {
	Cursor int64 `json:"cursor"`
	Ps     int64 `form:"ps,default=20"` // 每页大小
}

type RecommendResponse struct {
	Products      []*Product `json:"products"`
	IsEnd         bool       `json:"is_end"`         // 是否最后一页
	RecommendTime int64      `json:"recommend_time"` // 商品列表最后一个商品的推荐时间
}

type CategoryListRequest struct {
	Cursor   int64  `form:"cursor"`        // 分页游标
	Ps       int64  `form:"ps,default=20"` // 每页大小
	Category string `form:"category"`      // 分类
	Sort     string `form:"sort"`          // 排序
}

type CategoryListResponse struct {
	Products []*Product `json:"products"`
	IsEnd    bool       `json:"is_end"`
	LastVal  int64      `json:"last_val"`
}

type CartListRequest struct {
	UID int64 `form:"uid"`
}

type CartListResponse struct {
	Products []*CartProduct `json:"products"`
}

type CartProduct struct {
	Product *Product `json:"product"`
	Count   int64    `json:"count"` // 购买数量
}

type ProductCommentRequest struct {
	ProductID int64 `form:"product_id"`
	Cursor    int64 `form:"cursor"`
	Ps        int64 `form:"ps,default=20"`
}

type ProductCommentResponse struct {
	Comments    []*Comment `json:"comments"`
	IsEnd       bool       `json:"is_end"`       // 是否最后一页
	CommentTime int64      `json:"comment_time"` // 评论列表最后一个评论的时间
}

type Comment struct {
	ID         int64    `json:"id"`          // 评论ID
	ProductID  int64    `json:"product_id"`  // 商品ID
	Content    string   `json:"content"`     // 评论内容
	Images     []*Image `json:"images"`      // 评论图片
	User       *User    `json:"user"`        // 用户信息
	CreateTime int64    `json:"create_time"` // 评论时间
	UpdateTime int64    `json:"update_time"` // 更新时间
}

type User struct {
	ID     int64  `json:"id"`     // 用户ID
	Name   string `json:"name"`   // 用户名
	Avatar string `json:"avatar"` // 头像
}

type Image struct {
	ID  int64  `json:"id"`
	URL string `json:"url"`
}

type OrderListRequest struct {
	UID    int64 `form:"uid"`
	Status int32 `form:"status,optional"`
	Cursor int64 `form:"cursor,optional"`
	Ps     int64 `form:"ps,default=20"`
}

type OrderListResponse struct {
	Orders    []*Order `json:"orders"`
	IsEnd     bool     `json:"is_end"` // 是否最后一页
	OrderTime int64    `json:"order_time"`
}

type Order struct {
	OrderID            string  `json:"order_id"`
	Status             int32   `json:"status"`
	Quantity           int64   `json:"quantity"`
	Payment            float64 `json:"payment"`
	TotalPrice         float64 `json:"total_price"`
	CreateTime         int64   `json:"create_time"`
	ProductID          int64   `json:"product_id"`
	ProductName        string  `json:"product_name"`
	ProductImage       string  `json:"product_image"`
	ProductDescription string  `json:"product_description"`
}

type ProductDetailRequest struct {
	ProductID int64 `form:"product_id"`
}

type ProductDetailResponse struct {
	Product  *Product   `json:"product"`
	Comments []*Comment `json:"comments"`
}

type UserReceiveAddress struct {
	Id            int64  `json:"id"`
	Uid           uint64 `json:"uid"`            //用户id
	Name          string `json:"name"`           //收货人名称
	Phone         string `json:"phone"`          //手机号
	IsDefault     uint8  `json:"is_default"`     //是否为默认地址
	PostCode      string `json:"post_code"`      //邮政编码
	Province      string `json:"province"`       //省份/直辖市
	City          string `json:"city"`           //城市
	Region        string `json:"region"`         //区
	DetailAddress string `json:"detail_address"` //详细地址(街道)
	IsDelete      uint8  `json:"is_delete"`      //是否删除
	CreateTime    int64  `json:"create_time"`    //数据创建时间[禁止在代码中赋值]
	UpdateTime    int64  `json:"update_time"`    //数据更新时间[禁止在代码中赋值]
}

type UserReceiveAddressListReq struct {
}

type UserReceiveAddressListRes struct {
	List []UserReceiveAddress `json:"list"`
}

type UserReceiveAddressAddReq struct {
	UserReceiveAddress UserReceiveAddress `json:"UserReceiveAddress"`
}

type UserReceiveAddressAddRes struct {
}

type UserReceiveAddressEditReq struct {
	UserReceiveAddress UserReceiveAddress `json:"UserReceiveAddress"`
}

type UserReceiveAddressEditRes struct {
}

type UserReceiveAddressDelReq struct {
	Id int64 `json:"id"`
}

type UserReceiveAddressDelRes struct {
}

type UserInfo struct {
	Id         uint64 `json:"id"`          //用户ID
	Username   string `json:"username"`    //用户名
	Password   string `json:"password"`    //用户密码，MD5加密
	Phone      string `json:"phone"`       //手机号
	Question   string `json:"question"`    //找回密码问题
	Answer     string `json:"answer"`      //找回密码答案
	CreateTime int64  `json:"create_time"` //创建时间
	UpdateTime int64  `json:"update_time"` //更新时间
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResp struct {
	AccessToken  string `json:"accessToken"`
	AccessExpire int64  `json:"accessExpire"`
}

type UserInfoReq struct {
}

type UserInfoResp struct {
	UserInfo UserInfo `json:"userInfo"`
}
