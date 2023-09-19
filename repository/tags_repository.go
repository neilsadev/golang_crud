package repository

import "golang-crud-gin/model"

type TagsRepository interface {
	Save(tag model.Tags)
	Update(tag model.Tags)
	Delete(tagID int)
	FindById(tagID int) (tag model.Tags, err error)
	FindAll() []model.Tags
}
