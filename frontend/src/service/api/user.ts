import { request } from '../request';

export function fetchUserList(params: any) {
  return request.get<ApiUserManagement.UserListResponse>('/users/list', { params });
}

export function addUser(data: any) {
  return request.post('/users/add', data);
}

export function editUser(data: any) {
  return request.post('/users/edit', data);
}

export function delUser(data: any) {
  return request.post('/users/delete', data);
}
