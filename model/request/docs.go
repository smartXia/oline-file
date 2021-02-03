package request

import "online_file/model"

type DocsSearch struct {
	model.Docs
	PageInfo
}
