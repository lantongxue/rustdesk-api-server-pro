import { $t } from '@/locales';

export const UserStatus: Map<number, string> = new Map<number, string>([
  [0, $t('dataMap.user.statusLabel.disabled')],
  [-1, $t('dataMap.user.statusLabel.unverified')],
  [1, $t('dataMap.user.statusLabel.normal')]
]);
