import Cookies from 'js-cookie'

const TokenKey = 'x-Token'

export function GetToken() {
  return Cookies.get(TokenKey)
}

export function SetToken(token) {
  return Cookies.set(TokenKey, token)
}

export function removeToken() {
  return Cookies.remove(TokenKey)
}

export function removeSession() {
  return Cookies.remove("SessionId")
}