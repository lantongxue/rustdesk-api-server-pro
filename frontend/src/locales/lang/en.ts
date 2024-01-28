const locale: I18nType.Schema = {
  system: {
    title: 'Rustdesk Api Server'
  },
  common: {
    add: 'Add',
    addSuccess: 'Add Success',
    edit: 'Edit',
    editSuccess: 'Edit Success',
    delete: 'Delete',
    deleteSuccess: 'Delete Success',
    batchDelete: 'Batch Delete',
    confirmDelete: 'Confirm Delete?',
    confirm: 'Confirm',
    cancel: 'Cancel',
    close: 'Close',
    action: 'Action',
    search: 'Search',
    refreshTable: 'Refresh Table',
    yes: 'Yes',
    no: 'No',
    page: 'Page'
  },
  routes: {
    dashboard: {
      _value: 'Dashboard',
      analysis: 'Analysis'
    },
    user: {
      _value: 'User Management',
      list: 'User List',
      session: 'Sessions'
    }
  },
  dataMap: {
    user: {
      username: 'Username',
      password: 'Password',
      name: 'Nickname',
      email: 'Email',
      licensed_devices: 'Licensed Devices',
      status: 'Status',
      is_admin: 'Is Admin',
      created_at: 'Created At',
      statusLabel: {
        disabled: 'Disabled',
        unverified: 'Unverified',
        normal: 'Normal'
      }
    },
    session: {
      expired: 'Expired At',
      created_at: 'Created At'
    }
  },
  layout: {
    settingDrawer: {
      title: 'Theme configuration',
      themeModeTitle: 'Theme mode',
      darkMode: 'Dark mode',
      layoutModelTitle: 'Layout mode',
      systemThemeTitle: 'System theme',
      pageFunctionsTitle: 'Page functions',
      pageViewTitle: 'Page view',
      followSystemTheme: 'Follow the system',
      isCustomizeDarkModeTransition: 'Custom dark theme animation transition',
      scrollMode: 'scrollMode',
      scrollModeList: {
        wrapper: 'Outer layer scroll',
        content: 'Main body scroll'
      },
      fixedHeaderAndTab: 'Fixed header and multiple tabs',
      header: {
        inverted: 'darkHead',
        height: 'Head Height',
        crumb: {
          visible: 'Crumb',
          icon: 'Crumb icon'
        }
      },
      tab: {
        visible: 'Multi-page tab',
        height: 'Multiple tab height',
        modeList: {
          mode: 'Multi-tab style',
          chrome: 'Google style',
          button: 'Button style'
        },
        isCache: 'Multiple tab caching'
      },
      sider: {
        inverted: 'Dark sidebar',
        width: 'Sidebar expanded width',
        mixWidth: 'Left hybrid sidebar expanded width'
      },
      menu: {
        horizontalPosition: 'Top menu position',
        horizontalPositionList: {
          flexStart: 'Right',
          center: 'center',
          flexEnd: 'Left'
        }
      },
      footer: {
        inverted: 'Dark bottom',
        visible: 'Show bottom',
        fixed: 'Fixed bottom',
        right: 'Bottom to the right'
      },
      page: {
        animate: 'switch animation',
        animateMode: 'switch animation type',
        animateModeList: {
          zoomFade: 'Gradual change',
          zoomOut: 'Flash',
          fadeSlide: 'Slide',
          fade: 'Fade away',
          fadeBottom: 'Bottom fade',
          fadeScale: 'Resizing fade away'
        }
      },
      systemTheme: {
        moreColors: 'More colors'
      }
    },
    header: {
      search: 'Search'
    }
  },
  component: {
    userAvatar: {
      confirmLogout: 'Are you sure you want to log out?',
      userCenter: 'User Center',
      logout: 'Logout'
    },
    search: {
      enterKeywords: 'Please enter a keyword to search',
      noResult: 'No search results',
      enter: 'Enter',
      select: 'Select',
      close: 'Close'
    }
  },
  page: {
    login: {
      common: {
        userNamePlaceholder: 'Please enter user name',
        codePlaceholder: 'Please enter verification code',
        passwordPlaceholder: 'Please enter password',
        confirm: 'Confirm',
        loginSuccess: 'Login success',
        welcomeBack: 'Welcome back, {userName}!'
      },
      pwdLogin: {
        title: 'Password Login'
      }
    },
    dashboard: {
      friendlySponsorship: 'Friendly sponsorship',
      cupOfCoffee: 'Can you treat me to a cup of coffee?',
      thankYou: 'Thank you for your sponsorship',
      userCount: 'User Count',
      peerCount: 'Peer Count',
      onlineCount: 'Online Count',
      visitsCount: 'Visits Count',
      operatingSystem: 'Operating System',
      oneWeek: 'One Week'
    },
    users: {
      addUser: 'Add User',
      editUser: 'Edit User',
      inputUsername: 'Input Username',
      inputPassword: 'Input Password',
      inputNickname: 'Input Nickname',
      emailFormatError: 'Email format error',
      selectUserStatus: 'Please select user status',
      searchPlaceholder: 'Username\\Nickname\\Email'
    },
    session: {
      kill: 'Kill',
      confirmKill: 'Confirm Kill?'
    }
  },
  backend: {
    CaptchaError: 'CAPTCHA error',
    UserNotExists: 'The user does not exist',
    UsernameOrPasswordError: 'Incorrect account or password',
    UsernameEmpty: 'User name cannot be empty',
    PasswordEmpty: 'Password cannot be empty',
    UserAddSuccess: 'User created successfully',
    DataError: 'data error',
    UserUpdateSuccess: 'User modified successfully',
    UserDeleteSuccess: 'User deleted successfully',
    SessionKillSuccess: 'Session killed successfully'
  }
};

export default locale;
