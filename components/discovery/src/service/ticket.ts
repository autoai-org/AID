import axios from 'axios';
import { config } from '../config'

interface ticketRequest {
    name: string;
    email: string;
    subject: string;
    description: string;
    status: number;
    priority: number;
}

class FreshDesk {
    protected baseUrl: string;
    protected apiKey: string;
    private _auth: string;
    constructor(baseURL: string, apiKey: string) {
        this.baseUrl = baseURL;
        this._auth = 'Basic ' + new Buffer(`${apiKey}:X`).toString('base64');
    }
    public _request(method: string, url: string, data: any) {
        return axios.request({
            method,
            url,
            data,
            headers: {
                'Content-Type': 'application/json',
                'Authorization': this._auth
            }
        });
    }
    public createTicket(data: ticketRequest) {
        return this._request('post', `${this.baseUrl}/api/v2/tickets`, data);
    }
}

const freshdesk: FreshDesk = new FreshDesk(config.freshdeskDomain, config.freshdeskToken);

export {
    freshdesk
};