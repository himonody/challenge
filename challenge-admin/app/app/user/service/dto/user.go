package dto

import (
	commDto "challenge-admin/app/app/common/dto"
	"challenge-admin/core/dto"
	"time"

	"github.com/shopspring/decimal"
)

type UserQueryReq struct {
	dto.Pagination    `search:"-"`
	BeginCreatedAt    string          `form:"beginCreatedAt" search:"type:gte;column:created_at;table:app_user" comment:"创建时间"`
	EndCreatedAt      string          `form:"endCreatedAt" search:"type:lte;column:created_at;table:app_user" comment:"创建时间"`
	LevelId           int64           `form:"levelId"  search:"type:exact;column:level_id;table:app_user" comment:"用户等级编号"`
	LevelIds          []int64         `form:"levelId"  search:"type:exact;column:level_id;table:app_user" comment:"用户等级编号"`
	UserName          string          `form:"userName"  search:"type:exact;column:username;table:app_user" comment:"账号名称"`
	NickName          string          `form:"nickName"  search:"type:exact;column:nickname;table:app_user" comment:"用户昵称"`
	TrueName          string          `form:"trueName"  search:"type:exact;column:true_name;table:app_user" comment:"真实姓名"`
	Money             decimal.Decimal `form:"money"  search:"type:exact;column:money;table:app_user" comment:"余额"`
	FreezeMoney       decimal.Decimal `form:"freezeMoney"  search:"type:exact;column:freeze_money;table:app_user" comment:"冻结金额"`
	Email             string          `form:"email"  search:"type:exact;column:email;table:app_user" comment:"电子邮箱"`
	MobileTitle       string          `form:"mobileTitle"  search:"type:exact;column:mobile_title;table:app_user" comment:"国家区号"`
	MobileTitles      []string        `form:"-"  search:"type:in;column:mobile_title;table:app_user" comment:"国家区号列表"`
	Mobile            string          `form:"mobile"  search:"type:exact;column:mobile;table:app_user" comment:"手机号码"`
	Avatar            string          `form:"avatar"  search:"type:exact;column:avatar;table:app_user" comment:"头像"`
	PayPwd            string          `form:"payPwd"  search:"type:exact;column:pay_pwd;table:app_user" comment:"提现密码"`
	PayStatus         string          `form:"payStatus"  search:"type:exact;column:pay_status;table:app_user" comment:"提现状态"`
	Pwd               string          `form:"pwd"  search:"type:exact;column:pwd;table:app_user" comment:"登录密码"`
	RefCode           string          `form:"refCode"  search:"type:exact;column:ref_code;table:app_user" comment:"推荐码"`
	ParentId          int64           `form:"parentId"  search:"type:exact;column:parent_id;table:app_user" comment:"父级编号"`
	TreeSort          int64           `form:"treeSort"  search:"type:exact;column:tree_sort;table:app_user" comment:"本级排序号"`
	TreeLevel         int64           `form:"treeLevel"  search:"type:exact;column:tree_level;table:app_user" comment:"层级"`
	Status            string          `form:"status" search:"type:exact;column:status;table:app_user" comment:"状态"`
	RegisterAt        string          `form:"registerAt" search:"type:exact;column:register_at;table:app_user" comment:"注册时间"`
	RegisterIp        string          `form:"registerIp" search:"type:exact;column:register_ip;table:app_user" comment:"注册IP"`
	LastLoginAt       string          `form:"lastLoginAt" search:"type:exact;column:last_login_at;table:app_user" comment:"最后登录时间"`
	LastLoginIp       string          `form:"lastLoginIp" search:"type:exact;column:last_login_ip;table:app_user" comment:"最后登录IP"`
	ShowInfo          bool            `form:"-"  search:"-" comment:"是否明文显示加密信息"`
	commDto.LevelJoin `search:"type:inner;on:id:level_id;table:app_user;join:app_user_level"`
	//扩展
	ParentRefCode string `form:"parentRefCode"  search:"-" comment:"上级用户邀请码"`
	UserOrder
}

