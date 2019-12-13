import * as utility from 'utility';
import Axios from 'axios';
import { config } from '../config'

class GithubService {
    protected id:string
    protected secret:string
    constructor(id:string, secret:string) {
        this.id = id
        this.secret = secret        
    }
    public auth(redirectUrl: string, scope: string) {
        const state = utility.randomString()
        let formattedRedirectUrl = `https://github.com/login/oauth/authorize?client_id=${this.id}&redirect_uri=${redirectUrl}&scope=${scope}&state=${state}`
        return formattedRedirectUrl
    }
    public async finishAuth(code: string, redirectUrl: string) {
        const tokenUrl = 'https://github.com/login/oauth/access_token'
        return await Axios.post(tokenUrl, {
            client_id: this.id,
            client_secret: this.secret,
            code: code,
            redirectUrl: redirectUrl,
            grant_type: 'authorization_code',
            state: utility.randomString()
        }, {
            headers: {
                'Content-Type': 'application/json'
            }
        })
    }
}

let githubService = new GithubService(config.githubId, config.githubSecret)

export {
    githubService
}