import FlexSearch from 'flexsearch'

class SearchService {
  constructor () {
    self.index = new FlexSearch({
            // default values:
      encode: 'balance',
      tokenize: 'forward',
      async: false,
      worker: false,
      cache: false
    })
  }
  addItem (id, item) {
    self.index.add(id, item.Desc)
  }
  searchItems (kw) {
    return new Promise((resolve, reject) => {
      self.index.search({
        query: kw,
        limit: 1000,
        threshold: 5, // >= initial threshold
        depth: 3,     // <= initial depth
        callback: function (results) {
          console.log(results)
          resolve(results)
        }
      })
    })
  }
}

const searchService = new SearchService()

export {
    searchService
}
