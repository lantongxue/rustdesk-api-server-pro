import {defineStore} from "pinia";
import {System} from "@/typings/system";
import {SecureStorage} from "@/store/plugins";
import {darkTheme, GlobalThemeOverrides, useOsTheme} from 'naive-ui';
import {themeSetting} from "@/setting/theme";
import {getColorPalette} from "@/utils";

// 系统主题
export const systemThemeRef = useOsTheme()
export const useThemeStore = defineStore({
    id: 'theme-store',
    state: (): System.GlobalTheme => initTheme(),
    getters: {
        systemTheme(state) {
            if (state.mode === 'dark') {
                this.setThemeMode("dark")
                return darkTheme
            } else {
                this.setThemeMode("light")
                return undefined
            }
        },
        /**
         * naive-ui theme override
         *
         * @type import('naive-ui').GlobalThemeOverrides
         */
        getNaiveThemeOverrides(state) {
            const themeOverride: GlobalThemeOverrides = {
                common: {
                    primaryColor: state.color,
                    primaryColorHover: getColorPalette(state.color, 5),
                    primaryColorPressed: getColorPalette(state.color, 7),
                    primaryColorSuppl: state.color,
                    borderRadius: `${state.round}px`
                },
                LoadingBar: {
                    colorLoading: state.color
                }
            }
            return themeOverride
        }
    },
    actions: {
        /** 重置theme状态 */
        resetThemeStore() {
            this.$reset();
        },
        /** 设置主题模式 */
        setThemeMode(mode: System.GlobalTheme['mode']) {
            this.mode = mode;
        },
        /** 设置是否跟随系统 */
        setThemeAuto(auto: boolean) {
            this.auto = auto;
        },
        /** 设置主题色 */
        setThemeColor(color: string) {
            this.color = color;
        },
        /** 设置边框圆角 */
        setThemeRounded(round: number) {
            this.round = round;
        },
        /** 设置过渡动画类型 */
        setThemeAnimation(animation: System.GlobalTheme['animation']) {
            this.animation = animation;
        },
        /** 设置是否折叠菜单 */
        setThemeCollapsed(collapsed: boolean) {
            this.collapsed = collapsed;
        },
        /** 设置导航模式 */
        setThemeLayout(layout: System.GlobalTheme['layout']) {
            this.layout = layout;
        }
    },
    persist: {
        storage: SecureStorage
    }
})

function initTheme() {
    return themeSetting
}