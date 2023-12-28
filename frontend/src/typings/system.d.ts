// Desc: 系统相关的类型定义
import {IRouteMeta, RouteComponentType} from "@/typings/route";
import {RouteComponent} from "vue-router";


declare namespace System {

    // 菜单项
    type GlobalMenu = import('naive-ui').MenuOption & {
        key: string;
        label: string;
        routeName: string;
        routePath: string;
        meta: IRouteMeta;
        component: RouteComponent;
        icon?: () => import('vue').VNodeChild;
        children?: GlobalMenu[];
    }
    type GlobalBreadcrumb = {
        label: string;
        name: string;
        icon?: () => import('vue').VNodeChild;
        meta?: IRouteMeta;
        children?: GlobalBreadcrumb[]
        // 下拉中disabled
        disabled?: boolean;
    }
    // 下拉菜单
    type GlobalDropdown = {
        label: string;
        key: string;
        icon?: () => import('vue').VNodeChild;
        disabled?: boolean;
        render?: () => import('vue').VNodeChild;
        [x: string]: any;
    }

    type GlobalTheme = {
        // 主题模式
        mode: 'light' | 'dark';
        // 是否跟随系统
        auto: boolean;
        // 主题色
        color: string;
        // 边框圆角
        round: number;
        // 是否开启过渡动画
        isAnimate: boolean;
        /**
         * 过渡动画类型
         * - zoom-fade: 渐变
         * - zoom-out: 闪现
         * - fade-slide: 滑动
         * - fade: 消退
         * - fade-bottom: 底部消退
         * - fade-scale: 缩放消退
         */
        animation: 'zoom-fade' | 'zoom-out' | 'fade-slide' | 'fade' | 'fade-bottom' | 'fade-scale';
        // 是否折叠菜单
        collapsed: boolean;
        // 导航模式
        layout: 'vertical' | 'horizontal';
        // 主题色列表
        colorList: string[];
        // 头部圆角
        headerRound: number;
        // todo：扩展-留作业了。如设置菜单反转等
    }
}