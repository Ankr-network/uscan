import axios, { AxiosRequestConfig } from 'axios';
import { ResponseType } from '../model/index';

const instance = axios.create({
  baseURL: import.meta.env.VITE_BASE_URL,
});

const request = async <T = any>(config: AxiosRequestConfig): Promise<ResponseType<T>> => {
  try {
    const { data } = await instance.request<ResponseType<T>>(config);
    data.code === 200 ? console.log('request', data.msg) : console.error('request', data.msg);
    return data;
  } catch (err) {
    const msg = 'request fail';
    console.error(err);
    return {
      code: -1,
      msg,
      data: err as any,
    };
  }
};

export default request;
