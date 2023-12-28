import {RouteComponent, RouteMeta, RouteRecordRaw} from "vue-router";

/**
 *   路由的组件
 * - basic - 基础布局，具有公共部分的布局
 * - blank - 空白布局
 * - multi - 多级路由布局(三级路由或三级以上时，除第一级路由和最后一级路由，其余的采用该布局) 未用到哈
 * - self - 作为子路由，使用自身的布局(作为最后一级路由，没有子路由) 未用到哈
 * - iframe - 内嵌页面 href不能为空 () => import('vue').VNodeChild
 */
type RouteComponentType = 'iframe' | 'basic' | 'blank' |RouteComponent  ;

interface PageRoute {
    /** 路由名称(路由唯一标识) */
    name: string;
    /** 路由路径 */
    path: string;
    /** 路由重定向 */
    redirect?: string ;
    /**
     *   路由的组件
     * - basic - 基础布局，具有公共部分的布局
     * - blank - 空白布局
     * - multi - 多级路由布局(三级路由或三级以上时，除第一级路由和最后一级路由，其余的采用该布局)
     * - self - 作为子路由，使用自身的布局(作为最后一级路由，没有子路由)
     * - iframe - 内嵌页面 url不能为空
     */
    component?: RouteComponentType;
    /** 路由描述 */
    meta: IRouteMeta;
    /** 子路由 */
    children?: PageRoute[];
}

interface IRouteMeta extends RouteMeta {
    /** 路由标题(可用来作document.title或者菜单的名称) */
    title: string;
    /** 缓存页面 */
    keepAlive?: boolean;
    /** 菜单和面包屑对应的图标 */
    icon?: string;
    /** 使用本地svg作为的菜单和面包屑对应的图标(assets/svg-icon文件夹的的svg文件名) */
    localIcon?: string;
    /** 是否在菜单中隐藏(一些列表、表格的详情页面需要通过参数跳转，所以不能显示在菜单中) */
    hide?: boolean;
    /** 外链链接 */
    href?: string;
    permission?: string;
    /** 路由顺序，可用于菜单的排序 */
    order?: number;
    /** 当前路由需要选中的菜单项(用于跳转至不在左侧菜单显示的路由且需要高亮某个菜单的情况) */
    activeMenu?: string;
    /** 是否固定在tab卡不可关闭  */
    affix?: boolean;
}
