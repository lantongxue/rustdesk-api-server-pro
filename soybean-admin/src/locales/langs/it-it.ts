const local: App.I18n.Schema = {
  system: {
    title: 'Server API Rustdesk',
    updateTitle: 'Notifica aggiornamento versione di sistema',
    updateContent: 'È stata rilevata una nuova versione del sistema. Vuoi aggiornare subito la pagina?',
    updateConfirm: 'Aggiorna subito',
    updateCancel: 'Più tardi'
  },
  common: {
    action: 'Azione',
    add: 'Aggiungi',
    addSuccess: 'Aggiunto con successo',
    backToHome: 'Torna alla home',
    batchDelete: 'Elimina in blocco',
    cancel: 'Annulla',
    close: 'Chiudi',
    check: 'Controlla',
    expandColumn: 'Espandi colonna',
    columnSetting: 'Impostazioni colonna',
    config: 'Configura',
    confirm: 'Conferma',
    delete: 'Elimina',
    deleteSuccess: 'Eliminato con successo',
    confirmDelete: 'Sei sicuro di voler eliminare?',
    edit: 'Modifica',
    look: 'Visualizza',
    warning: 'Avviso',
    error: 'Errore',
    index: 'Indice',
    keywordSearch: 'Inserisci parola chiave',
    logout: 'Disconnetti',
    logoutConfirm: 'Sei sicuro di voler uscire?',
    lookForward: 'Prossimamente',
    modify: 'Modifica',
    modifySuccess: 'Modificato con successo',
    noData: 'Nessun dato',
    operate: 'Operazione',
    pleaseCheckValue: 'Controlla se il valore è valido',
    refresh: 'Aggiorna',
    reset: 'Reimposta',
    search: 'Cerca',
    switch: 'Cambia',
    tip: 'Suggerimento',
    trigger: 'Attiva',
    update: 'Aggiorna',
    updateSuccess: 'Aggiornato con successo',
    userCenter: 'Centro utente',
    yesOrNo: { yes: 'Sì', no: 'No' }
  },
  request: {
    logout: 'Disconnetti utente dopo errore richiesta',
    logoutMsg: 'Stato utente non valido, effettua nuovamente l’accesso',
    logoutWithModal: 'Mostra finestra dopo errore richiesta e poi disconnetti utente',
    logoutWithModalMsg: 'Stato utente non valido, effettua nuovamente l’accesso',
    refreshToken: 'Il token richiesto è scaduto, aggiorna il token',
    tokenExpired: 'Il token richiesto è scaduto'
  },
  theme: {
    themeSchema: {
      title: 'Schema tema',
      light: 'Chiaro',
      dark: 'Scuro',
      auto: 'Segui sistema'
    },
    grayscale: 'Scala di grigi',
    colourWeakness: 'Daltonismo',
    layoutMode: {
      title: 'Modalità layout',
      vertical: 'Menu verticale',
      horizontal: 'Menu orizzontale',
      'vertical-mix': 'Menu misto verticale',
      'horizontal-mix': 'Menu misto orizzontale',
      reverseHorizontalMix: 'Inverti posizione dei menu di primo e secondo livello'
    },
    recommendColor: 'Applica colore consigliato',
    recommendColorDesc: 'L’algoritmo del colore consigliato si riferisce a',
    themeColor: {
      title: 'Colore tema',
      primary: 'Primario',
      info: 'Informazione',
      success: 'Successo',
      warning: 'Avviso',
      error: 'Errore',
      followPrimary: 'Segui primario'
    },
    scrollMode: {
      title: 'Modalità scorrimento',
      wrapper: 'Contenitore',
      content: 'Contenuto'
    },
    page: {
      animate: 'Animazione pagina',
      mode: {
        title: 'Modalità animazione pagina',
        fade: 'Dissolvenza',
        'fade-slide': 'Scorrimento',
        'fade-bottom': 'Zoom dal basso',
        'fade-scale': 'Dissolvenza con scala',
        'zoom-fade': 'Zoom dissolvenza',
        'zoom-out': 'Zoom out',
        none: 'Nessuna'
      }
    },
    fixedHeaderAndTab: 'Intestazione e schede fisse',
    header: {
      height: 'Altezza intestazione',
      breadcrumb: {
        visible: 'Visibilità breadcrumb',
        showIcon: 'Mostra icone breadcrumb'
      }
    },
    tab: {
      visible: 'Schede visibili',
      cache: 'Cache schede',
      height: 'Altezza schede',
      mode: {
        title: 'Modalità scheda',
        chrome: 'Stile Chrome',
        button: 'Pulsante'
      }
    },
    sider: {
      inverted: 'Menu laterale scuro',
      width: 'Larghezza menu laterale',
      collapsedWidth: 'Larghezza menu collassato',
      mixWidth: 'Larghezza menu misto',
      mixCollapsedWidth: 'Larghezza menu misto collassato',
      mixChildMenuWidth: 'Larghezza sottomenu misto'
    },
    footer: {
      visible: 'Visibilità piè di pagina',
      fixed: 'Piè di pagina fisso',
      height: 'Altezza piè di pagina',
      right: 'Piè di pagina destro'
    },
    watermark: {
      visible: 'Filigrana visibile a schermo intero',
      text: 'Testo filigrana'
    },
    themeDrawerTitle: 'Configurazione tema',
    pageFunTitle: 'Funzioni pagina',
    configOperation: {
      copyConfig: 'Copia configurazione',
      copySuccessMsg:
        'Copiato con successo, sostituisci la variabile "themeSettings" in "src/theme/settings.ts"',
      resetConfig: 'Reimposta configurazione',
      resetSuccessMsg: 'Reimpostazione riuscita'
    }
  },
  route: {
    login: 'Accesso',
    403: 'Permesso negato',
    404: 'Pagina non trovata',
    500: 'Errore del server',
    'iframe-page': 'Iframe',
    home: 'Home',
    audit: 'Audit',
    user: 'Gestione utenti',
    user_list: 'Elenco utenti',
    user_sessions: 'Sessioni',
    system: 'Gestione sistema',
    system_mail_template: 'Modelli email',
    system_mail_logs: 'Log email',
    system_mail: 'Gestione email',
    audit_baselogs: 'Log di base',
    audit_filetransferlogs: 'Log trasferimenti file',
    devices: 'Dispositivi',
  },
  page: {
    login: {
      common: {
        loginOrRegister: 'Accedi / Registrati',
        userNamePlaceholder: 'Inserisci nome utente',
        phonePlaceholder: 'Inserisci numero di telefono',
        codePlaceholder: 'Inserisci codice di verifica',
        passwordPlaceholder: 'Inserisci password',
        confirmPasswordPlaceholder: 'Inserisci nuovamente la password',
        codeLogin: 'Accesso con codice di verifica',
        confirm: 'Conferma',
        back: 'Indietro',
        validateSuccess: 'Verifica riuscita',
        loginSuccess: 'Accesso effettuato con successo',
        welcomeBack: 'Bentornato, {userName}!'
      },
      pwdLogin: {
        title: 'Accesso con password',
        rememberMe: 'Ricordami'
      }
    },
    home: {
      greeting: 'Buongiorno, {userName}, oggi è un altro giorno pieno di energia!',
      friendlySponsorship: 'Sponsorizzazione amichevole',
      cupOfCoffee: 'Mi offri un caffè?',
      thankYou: 'Grazie per la tua sponsorizzazione',
      userCount: 'Numero utenti',
      deviceCount: 'Numero dispositivi',
      onlineCount: 'Online',
      visitsCount: 'Visite',
      operatingSystem: 'Sistema operativo',
      oneWeek: 'Una settimana',
      changeLogs: 'Registro modifiche'
    },
    user: {
      list: {
        addUser: 'Aggiungi utente',
        editUser: 'Modifica utente',
        inputUsername: 'Inserisci nome utente',
        inputPassword: 'Inserisci password',
        inputNickname: 'Inserisci nickname',
        emailFormatError: 'Formato email non valido',
        selectUserStatus: 'Seleziona stato utente',
        searchPlaceholder: 'Nome utente\\Nickname\\Email',
        tfa_secret_bind: 'Associa dispositivo 2FA',
        require2FASecret: 'Segreto 2FA mancante',
        require2FACode: 'Codice 2FA obbligatorio'
      },
      sessions: {
        kill: 'Termina',
        confirmKill: 'Confermare terminazione?'
      },
      audit: {
        logsSearchPlaceholder: 'Nome utente\\Azione\\RustdeskID\\IP'
      }
    },
    system: {
      mailTemplate: {
        addMailTemplate: 'Aggiungi modello',
        editMailTemplate: 'Modifica modello',
        inputName: 'Inserisci nome',
        inputSubject: 'Inserisci oggetto',
        inputContents: 'Inserisci contenuto',
        selectType: 'Seleziona tipo'
      },
      mailLog: {
        info: 'Informazioni'
      }
    }
  },
  dropdown: {
    closeCurrent: 'Chiudi corrente',
    closeOther: 'Chiudi altre',
    closeLeft: 'Chiudi a sinistra',
    closeRight: 'Chiudi a destra',
    closeAll: 'Chiudi tutte'
  },
  icon: {
    themeConfig: 'Configurazione tema',
    themeSchema: 'Schema tema',
    lang: 'Cambia lingua',
    fullscreen: 'Schermo intero',
    fullscreenExit: 'Esci da schermo intero',
    reload: 'Ricarica pagina',
    collapse: 'Comprimi menu',
    expand: 'Espandi menu',
    pin: 'Blocca',
    unpin: 'Sblocca'
  },
  datatable: {
    itemCount: 'Totale {total} elementi'
  },
  dataMap: {
    user: {
      username: 'Nome utente',
      password: 'Password',
      name: 'Nickname',
      email: 'Email',
      licensed_devices: 'Dispositivi autorizzati',
      login_verify: 'Verifica accesso',
      status: 'Stato',
      is_admin: 'È amministratore',
      tfa_secret: 'Segreto 2FA',
      tfa_code: 'Codice 2FA',
      created_at: 'Creato il',
      statusLabel: {
        disabled: 'Disabilitato',
        unverified: 'Non verificato',
        normal: 'Normale'
      },
      loginVerifyLabel: {
        none: 'Nessuna',
        emailCheck: 'Verifica email',
        tfaCheck: '2FA'
      }
    },
    session: {
      expired: 'Scade il',
      created_at: 'Creato il'
    },
    device: {
        username: 'Nome utente',
        hostname: 'Nome computer',
        version: 'Versione Rustdesk',
        memory: 'Memoria',
        os: 'Sistema Operativo',
        rustdesk_id: 'ID Rustdesk',
    },
    audit: {
      username: 'Nome utente',
      type: 'Tipo',
      conn_id: 'ID connessione',
      rustdesk_id: 'ID Rustdesk',
      ip: 'IP',
      session_id: 'ID sessione',
      uuid: 'UUID',
      created_at: 'Creato il',
      closed_at: 'Chiuso il',
      typeLabel: {
        remote_control: 'Controllo remoto',
        file_transfer: 'Trasferimento file',
        tcp_tunnel: 'Tunnel TCP'
      },
      fileTransferTypeLabel: {
        master_controlled: 'Master -> Controllato',
        controlled_master: 'Controllato -> Master'
      },
      peer_id: 'ID peer',
      path: 'Percorso'
    },
    mailTemplate: {
      name: 'Nome',
      type: 'Tipo',
      subject: 'Oggetto',
      contents: 'Contenuto',
      created_at: 'Creato il',
      typeLabel: {
        loginVerify: 'Verifica accesso',
        registerVerify: 'Verifica registrazione',
        other: 'Altro'
      }
    },
    mailLog: {
      username: 'Nome utente',
      uuid: 'UUID',
      from: 'Da',
      to: 'A',
      subject: 'Oggetto',
      contents: 'Contenuto',
      status: 'Stato',
      created_at: 'Ora invio',
      statusLabel: {
        ok: 'Successo',
        err: 'Errore'
      }
    }
  },
  api: {
    CaptchaError: 'Errore CAPTCHA',
    UserNotExists: "L'utente non esiste",
    UsernameOrPasswordError: 'Nome utente o password errati',
    UserExists: 'Il nome utente è già in uso',
    UsernameEmpty: 'Il nome utente non può essere vuoto',
    PasswordEmpty: 'La password non può essere vuota',
    UserAddSuccess: 'Utente creato con successo',
    DataError: 'Errore dati',
    UserUpdateSuccess: 'Utente modificato con successo',
    UserDeleteSuccess: 'Utente eliminato con successo',
    SessionKillSuccess: 'Sessione terminata con successo',
    MailTemplateNameEmpty: 'Il nome non può essere vuoto',
    MailTemplateSubjectEmpty: "L'oggetto non può essere vuoto",
    MailTemplateContentsEmpty: 'Il contenuto non può essere vuoto',
    MailTemplateAddSuccess: 'Modello email creato con successo',
    MailTemplateUpdateSuccess: 'Modello email modificato con successo',
    NoEmailAddress: 'Nessun indirizzo email impostato',
    VerificationCodeError: 'Errore codice di verifica',
    UUIDEmpty: "L'UUID non può essere vuoto"
  }
};

export default local;
