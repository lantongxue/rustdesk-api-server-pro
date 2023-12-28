<template>
  <base-dropdown-container :options="options" size="medium" @select="handleSelect" >
      <div class="flex-center cursor-pointer">
        <n-avatar :src="avatarUrl" round></n-avatar>
      </div>
  </base-dropdown-container>
</template>
<script setup lang="ts">
import localAvatarUrl from  "@/assets/images/linki.png"
import {useAuthStore} from "@/store";
import {System} from "@/typings/system";
import {renderIcon} from "@/utils";
import {useDialog} from "naive-ui";
import {useRouter} from "vue-router";
import {inject} from "vue";
const ua = useAuthStore()
const avatarUrl = ua.userInfo?.avatar || localAvatarUrl
// 个人中心 系统设置 退出登录
const options:System.GlobalDropdown[] = [
  {
    key: 'profile',
    label: '个人中心',
    icon: renderIcon("line-md:person"),
  },
  {
    key: 'setting',
    label: '系统设置',
    icon: renderIcon('line-md:cog-loop'),
  },
  {
    key: 'logout',
    label: '退出登录',
    icon: renderIcon('line-md:logout'),
  },
]
const dialog = useDialog()
const router = useRouter()
const handleSelect = (key: string) => {
  console.log(key)
  switch (key) {
    case 'profile':
      break
    case 'setting':
      break
    case 'logout':
      dialog.info({
        title: '提示',
        content: '确定退出登录吗？',
        positiveText: '确定',
        negativeText: '取消',
        onPositiveClick: () => {
          ua.resetAuthStore()
          router.push({name: 'login'})
        }
      })
      break
  }
}
</script>

<style scoped>

</style>