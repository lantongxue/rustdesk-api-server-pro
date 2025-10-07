import zhCN from './langs/zh-cn';
import enUS from './langs/en-us';
import ruRU from './langs/ru-ru';
import itIT from './langs/it-it';

const locales: Record<App.I18n.LangType, App.I18n.Schema> = {
  'zh-CN': zhCN,
  'en-US': enUS,
  'ru-RU': ruRU,
  'it-IT': itIT
};

export default locales;
