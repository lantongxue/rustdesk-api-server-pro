import { request } from '../request';

export function fetchUserList(params: any) {
  return request<Api.UserManagement.UserList>({ url: '/users/list', params });
}

export function addUser(data: Api.UserManagement.User) {
  return request({ url: '/users/add', method: 'post', data });
}

export function editUser(data: Api.UserManagement.User) {
  return request({ url: '/users/edit', method: 'post', data });
}

export function getTOTP(data: Api.UserManagement.User) {
  return request({ url: '/users/totp', method: 'post', data });
}

export function delUser(data: any) {
  return request({ url: '/users/delete', method: 'post', data });
}

export function fetchSessionList(params: any) {
  return request<Api.UserManagement.SessionList>({ url: '/sessions/list', params });
}

export function killSession(data: any) {
  return request({ url: '/sessions/kill', method: 'post', data });
}
