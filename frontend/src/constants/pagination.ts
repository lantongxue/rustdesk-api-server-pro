import { $t } from "@/locales";
import { PaginationSizeOption } from "naive-ui/es/pagination";

export const PageSizes: PaginationSizeOption[] = [
  {
    label: '10 / ' + $t('common.page'),
    value: 10
  },
  {
    label: '15 / ' + $t('common.page'),
    value: 15
  },
  {
    label: '20 / ' + $t('common.page'),
    value: 20
  },
  {
    label: '25 / ' + $t('common.page'),
    value: 25
  },
  {
    label: '30 / ' + $t('common.page'),
    value: 30
  },
  {
    label: '50 / ' + $t('common.page'),
    value: 50
  }
];
