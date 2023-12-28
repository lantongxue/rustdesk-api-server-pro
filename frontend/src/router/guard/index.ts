import {Router} from "vue-router";
import {PageRoute} from "@/typings/route";
import {routeModuleList, router} from "@/router";
import {useRouteStore} from "@/store/modules/route";
import {createDynamicRouteGuard} from "@/router/guard/dynamic";
import {useAuthStore} from "@/store";


export async function createRouterGuard(router: Router) {
    router.beforeEach(async (to, from, next) => {
        //@ts-ignore
        window.$loadingBar?.start();
        const ua = useAuthStore();
        const isCreate =  await createDynamicRouteGuard(to,from,next);
        if (!isCreate) return;
        // 权限
        next();
    });
    router.afterEach(to => {
        //@ts-ignore
        window.$loadingBar?.finish();
    });
    router.onError((error) => {
        console.log(error);
    })
}