<template>
 <div>
   <div>404</div>
   <n-button type="primary" @click="goHome">去首页</n-button>
 </div>
</template>
<script setup lang="ts">

import {useRouter} from "vue-router";
import {nextTick} from "vue";

const router = useRouter()
function goHome() {

  nextTick(() => {
    router.push({name: 'home'})
  })
}
</script>

<style scoped>

</style>