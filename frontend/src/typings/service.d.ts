// Desc:请求相关的类型定义


declare namespace Service {
    /**
     * 服务端返回的数据结构配置
     * @property codeKey - 表示后端请求状态码的属性字段
     * @property dataKey - 表示后端请求数据的属性字段
     * @property msgKey - 表示后端消息的属性字段
     * @property successCode - 后端业务上定义的成功请求的状态
     * @property loginCode - 后端业务上定义的登录失效的状态
     */
    interface ResultConfig {
        codeKey: string;
        dataKey: string;
        msgKey: string;
        successCode: number | string;
        loginCode: number | string;
    }




    /**
     * 样式配置
     */
    interface StyleConfig{
        /**
         * 是否开启loadingBar
         */
        loadingBar: boolean;
    }
    /**
     * 请求的错误类型：
     * - axios: axios错误：网络错误, 请求超时, 默认的兜底错误
     * - http: 请求成功，响应的http状态码非200的错误
     * - backend: 请求成功，响应的http状态码为200，由后端定义的业务错误
     */
    type RequestErrorType = 'axios' | 'http' | 'backend';
    interface RequestError {
        type: RequestErrorType;
        code: string | number;
        msg: string;

    }
    interface SuccessResult<T = any> {
        error: null;
        data: T;
    }
    interface FailedResult {
        error: RequestError;
        data: null;
    }

    type RestResponse<T = any> = SuccessResult<T> | FailedResult;
}