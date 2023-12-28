<template>
  <n-scrollbar style="max-height:calc(100vh - 108px)" :trigger="'none'">
    <div class="">
        <router-view v-slot="{ Component, route }" :key="key">
          <transition :name="theme.isAnimate?theme.animation:''" mode="out-in" :appear="true"  >
            <keep-alive :include="routeStore.cacheRoutes">
              <component :is="Component" :key="route.fullPath" v-if="routeStore.reloadFlag" />
            </keep-alive>
          </transition>
        </router-view>
      <n-back-top :right="30" class="z100" />
    </div>
  </n-scrollbar>
</template>
<script setup lang="ts">
import {useRouteStore, useThemeStore} from "@/store";
import {computed} from "vue";
import {useRoute} from "vue-router";

const theme = useThemeStore();
const routeStore = useRouteStore();
const route = useRoute();
const key = computed(() => {
  return route.fullPath+Math.random()
});
</script>

<style lang="scss" scoped>

</style>