package report

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/report"
)

// EcpmReport 快小游ECPM报表
func EcpmReport(clt *core.SDKClient, accessToken string, req *report.EcpmReportRequest) (*report.EcpmReportResponse, error) {
	var resp report.EcpmReportResponse
	err := clt.Post(accessToken, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
