const component: AuthRoute.Route = {
  name: 'audit',
  path: '/audit',
  component: 'self',
  meta: {
    title: '日志审计',
    singleLayout: 'basic',
    i18nTitle: 'routes.audit._value',
    icon: 'ph:warning-bold',
    order: 3
  }
};

export default component;
