import {defineStore} from "pinia";
import {SecureStorage} from "@/store/plugins";
import {useRouteStore} from "@/store/modules/route";


interface UserInfo {
    /** 用户id */
    id: string;
    /** 用户名 */
    name: string;
    [x: string]: any;
}
interface AuthState {
    /** 用户信息 */
    userInfo: UserInfo;
    /** 用户token */
    token: string;
}

export const useAuthStore = defineStore({
    id: 'auth-store',
    state: (): AuthState => ({
        userInfo: {
            id: '',
            name: ''
        },
        token: ''
    }),
    getters: {
        /** 是否登录 */
        isLogin(state) {
            return Boolean(state?.token);
        }
    },
    actions: {
        // 重置auth状态
        resetAuthStore() {
            this.$reset();
        },
        async login(){
            this.userInfo = {
                id: '1',
                name: 'admin'
            }
            this.token = '123456'
        }
    },
    persist:{
        // 持久化存储 使用加密的存储方式
        storage: SecureStorage
    }
})