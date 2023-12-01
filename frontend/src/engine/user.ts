import { ListRegisterApplicationRequest, ListRegisterApplicationResponse, LoginRequest } from "../pb/user/request.ts";
import { LoginResponse } from "../pb/user/response.ts";
import axios from './axios.ts';

export default class User {
  static loginWithPassword = (req: LoginRequest) => {
    return axios.post<LoginResponse>('/users/login', req)
  }

  static fetchRegisterApplications = (params: ListRegisterApplicationRequest) => {
    return axios.get<ListRegisterApplicationResponse>('/users/registerApplications', {params});
  }
}