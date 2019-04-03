import VueI18n from 'vue-i18n'
import Vue from 'vue'
import enUS from './en-US.json'
import zhCN from './zh-CN.json'

Vue.use(VueI18n)

const messages = {
  en: enUS,
  zh: zhCN
}

const i18n = new VueI18n({
  locale: 'zh',
  messages
})

export default i18n
