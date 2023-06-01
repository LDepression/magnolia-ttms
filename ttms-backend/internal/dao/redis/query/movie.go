/**
 * @Author: lenovo
 * @Description:
 * @File:  movie
 * @Version: 1.0.0
 * @Date: 2023/05/31 16:55
 */

package query

import (
	"context"
	"mognolia/internal/pkg/utils"
)

// 使用zset作为排行榜
const (
	MovieReadCountRank = "movieReadCountRank"
	MovieReadCountKey  = "movieReadCountKey"
	MAXNUM             = "99999999"
)

func (q *Queries) AddReadCountToMovie(ctx context.Context, movieID uint) error {
	return q.redis.ZIncrBy(ctx, MovieReadCountRank, 1, utils.LinkStr(MovieReadCountKey, utils.IDToSting(movieID))).Err()
}

// 将某电影的访问量清零
func (q *Queries) FlushDataByMovieID(ctx context.Context, movieID uint) error {
	return q.redis.ZRem(ctx, MovieReadCountRank, utils.LinkStr(MovieReadCountKey, utils.IDToSting(movieID))).Err()
}

// 将全部电影的访问量清零
func (q *Queries) FlushAllData(ctx context.Context) error {
	members := q.redis.ZRange(ctx, MovieReadCountRank, 0, -1).Val()
	for _, member := range members {
		if err := q.redis.ZRem(ctx, MovieReadCountRank, member).Err(); err != nil {
			return err
		}
	}
	return nil
}
