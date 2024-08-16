package report

import "encoding/json"

// EcpmReportRequest 快小游ECPM报表APIRequest
type EcpmReportRequest struct {
	// AdvertiserID 广告主ID（注：非账户快手ID），主要用于鉴权无实际含义
	AdvertiserID uint64 `json:"advertiser_id"`
	// AppID 游戏id
	AppID string `json:"app_id"`
	// OpenID 用户的 open_id或union_id，单次允许最多传入200个，传入为空字符串时，查询所有用户
	OpenID []string `json:"open_id,omitempty"`
	// Page 查询的页码，页码从1开始
	Page int `json:"page,omitempty"`
	// PageSize 单页的大小，最大支持500
	PageSize int `json:"page_size,omitempty"`
	// DataHour 时间范围, 根据传参，有两种时间范围可查： - 若传 YYYY-MM-DD，则查天级范围，即 YYYY-MM-DD 00:00:00 ～ YYYY-MM-DD 23:59:59 - 若传 YYYY-MM-DD hh:mm:ss，则查小时级范围，即 YYYY-MM-DD hh:00:00 ～ YYYY-MM-DD hh:59:59 举例：查 13 号：“2023-07-13” 查 13 号 14 点：“2023-07-13 14” 或 “2023-07-13 14:00:00” 或 “2023-07-13 14:13:24” PS：数据源仅保存近 168小时的数据，建议明确时间范围的情况下，指定小时级别只支持1天的查询
	DataHour string `json:"data_hour"`
}

// Url implement PostRequest interface
func (r EcpmReportRequest) Url() string {
	return "gw/dsp/v1/report/ecpm_report"
}

// Encode implement PostRequest interface
func (r EcpmReportRequest) Encode() []byte {
	buf, _ := json.Marshal(r)
	return buf
}
