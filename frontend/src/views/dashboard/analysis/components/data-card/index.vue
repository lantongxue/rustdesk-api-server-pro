<template>
  <n-card :bordered="false" class="h-full rounded-8px shadow-sm">
    <n-grid cols="s:1 m:2 l:4" responsive="screen" :x-gap="16" :y-gap="16">
      <n-grid-item>
        <gradient-bg class="h-100px" start-color="#ec4786" end-color="#b955a4">
          <h3 class="text-16px">{{ $t('page.dashboard.userCount') }}</h3>
          <div class="flex justify-between pt-12px">
            <svg-icon icon="gravity-ui:person" class="text-32px" />
            <count-to :start-value="0" :end-value="stat?.userCount" class="text-30px text-white dark:text-dark" />
          </div>
        </gradient-bg>
      </n-grid-item>
      <n-grid-item>
        <gradient-bg class="h-100px" start-color="#865ec0" end-color="#5144b4">
          <h3 class="text-16px">{{ $t('page.dashboard.peerCount') }}</h3>
          <div class="flex justify-between pt-12px">
            <svg-icon icon="fluent:desktop-32-regular" class="text-32px" />
            <count-to :start-value="0" :end-value="stat?.peerCount" class="text-30px text-white dark:text-dark" />
          </div>
        </gradient-bg>
      </n-grid-item>
      <n-grid-item>
        <gradient-bg class="h-100px" start-color="#56cdf3" end-color="#719de3">
          <h3 class="text-16px">{{ $t('page.dashboard.onlineCount') }}</h3>
          <div class="flex justify-between pt-12px">
            <svg-icon icon="fluent:desktop-checkmark-20-regular" class="text-32px" />
            <count-to :start-value="0" :end-value="stat?.onlineCount" class="text-30px text-white dark:text-dark" />
          </div>
        </gradient-bg>
      </n-grid-item>
      <n-grid-item>
        <gradient-bg class="h-100px" start-color="#fcbc25" end-color="#f68057">
          <h3 class="text-16px">{{ $t('page.dashboard.visitsCount') }}</h3>
          <div class="flex justify-between pt-12px">
            <svg-icon icon="ant-design:bar-chart-outlined" class="text-32px" />
            <count-to :start-value="0" :end-value="stat?.visitsCount" class="text-30px text-white dark:text-dark" />
          </div>
        </gradient-bg>
      </n-grid-item>
    </n-grid>
  </n-card>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { fetchStat } from '@/service';
import { $t } from '@/locales';
import { GradientBg } from './components';

defineOptions({ name: 'DashboardAnalysisDataCard' });

const stat = ref<ApiDashboard.Stat>({
  userCount: 110,
  peerCount: 0,
  onlineCount: 0,
  visitsCount: 0
});

onMounted(async () => {
  const d = (await fetchStat()).data;
  if (d !== null) {
    stat.value = d;
  }
});
</script>

<style scoped></style>
