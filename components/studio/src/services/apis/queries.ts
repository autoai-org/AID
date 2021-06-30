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
        }
      	image {
          id,
        }
    }
}
`

export {
    ALL_REPOSITORIES,
    ALL_SOLVERS
}