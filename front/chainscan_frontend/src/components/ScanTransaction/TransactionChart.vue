<template lang="">
  <div>
    <p class="chart-title">TRANSACTION HISTORY IN 14 DAYS</p>
    <div id="char" style="width: 700px; height: 200px"></div>
  </div>
</template>
<script lang="ts" setup>
import { onMounted } from 'vue';
import { GetTxTotal } from '../../script/service/transactionService';
import moment from 'moment';
import { ECharts, EChartsOption, init } from 'echarts';

const dataList: string[] = [];
const totalList: number[] = [];

onMounted(async () => {
  const res = await GetTxTotal(moment().subtract(14, 'days').format('YYYYMMDD'), moment().format('YYYYMMDD'));
  res.data.forEach((element) => {
    dataList.push(element.date.slice(0, 10));
    totalList.push(element.txCount);
  });
  // console.log('GetTxTotal', res.data.data);
  const charEle = document.getElementById('char') as HTMLElement;
  const charEch: ECharts = init(charEle);
  const option: EChartsOption = {
    tooltip: {
      trigger: 'axis',
    },
    xAxis: {
      type: 'category',
      data: dataList,
      // show:false,
      axisTick: {
        show: false,
      },
    },
    yAxis: {
      type: 'value',
      interval: 100,
      splitLine: {
        show: false,
      },
    },
    series: [
      {
        data: totalList,
        type: 'line',
        smooth: true,
      },
    ],
  };
  charEch.setOption(option);
  document.getElementById('char')!.setAttribute('_echarts_instance_', '');
});
</script>
<style lang="less" scoped>
.chart-title {
  display: flex;
  justify-content: center;
  font-size: 3.76562rem;
  color: #77838f;
}
</style>
