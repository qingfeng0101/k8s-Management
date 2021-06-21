<template>
    <div>
     <div id="myChart" style="width: 600px;height:400px;float:right;;"></div>
     <div id="myChart1"  style="width: 600px;height:400px;float:right;;"></div>
     
    </div>
   
</template>
<script>
export default {
    name: 'hello',
  data () {
    return {
      msg: 'Welcome to Your Vue.js App'
    }
  },
  mounted(){
    this.drawLine();
  },
  methods: {
    drawLine(){
        // 基于准备好的dom，初始化echarts实例
        var echarts = require('echarts');
        let myChart =echarts.init(document.getElementById('myChart'))
         let myChart1 =echarts.init(document.getElementById('myChart1'))
        // 绘制图表
        myChart.setOption({
            title: { text: '在Vue中使用echarts' },
            tooltip: {},
            xAxis: {
                data: ["衬衫","羊毛衫","雪纺衫","裤子","高跟鞋","袜子"]
            },
            yAxis: {},
            series: [{
                name: '销量',
                type: 'bar',
                data: [5, 20, 36, 10, 10, 20]
            }]
        });
        myChart1.setOption({
    series: [{
        type: 'gauge',
        axisLine: {
            lineStyle: {
                width: 30,
                color: [
                    [0.3, '#67e0e3'],
                    [0.7, '#37a2da'],
                    [1, '#fd666d']
                ]
            }
        },
        pointer: {
            itemStyle: {
                color: 'auto'
            }
        },
        axisTick: {
            distance: -30,
            length: 8,
            lineStyle: {
                color: '#fff',
                width: 2
            }
        },
        splitLine: {
            distance: -30,
            length: 30,
            lineStyle: {
                color: '#fff',
                width: 4
            }
        },
        axisLabel: {
            color: 'auto',
            distance: 40,
            fontSize: 20
        },
        detail: {
            valueAnimation: true,
            formatter: '{value} km/h',
            color: 'auto'
        },
        data: [{
            value: 70
        }]
    }]
}
)
    }
  }
}
</script>
<style lang="scss" scoped>
.div-inline{ 
    display:inline
    } 
</style>