package file

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/file"
)

// PicShare 图片库推送
func PicShare(clt *core.SDKClient, accessToken string, req *file.PicShareRequest) ([]file.PicShareResult, error) {
	var resp file.PicShareResponse
	if err := clt.Post(accessToken, req, &resp); err != nil {
		return nil, err
	}
	return resp.Details, nil
}
