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
import { ref, onMounted } from 'vue';
import type { Ref } from 'vue';
import { fetchPieCharts, fetchLineCharts } from '@/service';
import { type ECOption, useEcharts } from '@/composables';
import alipay from '@/assets/images/donet/alipay.jpeg';
import wechat from '@/assets/images/donet/wechat.jpeg';
import { $t } from '@/locales';

defineOptions({ name: 'DashboardAnalysisTopCard' });

const lineOptions = ref<ECOption>({
  title: {
    text: $t('page.dashboard.oneWeek')
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
    data: [$t('page.dashboard.userCount'), $t('page.dashboard.peerCount')]
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
      name: $t('page.dashboard.userCount'),
      type: 'line',
      data: []
    },
    {
      name: $t('page.dashboard.peerCount'),
      type: 'line',
      data: []
    }
  ]
}) as Ref<ECOption>;
const { domRef: lineRef } = useEcharts(lineOptions);

const pieOptions = ref<ECOption>({
  title: {
    text: $t('page.dashboard.operatingSystem'),
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
  }
}) as Ref<ECOption>;
const { domRef: pieRef } = useEcharts(pieOptions);

async function fetchChartsData() {
  const line = await fetchLineCharts();
  lineOptions.value.xAxis = [
    {
      type: 'category',
      boundaryGap: false,
      data: line.data?.xAxis
    }
  ];
  lineOptions.value.series = [
    {
      color: '#8e9dff',
      name: $t('page.dashboard.userCount'),
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
      data: line.data?.users
    },
    {
      color: '#26deca',
      name: $t('page.dashboard.peerCount'),
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
      data: line.data?.peer
    }
  ];

  const pie = await fetchPieCharts();
  pieOptions.value.series = [
    {
      name: $t('page.dashboard.operatingSystem'),
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
}

onMounted(() => {
  fetchChartsData();
});
</script>

<style scoped></style>
