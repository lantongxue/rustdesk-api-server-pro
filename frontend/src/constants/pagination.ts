import type { PaginationSizeOption } from 'naive-ui/es/pagination';
import { $t } from '@/locales';

const pageString = $t('common.page');
export const PageSizes: PaginationSizeOption[] = [
  {
    label: `10 / ${pageString}`,
    value: 10
  },
  {
    label: `15 / ${pageString}`,
    value: 15
  },
  {
    label: `20 / ${pageString}`,
    value: 20
  },
  {
    label: `25 / ${pageString}`,
    value: 25
  },
  {
    label: `30 / ${pageString}`,
    value: 30
  },
  {
    label: `50 / ${pageString}`,
    value: 50
  }
];
