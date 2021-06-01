export interface Repository {
    name: string,
    uid: string
}

export interface Solver {
    name: string
    uid: string
    status: string
    repository: Repository
}
