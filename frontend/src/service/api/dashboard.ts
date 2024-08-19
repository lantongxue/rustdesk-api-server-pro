import { request } from '../request';

export function fetchStat() {
  return request.get<ApiDashboard.Stat>('/dashboard/stat');
}

export function fetchLineCharts() {
  return request.get<ApiDashboard.LineChart>('/dashboard/line/charts');
}

export function fetchPieCharts() {
  return request.get<ApiDashboard.PieChart[]>('/dashboard/pie/charts');
}
