import { request } from '../request';

export function fetchStat() {
  return request<Api.Home.Stat>({ url: '/dashboard/stat' });
}

export function fetchLineCharts() {
  return request<Api.Home.LineChart>({ url: '/dashboard/line/charts' });
}

export function fetchPieCharts() {
  return request<Api.Home.PieChart[]>({ url: '/dashboard/pie/charts' });
}
