import { reloadAuthorized } from './Authorized';
import {w3cwebsocket as W3CWebSocket} from "websocket/lib/websocket"; // use localStorage to store the authority info, which might be sent from server in actual project.

const authKey = "FullStack-authority"
const userNameKey = "FullStack-username"
const userIdKey = "FullStack-user"
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
  } // preview.pro.ant.design only do not use in your production.
  // preview.pro.ant.design 专用环境变量，请不要在你的项目中使用它。

  if (!authority && ANT_DESIGN_PRO_ONLY_DO_NOT_USE_IN_YOUR_PRODUCTION === 'site') {
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


export function getLocalStorage(key) {
  const item = localStorage.getItem(key);

  return JSON.parse(item)
}

export function getToken () {
  return getLocalStorage(tokenKey);
};
