package handlers

type ConfBase struct {
	Host string `json:"host"`
	Name string `json:"name"`
	Conf string `json:"conf"`
}

type RuleBase struct {
	Id             int
	RuleName       string `yaml:"ruleName"`       // 必填
	SpecifyAppCode string `yaml:"specifyAppCode"` // 指定AppCode, 说明这条规则为单个服务配置
	//SpecifyDc      string   `yaml:"specifyDc"`
	//SpecifyEnv     string   `yaml:"specifyEnv"`
	SpecifyAppId string  `yaml:"specifyAppId"` // 可选 指定数据中心环境的时候必须填写
	Sla          float64 `yaml:"sla"`          // 必填 需要保证多少SLA OR AlarmCount满足
	AlarmCount   int     `yaml:"alarmCount"`   // 必填 警告数量达到多少告警 OR SLA不满足 默认1000000
	MinAlarm     int     `yaml:"minAlarm"`     // 必须满足最小的告警数据
	// MinRequests      int      `yaml:"minRequests"`      // 必填 需要最少多少次请求才能告警
	ExtendReceiver   []string      `yaml:"extendReceiver"`   // 可选 扩展指定接收者
	OnlyExtend       bool          `yaml:"onlyExtend"`       // 可选 只发送给指定接收者
	Interval         int64         `yaml:"interval"`         // 可选 告警检查间隔 默认60s
	Silences         []string      `yaml:"silences"`         // 可选 静音规则
	AlarmRule        string        `yaml:"alarmRule"`        // 必填 告警规则
	SpecialAlarmRule []SpecialRule `yaml:"specialAlarmRule"` // 可选
	RuleType         string        `yaml:"ruleType"`         // 必填 监控类型 slow or errorCode
	ContinuousTimes  int           `yaml:"continuousTimes"`  // 持续几次才告警
}

type SpecialRule struct {
	Filter    string `yaml:"filter"`
	AlarmRule string `yaml:"alarmRule"`
}

type AlertReceiver struct {
	KafkaAddr string `yaml:"kafkaAddr"`
	Topic     string `yaml:"topic"`
}
