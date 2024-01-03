const dashboard: AuthRoute.Route = {
  name: 'dashboard',
  path: '/dashboard',
  component: 'basic',
  children: [
    {
      name: 'dashboard_analysis',
      path: '/dashboard/analysis',
      component: 'self',
      meta: {
        title: '分析页',
        requiresAuth: true,
        icon: 'icon-park-outline:analysis',
        i18nTitle: 'routes.dashboard.analysis'
      }
    }
  ],
  meta: {
    title: '仪表盘',
    icon: 'mdi:monitor-dashboard',
    order: 1,
    i18nTitle: 'routes.dashboard._value'
  }
};

export default dashboard;
