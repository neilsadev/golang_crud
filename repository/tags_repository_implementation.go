package repository

import (
	"errors"
	"golang-crud-gin/data/request"
	"golang-crud-gin/helper"
	"golang-crud-gin/model"

	"gorm.io/gorm"
)

type TagsRepositoryImplementation struct {
	Db *gorm.DB
}

// Delete implements TagsRepository.
func (t *TagsRepositoryImplementation) Delete(tagID int) {
	var tags model.Tags
	result := t.Db.Where("id = ?", tagID).Delete(&tags)
	helper.ErrorPanic(result.Error)
}

// FindAll implements TagsRepository.
func (t *TagsRepositoryImplementation) FindAll() []model.Tags {
	var tags []model.Tags
	result := t.Db.Find(&tags)
	helper.ErrorPanic(result.Error)
	return tags
}

// FindById implements TagsRepository.
func (t *TagsRepositoryImplementation) FindById(tagID int) (tag model.Tags, err error) {
	var newTag model.Tags
	result := t.Db.Find(&newTag, tagID)
	if result != nil {
		return newTag, nil
	} else {
		return newTag, errors.New("tag is not found")
	}
}

// Save implements TagsRepository.
func (t *TagsRepositoryImplementation) Save(tag model.Tags) {
	result := t.Db.Create(&tag)
	helper.ErrorPanic(result.Error)
}

// Update implements TagsRepository.
func (t *TagsRepositoryImplementation) Update(tag model.Tags) {
	var updateTag = request.UpdateTagsRequest{
		Id:   tag.Id,
		Name: tag.Name,
	}
	result := t.Db.Model(&tag).Updates(updateTag)
	helper.ErrorPanic(result.Error)
}
