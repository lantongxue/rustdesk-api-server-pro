import type { RouteComponent } from 'vue-router';

export const views: Record<
  PageRoute.LastDegreeRouteKey,
  RouteComponent | (() => Promise<{ default: RouteComponent }>)
> = {
  403: () => import('./_builtin/403/index.vue'),
  404: () => import('./_builtin/404/index.vue'),
  500: () => import('./_builtin/500/index.vue'),
  'not-found': () => import('./_builtin/not-found/index.vue'),
  dashboard_analysis: () => import('./dashboard/analysis/index.vue'),
  login: () => import('./login/index.vue'),
  user_list: () => import('./user/list/index.vue'),
  user_sessions: () => import('./user/sessions/index.vue')
};
