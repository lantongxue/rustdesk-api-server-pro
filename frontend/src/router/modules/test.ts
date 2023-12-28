import {PageRoute} from "@/typings/route";

const router: PageRoute =
    {
        name: 'text',
        path: '/test',
        redirect: '/test_1',
        meta: {
            title: '测试',
            icon: 'mdi:monitor-dashboard',
        },
        component: () => import('@/layouts/index.vue'),
        children: [
            {
                name: 'test_1',
                path: '/test_1',
                meta: {
                    title: '表格测试',
                    icon: 'line-md:menu',
                },
                component: () => import('@/views/test/one/index.vue'),
            },{
                name: 'test_2',
                path: '/test_2',
                meta: {
                    title: '列表测试',
                    icon: 'line-md:menu',
                },
                component: () => import('@/views/test/two/index.vue'),
            },
        ]
    }

export default router;