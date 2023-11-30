import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import Cookies from "js-cookie";

interface UserState {
  token?: string;
  id?: string;
  email?: string;
  avatar?: string;
  name?: string;
}

interface LoginPayload {
  token: string;
}

type UserReducer = {
  login: (state: UserState, action: PayloadAction<LoginPayload>) => void;
}

const accessTokenCookieKey = 'accessToken';

const userSlice = createSlice<UserState, UserReducer, 'user'>({
  name: 'user',
  reducers: {
    login: (state, action) => {
      state.token = action.payload.token;
      Cookies.set(accessTokenCookieKey, action.payload.token)
    }
  },
  initialState: () => {
    return {
      token: Cookies.get(accessTokenCookieKey),
    }
  }
})

export default userSlice;