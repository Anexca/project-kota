import axios from "axios";
import { supabase } from "../supabase/client";

// Create an axios instance
const axiosBase = axios.create({
  baseURL: import.meta.env.VITE_SERVER_URL, // Replace with your API's base URL
  timeout: 120000, // Set a timeout (in milliseconds)
  headers: {
    Authorization: `Bearer`, // Optional: add authorization if required
  },
});

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
