// import { App } from 'vue';
import { useLocalStorage } from '@vueuse/core';
import { createI18n /*, VueMessageType*/ } from 'vue-i18n';
// import type { LocaleMessage } from '@intlify/core-base';

import _ from 'lodash-es';

const localePath = '../locales/';
const localeList = ['zh-CN', 'en-US'];

const getLocale = () => {
  const storage = useLocalStorage('DAuth_options', {}) as any;

  const params = new URL(window.location.href).searchParams;
  let locale = params.get('locale') || '';
  if (localeList.includes(locale)) {
    storage.value = {
      appearance: {
        language: locale,
      },
    };
  }

  locale = storage.value?.appearance?.language || '';
  if (localeList.includes(locale)) {
    return locale;
  }

  locale = navigator.language;
  if (locale === 'en') {
    locale = 'en-US';
  }

  if (localeList.includes(locale)) {
    return locale;
  }

  return 'zh-CN';
};

type LocaleMap = { [key: string]: string | LocaleMap };
const getMergedLocaleMessage = () => {
  const locales = Object.entries(import.meta.glob('../locales/**/*.json', { eager: true }));
  console.log(locales);

  const localeMap = {} as LocaleMap;
  locales.forEach((locale) => {
    let key = locale[0] as string;
    const value = locale[1] as LocaleMap;

    key = key.slice(localePath.length, -5);
    const section = key.split('/');
    if (section.length === 1) {
      localeMap[key] = _.merge(value.default, localeMap[key] || {});
    } else {
      const file = section.slice(-1)[0];
      const sectionName = section[0];
      const existed = (localeMap[file] || {}) as LocaleMap;
      localeMap[file] = {
        ...existed,
        [sectionName]: _.merge(value.default, existed[sectionName] || {}),
      };
    }
  });

  // return localeMap as { [x: string]: LocaleMessage<VueMessageType> };
  return localeMap as any;
};

const i18n = createI18n({
  legacy: false,
  locale: getLocale(),
  globalInjection: true,
  messages: getMergedLocaleMessage(),
  fallbackLocale: 'zh-CN',
});

// export const t = _t.t;
// export const te = i18n.global.te;
// export const curLocale = i18n.global.locale;

// const install = (app: App) => {
//   app.config.globalProperties.$t = i18n.global.t;
//   // app.use(i18n);
// };

// export default install;

export default i18n;
