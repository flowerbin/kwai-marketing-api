package file

import (
	"github.com/bububa/kwai-marketing-api/core"
	"github.com/bububa/kwai-marketing-api/model/file"
)

// UploadAtlasPhoto 上传图文视频
func UploadAtlasPhoto(clt *core.SDKClient, accessToken string, req *file.UploadAtlasPhotoRequest) (string, error) {
	var ret file.UploadAtlasPhotoResponse
	if err := clt.Post(accessToken, req, &ret); err != nil {
		return "", err
	}
	return ret.PhotoID, nil
}
