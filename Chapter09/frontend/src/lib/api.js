import axios from 'axios';

// Create our "axios" object and export
// to the general namespace. This lets us call it as
// api.post(), api.get() etc
export default axios.create({
  baseURL: import.meta.env.VITE_BASE_API_URL,
  withCredentials: true,
});
