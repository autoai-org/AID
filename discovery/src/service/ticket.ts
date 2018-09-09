import axios from 'axios'

interface ticketRequest {
    name: string
    email:string
    subject: string
    description: string
    status: number
    priority: number
}

class FreshDesk {
    protected baseUrl: string
    protected apiKey: string
    private _auth: string
    constructor(baseURL: string, apiKey: string) {
        this.baseUrl = baseURL
        this._auth = 'Basic ' + new Buffer(`${apiKey}:X`).toString('base64')
    }
    _request(method:string, url:string, data:any) {
        return axios.request({
            method: method,
            url: url,
            data: data,
            headers: {
                'Content-Type': 'application/json',
                'Authorization': this._auth
            }
        })
    }
    createTicket(data:ticketRequest) {
        return this._request('post',`${this.baseUrl}/api/v2/tickets`, data)
    }
}

const freshdesk:FreshDesk = new FreshDesk('https://autoai.freshdesk.com', 'a1DuC27QtPoQj61d2NFo')
export {
    freshdesk
}