package handlers

import (
	"encoding/json"
	"errors"
	"time"
)

type ConfBase struct {
	Host string `json:"host" required:"true"`
	Name string `json:"name"`
	Conf string `json:"conf"`
}

type RuleBase struct {
	Id               int
	RuleName         string        `yaml:"ruleName" json:"rule_name"`                  // 必填
	SpecifyAppCode   string        `yaml:"specifyAppCode" json:"specify_app_code"`     // 指定AppCode, 说明这条规则为单个服务配置
	SpecifyAppId     string        `yaml:"specifyAppId" json:"specify_app_id"`         // 可选 指定数据中心环境的时候必须填写
	Sla              float64       `yaml:"sla" json:"sla"`                             // 必填 需要保证多少SLA OR AlarmCount满足
	AlarmCount       int           `yaml:"alarmCount" json:"alarm_count"`              // 必填 警告数量达到多少告警 OR SLA不满足 默认1000000
	MinAlarm         int           `yaml:"minAlarm" json:"min_alarm"`                  // 必须满足最小的告警数据
	ExtendReceiver   []string      `yaml:"extendReceiver" json:"extend_receiver"`      // 可选 扩展指定接收者
	OnlyExtend       bool          `yaml:"onlyExtend" json:"only_extend"`              // 可选 只发送给指定接收者
	Interval         int64         `yaml:"interval" json:"interval"`                   // 可选 告警检查间隔 默认60s
	Silences         []string      `yaml:"silences" json:"silences"`                   // 可选 静音规则
	AlarmRule        string        `yaml:"alarmRule" json:"alarm_rule"`                // 必填 告警规则
	SpecialAlarmRule []SpecialRule `yaml:"specialAlarmRule" json:"special_alarm_rule"` // 可选
	RuleType         string        `yaml:"ruleType" json:"rule_type"`                  // 必填 监控类型 slow or errorCode
	ContinuousTimes  int           `yaml:"continuousTimes" json:"continuous_times"`    // 持续几次才告警
}

type SpecialRule struct {
	Filter    string `yaml:"filter" json:"filter"`
	AlarmRule string `yaml:"alarmRule" json:"alarm_rule"`
}

type SilenceRule struct {
	Expr       string    `yaml:"expr" json:"expr"`
	ExpireAt   time.Time `yaml:"expireAt" json:"expire_at"`
	RequireMan string    `yaml:"requireMan" json:"require_man"`
	Reason     string    `yaml:"reason" json:"reason"`
}

type AppBase struct {
	SpecifyAppCode string `yaml:"specifyAppCode" json:"specify_app_code"`
	SpecifyAppId   string `yaml:"specifyAppId" json:"specify_app_id"`
}

// bindReqToM will bind reqObj to obj
func BindReqToM(reqObj any, obj any) error {
	reqBytes, err := json.Marshal(&reqObj)
	if err != nil {
		return errors.New("Marshal失败")
	}
	err = json.Unmarshal(reqBytes, &obj)
	if err != nil {
		return errors.New("Unmarshal失败")
	}
	return nil
}
