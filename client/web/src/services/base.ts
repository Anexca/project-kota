import axios from "axios";
import { supabase } from "../supabase/client";

// Create an axios instance
const axiosBase = axios.create({
  baseURL: "https://server-hgkzytv7nq-el.a.run.app", // Replace with your API's base URL
  timeout: 10000, // Set a timeout (in milliseconds)
  headers: {
    Authorization: `Bearer`, // Optional: add authorization if required
  },
});

// const axiosInstance = async () => {
//   const { data } = await supabase.auth.getSession();
//   axiosBase.defaults.headers[
//     "Authorization"
//   ] = `Bearer ${data.session?.access_token}`;
//   return axiosBase;
// };
// Optionally, add request/response interceptors for additional processing

// Request interceptor
axiosBase.interceptors.request.use(
  async (config) => {
    // Modify the config before the request is sent (e.g., attach tokens, etc.)
    const { data } = await supabase.auth.getSession();
    config.headers["Authorization"] = `Bearer ${data.session?.access_token}`;
    return config;
  },
  (error) => {
    // Handle the request error
    return Promise.reject(error);
  }
);

// Response interceptor
axiosBase.interceptors.response.use(
  (response) => {
    // Handle response data
    return response;
  },
  (error) => {
    // Handle the response error
    return Promise.reject(error);
  }
);

export default axiosBase;
