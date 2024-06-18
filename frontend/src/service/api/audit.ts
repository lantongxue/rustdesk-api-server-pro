import { request } from '../request';

export function fetchLogs(params: any) {
  return request.get<ApiAudit.AuditListResponse>('/audit/list', { params });
}
