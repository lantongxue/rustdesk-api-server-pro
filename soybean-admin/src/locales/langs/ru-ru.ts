const local: App.I18n.Schema = {
  system: {
    title: 'Rustdesk Api Server',
    updateTitle: 'Уведомление об обновлении версии системы',
    updateContent: 'Обнаружена новая версия системы. Хотите обновить страницу немедленно?',
    updateConfirm: 'Обновить немедленно',
    updateCancel: 'Позже'
  },
  common: {
    action: 'Действие',
    add: 'Добавить',
    addSuccess: 'Успешно добавлено',
    backToHome: 'Назад к обзору',
    batchDelete: 'Удалить несколько',
    cancel: 'Отмена',
    close: 'Закрыть',
    check: 'Проверить',
    expandColumn: 'Развернуть столбец',
    columnSetting: 'Настройка столбцов',
    config: 'Настроить',
    confirm: 'Принять',
    delete: 'Удалить',
    deleteSuccess: 'Успешное удаление',
    confirmDelete: 'Вы уверены, что хотите удалить?',
    edit: 'Править',
    look: 'Смотреть',
    warning: 'Предупреждение',
    error: 'Ошибка',
    index: 'Индекс',
    keywordSearch: 'Пожалуйста, введите ключевое слово',
    logout: 'Выход',
    logoutConfirm: 'Вы уверены, что хотите выйти?',
    lookForward: 'Coming soon',
    modify: 'Изменить',
    modifySuccess: 'Изменение принято',
    noData: 'Нет данных',
    operate: 'Выполнить',
    pleaseCheckValue: 'Проверьте, является ли значение допустимым.',
    refresh: 'Обновить',
    reset: 'Сбросить',
    search: 'Поиск',
    switch: 'Сменить',
    tip: 'Tip',
    trigger: 'Переключить',
    update: 'Обновить',
    updateSuccess: 'Успешно обновлено',
    userCenter: 'User Center',
    yesOrNo: {
      yes: 'Да',
      no: 'Нет'
    }
  },
  request: {
    logout: 'Выйти из системы после неудачного запроса',
    logoutMsg: 'Статус пользователя недействителен, пожалуйста, войдите снова',
    logoutWithModal: 'Всплывающее модальное окно после сбоя запроса, а затем выход пользователя из системы',
    logoutWithModalMsg: 'Статус пользователя недействителен, пожалуйста, войдите снова',
    refreshToken: 'Срок действия запрошенного токена истек, обновите токен',
    tokenExpired: 'Срок действия запрошенного токена истек'
  },
  theme: {
    themeSchema: {
      title: 'Оформление темы',
      light: 'Светлое',
      dark: 'Темное',
      auto: 'Как в системе'
    },
    grayscale: 'Оттенки серого',
    colourWeakness: 'Цветовая слепота',
    layoutMode: {
      title: 'Режим макета',
      vertical: 'Вертикальное меню',
      horizontal: 'Горизонтальное меню',
      'vertical-mix': 'Mix вертикальное меню',
      'horizontal-mix': 'Mix горизонтальное меню',
      reverseHorizontalMix: 'Изменить положение меню первого уровня и меню дочернего уровня'
    },
    recommendColor: 'Применить рекомендуемую цветовую схему',
    recommendColorDesc: 'Рекомендуемая цветовая схема взята из',
    themeColor: {
      title: 'Цвет темы',
      primary: 'Основной',
      info: 'Информация',
      success: 'Успешно',
      warning: 'Предупреждение',
      error: 'Ошибка',
      followPrimary: 'как основной'
    },
    scrollMode: {
      title: 'Режим прокрутки',
      wrapper: 'Wrapper',
      content: 'Content'
    },
    page: {
      animate: 'Анимация страниц',
      mode: {
        title: 'Режим анимации',
        fade: 'Fade',
        'fade-slide': 'Slide',
        'fade-bottom': 'Fade Zoom',
        'fade-scale': 'Fade Scale',
        'zoom-fade': 'Zoom Fade',
        'zoom-out': 'Zoom Out',
        none: 'None'
      }
    },
    fixedHeaderAndTab: 'Фиксировать заголовок и вкладки',
    header: {
      height: 'Высота заголовка',
      breadcrumb: {
        visible: 'Включить цепочку',
        showIcon: 'Иконки цепочки'
      }
    },
    tab: {
      visible: 'Показать вкладки',
      cache: 'Кэшировать вкладки',
      height: 'Высота вкладок',
      mode: {
        title: 'Режим вкладок',
        chrome: 'Chrome',
        button: 'Button'
      }
    },
    sider: {
      inverted: 'Темный сайдер',
      width: 'Ширина сайдера',
      collapsedWidth: 'Ширина сложенного сайдера',
      mixWidth: 'Mix ширина сайдера',
      mixCollapsedWidth: 'Mix ширина сложенного сайдера',
      mixChildMenuWidth: 'Mix ширина дочернего меню'
    },
    footer: {
      visible: 'Показать подвал',
      fixed: 'Фиксировать подвал',
      height: 'Высота подвала',
      right: 'Подвал вправо'
    },
    watermark: {
      visible: 'Показать водяной знак',
      text: 'Текст водяного знака'
    },
    themeDrawerTitle: 'Настройка темы',
    pageFunTitle: 'Настройка страницы',
    configOperation: {
      copyConfig: 'Копировать',
      copySuccessMsg: 'Копирование конфигурации успешно, пожалуйста, замените переменную «themeSettings» в «src/theme/settings.ts»',
      resetConfig: 'Сбросить',
      resetSuccessMsg: 'Успешный сброс'
    }
  },
  route: {
    login: 'Войти',
    403: 'Нет доступа',
    404: 'Страница не найдена',
    500: 'Ошибка сервера',
    'iframe-page': 'Iframe',
    home: 'Обзор',
    audit: 'Аудит',
    user: 'Пользователи',
    user_list: 'Список пользователей',
    user_sessions: 'Сессии',
    system: 'Система',
    system_mailtemplate: 'Шаблоны почты',
    system_maillogs: 'Журнал почты'
  },
  page: {
    login: {
      common: {
        loginOrRegister: 'Вход / Регистрация',
        userNamePlaceholder: 'Введите имя пользователя',
        phonePlaceholder: 'Введите номер телефона',
        codePlaceholder: 'Введите проверочный код',
        passwordPlaceholder: 'Введите пароль',
        confirmPasswordPlaceholder: 'Введите пароль еще раз',
        codeLogin: 'Вход с кодом подтверждения',
        confirm: 'Войти',
        back: 'Назад',
        validateSuccess: 'Проверка пройдена',
        loginSuccess: 'Авторизация прошла успешно',
        welcomeBack: 'С возвращением, {userName} !'
      },
      pwdLogin: {
        title: 'Password Login',
        rememberMe: 'Запомнить меня'
      }
    },
    home: {
      greeting: 'Доброго утра, {userName}, сегодня еще один день, полный приключений!',
      friendlySponsorship: 'Дружеское спонсорство',
      cupOfCoffee: 'Как насчет чашечки кофе для меня?',
      thankYou: 'Спасибо за твоё спонсорство!',
      userCount: 'Пользователей',
      deviceCount: 'Устройств',
      onlineCount: 'Сейчас онлайн',
      visitsCount: 'Посещений',
      operatingSystem: 'Операционные системы',
      oneWeek: 'За неделю',
      changeLogs: 'Журнал изменений'
    },
    user: {
      list: {
        addUser: 'Добавить пользователя',
        editUser: 'Редактировать пользователя',
        inputUsername: 'Логин',
        inputPassword: 'Пароль',
        inputNickname: 'Имя',
        emailFormatError: 'Ошибка формата электронной почты',
        selectUserStatus: 'Выберите статус пользователя',
        searchPlaceholder: 'Username\\Nickname\\Email',
        tfa_secret_bind: '2FA привязка устройства',
        require2FASecret: '2FA секрет отсутствует',
        require2FACode: "2FA код не подходит"
      },
      sessions: {
        kill: 'Завершить',
        confirmKill: 'Точно завершить?'
      },
      audit: {
        logsSearchPlaceholder: 'Username\\Action\\RustdeskID\\IP'
      }
    },
    system: {
      mailTemplate: {
        addMailTemplate: 'Добавить шаблон',
        editMailTemplate: 'Редактировать шаблон',
        inputName: 'Имя',
        inputSubject: 'Тема письма',
        inputContents: 'Содержимое письма',
        selectType: 'Выбрать тип'
      },
      mailLog: {
        info: 'Информация'
      }
    }
  },
  dropdown: {
    closeCurrent: 'Закрыть текущую',
    closeOther: 'Закрыть другие',
    closeLeft: 'Закрыть левые',
    closeRight: 'Закрыть правые',
    closeAll: 'Закрыть все'
  },
  icon: {
    themeConfig: 'Настройка темы',
    themeSchema: 'Theme Schema',
    lang: 'Сменить язык',
    fullscreen: 'На весь экран',
    fullscreenExit: 'Вернуть меню',
    reload: 'Перезагрузить страницу',
    collapse: 'Свернуть меню',
    expand: 'Расширить меню',
    pin: 'Прикрепить',
    unpin: 'Открепить'
  },
  datatable: {
    itemCount: 'Всего {total} элементов'
  },
  dataMap: {
    user: {
      username: 'Логин',
      password: 'Пароль',
      name: 'Имя',
      email: 'Email',
      licensed_devices: 'Лицензий устройств',
      login_verify: 'Проверен',
      status: 'Статус',
      is_admin: 'Админ',
      tfa_secret: '2FA секрет',
      tfa_code: '2FA код',
      created_at: 'Создан',
      statusLabel: {
        disabled: 'Отключен',
        unverified: 'Не проверен',
        normal: 'Нормальный'
      },
      loginVerifyLabel: {
        none: 'None',
        emailCheck: 'Email',
        tfaCheck: '2FA'
      }
    },
    session: {
      expired: 'Истекает',
      created_at: 'Создана'
    },
    audit: {
      username: 'Имя',
      type: 'Тип',
      conn_id: 'Подключен',
      rustdesk_id: 'Rustdesk ID',
      ip: 'IP',
      session_id: 'ID сессии',
      uuid: 'UUID',
      created_at: 'Создана',
      closed_at: 'Закрыть',
      typeLabel: {
        remote_control: 'Удаленное управление',
        file_transfer: 'Передача файлов',
        tcp_tunnel: 'Туннель TCP'
      }
    },
    mailTemplate: {
      name: 'Имя',
      type: 'Тип',
      subject: 'Тема',
      contents: 'Письмо',
      created_at: 'Создана',
      typeLabel: {
        loginVerify: 'Логин проверен',
        registerVerify: 'Регистрация проверена',
        other: 'Другое'
      }
    },
    mailLog: {
      username: 'Имя',
      uuid: 'UUID',
      from: 'От кого',
      to: 'Кому',
      subject: 'Тема',
      contents: 'Письмо',
      status: 'Статус',
      created_at: 'Время отправки',
      statusLabel: {
        ok: 'Успешно',
        err: 'Ошибка'
      }
    }
  },
  api: {
    CaptchaError: 'Ошибка CAPTCHA',
    UserNotExists: 'Пользователь не существует',
    UsernameOrPasswordError: 'Неправильная учетная запись или пароль',
    UserExists: 'Имя пользователя уже используется',
    UsernameEmpty: 'Имя пользователя не может быть пустым',
    PasswordEmpty: 'Пароль не может быть пустым',
    UserAddSuccess: 'Пользователь успешно создан',
    DataError: 'ошибка данных',
    UserUpdateSuccess: 'Пользователь успешно изменен',
    UserDeleteSuccess: 'Пользователь успешно удален',
    SessionKillSuccess: 'Сессия успешно завершена',
    MailTemplateNameEmpty: 'Имя не может быть пустым',
    MailTemplateSubjectEmpty: 'Тема не может быть пустой',
    MailTemplateContentsEmpty: 'Письмо не может быть пустым',
    MailTemplateAddSuccess: 'Шаблон почты успешно создан',
    MailTemplateUpdateSuccess: 'Шаблон почты успешно изменен',
    NoEmailAddress: 'Адрес электронной почты не установлен',
    VerificationCodeError: 'Ошибка проверочного кода',
    UUIDEmpty: 'UUID не может быть пустым'
  }
};

export default local;
