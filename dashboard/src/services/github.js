import { _get } from './base'

class GithubService {
  constructor (repo) {
    this.repo = repo
  }

  fetchCVPMConfig () {
    return _get('https://api.github.com/repos/' + this.repo + '/contents/cvpm.toml')
  }

  fetchDependency () {
    return _get('https://api.github.com/repos/' + this.repo + '/contents/requirements.txt')
  }

  fetchReadme () {
    return _get('https://api.github.com/repos/' + this.repo + '/readme')
  }
}

export {
  GithubService
}
