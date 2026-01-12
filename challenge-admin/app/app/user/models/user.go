package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type User struct {
	Id          int64           `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	LevelId     int64           `json:"levelId" gorm:"column:level_id;type:int;comment:用户等级编号"`
	UserName    string          `json:"userName" gorm:"column:username;type:varchar(100);comment:账号名称"`
	NickName    string          `json:"nickName" gorm:"column:nickname;type:varchar(100);comment:用户昵称"`
	TrueName    string          `json:"trueName" gorm:"column:true_name;type:varchar(100);comment:真实姓名"`
	Money       decimal.Decimal `json:"money" gorm:"column:money;type:decimal(30,2);comment:余额"`
	FreezeMoney decimal.Decimal `json:"freezeMoney" gorm:"column:freeze_money;type:decimal(30,2);comment:冻结金额"`
	Email       string          `json:"email" gorm:"column:email;type:varchar(300);comment:电子邮箱"`
	MobileTitle string          `json:"mobileTitle" gorm:"column:mobile_title;type:varchar(255);comment:用户手机号国家前缀"`
	Mobile      string          `json:"mobile" gorm:"column:mobile;type:varchar(100);comment:手机号码"`
	Avatar      string          `json:"avatar" gorm:"column:avatar;type:varchar(1000);comment:头像路径"`
	PayPwd      string          `json:"payPwd" gorm:"column:pay_pwd;type:varchar(100);comment:提现密码"`
	PayStatus   string          `json:"payStatus" gorm:"column:pay_status;type:char(1);comment:提现状态(1-启用 2-禁用)"`
	Pwd         string          `json:"pwd" gorm:"column:pwd;type:varchar(100);comment:登录密码"`
	RefCode     string          `json:"refCode" gorm:"column:ref_code;type:varchar(255);comment:推荐码"`
	ParentId    int64           `json:"parentId" gorm:"column:parent_id;type:int;comment:父级编号"`
	ParentIds   string          `json:"parentIds" gorm:"column:parent_ids;type:varchar(1000);comment:所有父级编号"`
	TreeSort    int64           `json:"treeSort" gorm:"column:tree_sort;type:int;comment:本级排序号（升序）"`
	TreeSorts   string          `json:"treeSorts" gorm:"column:tree_sorts;type:varchar(1000);comment:所有级别排序号"`
	TreeLeaf    string          `json:"treeLeaf" gorm:"column:tree_leaf;type:char(1);comment:是否最末级"`
	TreeLevel   int64           `json:"treeLevel" gorm:"column:tree_level;type:int;comment:层次级别"`
	Status      string          `json:"status" gorm:"column:status;type:char(1);comment:状态(1-正常 2-异常)"`
	Remark      string          `json:"remark" gorm:"column:remark;type:varchar(500);comment:备注信息"`
	RegisterAt  *time.Time      `json:"registerAt" gorm:"column:register_at;type:datetime;comment:注册时间"`
	RegisterIp  string          `json:"registerIp" gorm:"column:register_ip;type:varchar(45);comment:注册IP"`
	LastLoginAt *time.Time      `json:"lastLoginAt" gorm:"column:last_login_at;type:datetime;comment:最后登录时间"`
	LastLoginIp string          `json:"lastLoginIp" gorm:"column:last_login_ip;type:varchar(45);comment:最后登录IP"`
	CreateBy    int64           `json:"createBy" gorm:"column:create_by;type:int;comment:创建者"`
	UpdateBy    int64           `json:"updateBy" gorm:"column:update_by;type:int;comment:更新者"`
	CreatedAt   *time.Time      `json:"createdAt" gorm:"column:created_at;type:datetime;comment:创建时间"`
	UpdatedAt   *time.Time      `json:"updatedAt" gorm:"column:updated_at;type:datetime;comment:更新时间"`

	// 扩展
	UserLevel      *UserLevel `json:"userLevel" gorm:"foreignkey:level_id"`
	ParentUserName string     `json:"parentUserName" gorm:"-"`
	ParentRefCode  string     `json:"parentRefCode" gorm:"-"`
}

func (User) TableName() string {
	return "app_user"
}
