import { LoginResponse } from "../pb/user/response.ts";

export interface LoginProps {
  onSuccess: (resp: LoginResponse) => void;
  onLoading: () => void;
  onFailure: (message: string) => void;
}