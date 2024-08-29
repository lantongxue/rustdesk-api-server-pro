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
