// Copyright (c) 2021 Xiaozhe Yao et al.
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT


const ALL_REPOSITORIES = `
query AllRepositories {
repositories {
    id,
    name,
    uid,
    localpath,
    status
}
}
`

const ALL_SOLVERS = `
query AllSolvers {
    solvers {
        uid,
      	name,
      	status,
      	repository {
          uid,
          name,
          created_at,
          remote_url,
          vendor
        }
      	image {
          id,
          createdAt,
          container {
              uid,
              running
          }
        }
    }
}
`

const SOLVER_AND_IMAGES = `
query solvers {
    solvers {
        uid,
        image {
        id,
        createdAt,
        container {
            uid,
            running
        }
    }
  }
}
`
const ALL_CONTAINERS = `
query ALLContainers {
    containers {
        port,
    }
}
`

const ALL_IMAGES = `
query Allimages {
	images {
    title
    container {
      port
      running
    }
  }
}
`
export {
    ALL_REPOSITORIES,
    ALL_SOLVERS,
    SOLVER_AND_IMAGES,
    ALL_CONTAINERS,
    ALL_IMAGES
}