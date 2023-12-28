import axios, {AxiosInstance, AxiosRequestConfig} from "axios";
import {REFRESH_LOGIN_CODE, RESULT_CONFIG, STYLE_CONFIG} from "@/config";
import {useAuthStore} from "@/store";
import {handelAxiosError, handelBackendError, handleResult} from "@/service/http/helper";


export default class CustomHttpInstance {

    /** axios实例 */
    instance: AxiosInstance;
    /** 后端返回的数据配置 */
    config: Service.ResultConfig;
    /** 样式配置 */
    styleConfig: Service.StyleConfig;


    constructor(axiosConfig: AxiosRequestConfig, resultConfig?: Service.ResultConfig, styleConfig?: Service.StyleConfig) {
        this.instance = axios.create(axiosConfig);
        // 后端返回的数据配置 如果不传入取默认的
        if (resultConfig) {
            this.config = resultConfig;
        } else {
            this.config = RESULT_CONFIG
        }
        // 样式配置 如果不传入取默认的
        if (styleConfig) {
            this.styleConfig = styleConfig;
        } else {
            this.styleConfig = STYLE_CONFIG
        }
        this.setInterceptor();
    }

    setInterceptor() {
        // 添加请求拦截器
        this.instance.interceptors.request.use((config) => {
            // 在发送请求之前做些什么
            if (this.styleConfig.loadingBar) {
                //@ts-ignore
                window.$loadingBar?.start();
            }
            const {token} = useAuthStore();
            config.headers.Authorization = token || '';
            return config;
        }, (error) => {
            if (this.styleConfig.loadingBar) {
                //@ts-ignore
                window.$loadingBar?.error();
            }
            // 对请求错误做些什么
            return handelAxiosError(error);
        });

        // 添加响应拦截器
        // @ts-ignore
        this.instance.interceptors.response.use(response => {
            if (this.styleConfig.loadingBar) {
                //@ts-ignore
                window.$loadingBar?.finish();

            }
            const {codeKey, dataKey, successCode,msgKey} = this.config;
            const {status} = response;
            const resData = response.data;
            if (status === 200) {
                if (resData[codeKey] === successCode) {
                    return handleResult(null, resData[dataKey]);
                }else {
                    return handleResult({code: resData[codeKey] , msg: resData[msgKey],type:'backend' }, null);
                }
                // todo 重新登录
            } else {
                // 不是200的错误
                return handelBackendError(resData, this.config);
            }
        }, (error) => {
            if (this.styleConfig.loadingBar) {
                //@ts-ignore
                window.$loadingBar?.error();
            }
            // todo 401 重新登录
            // 超出 2xx 范围的状态码都会触发该函数。
            // 对响应错误做点什么
            return handelAxiosError(error);
        });
    }
}