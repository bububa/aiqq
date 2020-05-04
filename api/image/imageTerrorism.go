package image

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/url"

	"github.com/bububa/aiqq"
)

type ImageTerrorismRequest struct {
	Image    string `json:"image,omitempty"`     // 原始图片的base64编码数据（原图大小上限1MB，支持JPG、PNG、BMP格式），image和image_url必须至少提供一个
	ImageUrl string `json:"image_url,omitempty"` // 如果image和image_url都提供，仅支持image_url，image和image_url必须至少提供一个
}

func (this *ImageTerrorismRequest) Path() string {
	return "image/image_terrorism"
}

func (this *ImageTerrorismRequest) Method() string {
	return "post"
}

func (this *ImageTerrorismRequest) Values() url.Values {
	values := url.Values{}
	if this.Image != "" {
		values.Set("image", this.Image)
	} else if this.ImageUrl != "" {
		values.Set("image_url", this.ImageUrl)
	}
	return values
}

// 参考判断标准如下
// 1. 色情图片：porn值 > 83；
// 2. 性感图片：hot值 > normal值；
// 3. 其他情况认为是正常图片。
type ImageTerrorismResponse struct {
	Tags []ImageTerrorismTag `json:"tag_list,omitempty"` // 图像的分类标签
}

type ImageTerrorismTag struct {
	Name        ImageTerrorismName `json:"tag_name,omitempty"`         // 返回图像标签的名字
	Confidence  int                `json:"tag_confidence,omitempty"`   // 图像标签的置信度,范围0-100,越大置信度越高
	ConfidenceF float64            `json:"tag_confidence_f,omitempty"` // 图像标签的置信度,浮点数范围0-1,越大置信度越高
}

func ImageTerrorismFromImageUrl(clt *aiqq.Client, imageUrl string) ([]ImageTerrorismTag, error) {
	resp, err := clt.Do(&ImageTerrorismRequest{ImageUrl: imageUrl})
	if err != nil {
		return nil, err
	}
	var ret ImageTerrorismResponse
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return ret.Tags, nil
}

func ImageTerrorismFromImageData(clt *aiqq.Client, data []byte) ([]ImageTerrorismTag, error) {
	imageData := base64.URLEncoding.EncodeToString(data)
	resp, err := clt.Do(&ImageTerrorismRequest{Image: imageData})
	if err != nil {
		return nil, err
	}
	var ret ImageTerrorismResponse
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return ret.Tags, nil
}

func ImageTerrorismFromReader(clt *aiqq.Client, r io.Reader) ([]ImageTerrorismTag, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	imageData := base64.URLEncoding.EncodeToString(data)
	resp, err := clt.Do(&ImageTerrorismRequest{Image: imageData})
	if err != nil {
		return nil, err
	}
	var ret ImageTerrorismResponse
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return ret.Tags, nil
}
