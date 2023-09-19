package service

import (
	"golang-crud-gin/repository"
)

type TagsServiceImplementation struct {
	TagsRepoitory repository.TagsRepository
}
