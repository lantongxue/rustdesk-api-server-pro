import { request } from '../request';

export function fetchAuditLogList(params: any) {
  return request<Api.Audit.AuditLogList>({ url: '/audit/list', params });
}

export function fetchAuditFileTransferLogList(params: any) {
  return request<Api.Audit.AuditFileTransferList>({ url: '/audit/file-transfer-list', params });
}
