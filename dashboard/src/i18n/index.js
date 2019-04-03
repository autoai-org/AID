import VueI18n from 'vue-i18n'
import Vue from 'vue'
import enUS from './en-US.json'
import zhCN from './zh-CN.json'
import deBRD from './de-BRD.json'

Vue.use(VueI18n)

const messages = {
  en: enUS,
  zh: zhCN,
  de: deBRD
}

const i18n = new VueI18n({
  locale: 'en',
  messages
})

export default i18n
