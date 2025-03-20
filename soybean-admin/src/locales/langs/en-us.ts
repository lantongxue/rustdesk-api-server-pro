const local: App.I18n.Schema = {
  system: {
    title: 'Rustdesk Api Server',
    updateTitle: 'System Version Update Notification',
    updateContent: 'A new version of the system has been detected. Do you want to refresh the page immediately?',
    updateConfirm: 'Refresh immediately',
    updateCancel: 'Later'
  },
  common: {
    action: 'Action',
    add: 'Add',
    addSuccess: 'Add Success',
    backToHome: 'Back to home',
    batchDelete: 'Batch Delete',
    cancel: 'Cancel',
    close: 'Close',
    check: 'Check',
    expandColumn: 'Expand Column',
    columnSetting: 'Column Setting',
    config: 'Config',
    confirm: 'Confirm',
    delete: 'Delete',
    deleteSuccess: 'Delete Success',
    confirmDelete: 'Are you sure you want to delete?',
    edit: 'Edit',
    look: 'Look',
    warning: 'Warning',
    error: 'Error',
    index: 'Index',
    keywordSearch: 'Please enter keyword',
    logout: 'Logout',
    logoutConfirm: 'Are you sure you want to log out?',
    lookForward: 'Coming soon',
    modify: 'Modify',
    modifySuccess: 'Modify Success',
    noData: 'No Data',
    operate: 'Operate',
    pleaseCheckValue: 'Please check whether the value is valid',
    refresh: 'Refresh',
    reset: 'Reset',
    search: 'Search',
    switch: 'Switch',
    tip: 'Tip',
    trigger: 'Trigger',
    update: 'Update',
    updateSuccess: 'Update Success',
    userCenter: 'User Center',
    yesOrNo: {
      yes: 'Yes',
      no: 'No'
    }
  },
  request: {
    logout: 'Logout user after request failed',
    logoutMsg: 'User status is invalid, please log in again',
    logoutWithModal: 'Pop up modal after request failed and then log out user',
    logoutWithModalMsg: 'User status is invalid, please log in again',
    refreshToken: 'The requested token has expired, refresh the token',
    tokenExpired: 'The requested token has expired'
  },
  theme: {
    themeSchema: {
      title: 'Theme Schema',
      light: 'Light',
      dark: 'Dark',
      auto: 'Follow System'
    },
    grayscale: 'Grayscale',
    colourWeakness: 'Colour Weakness',
    layoutMode: {
      title: 'Layout Mode',
      vertical: 'Vertical Menu Mode',
      horizontal: 'Horizontal Menu Mode',
      'vertical-mix': 'Vertical Mix Menu Mode',
      'horizontal-mix': 'Horizontal Mix menu Mode',
      reverseHorizontalMix: 'Reverse first level menus and child level menus position'
    },
    recommendColor: 'Apply Recommended Color Algorithm',
    recommendColorDesc: 'The recommended color algorithm refers to',
    themeColor: {
      title: 'Theme Color',
      primary: 'Primary',
      info: 'Info',
      success: 'Success',
      warning: 'Warning',
      error: 'Error',
      followPrimary: 'Follow Primary'
    },
    scrollMode: {
      title: 'Scroll Mode',
      wrapper: 'Wrapper',
      content: 'Content'
    },
    page: {
      animate: 'Page Animate',
      mode: {
        title: 'Page Animate Mode',
        fade: 'Fade',
        'fade-slide': 'Slide',
        'fade-bottom': 'Fade Zoom',
        'fade-scale': 'Fade Scale',
        'zoom-fade': 'Zoom Fade',
        'zoom-out': 'Zoom Out',
        none: 'None'
      }
    },
    fixedHeaderAndTab: 'Fixed Header And Tab',
    header: {
      height: 'Header Height',
      breadcrumb: {
        visible: 'Breadcrumb Visible',
        showIcon: 'Breadcrumb Icon Visible'
      }
    },
    tab: {
      visible: 'Tab Visible',
      cache: 'Tab Cache',
      height: 'Tab Height',
      mode: {
        title: 'Tab Mode',
        chrome: 'Chrome',
        button: 'Button'
      }
    },
    sider: {
      inverted: 'Dark Sider',
      width: 'Sider Width',
      collapsedWidth: 'Sider Collapsed Width',
      mixWidth: 'Mix Sider Width',
      mixCollapsedWidth: 'Mix Sider Collapse Width',
      mixChildMenuWidth: 'Mix Child Menu Width'
    },
    footer: {
      visible: 'Footer Visible',
      fixed: 'Fixed Footer',
      height: 'Footer Height',
      right: 'Right Footer'
    },
    watermark: {
      visible: 'Watermark Full Screen Visible',
      text: 'Watermark Text'
    },
    themeDrawerTitle: 'Theme Configuration',
    pageFunTitle: 'Page Function',
    configOperation: {
      copyConfig: 'Copy Config',
      copySuccessMsg: 'Copy Success, Please replace the variable "themeSettings" in "src/theme/settings.ts"',
      resetConfig: 'Reset Config',
      resetSuccessMsg: 'Reset Success'
    }
  },
  route: {
    login: 'Login',
    403: 'No Permission',
    404: 'Page Not Found',
    500: 'Server Error',
    'iframe-page': 'Iframe',
    home: 'Home',
    audit: 'Audit',
    user: 'User Management',
    user_list: 'User List',
    user_sessions: 'Sessions',
    system: 'System Management',
    system_mail_template: 'Mail Template',
    system_mail_logs: 'Mail Logs',
    system_mail: 'Mail Managment'
  },
  page: {
    login: {
      common: {
        loginOrRegister: 'Login / Register',
        userNamePlaceholder: 'Please enter user name',
        phonePlaceholder: 'Please enter phone number',
        codePlaceholder: 'Please enter verification code',
        passwordPlaceholder: 'Please enter password',
        confirmPasswordPlaceholder: 'Please enter password again',
        codeLogin: 'Verification code login',
        confirm: 'Confirm',
        back: 'Back',
        validateSuccess: 'Verification passed',
        loginSuccess: 'Login successfully',
        welcomeBack: 'Welcome back, {userName} !'
      },
      pwdLogin: {
        title: 'Password Login',
        rememberMe: 'Remember me'
      }
    },
    home: {
      greeting: 'Good morning, {userName}, today is another day full of vitality!',
      friendlySponsorship: 'Friendly sponsorship',
      cupOfCoffee: 'Can you treat me to a cup of coffee?',
      thankYou: 'Thank you for your sponsorship',
      userCount: 'User Count',
      deviceCount: 'Device Count',
      onlineCount: 'Online Count',
      visitsCount: 'Visits Count',
      operatingSystem: 'Operating System',
      oneWeek: 'One Week',
      changeLogs: 'Change Logs'
    },
    user: {
      list: {
        addUser: 'Add User',
        editUser: 'Edit User',
        inputUsername: 'Input Username',
        inputPassword: 'Input Password',
        inputNickname: 'Input Nickname',
        emailFormatError: 'Email format error',
        selectUserStatus: 'Please select user status',
        searchPlaceholder: 'Username\\Nickname\\Email',
        tfa_secret_bind: '2FA Device Bind',
        require2FASecret: '2FA Secret Empty',
        require2FACode: "2FA Code Can't Empty"
      },
      sessions: {
        kill: 'Kill',
        confirmKill: 'Confirm Kill?'
      },
      audit: {
        logsSearchPlaceholder: 'Username\\Action\\RustdeskID\\IP'
      }
    },
    system: {
      mailTemplate: {
        addMailTemplate: 'Add Template',
        editMailTemplate: 'Edit Template',
        inputName: 'Input Name',
        inputSubject: 'Input Subject',
        inputContents: 'Input Contents',
        selectType: 'Please select type'
      },
      mailLog: {
        info: 'Info'
      }
    }
  },
  dropdown: {
    closeCurrent: 'Close Current',
    closeOther: 'Close Other',
    closeLeft: 'Close Left',
    closeRight: 'Close Right',
    closeAll: 'Close All'
  },
  icon: {
    themeConfig: 'Theme Configuration',
    themeSchema: 'Theme Schema',
    lang: 'Switch Language',
    fullscreen: 'Fullscreen',
    fullscreenExit: 'Exit Fullscreen',
    reload: 'Reload Page',
    collapse: 'Collapse Menu',
    expand: 'Expand Menu',
    pin: 'Pin',
    unpin: 'Unpin'
  },
  datatable: {
    itemCount: 'Total {total} items'
  },
  dataMap: {
    user: {
      username: 'Username',
      password: 'Password',
      name: 'Nickname',
      email: 'Email',
      licensed_devices: 'Licensed Devices',
      login_verify: 'Login Verify',
      status: 'Status',
      is_admin: 'Is Admin',
      tfa_secret: '2FA Secret',
      tfa_code: '2FA Code',
      created_at: 'Created At',
      statusLabel: {
        disabled: 'Disabled',
        unverified: 'Unverified',
        normal: 'Normal'
      },
      loginVerifyLabel: {
        none: 'None',
        emailCheck: 'Email Check',
        tfaCheck: '2FA'
      }
    },
    session: {
      expired: 'Expired At',
      created_at: 'Created At'
    },
    audit: {
      username: 'Username',
      type: 'Type',
      conn_id: 'Connect Id',
      rustdesk_id: 'Rustdesk ID',
      ip: 'IP',
      session_id: 'Session Id',
      uuid: 'UUID',
      created_at: 'Created At',
      closed_at: 'Closed At',
      typeLabel: {
        remote_control: 'Remote Control',
        file_transfer: 'File Transfer',
        tcp_tunnel: 'TCP Tunnel'
      }
    },
    mailTemplate: {
      name: 'Name',
      type: 'Type',
      subject: 'Subject',
      contents: 'Content',
      created_at: 'Created At',
      typeLabel: {
        loginVerify: 'Login Verify',
        registerVerify: 'Register Verify',
        other: 'Other'
      }
    },
    mailLog: {
      username: 'Username',
      uuid: 'UUID',
      from: 'From',
      to: 'To',
      subject: 'Subject',
      contents: 'Content',
      status: 'Status',
      created_at: 'Send Time',
      statusLabel: {
        ok: 'Success',
        err: 'Error'
      }
    }
  },
  api: {
    CaptchaError: 'CAPTCHA error',
    UserNotExists: 'The user does not exist',
    UsernameOrPasswordError: 'Incorrect account or password',
    UserExists: 'The username already used',
    UsernameEmpty: 'User name cannot be empty',
    PasswordEmpty: 'Password cannot be empty',
    UserAddSuccess: 'User created successfully',
    DataError: 'data error',
    UserUpdateSuccess: 'User modified successfully',
    UserDeleteSuccess: 'User deleted successfully',
    SessionKillSuccess: 'Session killed successfully',
    MailTemplateNameEmpty: 'Name cannot be empty',
    MailTemplateSubjectEmpty: 'Subject cannot be empty',
    MailTemplateContentsEmpty: 'Contents cannot be empty',
    MailTemplateAddSuccess: 'Mail template created successfully',
    MailTemplateUpdateSuccess: 'Mail template modified successfully',
    NoEmailAddress: 'No e-mail address set',
    VerificationCodeError: 'Verification Code Error',
    UUIDEmpty: 'UUID cannot be empty'
  }
};

export default local;
