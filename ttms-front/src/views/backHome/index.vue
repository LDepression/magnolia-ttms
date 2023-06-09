<template>
  <div class="M">
    <div class="top">
      <h2>mogalian电影院电影播放量TOP8</h2>
    </div>
    <div ref="chart" class="echarts" style="width:700px;height:600px"></div>
  </div>
</template>

<script>
import { onMounted, ref } from 'vue';
import * as echarts from 'echarts';
// import ECharts from 'vue-echarts';
export default {
  setup() {
    const chart = ref(null);

    onMounted(() => {
        const myChart = echarts.init(chart.value);
        const option = {
          legend: {
            top: 'bottom'
          },
          toolbox: {
            show: true,
            feature: {
              mark: { show: true },
              dataView: { show: true, readOnly: false },
              restore: { show: true },
              saveAsImage: { show: true }
            }
          },
          series: [
            {
              name: '',
              type: 'pie',
              radius: [50, 250],
              center: ['50%', '50%'],
              roseType: 'area',
              itemStyle: {
                borderRadius: 10
              },
              data: [
                { value: 40, name: '我不是药神' },
                { value: 38, name: '肖申克的救赎' },
                { value: 32, name: '海上钢琴师' },
                { value: 32, name: '绿皮书' },
                { value: 30, name: '霸王别姬' },
                { value: 26, name: '美丽人生' },
                { value: 22, name: '这个杀手不太冷' },
                { value: 18, name: '哪吒之魔童降世' },
              ]
            }
          ]
        };
        myChart.setOption(option);
    });

    return {
      chart,
    };
  },
};
</script>

<style>
.echarts{
  position: relative;
  left: 250px;
  top: -50px;
}
  .top{
    /* margin-left: 50px; */
    text-align: center;
    position: relative;
    left: 270px;
    color: black;
}
  .top h2 {
    padding-bottom: 20px;
  }
  #main {
    float: left;
    margin-top: 50px;
    margin-bottom: 50px;
    margin-left: 200px;
  }
</style>
