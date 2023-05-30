/**
 * @Author: lenovo
 * @Description:
 * @File:  common
 * @Version: 1.0.0
 * @Date: 2023/05/29 10:49
 */

package automigrate

import (
	"database/sql/driver"
	"encoding/json"
)

func (g Roler) Value() (driver.Value, error) {
	return json.Marshal(g)
}

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb

func (g *Roler) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &g)
}

func (g Gend) Value() (driver.Value, error) {
	return json.Marshal(g)
}

// 实现 sql.Scanner 接口，Scan 将 value 扫描至 Jsonb

func (g *Gend) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), &g)
}
