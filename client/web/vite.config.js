import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
// https://vitejs.dev/config/
export default defineConfig({
    // resolve: {
    //   alias: {
    //     '@': path.resolve(__dirname, './src'),
    //     '@assets': path.resolve(__dirname, './src/assets'),
    //     '@componnets': path.resolve(__dirname, './src/componnets'),
    //   },
    // },
    plugins: [react()],
});
