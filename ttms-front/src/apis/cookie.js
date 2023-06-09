import Cookies from 'js-cookie'

// 将token写入cookie
const AccessToken = ''

export function getToken() {
    return Cookies.get('AccessToken')
}


export function setExpires(expiresAt, expires) {
    return Cookies.set(expiresAtKey, expiresAt, { expires: expires })
}

export function setToken(token, expires = 24) {
    return Cookies.set('AccessToken', token, { expires: expires })
}

export function removeToken() {
    return Cookies.remove('AccessToken')
}


// 将expiresAt过期时间写入cookie
const expiresAtKey = 'ExpiresAt'

export function getExpires() {
    return Cookies.get(expiresAtKey)
}