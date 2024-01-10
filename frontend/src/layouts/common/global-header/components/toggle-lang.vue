<template>
  <hover-container class="w-40px h-full" :inverted="theme.header.inverted">
    <n-dropdown :options="options" trigger="click" :value="language" @select="handleSelect">
      <icon-cil:language class="text-18px outline-transparent" />
    </n-dropdown>
  </hover-container>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { useI18n } from 'vue-i18n';
import { useThemeStore } from '@/store';
import { localStg } from '@/utils';
import { useAppVueContext } from '@/context';

const theme = useThemeStore();
const { locale } = useI18n();

const language = ref<I18nType.LangType>(localStg.get('lang') || 'zh-CN');
const options = [
  {
    label: '中文',
    key: 'zh-CN'
  },
  {
    label: 'English',
    key: 'en'
  }
];
const { useInject } = useAppVueContext();
const { reload } = useInject();
const handleSelect = (key: string) => {
  language.value = key as I18nType.LangType;
  locale.value = key;
  localStg.set('lang', key as I18nType.LangType);
  reload();
};
</script>
<style scoped></style>
