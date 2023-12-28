import {PageRoute} from "@/typings/route";
import {defineStore} from "pinia";
import {constantRoutes} from "@/router/common";
import {SecureStorage} from "@/store/plugins";


export const useTabsStore = defineStore({
    id: 'app-tabs',
    state: (): {
        tabList: PageRoute[]
    } => ({
        tabList: []
    }),
    actions: {
        // 添加标签页
        addTab(route: PageRoute) {
            // 排除固定路由
            if (constantRoutes.some((item: PageRoute) => {
                if (item.name == route.name) return true;
                // 递归子路由
                if (item.children?.some((child: PageRoute) => child.name == route.name)) return true;
            })) return;
            const isExists = this.tabList.some((item: PageRoute) => item.name == route.name);
            if (!isExists) {
                this.tabList.push(route);
            }
        },
        // 关闭左侧
        closeLeftTabs(route: PageRoute) {
            const index = this.tabList.findIndex((item: PageRoute) => item.name == route.name);
            this.tabList = this.tabList.filter((item: PageRoute, i: number) => i >= index || (item?.meta?.affix ?? false));
        },
        // 关闭右侧
        closeRightTabs(route: PageRoute) {
            const index = this.tabList.findIndex((item: PageRoute) => item.name == route.name);
            this.tabList = this.tabList.filter((item: PageRoute, i: number) => i <= index || (item?.meta?.affix ?? false));
        },
        // 关闭其他
        closeOtherTabs(route: PageRoute) {
            this.tabList = this.tabList.filter(
                (item: PageRoute) => item.name == route.name || (item?.meta?.affix ?? false)
            );
        },
        // 关闭当前页
        closeCurrentTab(route: PageRoute) {
            const index = this.tabList.findIndex((item: PageRoute) => item.name == route.name);
            this.tabList.splice(index, 1);
        },
        // 关闭全部
        closeAllTabs() {
            // 保留固定路由
            this.tabList = this.tabList.filter((item: PageRoute) => item?.meta?.affix ?? false);
        },

    },
    persist: {
        storage: SecureStorage
    }
})