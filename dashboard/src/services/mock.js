function makeid () {
  var text = ''
  var possible = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'

  for (let i = 0; i < 14; i++) { text += possible.charAt(Math.floor(Math.random() * possible.length)) }
  return text
}

function getNews (num) {
  const news = []
  for (let i = 0; i < num; i++) {
    const item = {
      'url': 'https://' + makeid() + '.com'
    }
    news.push(item)
  }
  return news
}

function getStatus () {
  const status = {
    'status': 'running',
    'installed': 5,
    'running': 3
  }
  return status
}

export {
    getNews,
    getStatus
}
