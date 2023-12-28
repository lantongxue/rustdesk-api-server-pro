import {PageRoute} from "@/typings/route";

const {VITE_ROUTE_HOME_PATH} = import.meta.env;

/** 根路由 */
export const ROOT_ROUTE: PageRoute = {
    name: 'root',
    path: '/',
    redirect: VITE_ROUTE_HOME_PATH,
    meta: {
        title: 'Root'
    }
}
export  const basicLayout = () => import('@/layouts/index.vue')
export  const blankLayout = () => import('@/layouts/blank.vue')
export  const iframeLayout = () => import('@/layouts/iframe.vue')

/** 固定的路由 */
export const constantRoutes: PageRoute[] = [
    ROOT_ROUTE,
    {
        name: 'login',
        path: '/login',
        component: () => import('@/views/login/index.vue'),
        meta: {
            title: '登录',
        },
    },
    {
        name: 'comm',
        path: '/comm',
        component: () => basicLayout(),
        meta: {
            title: '未找到'
        },
        children: [
            {
                name: '404',
                path: '/404',
                component: () => import('@/views/comm/404.vue'),
                meta: {
                    title: '未找到'
                },
            }

        ]
    },
    {
        name: 'not-found',
        path: '/:path(.*)*',
        meta: {
            title: '未找到'
        },
        component: () => import('@/views/comm/404.vue')
    }
]