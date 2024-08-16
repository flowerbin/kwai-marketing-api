package report

// EcpmReportResponse 快小游ECPM报表APIResponse
type EcpmReportResponse struct {
	// TotalCount 数据的总行数
	TotalCount int `json:"total_count,omitempty"`
	// Details 数据明细信息
	Details []EcpmStat `json:"details,omitempty"`
}
