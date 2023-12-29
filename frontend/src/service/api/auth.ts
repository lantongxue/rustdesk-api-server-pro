import { request } from '../request';

/**
 * 登录
 * @param userName - 用户名
 * @param password - 密码
 */
export function fetchLogin(userName: string, password: string) {
  return request.post<ApiAuth.Token>('/login', { userName, password });
}

/** 获取用户信息 */
export function fetchUserInfo() {
  return request.get<ApiAuth.UserInfo>('/getUserInfo');
}

/**
 * 获取登录验证码
 */
export function fetchCaptcha() {
  return request.get<ApiAuth.Captcha>('/auth/captcha');
}
