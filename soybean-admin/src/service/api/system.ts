import { request } from '../request';

export function fetchMailTemplateList(params: any) {
  return request<Api.System.MailTemplateList>({ url: '/mail/templates/list', params });
}

export function addMailTemplate(data: Api.System.MailTemplate) {
  return request({ url: '/mail/templates/add', method: 'post', data });
}

export function editMailTemplate(data: Api.System.MailTemplate) {
  return request({ url: '/mail/templates/edit', method: 'post', data });
}
