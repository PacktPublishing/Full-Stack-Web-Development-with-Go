import api from '@/lib/api';



export function getFromServer() {
    return api.get(`/`);
}

export function postToServer(data) {
    return api.post(`/`, data );
}

export function doLogin(username, password) {
    return api.post(`/login`, { username, password} );
}


