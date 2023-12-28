import {PageRoute} from "@/typings/route";
import {basicLayout} from "@/router/common";

const router: PageRoute =
    {
        name: 'home',
        path: '/home',
        redirect: '/home/index',
        meta: {
            title: '首页',
            icon: 'ri:dashboard-2-fill',
            keepAlive: true,
        },
        component: () => basicLayout(),
        children: [
            {
                name: 'home_index',
                path: '/home/index',
                meta: {
                    title: '首页',
                    icon: 'ri:dashboard-2-fill',
                    affix: true,
                    keepAlive: true,
                },
                component: () => import('@/views/home/index.vue'),
            },
        ]
    }

export default router;