import { $t } from '@/locales';

export const UserStatus: Map<number, string> = new Map<number, string>([
  [0, $t('dataMap.user.statusLabel.disabled')],
  [-1, $t('dataMap.user.statusLabel.unverified')],
  [1, $t('dataMap.user.statusLabel.normal')]
]);

export const UserStatusOptions = [
  {
    value: -1,
    label: $t('dataMap.user.statusLabel.unverified')
  },
  {
    value: 0,
    label: $t('dataMap.user.statusLabel.disabled')
  },
  {
    value: 1,
    label: $t('dataMap.user.statusLabel.normal')
  }
];

export const UserIsAdminOptions = [
  {
    value: 0,
    label: $t('common.yesOrNo.no')
  },
  {
    value: 1,
    label: $t('common.yesOrNo.yes')
  }
];
