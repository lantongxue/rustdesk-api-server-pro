const locale: I18nType.Schema = {
  system: {
    title: 'Rustdesk Api Server'
  },
  common: {
    add: '添加',
    addSuccess: '添加成功',
    edit: '修改',
    editSuccess: '修改成功',
    delete: '删除',
    deleteSuccess: '删除成功',
    batchDelete: '批量删除',
    confirm: '确认',
    cancel: '取消',
    pleaseCheckValue: '请检查输入的值是否合法',
    action: '操作',
    refreshTable: '刷新表格',
    yes: '是',
    no: '否',
    page: '页'
  },
  routes: {
    dashboard: {
      _value: '仪表盘',
      analysis: '分析页'
    },
    user: {
      _value: '用户管理',
      list: '用户列表'
    }
  },
  dataMap: {
    user: {
      username: '用户名',
      name: '昵称',
      email: '邮箱',
      licensed_devices: '授权设备数量',
      status: '状态',
      is_admin: '管理员',
      created_at: '创建时间',
      statusLabel: {
        disabled: '禁用',
        unverified: '未验证',
        normal: '正常'
      }
    }
  },
  layout: {
    settingDrawer: {
      title: '主题配置',
      themeModeTitle: '主题模式',
      darkMode: '深色主题',
      layoutModelTitle: '布局模式',
      systemThemeTitle: '系统主题',
      pageFunctionsTitle: '界面功能',
      pageViewTitle: '界面显示',
      followSystemTheme: '跟随系统',
      isCustomizeDarkModeTransition: '自定义暗黑主题动画过渡',
      scrollMode: '滚动模式',
      scrollModeList: {
        wrapper: '外层滚动',
        content: '主体滚动'
      },
      fixedHeaderAndTab: '固定头部和多页签',
      header: {
        inverted: '头部深色',
        height: '头部高度',
        crumb: {
          visible: '面包屑',
          icon: '面包屑图标'
        }
      },
      tab: {
        visible: '多页签',
        height: '多页签高度',
        modeList: {
          mode: '多页签风格',
          chrome: '谷歌风格',
          button: '按钮风格'
        },
        isCache: '多页签缓存'
      },
      sider: {
        inverted: '侧边栏深色',
        width: '侧边栏展开宽度',
        mixWidth: '左侧混合侧边栏展开宽度'
      },
      menu: {
        horizontalPosition: '顶部菜单位置',
        horizontalPositionList: {
          flexStart: '居左',
          center: '居中',
          flexEnd: '居右'
        }
      },
      footer: {
        inverted: '底部深色',
        visible: '显示底部',
        fixed: '固定底部',
        right: '底部居右'
      },
      page: {
        animate: '页面切换动画',
        animateMode: '页面切换动画类型',
        animateModeList: {
          zoomFade: '渐变',
          zoomOut: '闪现',
          fadeSlide: '滑动',
          fade: '消退',
          fadeBottom: '底部消退',
          fadeScale: '缩放消退'
        }
      },
      systemTheme: {
        moreColors: '更多颜色'
      },
      themeConfiguration: {
        title: '主题配置',
        copy: '拷贝当前配置',
        reset: '重置当前配置',
        resetSuccess: '已重置配置，请重新拷贝！',
        operateSuccess: '操作成功',
        copySuccess: '复制成功,请替换 src/settings/theme.json的内容！',
        confirmCopy: '确认'
      }
    },
    header: {
      search: '搜索'
    }
  },
  component: {
    userAvatar: {
      confirmLogout: '您确定要退出登录吗？',
      userCenter: '用户中心',
      logout: '退出登录'
    },
    search: {
      enterKeywords: '请输入关键词搜索',
      noResult: '暂无搜索结果',
      enter: '确认',
      select: '选择',
      close: '关闭'
    }
  },
  page: {
    login: {
      common: {
        userNamePlaceholder: '请输入用户名',
        codePlaceholder: '请输入验证码',
        passwordPlaceholder: '请输入密码',
        confirm: '确定',
        loginSuccess: '登录成功',
        welcomeBack: '欢迎回来，{userName}!'
      },
      pwdLogin: {
        title: '密码登录'
      }
    },
    dashboard: {
      friendlySponsorship: '友情赞助',
      cupOfCoffee: '可以请我喝杯咖啡吗？',
      thankYou: '感谢您的赞助',
      userCount: '用户数量',
      peerCount: '主机数量',
      onlineCount: '在线数量',
      visitsCount: '访问次数',
      operatingSystem: '操作系统',
      oneWeek: '一周内'
    }
  },
  backend: {
    CaptchaError: '验证码错误',
    UserNotExists: '用户不存在',
    UsernameOrPasswordError: '账号或密码错误'
  }
};

export default locale;
