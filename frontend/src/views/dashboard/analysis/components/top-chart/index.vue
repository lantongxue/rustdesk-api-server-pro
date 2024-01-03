<template>
  <n-grid :x-gap="16" :y-gap="16" :item-responsive="true">
    <n-grid-item span="0:24 640:24 1024:6">
      <n-card :bordered="false" class="rounded-8px shadow-sm">
        <div class="w-full h-360px py-12px">
          <h3 class="text-16px font-bold">{{ $t('page.dashboard.friendlySponsorship') }}</h3>
          <p class="text-#aaa">{{ $t('page.dashboard.cupOfCoffee') }}</p>
          <n-image-group>
            <n-space>
              <n-image :src="alipay" width="150" />
              <n-image :src="wechat" width="150" />
            </n-space>
          </n-image-group>
          <p class="text-center text-24px font-bold">{{ $t('page.dashboard.thankYou') }}</p>
        </div>
      </n-card>
    </n-grid-item>
    <n-grid-item span="0:24 640:24 1024:10">
      <n-card :bordered="false" class="rounded-8px shadow-sm">
        <div ref="lineRef" class="w-full h-360px"></div>
      </n-card>
    </n-grid-item>
    <n-grid-item span="0:24 640:24 1024:8">
      <n-card :bordered="false" class="rounded-8px shadow-sm">
        <div ref="pieRef" class="w-full h-360px"></div>
      </n-card>
    </n-grid-item>
  </n-grid>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import type { Ref } from 'vue';
import { type ECOption, useEcharts } from '@/composables';
import alipay from '@/assets/images/donet/alipay.jpeg';
import wechat from '@/assets/images/donet/wechat.jpeg';
import { $t } from '../../../../../locales';

defineOptions({ name: 'DashboardAnalysisTopCard' });

const lineOptions = ref<ECOption>({
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
    data: ['用户数', '主机数']
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
      data: ['06:00', '08:00', '10:00', '12:00', '14:00', '16:00', '18:00', '20:00', '22:00', '24:00']
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
      name: '用户数',
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
      data: [4623, 6145, 6268, 6411, 1890, 4251, 2978, 3880, 3606, 4311]
    },
    {
      color: '#26deca',
      name: '主机数',
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
      data: [2208, 2016, 2916, 4512, 8281, 2008, 1963, 2367, 2956, 678]
    }
  ]
}) as Ref<ECOption>;
const { domRef: lineRef } = useEcharts(lineOptions);

const pieOptions = ref<ECOption>({
  title: {
    text: '系统分布',
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
      name: '系统分布',
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
      data: [
        { value: 20, name: 'Windows' },
        { value: 10, name: 'macOS' },
        { value: 30, name: 'Linux' },
        { value: 40, name: 'Android' },
        { value: 40, name: 'iOS' }
      ]
    }
  ]
}) as Ref<ECOption>;
const { domRef: pieRef } = useEcharts(pieOptions);
</script>

<style scoped></style>