type UserOrder struct {
	IdOrder          int64           `form:"idOrder"  search:"type:order;column:id;table:app_user"`
	LevelIdOrder     int64           `form:"levelIdOrder"  search:"type:order;column:level_id;table:app_user"`
	UserNameOrder    string          `form:"userNameOrder"  search:"type:order;column:username;table:app_user"`
	NickNameOrder    string          `form:"nickNameOrder"  search:"type:order;column:nickname;table:app_user"`
	TrueNameOrder    string          `form:"trueNameOrder"  search:"type:order;column:true_name;table:app_user"`
	MoneyOrder       decimal.Decimal `form:"moneyOrder"  search:"type:order;column:money;table:app_user"`
	FreezeMoneyOrder decimal.Decimal `form:"freezeMoneyOrder"  search:"type:order;column:freeze_money;table:app_user"`
	EmailOrder       string          `form:"emailOrder"  search:"type:order;column:email;table:app_user"`
	MobileTitleOrder string          `form:"mobileTitleOrder"  search:"type:order;column:mobile_title;table:app_user"`
	MobileOrder      string          `form:"mobileOrder"  search:"type:order;column:mobile;table:app_user"`
	AvatarOrder      string          `form:"avatarOrder"  search:"type:order;column:avatar;table:app_user"`
	PayPwdOrder      string          `form:"payPwdOrder"  search:"type:order;column:pay_pwd;table:app_user"`
	PayStatusOrder   string          `form:"payStatusOrder"  search:"type:order;column:pay_status;table:app_user"`
	PwdOrder         string          `form:"pwdOrder"  search:"type:order;column:pwd;table:app_user"`
	RefCodeOrder     string          `form:"refCodeOrder"  search:"type:order;column:ref_code;table:app_user"`
	ParentIdOrder    int64           `form:"parentIdOrder"  search:"type:order;column:parent_id;table:app_user"`
	ParentIdsOrder   string          `form:"parentIdsOrder"  search:"type:order;column:parent_ids;table:app_user"`
	TreeSortOrder    int64           `form:"treeSortOrder"  search:"type:order;column:tree_sort;table:app_user"`
	TreeSortsOrder   string          `form:"treeSortsOrder"  search:"type:order;column:tree_sorts;table:app_user"`
	TreeLeafOrder    string          `form:"treeLeafOrder"  search:"type:order;column:tree_leaf;table:app_user"`
	TreeLevelOrder   int64           `form:"treeLevelOrder"  search:"type:order;column:tree_level;table:app_user"`
	StatusOrder      string          `form:"statusOrder"  search:"type:order;column:status;table:app_user"`
	RemarkOrder      string          `form:"remarkOrder"  search:"type:order;column:remark;table:app_user"`
	RegisterAtOrder  *time.Time      `form:"registerAtOrder"  search:"type:order;column:register_at;table:app_user"`
	RegisterIpOrder  string          `form:"registerIpOrder"  search:"type:order;column:register_ip;table:app_user"`
	LastLoginAtOrder *time.Time      `form:"lastLoginAtOrder"  search:"type:order;column:last_login_at;table:app_user"`
	LastLoginIpOrder string          `form:"lastLoginIpOrder"  search:"type:order;column:last_login_ip;table:app_user"`
	CreateByOrder    int64           `form:"createByOrder"  search:"type:order;column:create_by;table:app_user"`
	UpdateByOrder    int64           `form:"updateByOrder"  search:"type:order;column:update_by;table:app_user"`
	CreatedAtOrder   *time.Time      `form:"createdAtOrder"  search:"type:order;column:created_at;table:app_user"`
	UpdatedAtOrder   *time.Time      `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:app_user"`
}

func (m *UserQueryReq) GetNeedSearch() interface{} {
	return *m
}

