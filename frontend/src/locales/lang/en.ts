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
    confirm: 'Confirm',
    cancel: 'Cancel',
    pleaseCheckValue: 'Please check the value is valid',
    action: 'Action'
  },
  routes: {
    dashboard: {
      _value: 'Dashboard',
      analysis: 'Analysis'
    },
    component: {
      _value: 'Component',
      button: 'Button',
      card: 'Card',
      table: 'Table'
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
      },
      themeConfiguration: {
        title: 'Theme configuration',
        copy: 'Copy the current configuration',
        reset: 'Reset the current configuration',
        resetSuccess: 'The configuration has been reset, please copy it again!',
        operateSuccess: 'Successful operation',
        copySuccess: 'Copy success, please replace the content of src/settings/theme.json!',
        confirmCopy: 'Confirm'
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
    }
  },
  backend: {
    CaptchaError: 'CAPTCHA error',
    UserNotExists: 'The user does not exist',
    UsernameOrPasswordError: 'Incorrect account or password'
  }
};

export default locale;
