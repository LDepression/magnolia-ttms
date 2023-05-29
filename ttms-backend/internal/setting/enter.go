/**
 * @Author: lenovo
 * @Description:
 * @File:  enter
 * @Version: 1.0.0
 * @Date: 2023/05/29 8:11
 */

package setting

type group struct {
	Log    log
	Va     va
	Dao    mdao
	Worker worker
}

var Group = new(group)

func AllInit() {
	Group.Log.Init()
	Group.Va.InitTrans("zh")
	Group.Dao.Init()
	Group.Worker.Init()
}
