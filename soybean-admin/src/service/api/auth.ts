import { request } from '../request';

/**
 * Login
 *
 * @param userName User name
 * @param password Password
 */
export function fetchLogin(model: Api.Form.LoginForm) {
  return request<Api.Auth.LoginToken>({
    url: '/auth/login',
    method: 'post',
    data: {
      username: model.username,
      password: model.password,
      code: model.code,
      captchaId: model.captchaId
    }
  });
}

export function fetchCaptcha() {
  return request<Api.Auth.Captcha>({ url: '/auth/captcha' });
}

/** Get user info */
export function fetchGetUserInfo() {
  return request<Api.Auth.UserInfo>({ url: '/userinfo' });
}
