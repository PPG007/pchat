import { LoginRequest } from "../pb/user/request.ts";
import { LoginResponse } from "../pb/user/response.ts";
import axios from './axios.ts';

export default class User {
  static loginWithPassword = (req: LoginRequest) => {
    return axios.post<LoginResponse>('/users/login', req)
  }
}