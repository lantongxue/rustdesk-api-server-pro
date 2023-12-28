<template>
  <app-container height="100vh">
    <n-grid cols="5" item-responsive>
      <n-grid-item span="0 400:5 600:5 800:5 1000:3" :offset="1">
        <div class="login">
          <div class="absolute right-10 top-10">
            <dark-switch/>
          </div>
          <div class="login-item ">
            <icon-local-login-bg class="text-500px text-primary  px-20px"/>
          </div>
          <div class="login-item px-4 w-360px ">
            <PageLogo/>
            <p class="text-20px my-4">登录</p>
            <n-tabs type="segment">
              <n-tab-pane name="chap1" tab="账号密码">
               <transition name="fade-slide" mode="out-in" appear>
                 <LoginPassword/>
                </transition>
              </n-tab-pane>
            </n-tabs>
          </div>
        </div>
      </n-grid-item>
    </n-grid>
  </app-container>
</template>
<script setup lang="ts">

import {useAuthStore} from "@/store";
import {useRouter} from "vue-router";
import {router as globalRouter} from "@/router";
import PageLogo from "@/layouts/components/logo/index.vue";
import DarkSwitch from "@/components/common/dark-switch.vue";
import LoginPassword from "@/views/login/components/login-password.vue";

const ua = useAuthStore()

const router = useRouter()
const login = async () => {
  await ua.login()
}

function goHome() {
  const {query} = router.currentRoute.value
  if (query?.redirect) {
    console.log(query.redirect)
    router.push({path: query.redirect as string})
  } else {
    router.push({name: 'home'})
  }
}
</script>

<style scoped lang="scss">
.login {
  box-sizing: border-box;
  display: flex;
  flex-flow: row wrap;
  width: 100%;
  height: 100vh;
  align-items: center;
  justify-content: center;

  &-item {
    display: flex;
    flex-direction: column;
    justify-content: center;
  }

  &--left {
    background: linear-gradient(140.02deg, #f9f9fa, #f7f8ff);
  }
}
</style>