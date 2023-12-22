package models

import (
	"time"
)

// IngMonitorConf 全局的一些配置, json或者yaml格式存储，主要是变化不大的配置
type IngMonitorConf struct {
	ID   uint
	Host string `gorm:"size:32"`
	Name string `gorm:"size:32"`
	Conf string
}

// IngMonitorSilence 静音配置
type IngMonitorSilence struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreateUser  string    `gorm:"size:128;comments:'创建人'"`
	UpdatedUser string    `gorm:"size:128;comments:'最后更新人员'"`
	ConfType    int       `gorm:"comments:'这条配置属于配置'"`
	ConfId      int       `gorm:"comments:'这条配置属于那条配置的什么ID'"`
	Expr        string    `gorm:"size:2048;comments:'规则表达式'"`
	ExpireAt    time.Time `gorm:"comments:'这条配置什么时候过期'"`
	RequireMan  string    `gorm:"size:128;comments:'配置谁要求加的'"`
	Reason      string    `gorm:"size:256;comments:'配置为什么添加'"`
}

// IngMonitorRule 告警规则
type IngMonitorRule struct {
	ID             uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	CreateUser     string  `gorm:"size:128;comments:'创建人'"`
	UpdatedUser    string  `gorm:"size:128;comments:'最后更新人员'"`
	RuleName       string  `gorm:"size:128;comments:'规则名称'"`
	RuleType       string  `gorm:"size:32;comments:'告警配置类型'"`
	SpecifyAppCode string  `gorm:"size:256;comments:'指定appCode'"`
	SpecifyAppId   string  `gorm:"size:256;comments:'指定appID'"`
	Dc             string  `gorm:"size:64;comments:'指定数据中心'"`
	Env            string  `gorm:"size:64;comments:'指定环境'"`
	Sla            float64 `gorm:"comments:'sla小于多少时候告警'"`
	MinAlarm       int     `gorm:"comments:'sla小于多少时候告警并且日志数量大于该值'"`
	AlarmCount     int     `gorm:"comments:'异常日志达到多少的时候必定告警'"`
	Interval       int     `gorm:"comments:'检查企业间'"`
	AlarmRule      string  `gorm:"size:512;comments:'告警规则'"`
	ExtendReceiver string  `gorm:"size:258;comments:'自定义接受组, im group id, 多个逗号分隔'"`
}

// IngMonitorSpecialRule 特殊规则
type IngMonitorSpecialRule struct {
	ID          uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CreateUser  string `gorm:"size:128;comments:'创建人'"`
	UpdatedUser string `gorm:"size:128;comments:'最后更新人员'"`
	Filter      string `gorm:"size:2048;comments:'匹配规则'"`
	AlarmRule   string `gorm:"size:2048;comments:'告警规则'"`
}
