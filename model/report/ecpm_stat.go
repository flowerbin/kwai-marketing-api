package report

// EcpmStat 快小游ECPM数据
type EcpmStat struct {
	// EventTime 广告计费事件发生时间，格式：YYYY-MM-DD HH:mm:ss
	EventTime string `json:"event_time,omitempty"`
	// ID 记录唯一标识（返回记录按 ID 递增排序）
	ID string `json:"id,omitempty"`
	// OpenID 用户的 open_id或union_id
	OpenID string `json:"open_id,omitempty"`
	// 消耗：单位厘
	Cost int64 `json:"cost,omitempty"`
}
