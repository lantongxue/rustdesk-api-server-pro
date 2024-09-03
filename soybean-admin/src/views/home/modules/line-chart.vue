<script setup lang="ts">
import { $t } from '@/locales';
import { useEcharts } from '@/hooks/common/echarts';
import { fetchLineCharts } from '@/service/api/home';

defineOptions({
  name: 'LineChart'
});

const { domRef, updateOptions } = useEcharts(() => ({
  title: {
    text: $t('page.home.oneWeek')
  },
  tooltip: {
    trigger: 'axis',
    axisPointer: {
      type: 'cross',
      label: {
        backgroundColor: '#6a7985'
      }
    }
  },
  legend: {
    data: [$t('page.home.userCount'), $t('page.home.deviceCount')]
  },
  grid: {
    left: '3%',
    right: '4%',
    bottom: '3%',
    containLabel: true
  },
  xAxis: [
    {
      type: 'category',
      boundaryGap: false,
      data: []
    }
  ],
  yAxis: [
    {
      type: 'value'
    }
  ],
  series: [
    {
      color: '#8e9dff',
      smooth: true,
      stack: 'Total',
      areaStyle: {
        color: {
          type: 'linear',
          x: 0,
          y: 0,
          x2: 0,
          y2: 1,
          colorStops: [
            {
              offset: 0.25,
              color: '#8e9dff'
            },
            {
              offset: 1,
              color: '#fff'
            }
          ]
        }
      },
      emphasis: {
        focus: 'series'
      },
      name: $t('page.home.userCount'),
      type: 'line',
      data: []
    },
    {
      name: $t('page.home.deviceCount'),
      type: 'line',
      data: []
    }
  ]
}));

async function fetchChartsData() {
  const line = await fetchLineCharts();
  updateOptions(opt => {
    opt.xAxis = [
      {
        type: 'category',
        boundaryGap: false,
        data: line.data?.xAxis as []
      }
    ];
    opt.series = [
      {
        color: '#8e9dff',
        name: $t('page.home.userCount'),
        type: 'line',
        smooth: true,
        stack: 'Total',
        areaStyle: {
          color: {
            type: 'linear',
            x: 0,
            y: 0,
            x2: 0,
            y2: 1,
            colorStops: [
              {
                offset: 0.25,
                color: '#8e9dff'
              },
              {
                offset: 1,
                color: '#fff'
              }
            ]
          }
        },
        emphasis: {
          focus: 'series'
        },
        data: line.data?.users as []
      },
      {
        color: '#26deca',
        name: $t('page.home.deviceCount'),
        type: 'line',
        smooth: true,
        stack: 'Total',
        areaStyle: {
          color: {
            type: 'linear',
            x: 0,
            y: 0,
            x2: 0,
            y2: 1,
            colorStops: [
              {
                offset: 0.25,
                color: '#26deca'
              },
              {
                offset: 1,
                color: '#fff'
              }
            ]
          }
        },
        emphasis: {
          focus: 'series'
        },
        data: line.data?.peer as []
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
