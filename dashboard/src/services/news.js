import axios from 'axios'

class NewsService {
  fetchNews () {
    return new Promise((resolve, reject) => {
      axios.get('https://write.as/api/collections/autoai/posts', {
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
