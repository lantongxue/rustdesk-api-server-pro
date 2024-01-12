const component: AuthRoute.Route = {
  name: 'user',
  path: '/user',
  component: 'basic',
  children: [
    {
      name: 'user_list',
      path: '/user/list',
      component: 'self',
      meta: {
        title: '用户列表',
        i18nTitle: 'routes.user.list',
        requiresAuth: true,
        icon: 'gravity-ui:person'
      }
    }
  ],
  meta: {
    title: '用户管理',
    i18nTitle: 'routes.user._value',
    icon: 'majesticons:users-line',
    order: 2
  }
};

export default component;
