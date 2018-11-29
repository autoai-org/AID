import i18n from '@/i18n'
import map from '@/i18n/map'

function setLang (lang) {
  localStorage.setItem('lang', lang)
  console.log(map[lang])
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

export {
  getLang,
  loadDefautlLang,
  setLang
}
