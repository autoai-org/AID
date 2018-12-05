import axios from 'axios'
class GithubService {
  constructor (repo) {
    this.repo = repo
  }
  fetchCVPMConfig () {
    return new Promise((resolve, reject) => {
      axios.get('https://api.github.com/repos/' + this.repo + '/contents/cvpm.toml').then(function (res) {
        resolve(res)
      }).catch(function (err) {
        reject(err)
      })
    })
  }
  fetchDependency () {
    return new Promise((resolve, reject) => {
      axios.get('https://api.github.com/repos/' + this.repo + '/contents/requirements.txt').then(function (res) {
        resolve(res)
      }).catch(function (err) {
        reject(err)
      })
    })
  }
  fetchReadme () {
    return new Promise((resolve, reject) => {
      axios.get('https://api.github.com/repos/' + this.repo + '/readme').then(function (res) {
        resolve(res)
      }).catch(function (err) {
        reject(err)
      })
    })
  }
}

export {
    GithubService
}
