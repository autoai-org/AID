import {plain_get, plain_post} from '@/middlewares/api.mdw'

// const discovery_endpoint = "https://discovery.autoai.org"
const discovery_endpoint = "http://127.0.0.1:3000"
const google_profile_endpoint = "https://www.googleapis.com/userinfo/v2/me"

function get_discovery_version() {
    return plain_get(discovery_endpoint+"/version")
}

function login_with_pass(username:string, password:string) {
    return plain_post(discovery_endpoint+"/auth/login", {
        'username': username,
        'password': password
    })
}

function login_with_google() {
    window.open(discovery_endpoint +"/user/google", '_blank')
}

function get_profile_google(access_token:string) {
    return plain_get(google_profile_endpoint, {
        'Authorization': 'Bearer '+ access_token
    })
}

export {
    get_profile_google,
    get_discovery_version,
    login_with_pass,
    login_with_google
}