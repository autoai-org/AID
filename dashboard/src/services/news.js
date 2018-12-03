import axios from 'axios'

class NewsService {
  fetchNews () {
    return new Promise((resolve, reject) => {
      axios.get('http://127.0.0.1:3000/news', {
        headers: { 'Content-Type': 'application/json' }
      }).then(function (res) {
        resolve(res)
      }).then(function (err) {
        reject(err)
      })
    })
  }
}

const newsService = new NewsService()

export {
    newsService
}
