import axios from 'axios';

const requestNode = axios.create({
  baseURL: import.meta.env.VITE_NODE_URL,
});

export default requestNode;
