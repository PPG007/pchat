import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import {setTokenToCookie, getTokenFromCookie} from "../utils";

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



const userSlice = createSlice<UserState, UserReducer, 'user'>({
  name: 'user',
  reducers: {
    login: (state, action) => {
      state.token = action.payload.token;
      setTokenToCookie(action.payload.token);
    }
  },
  initialState: () => {
    return {
      token: getTokenFromCookie(),
    }
  }
})

export default userSlice;