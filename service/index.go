package service

import (
	"hao/config"
	"hao/dao"
	"hao/models"
	"html/template"
)

func GetAllIndexInfo(page, pageSize int) (*models.HomeResponse, error) {
	//页面上涉及到的所有数据，必须有定义
	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}

	posts, err := dao.GetPostPage(page, pageSize)
	var postMores []models.PostMore
	for _, post := range posts {
		categoryName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[0:100]
		}
		postMore := models.PostMore{
			post.Pid,
			post.Title,
			post.Slug,
			template.HTML(content),
			post.CategoryId,
			categoryName,
			post.UserId,
			userName,
			post.ViewCount,
			post.Type,
			models.DateDay(post.CreateAt),
			models.DateDay(post.UpdateAt),
		}
		postMores = append(postMores, postMore)

	}
	total := dao.CountGetAllPost()
	pagesCount := (total-1)/10 + 1
	var pages []int
	for i := 0; i < pagesCount; i++ {
		pages = append(pages, i+1)
	}
	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		postMores,
		total,
		page,
		pages,
		page != pagesCount,
	}

	return hr, nil
}
