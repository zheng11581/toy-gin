package models

import (
	"time"
)

// 全局的一些配置, json或者yaml格式存储，主要是变化不大的配置
type Conf struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"update_at"`
	Host      string    `gorm:"size:32;comments:'监控客户端的IP地址'" json:"host"`
	Name      string    `gorm:"size:32;comments:'监控客户端的名字'" json:"name"`
	Conf      string    `gorm:"size:2048;comments:'基础配置'" json:"conf"`
}

func (Conf) TableName() string {
	return "ingress_monitor_conf"
}

// 告警规则
type Rule struct {
	ID              uint      `json:"id"`
	CreatedAt       time.Time `json:"create_at"`
	UpdatedAt       time.Time `json:"update_at"`
	CreateUser      string    `gorm:"size:128;comments:'创建人'" json:"create_user"`
	UpdatedUser     string    `gorm:"size:128;comments:'最后更新人员'" json:"update_user"`
	RuleName        string    `gorm:"size:128;comments:'规则名称'" json:"rule_name"`
	RuleType        string    `gorm:"size:32;comments:'告警配置类型'" json:"rule_type"`
	SpecifyAppCode  string    `gorm:"size:256;comments:'指定appCode'" json:"specify_app_code"`
	SpecifyAppId    string    `gorm:"size:256;comments:'指定appID'" json:"specify_app_id"`
	Sla             float64   `gorm:"comments:'sla小于多少时候告警'" json:"sla"`
	MinAlarm        int       `gorm:"comments:'sla小于多少时候告警并且日志数量大于该值'" json:"min_alarm"`
	AlarmCount      int       `gorm:"comments:'异常日志达到多少的时候必定告警'" json:"alarm_count"`
	ContinuousTimes int       `gorm:"comments:'异常触发多少次在告警'" json:"continuous_times"`
	Interval        int       `gorm:"comments:'检查企业间'" json:"interval"`
	AlarmRule       string    `gorm:"size:512;comments:'告警规则'" json:"alarm_rule"`
	ExtendReceiver  string    `gorm:"size:258;comments:'自定义接受组, im group id, 多个逗号分隔'" json:"extend_receiver"`
	ConfID          int       `gorm:"comments:'属于哪个基础配置，外键'" json:"conf_id"`
	Conf            Conf      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;comments:'属于哪个基础配置'"`
}

func (Rule) TableName() string {
	return "ingress_monitor_rule"
}

// 特殊规则
type SpecialRule struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"create_at"`
	UpdatedAt   time.Time `json:"update_at"`
	CreateUser  string    `gorm:"size:128;comments:'创建人'" json:"create_user"`
	UpdatedUser string    `gorm:"size:128;comments:'最后更新人员'" json:"update_user"`
	Filter      string    `gorm:"size:2048;comments:'匹配规则'" json:"filter"`
	AlarmRule   string    `gorm:"size:2048;comments:'告警规则'" json:"alarm_rule"`
	RuleID      int       `gorm:"comments:'属于哪个告警规则，外键'" json:"rule_id"`
	Rule        Rule      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;comments:'属于哪个告警规则"`
}

func (SpecialRule) TableName() string {
	return "ingress_monitor_special"
}

// 静音配置
type SilenceRule struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"create_at"`
	UpdatedAt   time.Time `json:"update_at"`
	CreateUser  string    `gorm:"size:128;comments:'创建人'"`
	UpdatedUser string    `gorm:"size:128;comments:'最后更新人员'"`
	Expr        string    `gorm:"size:2048;comments:'规则表达式'"`
	ExpireAt    time.Time `gorm:"comments:'这条配置什么时候过期'"`
	RequireMan  string    `gorm:"size:128;comments:'配置谁要求加的'"`
	Reason      string    `gorm:"size:256;comments:'配置为什么添加'"`
	RuleID      int       `gorm:"comments:'属于哪个告警规则，外键'"`
	Rule        Rule      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;comments:'属于哪个告警规则"`
}

func (SilenceRule) TableName() string {
	return "ingress_monitor_silence"
}

type App struct {
	ID             uint      `json:"id"`
	CreatedAt      time.Time `json:"create_at"`
	UpdatedAt      time.Time `json:"update_at"`
	CreateUser     string    `gorm:"size:128;comments:'创建人'" json:"create_user"`
	UpdatedUser    string    `gorm:"size:128;comments:'最后更新人员'" json:"update_user"`
	SpecifyAppCode string    `gorm:"size:256;comments:'指定appCode';unique:true" json:"specify_app_code"`
	SpecifyAppName string    `gorm:"size:256;comments:'指定appID'" json:"specify_app_id"`
}

func (App) TableName() string {
	return "ingress_monitor_app"
}