type UserInsertReq struct {
	LevelId     int64           `json:"levelId" comment:"用户等级编号"`
	UserName    string          `json:"userName" comment:"账号名称"`
	NickName    string          `json:"nickName" comment:"用户昵称"`
	TrueName    string          `json:"trueName" comment:"真实姓名"`
	Money       decimal.Decimal `json:"money" comment:"余额"`
	FreezeMoney decimal.Decimal `json:"freezeMoney" comment:"冻结金额"`
	Email       string          `json:"email" comment:"电子邮箱"`
	MobileTitle string          `json:"mobileTitle" comment:"用户手机号国家前缀"`
	Mobile      string          `json:"mobile" comment:"手机号码"`
	Avatar      string          `json:"avatar" comment:"头像路径"`
	PayPwd      string          `json:"payPwd" comment:"提现密码"`
	PayStatus   string          `json:"payStatus" comment:"提现状态(1-启用 2-禁用)"`
	Pwd         string          `json:"pwd" comment:"登录密码"`
	RefCode     string          `json:"refCode" comment:"推荐码"`
	Emails      string          `json:"emails" comment:"邮箱集合"`
	Mobiles     string          `json:"mobiles" comment:"手机集合"`
	CurrUserId  int64           `json:"-" comment:"当前登陆用户"`
}

type UserUpdateReq struct {
	Id          int64           `json:"-" uri:"id" comment:"用户编号"` // 用户编号
	LevelId     int64           `json:"levelId" comment:"用户等级编号"`
	UserName    string          `json:"userName" comment:"账号名称"`
	NickName    string          `json:"nickName" comment:"用户昵称"`
	TrueName    string          `json:"trueName" comment:"真实姓名"`
	Money       decimal.Decimal `json:"money" comment:"余额"`
	FreezeMoney decimal.Decimal `json:"freezeMoney" comment:"冻结金额"`
	Email       string          `json:"email" comment:"电子邮箱"`
	MobileTitle string          `json:"mobileTitle" comment:"用户手机号国家前缀"`
	Mobile      string          `json:"mobile" comment:"手机号码"`
	Avatar      string          `json:"avatar" comment:"头像路径"`
	PayPwd      string          `json:"payPwd" comment:"提现密码"`
	PayStatus   string          `json:"payStatus" comment:"提现状态(1-启用 2-禁用)"`
	Pwd         string          `json:"pwd" comment:"登录密码"`
	RefCode     string          `json:"refCode" comment:"推荐码"`
	CurrUserId  int64           `json:"-" comment:"当前登陆用户"`
}
type UserRechargeReq struct {
	Id         int64           `json:"-" uri:"id" comment:"用户编号"` // 用户编号
	Amount     decimal.Decimal `json:"amount" comment:"金额"`
	CurrUserId int64           `json:"-" comment:"当前登陆用户"`
}
type UserDeductReq struct {
	Id         int64           `json:"-" uri:"id" comment:"用户编号"` // 用户编号
	Amount     decimal.Decimal `json:"amount" comment:"金额"`
	CurrUserId int64           `json:"-" comment:"当前登陆用户"`
}
type UserPasswordReq struct {
	Id         int64 `json:"-" uri:"id" comment:"用户编号"` // 用户编号
	CurrUserId int64 `json:"-" comment:"当前登陆用户"`
}
type UserPayPasswordReq struct {
	Id         int64 `json:"-" uri:"id" comment:"用户编号"` // 用户编号
	CurrUserId int64 `json:"-" comment:"当前登陆用户"`
}
type UserStatusUpdateReq struct {
	Id         int64  `json:"-" uri:"id" comment:"用户ID"` // 用户ID
	Status     string `json:"status" comment:"状态"`
	CurrUserId int64  `json:"-" comment:""`
}
type UserPayStatusUpdateReq struct {
	Id         int64  `json:"-" uri:"id" comment:"用户ID"` // 用户ID
	PayStatus  string `json:"pay_status" comment:"支付状态"`
	CurrUserId int64  `json:"-" comment:""`
}

// UserGetReq 功能获取请求参数
type UserGetReq struct {
	Id int64 `uri:"id"`
}
