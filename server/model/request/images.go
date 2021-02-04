package request

import "online_file/model"

type ImagesSearch struct {
	model.Images
	PageInfo
}
