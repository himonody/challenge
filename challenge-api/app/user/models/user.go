package models

import (
	"time"

	"github.com/shopspring/decimal"
)

// AppUser 用户管理
type AppUser struct {
	ID          int             `gorm:"column:id;primaryKey;autoIncrement;comment:用户编码" json:"id"`
	LevelID     int             `gorm:"column:level_id;not null;default:1;comment:用户等级编号" json:"level_id"`
	Username    string          `gorm:"column:username;type:varchar(100);not null;default:'';comment:账号名称/用户名" json:"username"`
	Nickname    string          `gorm:"column:nickname;type:varchar(100);not null;default:'';comment:用户昵称" json:"nickname"`
	TrueName    string          `gorm:"column:true_name;type:varchar(100);not null;default:'';comment:真实姓名" json:"true_name"`
	Money       decimal.Decimal `gorm:"column:money;type:decimal(30,2);not null;default:0.00;comment:余额" json:"money"`
	FreezeMoney decimal.Decimal `gorm:"column:freeze_money;type:decimal(30,2);not null;default:0.00;comment:冻结金额" json:"freeze_money"`
	Email       string          `gorm:"column:email;type:varchar(300);comment:电子邮箱" json:"email"`
	MobileTitle string          `gorm:"column:mobile_title;type:varchar(255);default:'+86';comment:用户手机号国家前缀" json:"mobile_title"`
	Mobile      string          `gorm:"column:mobile;type:varchar(100);comment:手机号码" json:"mobile"`
	Avatar      string          `gorm:"column:avatar;type:varchar(1000);comment:头像路径" json:"avatar"`
	PayPwd      string          `gorm:"column:pay_pwd;type:varchar(100);not null;default:'';comment:提现密码" json:"-"`
	PayStatus   string          `gorm:"column:pay_status;type:char(1);not null;default:'1';comment:提现状态(1-启用 2-禁用)" json:"pay_status"`
	Pwd         string          `gorm:"column:pwd;type:varchar(100);not null;default:'';comment:登录密码" json:"-"`
	RefCode     string          `gorm:"column:ref_code;type:varchar(255);comment:推荐码" json:"ref_code"`
	ParentID    int             `gorm:"column:parent_id;not null;default:0;comment:父级编号" json:"parent_id"`
	ParentIDs   string          `gorm:"column:parent_ids;type:varchar(1000);not null;default:'';comment:所有父级编号" json:"parent_ids"`
	TreeSort    int             `gorm:"column:tree_sort;not null;default:0;comment:本级排序号" json:"tree_sort"`
	TreeSorts   string          `gorm:"column:tree_sorts;type:varchar(1000);not null;default:'0';comment:所有级别排序号" json:"tree_sorts"`
	TreeLeaf    string          `gorm:"column:tree_leaf;type:char(1);not null;default:'0';comment:是否最末级" json:"tree_leaf"`
	TreeLevel   int             `gorm:"column:tree_level;not null;default:0;comment:层次级别" json:"tree_level"`
	Status      string          `gorm:"column:status;type:char(1);not null;default:'1';comment:状态(1-正常 2-异常)" json:"status"`
	Remark      string          `gorm:"column:remark;type:varchar(255);comment:备注信息" json:"remark"`
	CreateBy    int             `gorm:"column:create_by;not null;default:0;comment:创建者" json:"create_by"`
	UpdateBy    int             `gorm:"column:update_by;not null;default:0;comment:更新者" json:"update_by"`
	CreatedAt   time.Time       `gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"created_at"`
	UpdatedAt   time.Time       `gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP;comment:更新时间" json:"updated_at"`
}

// TableName 指定表名
func (AppUser) TableName() string {
	return "app_user"
}
