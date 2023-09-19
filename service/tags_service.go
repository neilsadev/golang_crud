package service

import (
	"golang-crud-gin/data/request"
	"golang-crud-gin/data/response"
)

type TagService interface {
	Create(tags request.CreateTagsRequest)
	Update(tags request.UpdateTagsRequest)
	Delete(tagsId int)
	FindById(tagsId int) response.TagsResponse
	FindAll() []response.TagsResponse
}
