<template>
  <n-spin :show="state.loading" class="w-full" style="height:calc(100vh - 122px)">
    <div class="wh-full" style="height:calc(100vh - 122px)">
      <iframe :src="route.meta?.href as any" class="wh-full box-border " ref="iframeRef"></iframe>
    </div>
  </n-spin>
</template>
<script setup lang="ts">
import {computed, onMounted, reactive, ref} from "vue";
import {useRoute} from "vue-router";

const route = useRoute();

const state = reactive({
  loading: false,
  href: computed(()=> route.meta?.href)
})
const iframeRef = ref<HTMLIFrameElement | null>(null);
function init(){
  if (!iframeRef.value) return;
  const _frame = iframeRef.value as any;
  if (_frame.attachEvent) {
    _frame.attachEvent('onload', () => {
     state.loading = false;
    });
  } else {
    iframeRef.value.onload = () => {
      state.loading = false;
    };
  }
}
onMounted(() => {
  state.loading = true;
  init();
});
</script>

<style lang="scss" scoped>

</style>