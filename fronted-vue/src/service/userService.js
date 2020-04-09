import { request } from '@/util/request';

export const register = ({ name, telephone, password }) => {
  return request.post('/api/auth/register', { name, telephone, password });
};

export const login = ({ telephone, password }) => {
  return request.post('/api/auth/login', { telephone, password });
};

export const info = () => {
  return request.post('/api/auth/info', {});
};

export const userService = {
  register,
  login,
  info,
};

export default userService;
