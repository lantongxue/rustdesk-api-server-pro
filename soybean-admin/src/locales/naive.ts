import { dateEnUS, dateItIT, dateRuRU, dateZhCN, enUS, itIT, ruRU, zhCN } from 'naive-ui';
import type { NDateLocale, NLocale } from 'naive-ui';

export const naiveLocales: Record<App.I18n.LangType, NLocale> = {
  'zh-CN': zhCN,
  'en-US': enUS,
  'ru-RU': ruRU,
  'it-IT': itIT
};

export const naiveDateLocales: Record<App.I18n.LangType, NDateLocale> = {
  'zh-CN': dateZhCN,
  'en-US': dateEnUS,
  'ru-RU': dateRuRU,
  'it-IT': dateItIT
};
