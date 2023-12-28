<template>
  <div class="page-header bg-#ffffff dark:bg-dark " :style="[getHeaderStyle]">
    <!--  todo 是否开启顶部菜单   -->
    <div class="page-header-left" v-if="false">
      <page-menu :mode="'horizontal'"/>
    </div>
    <n-space align="center"  v-else>
      <!--  菜单收起  todo 可以配置 显示其他的  -->
      <header-menu-collapse :collapsed="props.collapsed" @click="(e:boolean)=>{ emits('update:collapsed',e) }" v-if="true" />
      <div class="h-full flex-center" v-else>
        <n-button round secondary type="primary">常用</n-button>
      </div>
      <!--   面包屑   -->
      <header-breadcrumb  />
    </n-space>
    <n-space align="center" >
      <header-theme/>
      <n-divider vertical />
      <header-avatar/>
    </n-space>

  </div>
</template>
<script setup lang="ts">
import PageMenu from "@/layouts/components/menu/index.vue";
import HeaderMenuCollapse from "@/layouts/components/header/header-menu-collapse.vue";
import HeaderBreadcrumb from "@/layouts/components/header/header-breadcrumb.vue";
import HeaderAvatar from "@/layouts/components/header/header-avatar.vue";
import HeaderTheme from "@/layouts/components/header/header-theme.vue";
import {computed} from "vue";
import {useThemeStore} from "@/store";
const props = defineProps({
  collapsed: {
    // 侧边栏菜单是否收起
    type: Boolean,
  },
})
const emits = defineEmits(['update:collapsed', 'clickMenuItem']);
const theme = useThemeStore()
const getHeaderStyle = computed(() => {
  return {
    borderBottomLeftRadius: theme.headerRound+'px',
    borderBottomRightRadius: theme.headerRound+'px',
  }
})
</script>

<style scoped lang="scss">
.page-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 64px;
  z-index: 10;
  position: absolute;
  top: 0;
  left: 10px;
  width: calc(100% - 20px);
  box-shadow: 0 1px 4px rgb(0 21 41 / 8%);
  padding: 0 10px;

}

.page-header-fix {
  position: fixed;
  top: 0;
  right: 0;
  left: 200px;
  z-index: 11;
}
</style>