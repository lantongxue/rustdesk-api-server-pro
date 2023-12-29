<template>
  <div class="relative flex-center wh-full" :style="{ backgroundColor: bgColor }">
    <dark-mode-switch
      :dark="theme.darkMode"
      class="absolute left-48px top-24px z-3 text-20px"
      @update:dark="theme.setDarkMode"
    />
    <n-card :bordered="false" size="large" class="z-4 !w-auto rounded-20px shadow-sm">
      <div class="w-300px sm:w-360px">
        <header class="flex-y-center justify-between">
          <system-logo class="text-64px text-primary" />
          <n-gradient-text type="primary" :size="28">{{ $t('system.title') }}</n-gradient-text>
        </header>
        <main class="pt-24px">
          <h3 class="text-18px text-primary font-medium">{{ $t('page.login.pwdLogin.title') }}</h3>
          <div class="pt-24px">
            <transition name="fade-slide" mode="out-in" appear>
              <pwd-login />
            </transition>
          </div>
        </main>
      </div>
    </n-card>
    <login-bg :theme-color="bgThemeColor" />
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useThemeStore } from '@/store';
import { getColorPalette, mixColor } from '@/utils';
import { $t } from '@/locales';
import { LoginBg, PwdLogin } from './components';

const theme = useThemeStore();

const bgThemeColor = computed(() => (theme.darkMode ? getColorPalette(theme.themeColor, 7) : theme.themeColor));

const bgColor = computed(() => {
  const COLOR_WHITE = '#ffffff';
  const ratio = theme.darkMode ? 0.5 : 0.2;
  return mixColor(COLOR_WHITE, theme.themeColor, ratio);
});
</script>

<style scoped></style>
