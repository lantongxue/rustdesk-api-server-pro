/** 请求超时时间 */
export const REQUEST_TIMEOUT = 60 * 1000;

/** 错误信息的显示时间 */
export const ERROR_MSG_DURATION = 3 * 1000;

/** 默认的请求错误文本 */
export const DEFAULT_REQUEST_ERROR_MSG = '请求错误~';

/** 请求超时的错误文本 */
export const REQUEST_TIMEOUT_MSG = '请求超时~';


/** 请求不成功各种状态的错误  和服务端的错误同步即可*/
export const ERROR_STATUS = {
    400: '请求出现语法错误~',
    401: '用户未登录',
    403: '无访问权限',
    404: '请求的资源不存在~',
    405: '请求方法不存在~',
    408: '请求超时~',
    500: '服务器错误~',
    501: '网络未实现~',
    502: '网络错误~',
    503: '服务不可用~',
    504: '网络超时~',
    505: 'http版本不支持该请求~'
};

/** 不弹出错误信息的code  哪些信息不弹出*/
export const NO_ERROR_MSG_CODE: (string | number)[] = [];

/** 刷新登录的code  哪些信息需要刷新登录*/
export const REFRESH_LOGIN_CODE: (string | number)[] = [401];

/** 样式配置 */
export const STYLE_CONFIG: Service.StyleConfig = {
    loadingBar: true
}
/** 后端接口返回的数据结构配置 */
export const RESULT_CONFIG: Service.ResultConfig = {
    codeKey: 'code',
    dataKey: 'data',
    msgKey: 'msg',
    successCode: 0,
    loginCode: 401
}