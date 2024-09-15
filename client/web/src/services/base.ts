import axios from "axios";
import { supabase } from "../supabase/client";
import useSessionStore from "../store/auth-store";
import { paths } from "../routes/route.constant";

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
  async (error) => {
    // Handle the response error
    if (error.response && error.response.status === 401) {
      await useSessionStore.getState().logout();
      return;
    }
    if (error.response && error.response.status === 403) {
      window.location.href = `/${paths.HOMEPAGE}`;
      return;
    }
    return Promise.reject(error);
  }
);

export default axiosBase;
