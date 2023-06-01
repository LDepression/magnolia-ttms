/**
 * @Author: lenovo
 * @Description:
 * @File:  movie_test
 * @Version: 1.0.0
 * @Date: 2023/05/31 17:26
 */

package query

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueries_AddReadCountToMovie(t *testing.T) {
	err := client.AddReadCountToMovie(context.Background(), 1)
	assert.NoError(t, err)
}

func TestFlushAllData(t *testing.T) {
	err := client.FlushAllData(context.Background())
	assert.NoError(t, err)
}

func TestQueries_FlushDataByMovieID(t *testing.T) {
	err := client.FlushDataByMovieID(context.Background(), 1)
	assert.NoError(t, err)
}
