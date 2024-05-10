<template>
   <div id="main" style="width: 600px; height: 400px;"></div>
</template>

<script>
   import * as echarts from 'echarts';

   export default {
      data() {
         return {
            myChart: null,
            option: {
               title: {
                  text: '房产交易图',
                  subtext: '2024.5.10',
                  left: 'center'
               },
               tooltip: {
                  trigger: 'item'
               },
               legend: {
                  orient: 'vertical',
                  left: 'left'
               },
               series: [
                  {
                     name: 'Access From',
                     type: 'pie',
                     radius: '50%',
                     data: [
                        { value: 1, name: '已出售' },
                        { value: 2, name: '未出售' },
                     ],
                     emphasis: {
                        itemStyle: {
                           shadowBlur: 10,
                           shadowOffsetX: 0,
                           shadowColor: 'rgba(0, 0, 0, 0.5)'
                        }
                     }
                  }
               ]
            }
         };
      },
      mounted() {
         this.myChart = echarts.init(document.getElementById('main'));
         this.myChart.setOption(this.option);
      }
   };
</script>

<style scoped>
   /* 根据需要添加样式 */
</style>
