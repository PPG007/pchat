import axios, { AxiosError } from "axios";
import i18n from "../../i18n";

const axiosInstance = axios.create({
  timeout: 3000,
  baseURL: '/api',
})
// TODO: add interceptor

axiosInstance.interceptors.response.use((resp) => resp, (err) => {
  const errDetail = err as AxiosError;
  if (errDetail.response?.status !== 400) {
    return Promise.reject(i18n.t('errorMessage.network'));
  }
  return Promise.reject(err);
})
export default axiosInstance;