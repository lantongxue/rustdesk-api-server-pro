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
    },
    {
      name: 'user_sessions',
      path: '/user/sessions',
      component: 'self',
      meta: {
        title: '会话列表',
        i18nTitle: 'routes.user.session',
        requiresAuth: true,
        icon: 'carbon:mobile-session'
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
