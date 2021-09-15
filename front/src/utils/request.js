import axios from 'axios';
import storeService from '../service/storeService';

const server = axios.create({
  baseURL: 'http://localhost:9090/api/v1',
  timeout: 1000 * 5,
});

// Add a request interceptor
server.interceptors.request.use((config) => {
  // Do something before request is sent
  Object.assign(config.headers, { Authorization: `Bearer ${storeService.get(storeService.USER_TOKEN)}` });
  return config;
}, (error) => Promise.reject(error));

export default server;
