package ocr

import (
	"fmt"
	"net/url"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// imgAiCropURL 图片智能裁剪
	imgAiCropURL = "https://api.weixin.qq.com/cv/img/aicrop"
	// imgQrcodeURL 二维码/条码识别
	imgQrcodeURL = "https://api.weixin.qq.com/cv/img/qrcode"
	// imgSuperResolutionURL 图片高清化
	imgSuperResolutionURL = "https://api.weixin.qq.com/cv/img/superresolution"
)

// ImgAiCropRequest 图片智能裁剪请求
type ImgAiCropRequest struct {
	ImgURL string `json:"img_url"` // 图片链接
	Img    string `json:"img"`     // 图片文件，小于2M
	Ratios string `json:"ratios"`  // 宽高比；如果提供多个宽高比，请以英文逗号","分隔，最多支持5个宽高比
}

// ImgAiCropResponse 图片智能裁剪响应
type ImgAiCropResponse struct {
	util.CommonError
	Results []AiCropResult `json:"results"`  // 智能裁剪结果
	ImgSize imageSize      `json:"img_size"` // 图片大小
}

// AiCropResult 智能裁剪结果
type AiCropResult struct {
	CropLeft   int64 `json:"crop_left"`   // 左上角x
	CropTop    int64 `json:"crop_top"`    // 左上角y
	CropRight  int64 `json:"crop_right"`  // 右下角x
	CropBottom int64 `json:"crop_bottom"` // 右下角y
}

// ImgAiCrop 图片智能裁剪
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/img-ocr/img/api_imgaicrop
func (ocr *OCR) ImgAiCrop(req *ImgAiCropRequest) (res ImgAiCropResponse, err error) {
	accessToken, err := ocr.GetAccessToken()
	if err != nil {
		return
	}

	params := map[string]interface{}{
		"access_token": accessToken,
	}
	if req.ImgURL != "" {
		params["img_url"] = url.QueryEscape(req.ImgURL)
	}
	query := util.Query(params)

	fields := []util.MultipartFormField{}
	if req.Img != "" {
		fields = append(fields, util.MultipartFormField{
			IsFile:    true,
			Fieldname: "img",
			FilePath:  req.Img,
			Filename:  req.Img,
		})
	}
	if req.Ratios != "" {
		fields = append(fields, util.MultipartFormField{
			IsFile:    false,
			Fieldname: "ratios",
			Filename:  "ratios",
			Value:     []byte(req.Ratios),
		})
	}

	var response []byte
	if response, err = util.PostMultipartForm(fields, fmt.Sprintf("%s?%s", imgAiCropURL, query)); err != nil {
		return
	}

	err = util.DecodeWithError(response, &res, "ImgAiCrop")
	return
}

// ImgQrcodeRequest 二维码/条码识别请求
type ImgQrcodeRequest struct {
	ImgURL string `json:"img_url"` // 图片链接
	Img    string `json:"img"`     // 图片文件，限制小于 2 M
}

// ImgQrcodeResponse 二维码/条码识别响应
type ImgQrcodeResponse struct {
	util.CommonError
	CodeResults []QrcodeResult `json:"code_results"` // 处理结果
	ImgSize     imageSize      `json:"img_size"`     // 图片大小
}

// QrcodeResult 码的识别结果
type QrcodeResult struct {
	TypeName string   `json:"type_name"` // 码的类型
	Data     string   `json:"data"`      // 码的信息
	Pos      position `json:"pos"`       // 码的坐标
}

// ImgQrcode 二维码/条码识别
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/img-ocr/img/api_imgqrcode
func (ocr *OCR) ImgQrcode(req *ImgQrcodeRequest) (res ImgQrcodeResponse, err error) {
	accessToken, err := ocr.GetAccessToken()
	if err != nil {
		return
	}

	params := map[string]interface{}{
		"access_token": accessToken,
	}
	if req.ImgURL != "" {
		params["img_url"] = url.QueryEscape(req.ImgURL)
	}
	query := util.Query(params)

	fields := []util.MultipartFormField{}
	if req.Img != "" {
		fields = append(fields, util.MultipartFormField{
			IsFile:    true,
			Fieldname: "img",
			FilePath:  req.Img,
			Filename:  req.Img,
		})
	}

	var response []byte
	if response, err = util.PostMultipartForm(fields, fmt.Sprintf("%s?%s", imgQrcodeURL, query)); err != nil {
		return
	}

	err = util.DecodeWithError(response, &res, "ImgQrcode")
	return
}

// ImgSuperResolutionRequest 图片高清化请求
type ImgSuperResolutionRequest struct {
	ImgURL string `json:"img_url"` // 图片链接
	Img    string `json:"img"`     // 图片文件
}

// ImgSuperResolutionResponse 图片高清化响应
type ImgSuperResolutionResponse struct {
	util.CommonError
	MediaID string `json:"media_id"` // media id
}

// ImgSuperResolution 图片高清化
// see https://developers.weixin.qq.com/miniprogram/dev/server/API/img-ocr/img/api_imgsuperresolution
func (ocr *OCR) ImgSuperResolution(req *ImgSuperResolutionRequest) (res ImgSuperResolutionResponse, err error) {
	accessToken, err := ocr.GetAccessToken()
	if err != nil {
		return
	}

	params := map[string]interface{}{
		"access_token": accessToken,
	}
	if req.ImgURL != "" {
		params["img_url"] = url.QueryEscape(req.ImgURL)
	}
	query := util.Query(params)

	fields := []util.MultipartFormField{}
	if req.Img != "" {
		fields = append(fields, util.MultipartFormField{
			IsFile:    true,
			Fieldname: "img",
			FilePath:  req.Img,
			Filename:  req.Img,
		})
	}

	var response []byte
	if response, err = util.PostMultipartForm(fields, fmt.Sprintf("%s?%s", imgSuperResolutionURL, query)); err != nil {
		return
	}

	err = util.DecodeWithError(response, &res, "ImgSuperResolution")
	return
}
