package article

import (
	"goblog/pkg/logger"
	"goblog/pkg/types"
	"goblog/tests/app/models"
)

// GetAll 获取全部文章
func GetAll() ([]Article, error) {
	var articles []Article
	if err := models.DB.Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}

func Get(idstr string) (Article, error) {
	var article Article
	id := types.StringToInt(idstr)
	if err := models.DB.First(&article, id).Error; err != nil {
		return article, err
	}
	return article, nil
}

// Create 创建文章，通过 article.ID 来判断是否创建成功
func (article *Article) Create() (err error) {
	result := models.DB.Create(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return err
	}

	return nil
}

// Update 更新文章
func (article *Article) Update() (rowsAffected int64, err error) {
	result := models.DB.Save(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}

	return result.RowsAffected, nil
}

// Delete 删除文章
func (article *Article) Delete() (rowsAffected int64, err error) {
	result := models.DB.Delete(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}

	return result.RowsAffected, nil
}
