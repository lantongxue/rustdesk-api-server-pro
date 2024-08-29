import { request } from '../request';

export function fetchAuditLogList(params: any) {
  return request<Api.Audit.AuditLogList>({ url: '/audit/list', params });
}
