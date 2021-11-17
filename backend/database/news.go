package database

import (
	"context"
	"fmt"
	"myNewFeed/internal/log"
	"myNewFeed/model"
	"time"
)

func GetNews(ctx context.Context) ([]*model.News, error) {
	news := make([]*model.News, 0, 1024)
	if err := DB.Order("id desc").Limit(1000).Find(&news).Error; err != nil {
		log.Sugar.Errorw("get news error", "error", err)
		return nil, fmt.Errorf("get news error: %v", err)
	}

	return news, nil
}

func AddNews(ctx context.Context, news ...*model.News) error {
	if len(news) == 0 {
		return nil
	}

	if err := DB.Create(news).Error; err != nil {
		log.Sugar.Errorw("add news error", "error", err)
		return fmt.Errorf("add news error: %v", err)
	}
	return nil
}

func GetLastNewsTime(ctx context.Context, feedID int) (time.Time, error) {
	news := &model.News{}

	if err := DB.Model(&model.News{}).Select("max(publish_time) as publish_time").
		Where("feed_id = ?", feedID).First(&news).Error; err != nil {
		log.Sugar.Errorw("get last news time error", "error", err)
		return time.Time{}, fmt.Errorf("get last news time error: %v", err)
	}

	return news.PublishTime, nil
}
