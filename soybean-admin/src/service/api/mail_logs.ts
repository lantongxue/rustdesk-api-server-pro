import { request } from '../request';

export function fetchMailLogList(params: any) {
  return request<Api.System.MailLogList>({ url: '/mail/logs/list', params });
}

export function fetchMailLogInfo(params: any) {
  return request<any>({ url: '/mail/logs/info', params });
}
