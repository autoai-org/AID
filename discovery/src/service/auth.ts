import axios from 'axios';

interface AuthConfig {
    endpoint: string;
}

const authConfig: AuthConfig = {
    endpoint: 'https://oauth.autoai.org'
};

function getUserInfo(accessToken: string) {
    axios.get(authConfig.endpoint + '/oauth2/userInfo', {
        headers: { 'Authorization': 'Bearer ' + accessToken
    }
    }).then((response) => {
    }).catch((error) => {
    });
}

export {
    getUserInfo
};