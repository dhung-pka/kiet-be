package service

import (
	"service/config"
	"service/model"
	"service/repo"
	"service/utils"

	"gorm.io/gorm"
)

type basicQueryService struct {
	db             *gorm.DB
	basicQueryRepo repo.BasicQueryRepo
}

type BasicQueryService interface {
	BasicQuery(req model.BasicQueryRequest) (interface{}, error)
}

func (b *basicQueryService) BasicQuery(req model.BasicQueryRequest) (interface{}, error) {
	dataQuery := req.Data.(map[string]interface{})
	data := model.TypeModel[req.ModelType]

	utils.MapInterface(dataQuery, &data)

	switch option := req.Option; option {
	case model.INSERT:
		goto INSERT
	}

INSERT:
	res, err := b.basicQueryRepo.BasicQueryInsertData(data, req.ModelType)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewBasicQueryService() basicQueryService {
	db := config.GetDB()
	basicQueryRepo := repo.NewBasicQueryRepo()
	return basicQueryService{
		db:             db,
		basicQueryRepo: basicQueryRepo,
	}
}
