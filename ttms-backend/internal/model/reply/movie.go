/**
 * @Author: lenovo
 * @Description:
 * @File:  movie
 * @Version: 1.0.0
 * @Date: 2023/05/31 8:53
 */

package reply

import "time"

type CreateMovieRly struct {
	MovieID uint `json:"movieID"`
}
type GetMovieRly struct {
	MovieInfo []MovieInfo `json:"movieInfo"`
	Total     int64       `json:"total"`
}

type MovieInfo struct {
	MovieID    uint      `json:"movieID"`
	Name       string    `json:"name"`
	Area       string    `json:"area"`
	Actors     []string  `json:"actors"`
	Content    string    `json:"content"`
	Duration   int64     `json:"duration"`
	ShowTime   time.Time `json:"showTime"`
	Director   string    `json:"director"`
	Score      float32   `json:"score"`
	BoxOffice  float32   `json:"boxOffice"`
	VisitCount int       `json:"visitCount"`
}

type GetMovieDetails struct {
	MovieInfoRow MovieInfo `json:"movieInfo"`   //电影的基本信息
	IsComment    bool      `json:"isComment" `  //用户是否评论过
	IsFavorite   bool      `json:"isFavorite" ` //用户是否兴趣
	Tag          []string  `json:"tag"`         //电影的标签
}
