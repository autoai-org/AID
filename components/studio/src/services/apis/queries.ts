// Copyright (c) 2021 Xiaozhe Yao et al.
// 
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

import { gql } from '@apollo/client';

const ALL_REPOSITORIES = gql`
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

const ALL_SOLVERS = gql`
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

const SOLVER_AND_IMAGES = gql`
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

export {
    ALL_REPOSITORIES,
    ALL_SOLVERS,
    SOLVER_AND_IMAGES
}