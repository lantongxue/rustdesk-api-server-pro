import { request } from '../request';

export function fetchUserList(params: any) {
  return request.get<ApiUserManagement.UserListResponse>('/user/list', {params: params});
}
