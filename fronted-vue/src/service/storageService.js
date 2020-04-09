const PREFIX = 'frontvue_';

const USER_PREFIX = `${PREFIX}user_`;
const USER_TOKEN = `${USER_PREFIX}token_`;
const USER_INFO = `${USER_PREFIX}info_`;


const set = (key, val) => {
  localStorage.setItem(key, val);
};

const get = (key) => localStorage.getItem(key);

export const storageService = {
  set,
  get,
  USER_TOKEN,
  USER_INFO,
};

export default storageService;
