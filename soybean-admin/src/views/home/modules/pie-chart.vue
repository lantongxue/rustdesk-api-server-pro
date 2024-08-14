<script setup lang="ts">
import { $t } from '@/locales';
import { useEcharts } from '@/hooks/common/echarts';
import { fetchPieCharts } from '@/service/api/home';

defineOptions({
  name: 'PieChart'
});

const { domRef, updateOptions } = useEcharts(() => ({
  title: {
    text: $t('page.home.operatingSystem'),
    left: 'center'
  },
  tooltip: {
    trigger: 'item'
  },
  legend: {
    bottom: '1%',
    left: 'center',
    itemStyle: {
      borderWidth: 0
    }
  },
  series: [
    {
      name: $t('page.home.operatingSystem'),
      type: 'pie',
      radius: ['45%', '75%'],
      avoidLabelOverlap: false,
      itemStyle: {
        borderRadius: 10,
        borderColor: '#fff',
        borderWidth: 1
      },
      emphasis: {
        label: {
          show: true,
          fontSize: '12'
        }
      },
      data: []
    }
  ]
}));
async function fetchChartsData() {
  const pie = await fetchPieCharts();
  updateOptions(opt => {
    opt.series = [
      {
        name: $t('page.home.operatingSystem'),
        type: 'pie',
        radius: ['45%', '75%'],
        avoidLabelOverlap: false,
        itemStyle: {
          borderRadius: 10,
          borderColor: '#fff',
          borderWidth: 1
        },
        emphasis: {
          label: {
            show: true,
            fontSize: '12'
          }
        },
        data: pie.data as []
      }
    ];
    return opt;
  });
}
async function init() {
  fetchChartsData();
}
// init
init();
</script>

<template>
  <NCard :bordered="false" class="card-wrapper">
    <div ref="domRef" class="h-360px overflow-hidden"></div>
  </NCard>
</template>

<style scoped></style>
