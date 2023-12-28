import {NavigationGuardNext, RouteLocationNormalized} from "vue-router";
import {useRouteStore} from "@/store/modules";
import {useAuthStore} from "@/store";

export async function createDynamicRouteGuard(to: RouteLocationNormalized, _from: RouteLocationNormalized, next: NavigationGuardNext){
    const ur = useRouteStore();
   const ua = useAuthStore();
   const isLogin = ua.isLogin;
   // 初始化权限路由
    if(!ur.isInitRoute){
        // 未登录情况下直接回到登录页，登录成功后再加载权限路由
        if(!isLogin){
            if (to.name === 'login') {
                next();
            }else {
                next({name: 'login', query: {redirect: to.fullPath}});
            }
            return false;
        }
        await ur.initRoute();
        if (to.name === 'not-found') {
            // 动态路由没有加载导致被not-found路由捕获，等待权限路由加载好了，回到之前的路由
            // 若路由是从根路由重定向过来的，重新回到根路由
            next({path: to.redirectedFrom?.name === 'root' ? '/' : to.fullPath, replace: true, query: to.query, hash: to.hash});
            return false;
        }
    }
    // 权限路由已经加载，仍然未找到，重定向到404
    if (to.name === 'not-found') {
        next({ name: '404', replace: true });
        return false;
    }
    return true;
}