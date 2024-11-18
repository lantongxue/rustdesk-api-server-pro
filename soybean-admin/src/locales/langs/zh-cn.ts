const local: App.I18n.Schema = {
  system: {
    title: 'Rustdesk Api Server',
    updateTitle: '系统版本更新通知',
    updateContent: '检测到系统有新版本发布，是否立即刷新页面？',
    updateConfirm: '立即刷新',
    updateCancel: '稍后再说'
  },
  common: {
    action: '操作',
    add: '新增',
    addSuccess: '添加成功',
    backToHome: '返回首页',
    batchDelete: '批量删除',
    cancel: '取消',
    close: '关闭',
    check: '勾选',
    expandColumn: '展开列',
    columnSetting: '列设置',
    config: '配置',
    confirm: '确认',
    delete: '删除',
    deleteSuccess: '删除成功',
    confirmDelete: '确认删除吗？',
    edit: '编辑',
    look: '查看',
    warning: '警告',
    error: '错误',
    index: '序号',
    keywordSearch: '请输入关键词搜索',
    logout: '退出登录',
    logoutConfirm: '确认退出登录吗？',
    lookForward: '敬请期待',
    modify: '修改',
    modifySuccess: '修改成功',
    noData: '无数据',
    operate: '操作',
    pleaseCheckValue: '请检查输入的值是否合法',
    refresh: '刷新',
    reset: '重置',
    search: '搜索',
    switch: '切换',
    tip: '提示',
    trigger: '触发',
    update: '更新',
    updateSuccess: '更新成功',
    userCenter: '个人中心',
    yesOrNo: {
      yes: '是',
      no: '否'
    }
  },
  request: {
    logout: '请求失败后登出用户',
    logoutMsg: '用户状态失效，请重新登录',
    logoutWithModal: '请求失败后弹出模态框再登出用户',
    logoutWithModalMsg: '用户状态失效，请重新登录',
    refreshToken: '请求的token已过期，刷新token',
    tokenExpired: 'token已过期'
  },
  theme: {
    themeSchema: {
      title: '主题模式',
      light: '亮色模式',
      dark: '暗黑模式',
      auto: '跟随系统'
    },
    grayscale: '灰色模式',
    colourWeakness: '色弱模式',
    layoutMode: {
      title: '布局模式',
      vertical: '左侧菜单模式',
      'vertical-mix': '左侧菜单混合模式',
      horizontal: '顶部菜单模式',
      'horizontal-mix': '顶部菜单混合模式',
      reverseHorizontalMix: '一级菜单与子级菜单位置反转'
    },
    recommendColor: '应用推荐算法的颜色',
    recommendColorDesc: '推荐颜色的算法参照',
    themeColor: {
      title: '主题颜色',
      primary: '主色',
      info: '信息色',
      success: '成功色',
      warning: '警告色',
      error: '错误色',
      followPrimary: '跟随主色'
    },
    scrollMode: {
      title: '滚动模式',
      wrapper: '外层滚动',
      content: '主体滚动'
    },
    page: {
      animate: '页面切换动画',
      mode: {
        title: '页面切换动画类型',
        'fade-slide': '滑动',
        fade: '淡入淡出',
        'fade-bottom': '底部消退',
        'fade-scale': '缩放消退',
        'zoom-fade': '渐变',
        'zoom-out': '闪现',
        none: '无'
      }
    },
    fixedHeaderAndTab: '固定头部和标签栏',
    header: {
      height: '头部高度',
      breadcrumb: {
        visible: '显示面包屑',
        showIcon: '显示面包屑图标'
      }
    },
    tab: {
      visible: '显示标签栏',
      cache: '缓存标签页',
      height: '标签栏高度',
      mode: {
        title: '标签栏风格',
        chrome: '谷歌风格',
        button: '按钮风格'
      }
    },
    sider: {
      inverted: '深色侧边栏',
      width: '侧边栏宽度',
      collapsedWidth: '侧边栏折叠宽度',
      mixWidth: '混合布局侧边栏宽度',
      mixCollapsedWidth: '混合布局侧边栏折叠宽度',
      mixChildMenuWidth: '混合布局子菜单宽度'
    },
    footer: {
      visible: '显示底部',
      fixed: '固定底部',
      height: '底部高度',
      right: '底部局右'
    },
    watermark: {
      visible: '显示全屏水印',
      text: '水印文本'
    },
    themeDrawerTitle: '主题配置',
    pageFunTitle: '页面功能',
    configOperation: {
      copyConfig: '复制配置',
      copySuccessMsg: '复制成功，请替换 src/theme/settings.ts 中的变量 themeSettings',
      resetConfig: '重置配置',
      resetSuccessMsg: '重置成功'
    }
  },
  route: {
    login: '登录',
    403: '无权限',
    404: '页面不存在',
    500: '服务器错误',
    'iframe-page': '外链页面',
    home: '首页',
    audit: '日志审计',
    user: '用户管理',
    user_list: '用户列表',
    user_sessions: '会话管理',
    system: '系统管理',
    system_mailtemplate: '邮件模板',
    system_maillogs: '邮件日志'
  },
  page: {
    login: {
      common: {
        loginOrRegister: '登录 / 注册',
        userNamePlaceholder: '请输入用户名',
        phonePlaceholder: '请输入手机号',
        codePlaceholder: '请输入验证码',
        passwordPlaceholder: '请输入密码',
        confirmPasswordPlaceholder: '请再次输入密码',
        codeLogin: '验证码登录',
        confirm: '确定',
        back: '返回',
        validateSuccess: '验证成功',
        loginSuccess: '登录成功',
        welcomeBack: '欢迎回来，{userName} ！'
      },
      pwdLogin: {
        title: '密码登录',
        rememberMe: '记住我'
      }
    },
    home: {
      greeting: '你好，{userName}, 今天又是充满活力的一天!',
      friendlySponsorship: '友情赞助',
      cupOfCoffee: '可以请我喝杯咖啡吗？',
      thankYou: '感谢您的赞助',
      userCount: '用户数量',
      deviceCount: '主机数量',
      onlineCount: '在线数量',
      visitsCount: '访问次数',
      operatingSystem: '操作系统',
      oneWeek: '一周内',
      changeLogs: '更新日志'
    },
    user: {
      list: {
        addUser: '添加用户',
        editUser: '编辑用户',
        inputUsername: '请输入用户名',
        inputPassword: '请输入密码',
        inputNickname: '请输入昵称',
        emailFormatError: '邮箱格式错误',
        selectUserStatus: '请选择用户状态',
        searchPlaceholder: '用户名、昵称、邮箱'
      },
      sessions: {
        kill: '终结',
        confirmKill: '是否终结该会话？'
      },
      audit: {
        logsSearchPlaceholder: '用户名、行为、RustdeskID、IP'
      }
    },
    system: {
      mailTemplate: {
        addMailTemplate: '添加模板',
        editMailTemplate: '编辑模板',
        inputName: '请输入名称',
        inputSubject: '请输入主题',
        inputContents: '请输入模板内容',
        selectType: '请选择类型'
      },
      mailLog: {
        info: '详情'
      }
    }
  },
  dropdown: {
    closeCurrent: '关闭',
    closeOther: '关闭其它',
    closeLeft: '关闭左侧',
    closeRight: '关闭右侧',
    closeAll: '关闭所有'
  },
  icon: {
    themeConfig: '主题配置',
    themeSchema: '主题模式',
    lang: '切换语言',
    fullscreen: '全屏',
    fullscreenExit: '退出全屏',
    reload: '刷新页面',
    collapse: '折叠菜单',
    expand: '展开菜单',
    pin: '固定',
    unpin: '取消固定'
  },
  datatable: {
    itemCount: '共 {total} 条'
  },
  dataMap: {
    user: {
      username: '用户名',
      password: '密码',
      name: '昵称',
      email: '邮箱',
      licensed_devices: '授权设备数量',
      login_verify: '登录验证',
      status: '状态',
      is_admin: '管理员',
      created_at: '创建时间',
      statusLabel: {
        disabled: '禁用',
        unverified: '未验证',
        normal: '正常'
      },
      loginVerifyLabel: {
        none: '无需验证',
        emailCheck: '邮件验证',
        tfaCheck: '双重认证'
      }
    },
    session: {
      expired: '过期时间',
      created_at: '创建时间'
    },
    audit: {
      username: '用户名',
      action: '行为',
      conn_id: '连接ID',
      rustdesk_id: 'RustdeskID',
      ip: 'IP',
      session_id: '会话ID',
      uuid: 'UUID',
      created_at: '创建时间'
    },
    mailTemplate: {
      name: '名称',
      type: '类型',
      subject: '主题',
      contents: '内容',
      created_at: '创建时间',
      typeLabel: {
        loginVerify: '登录验证',
        registerVerify: '注册验证',
        other: '其他'
      }
    },
    mailLog: {
      username: '用户名',
      uuid: 'UUID',
      from: '发送地址',
      to: '接收地址',
      subject: '主题',
      contents: '内容',
      status: '状态',
      created_at: '发送时间',
      statusLabel: {
        ok: '成功',
        err: '失败'
      }
    }
  },
  api: {
    CaptchaError: '验证码错误',
    UserNotExists: '用户不存在',
    UsernameOrPasswordError: '账号或密码错误',
    UserExists: '用户名已使用',
    UsernameEmpty: '用户名不能为空',
    PasswordEmpty: '密码不能为空',
    UserAddSuccess: '用户新建成功',
    DataError: '数据错误',
    UserUpdateSuccess: '用户修改成功',
    UserDeleteSuccess: '用户删除成功',
    SessionKillSuccess: '会话终结成功',
    MailTemplateNameEmpty: '名称不能为空',
    MailTemplateSubjectEmpty: '主题不能为空',
    MailTemplateContentsEmpty: '内容不能为空',
    MailTemplateAddSuccess: '邮件模板添加成功',
    MailTemplateUpdateSuccess: '邮件模板修改成功',
    NoEmailAddress: '未设置邮箱地址',
    VerificationCodeError: '验证码错误',
    UUIDEmpty: 'UUID不能为空'
  }
};

export default local;
