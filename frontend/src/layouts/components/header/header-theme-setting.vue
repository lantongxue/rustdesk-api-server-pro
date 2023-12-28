<template>
  <n-drawer v-model:show="drawerShow" :width="props.width" :placement="props.placement">
    <n-drawer-content title="主题设置">
      <n-space vertical>
        <n-divider><p>主题模式</p></n-divider>
        <div class="flex-col items-center">
          <n-button-group>
            <n-button round :type="theme.mode=='light'?'primary':'default'" @click="theme.setThemeMode('light')">
              <template #icon>
                <icon-line-md:sunny-filled-loop class="text-#FFCC00"/>
              </template>
              浅色
            </n-button>
            <n-button round :type="theme.mode=='dark'?'primary':'default'" @click="theme.setThemeMode('dark')">
              <template #icon>
                <icon-line-md:moon-filled-loop />
              </template>
              深色
            </n-button>
          </n-button-group>

        </div>
        <n-divider><p>跟随系统</p></n-divider>
        <div class="flex items-center justify-between">
          <p class="text-16px text-#666">是否跟随</p>
          <n-switch :default-value="theme.auto" size="large"  @update:value="theme.setThemeAuto">
            <template #checked-icon >
             <icon-line-md:watch-loop class="text-primary"/>
            </template>
            <template #unchecked-icon>
              <icon-line-md:light-dark-loop class="text-primary"/>
            </template>
            <template #checked>
              自动
            </template>
            <template #unchecked>
              手动
            </template>
          </n-switch>
        </div>
        <n-divider><p>主题颜色</p></n-divider>
        <div class="flex-col w-full">
          <n-grid :cols="8" :x-gap="8" :y-gap="12">
            <n-grid-item v-for="color in theme.colorList" :key="color" class="flex-x-center" @click="theme.setThemeColor(color)">
              <div class="flex-center w-20px h-20px rounded-2px shadow cursor-pointer" :style="{backgroundColor:color}">
                <icon-line-md:confirm  v-if="color === theme.color" class="text-white"/>
              </div>
            </n-grid-item>
          </n-grid>
          <n-color-picker class="mt-24px" :value="theme.color" :show-alpha="false" @update:value="theme.setThemeColor" />
        </div>
        <n-divider><p>界面显示</p></n-divider>
        <div class="flex items-center justify-between">
          <p class="text-16px text-#666">圆角大小(px)</p>
          <n-input-number class="w-80px text-center" v-model:value="theme.round" :min="0"/>
        </div>
        <div class="flex items-center justify-between mt-12px">
          <p class="text-16px text-#666">头部圆角(px)</p>
          <n-input-number class="w-80px text-center" v-model:value="theme.headerRound" :step="2" :min="0"/>
        </div>
        <div class="flex items-center justify-between mt-12px">
          <p class="text-16px text-#666">折叠菜单</p>
          <n-switch v-model:value="theme.collapsed" size="large" >
            <template #unchecked-icon  >
              <icon-line-md:menu-fold-right class="text-primary"/>
            </template>
            <template #checked-icon>
              <icon-line-md:menu-fold-left class="text-primary"/>
            </template>
            <template #unchecked>
              展开
            </template>
            <template #checked>
              收起
            </template>
          </n-switch>
        </div>
        <n-divider><p>动画</p></n-divider>
        <div class="flex items-center justify-between">
          <p class="text-16px text-#666">过渡动画</p>
          <n-switch v-model:value="theme.isAnimate" size="large" >
            <template #checked>
              开启
            </template>
            <template #unchecked>
              关闭
            </template>
            <template #checked-icon >
              <icon-line-md:lightbulb-twotone class="text-primary" />
            </template>
            <template #unchecked-icon>
              <icon-line-md:lightbulb-off-loop class="text-primary"/>
            </template>
          </n-switch>
        </div>
        <div class="flex items-center justify-between mt-12px">
          <p class="text-16px text-#666">动画效果(开启生效)</p>
          <n-select class="w-120px text-center" v-model:value="theme.animation" :options="animationOptions" placeholder="选择效果" ></n-select>
        </div>
      </n-space>
    </n-drawer-content>
  </n-drawer>
</template>
<script setup lang="ts">
import {computed} from "vue";
import {useThemeStore} from "@/store";
import {isWhiteColor} from "@/utils";

const props = defineProps({
  width: {
    type: String,
    default: '360px'
  },
  placement: {
    type: String,
    default: 'right'
  },
  show: {
    type: Boolean,
    default: false
  }
})
const emits = defineEmits(['update:show'])
const drawerShow = computed({
  get() {
    return props.show
  },
  set(val) {
    emits('update:show', val)
  }
})
const theme = useThemeStore()
/**
 * 过渡动画类型
 * - zoom-fade: 渐变
 * - zoom-out: 闪现
 * - fade-slide: 滑动
 * - fade: 消退
 * - fade-bottom: 底部消退
 * - fade-scale: 缩放消退
 */
const animationOptions = [
  {label: '渐变', value: 'zoom-fade'},
  {label: '闪现', value: 'zoom-out'},
  {label: '滑动', value: 'fade-slide'},
  {label: '消退', value: 'fade'},
  {label: '底部消退', value: 'fade-bottom'},
  {label: '缩放消退', value: 'fade-scale'},
]
</script>

<style scoped>

</style>