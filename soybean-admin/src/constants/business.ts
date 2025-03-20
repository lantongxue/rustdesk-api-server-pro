import { $t } from '@/locales';

export function translateOptions(options: CommonType.Option<any>[]) {
  return options.map(option => ({
    ...option,
    label: $t(option.label as App.I18n.I18nKey)
  }));
}

export const UserStatusOptions = [
  {
    value: 0,
    label: 'dataMap.user.statusLabel.disabled'
  },
  {
    value: 1,
    label: 'dataMap.user.statusLabel.normal'
  },
  {
    value: -1,
    label: 'dataMap.user.statusLabel.unverified'
  }
];

export const UserLoginVerifyOptions = [
  {
    value: 'access_token',
    label: 'dataMap.user.loginVerifyLabel.none'
  },
  {
    value: 'email_check',
    label: 'dataMap.user.loginVerifyLabel.emailCheck'
  },
  {
    value: 'tfa_check',
    label: 'dataMap.user.loginVerifyLabel.tfaCheck'
  }
];

export const UserIsAdminOptions = [
  {
    value: 0,
    label: 'common.yesOrNo.no'
  },
  {
    value: 1,
    label: 'common.yesOrNo.yes'
  }
];

export const AuditTypeOptions = [
  {
    value: 0,
    label: 'dataMap.audit.typeLabel.remote_control'
  },
  {
    value: 1,
    label: 'dataMap.audit.typeLabel.file_transfer'
  },
  {
    value: 2,
    label: 'dataMap.audit.typeLabel.tcp_tunnel'
  }
];

export const AuditFileTransferTypeOptions = [
  {
    value: 0,
    label: 'dataMap.audit.fileTransferTypeLabel.master_controlled'
  },
  {
    value: 1,
    label: 'dataMap.audit.fileTransferTypeLabel.controlled_master'
  }
];

export const MailTemplateOptions = [
  {
    value: 1,
    label: 'dataMap.mailTemplate.typeLabel.loginVerify'
  },
  {
    value: 2,
    label: 'dataMap.mailTemplate.typeLabel.registerVerify'
  },
  {
    value: 3,
    label: 'dataMap.mailTemplate.typeLabel.other'
  }
];

export const MailLogStatusOptions = [
  {
    value: 1,
    label: 'dataMap.mailLog.statusLabel.ok'
  },
  {
    value: 2,
    label: 'dataMap.mailLog.statusLabel.err'
  }
];
