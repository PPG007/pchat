import i18n, { InitOptions } from "i18next";
import I18nextBrowserLanguageDetector from "i18next-browser-languagedetector";
import { initReactI18next } from "react-i18next";
import en_us from "./locales/en_us.ts";
import zh_cn from "./locales/zh_cn.ts";

const debug = process.env.isLocal;

const initOption: InitOptions = {
  debug: !!debug,
  resources: {
    zh: {
      translation: zh_cn,
    },
    en: {
      translation: en_us,
    },
  },
  interpolation: {
    escapeValue: false,
  },
}
const init = () => {
  i18n
    .use(I18nextBrowserLanguageDetector)
    .use(initReactI18next)
    .init(initOption);
}

export default init;