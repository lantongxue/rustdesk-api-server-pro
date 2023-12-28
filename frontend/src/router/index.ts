import {createRouter, createWebHashHistory, createWebHistory, RouteRecordRaw} from "vue-router";
import {App} from "vue";
import {constantRoutes} from "@/router/common";
import {PageRoute} from "@/typings/route";
import {createRouterGuard} from "@/router/guard";
import {useAuthStore} from "@/store";

const {VITE_BASE_URL} = import.meta.env;

export const router = createRouter({
    history: createWebHistory(VITE_BASE_URL),
    routes: constantRoutes as any, // `routes: routes` 的缩写
    strict: true,
    scrollBehavior: () => ({left: 0, top: 0}),
})

export async function setupRouter(app: App) {
    app.use(router);
    await createRouterGuard(router);
    await router.isReady();
}

// 加载需要验证的路由
const modules = import.meta.glob<any>('./modules/**/*.ts', {eager: true});
export const routeModuleList: any = Object.keys(modules).reduce((list: any, key) => {
    const mod: PageRoute | PageRoute[] = modules[key].default ?? {};
    const modList = Array.isArray(mod) ? [...mod] : [mod];
    return [...list, ...modList];
}, [])
