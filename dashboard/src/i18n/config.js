import i18n from '@/i18n'
import map from '@/i18n/map.json'
import vuetifyMap from '@/i18n/vuetify-map.json'

function setLang (lang) {
  localStorage.setItem('lang', lang)
  i18n.locale = map[lang]
}

function loadDefautlLang (lang) {
  if (localStorage.getItem('lang') === null) {
    localStorage.setItem('lang', 'English')
  }
  i18n.locale = map[localStorage.getItem('lang')]
}

function getLang () {
  return localStorage.getItem('lang')
}

function getVuetifyLang (lang) {
  return vuetifyMap[lang]
}

export {
  getLang,
  loadDefautlLang,
  setLang,
  getVuetifyLang
}
