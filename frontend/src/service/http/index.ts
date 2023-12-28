import {getServiceEnvConfig} from "~/env.config";
import {AxiosInstance, AxiosRequestConfig} from "axios";
import CustomHttpInstance from "@/service/http/request";

const {url, proxy} = getServiceEnvConfig(import.meta.env);

const isHttpProxy = import.meta.env.VITE_HTTP_PROXY === 'Y';

// 导出请求实例
export const request = createRequest({baseURL: isHttpProxy ? proxy : url});

type RequestMethod = 'get' | 'post' | 'put' | 'delete';

interface RequestParam {
    url: string;
    method?: RequestMethod;
    data?: any;
    axiosConfig?: AxiosRequestConfig;
}

function createRequest(axiosConfig: AxiosRequestConfig, resultConfig?: Service.ResultConfig, styleConfig?: Service.StyleConfig) {
    const customInstance = new CustomHttpInstance(axiosConfig, resultConfig, styleConfig);


    /**
     * 异步promise请求
     * @param param - 请求参数
     * - url: 请求地址
     * - method: 请求方法 get post put delete (默认get)
     * - data: 请求的body的data
     * - axiosConfig: axios配置
     * @returns {Promise<Service.RestResponse<T>>}
     */
    async function asyncRequest<T>(param: RequestParam): Promise<Service.RestResponse<T>> {
        const {url} = param;
        const method = param.method || 'get';
        const {instance} = customInstance;
        return (await handelRequest({
            instance,
            method,
            url,
            data: param.data,
            config: param.axiosConfig
        })) as Service.RestResponse<T>;
    }

    /**
     * get请求
     * @param url - 请求地址
     * @param config - axios配置
     */
    async function get<T>(url: string, config?: AxiosRequestConfig) {
        return asyncRequest<T>({url, method: 'get', axiosConfig: config});
    }

    /**
     * post请求
     * @param url
     * @param data
     * @param config
     */
    async function post<T>(url: string, data?: any, config?: AxiosRequestConfig) {
        return asyncRequest<T>({url, method: 'post', data, axiosConfig: config});
    }
    async function put<T>(url: string, data?: any, config?: AxiosRequestConfig) {
        return asyncRequest<T>({url, method: 'put', data, axiosConfig: config});
    }
    async function del<T>(url: string, data?: any, config?: AxiosRequestConfig) {
        return asyncRequest<T>({url, method: 'delete', data, axiosConfig: config});
    }
    return {
        get,
        post,
        put,
        del
    };
}

/**
 * 处理请求
 * @param params
 */
async function handelRequest(params: {
    instance: AxiosInstance;
    method: RequestMethod;
    url: string;
    data?: any;
    config?: AxiosRequestConfig;
}) {
    const {instance, method, url, data, config} = params;
    let res: any;
    if (method === 'get' || method === 'delete') {
        res = await instance[method](url, config);
    } else {
        res = await instance[method](url, data, config);
    }
    return res;
}

