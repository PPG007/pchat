import Cookies from "js-cookie";

export const getMessage = (e: any, fallback: string) => {
  if (typeof e === 'string') {
    return e
  }
  return fallback;
}

const accessTokenCookieKey = 'accessToken';

export const setTokenToCookie = (token: string) => {
  Cookies.set(accessTokenCookieKey, token);
}

export const getTokenFromCookie = () => {
  return Cookies.get(accessTokenCookieKey);
}