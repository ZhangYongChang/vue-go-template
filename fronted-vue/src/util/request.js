import axios from 'axios';
import {
  storageService,
} from '@/service/storageService';

export const request = axios.create({
  baseURL: 'http://192.168.43.184:8000',
  timeout: 1000,
});

request.interceptors.request.use((config) => {
  Object.assign(config.headers, {
    Authorization: `Bearer ${storageService.get(storageService.USER_TOKEN)}`,
  });
  return config;
}, (error) => {
  return Promise.reject(error);
});

export default request;
