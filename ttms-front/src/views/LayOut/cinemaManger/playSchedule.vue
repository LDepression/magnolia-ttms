<template>
  <div>
    <div class="detail">
          <!-- <h3 class="labelname">{{$route.query.movieInfo.name}}</h3> -->
          <hr/>
        <div>
          <div class="viewing">
            <span>观影时间：</span>
            <!-- <el-button
          size="small "
          class="btn2"
          type="primary" 
          @click="delSeat(row.id)">增加演出</el-button> -->
            <div @click="handletime" :class="{time:true,isSelected:currentTime==timelabel[0]}">{{timelabel[0]}}</div>
            <div @click="handletime" :class="{time:true,isSelected:currentTime==timelabel[1]}">{{timelabel[1]}}</div>
            <div @click="handletime" :class="{time:true,isSelected:currentTime==timelabel[2]}">{{timelabel[2]}}</div>
          </div>
        </div>
        <br>
        <br>
        <el-table
        :data="tableData.data"
        stripe
        style="width: 100%">
        <el-table-column
          prop="movieName"
          label="电影名字"
          width="180">
        </el-table-column>
        <el-table-column
          prop="startAt"
          label="放映时间"
          width="180">
        </el-table-column>
        <el-table-column
          prop="duration"
          label="放映时常"
          width="180">
        </el-table-column>
        <el-table-column
          prop="cinemaName"
          label="放映厅">
        </el-table-column>
        <el-table-column
          prop="price"
          label="售价（元）">
        </el-table-column>
        <el-table-column
          label="选座购票"
          prop="id">
          <template v-slot="{row}">
         
          <el-button
          size="small "
          class="btn"
          type="primary" round
          @click="delSeat(row.id)">删除演出</el-button>
         
        </template>
        </el-table-column>
        
      </el-table>
    </div>
  </div>
</template>

<script>
import { reactive, onMounted,ref } from 'vue';
// import request from '../api/index';
import { onBeforeRouteLeave } from 'vue-router'
import { getPlan,deletePlan} from '../../../apis/plan'
import Cookies from 'js-cookie'
import {useRouter} from 'vue-router'
export default {
  setup(props) {
    const tableData = reactive({
      data:[]
    });
    const movieInfo = reactive([])
    const router = useRouter()
    const currentTime = ref('今天');
    const timelabel = reactive(['今天','明天','后天'])
    const period = reactive({
      start:'',
      end:''
    })
    const cinemaID = JSON.parse(localStorage.getItem("cinemaInfo")).ID
    const getPlanList = async (page = 1) => {
      console.log(11);
      let res = await getPlan(page)
      console.log(res.data.data);
      if(res.data.status_code == 0){
        tableData.data.splice(0,tableData.data.length,...res.data.data.Data)
      }
      console.log(res);
    }
    const handletime = async (e,page = 1) => {
      console.log(e.target.innerHTML);
      currentTime.value = e.target.innerHTML;
      switch (currentTime) {
        case '今天':
          period.start = new Date().setHours(0, 0, 0, 0);
          period.end = new Date().setHours(0, 0, 0, 0) + 24 * 60 * 60 * 1000;
          break;
        case '明天':
          period.start = new Date().setHours(0, 0, 0, 0) + 24 * 60 * 60 * 1000;
          period.end = new Date().setHours(0, 0, 0, 0) + 2 * 24 * 60 * 60 * 1000;
          break;
        case '后天':
          period.start = new Date().setHours(0, 0, 0, 0) + 2 * 24 * 60 * 60 * 1000;
          period.end = new Date().setHours(0, 0, 0, 0) + 3 * 24 * 60 * 60 * 1000;
          break;
      }
      let res = await getPlan(page = 1)
      console.log(res.data.data);
      if(res.data.status_code == 0){
        tableData.data.splice(0,tableData.data.length,...res.data.data.Data)
      }
    };
    const delSeat = async (x) => {
    
      console.log(x);
      let res = await deletePlan(x)
      getPlanList()
    }
    const chooseSeat = (x) => {
      for(let i=0;i<tableData.data.length;i++){
        
        if(tableData.data[i].id == x){
          console.log(tableData.data[0].id);
          movieInfo.splice(0,movieInfo.length,tableData.data[i])
        }
      }
      console.log(movieInfo);
      router.push({
        name:'chooseSeat',
        query:{planID:x,movieInfo:JSON.stringify(movieInfo)}
      })
      // let token = Cookies.get('AccessToken');
      // if (token) {
      //   router.push({
      //     name: 'chooseSeat',
      //     query: { ticket: JSON.stringify(state.tableData[x]), movieInfo: JSON.stringify(props.route.query.movieInfo) },
      //   });
      // } else {
      //   router.push({ path: '/login' });
      // }
    };

    onMounted(() => {
      getPlanList(1)
      period.start = new Date().setHours(0, 0, 0, 0);
      period.end = new Date().setHours(0, 0, 0, 0) + 24 * 60 * 60 * 1000;
    
      // request
      //   .getPlanlist({
      //     movie_id: props.route.query.movie_id,
      //     start_time: new Date(state.period.start).toISOString().slice(0, 10),
      //     end_time: new Date(state.period.end).toISOString().slice(0, 10),
      //   })
      //   .then(
      //     (res) => {
      //       for (let i of res.data.data.list) {
      //         i.start_at = i.start_at.slice(11, 16);
      //       }
      //       state.tableData = res.data.data.list;
      //     },
      //     (err) => {
      //       console.log(err);
      //     }
      //   );
    });

    onBeforeRouteLeave((to, from, next) => {
      let token = Cookies.get('AccessToken');
      if (to.path == '/chooseSeat') {
        if (token) {
          next();
        } else {
          router.push({ path: '/login' });
        }
      }
      next();
    });
    return {tableData, 
   currentTime ,
    timelabel ,
     period, handletime, chooseSeat,getPlanList,delSeat };
  },
};
</script>

<style scoped>
.time {
  border-radius:14px;
  padding: 3px 9px;
  display: inline-block;
  margin-left: 12px;
  cursor: pointer;
}
.btn{
  position: relative;
  left: -12px;
}
.isSelected{
  background-color: #75c4ff;
  color: #fff;
}
.btn2{
  height: 40px;
}
.labelname::before{
    content: "";
    display: inline-block;
    width: 4px;
    height: 18px;
    margin-right: 6px;
    background-color: #11a8cd;
}
.detail{
    width:1000px;
    margin:50px auto;
    padding:10px;
}
</style>