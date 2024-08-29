<script setup lang="ts">
import { $t } from '@/locales';
import alipay from '@/assets/imgs/alipay.png';
import wxpay from '@/assets/imgs/wechat.png';

defineOptions({
  name: 'CreativityBanner'
});
</script>

<template>
  <NCard :title="$t('page.home.friendlySponsorship')" :bordered="false" size="small" class="h-full card-wrapper">
    <template #header-extra>
      <div>{{ $t('page.home.cupOfCoffee') }}</div>
    </template>
    <template #footer>
      <div>{{ $t('page.home.thankYou') }}</div>
    </template>
    <div class="h-full flex-center">
      <NSpace>
        <NImage width="200" :src="alipay"></NImage>
        <NImage width="200" :src="wxpay"></NImage>
      </NSpace>
    </div>
  </NCard>
</template>

<style scoped></style>
