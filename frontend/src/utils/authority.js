import {reloadAuthorized} from './Authorized';

const authKey = "FullStack-authority"
const userNameKey = "FullStack-username"
const tokenKey = "FullStack-token"


export function getAuthority(str) {
  const authorityString =
    typeof str === 'undefined' && localStorage ? localStorage.getItem(authKey) : str; // authorityString could be admin, "admin", ["admin"]

  let authority;

  try {
    if (authorityString) {
      authority = JSON.parse(authorityString);
    }
  } catch (e) {
    authority = authorityString;
  }
  if (typeof authority === 'string') {
    return [authority];
  }

  if (!authority) {
    return ['admin'];
  }

  return authority;
}

export function setAuthority(authority, data) {
  const proAuthority = typeof authority === 'string' ? [authority] : authority;
  localStorage.setItem(authKey, JSON.stringify(proAuthority)); // auto reload
  localStorage.setItem(userNameKey, JSON.stringify(data.username)); // auto reload
  localStorage.setItem(tokenKey, JSON.stringify(data.token)); // auto reload
  reloadAuthorized();
}

export function delAuthority() {
  localStorage.removeItem(authKey);
  localStorage.removeItem(userNameKey);
  localStorage.removeItem(tokenKey);

  reloadAuthorized();
}


export function getLocalStorage(key) {
  const item = localStorage.getItem(key);
  return JSON.parse(item)
}

export function getToken() {
  return getLocalStorage(tokenKey);
};
