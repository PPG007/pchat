import { configureStore } from "@reduxjs/toolkit";
import user from "./user.ts";

const store = configureStore({
  reducer: {
    user: user.reducer,
  }
})

export default store;

export type Dispatch = typeof store.dispatch;

export type State = ReturnType<typeof store.getState>

export const {
  login,
} = user.actions;