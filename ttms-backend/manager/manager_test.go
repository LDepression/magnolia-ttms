/**
 * @Author: lenovo
 * @Description:
 * @File:  manager_test
 * @Version: 1.0.0
 * @Date: 2023/06/06 22:39
 */

package manager

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"mognolia/internal/pkg/utils"
	"sync"
	"testing"
	"time"
)

func TestUtils(t *testing.T) {
	var (
		allNum   int64
		falseNum int64
		okNum    int64
	)

	n := int(utils.RandomInt(1, 1))
	fmt.Println(n)
	testPlans := make([]int64, n)
	for i := range testPlans {
		tickets := utils.RandomInt(1, 20)
		testPlans[i] = tickets
		allNum += tickets
	}
	m := Manager()
	for i := range testPlans {
		m.Set(uint(i), time.Hour)
	}
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		j := i
		go func() {
			defer wg.Done()
			userID := j
			planID := utils.RandomInt(0, int64(len(testPlans)-1))
			seatID := utils.RandomInt(1, testPlans[planID])

			ok := m.PlanMgr[uint(planID)].LockTicket(uint(userID), uint(seatID))
			if ok {
				okNum++
			} else {
				falseNum++
			}
		}()
	}
	wg.Wait()
	assert.Equal(t, okNum, allNum)
}
