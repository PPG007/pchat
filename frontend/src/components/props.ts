export interface LoginProps {
  onSuccess: (token: string) => void;
  onLoading: () => void;
  onFailure: (message: string) => void;
}