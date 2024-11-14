import { request } from '../request';

export function fetchMailLogList(params: any) {
  return request<Api.System.MailLogList>({ url: '/mail/logs/list', params });
}
