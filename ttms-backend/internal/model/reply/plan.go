/**
 * @Author: lenovo
 * @Description:
 * @File:  plan
 * @Version: 1.0.0
 * @Date: 2023/06/04 5:17
 */

package reply

import "time"

type CreatePlanRly struct {
	PlanID  uint      `json:"planID"`
	StartAt time.Time `json:"startAt"`
	EndAt   time.Time `json:"endAt"`
}
