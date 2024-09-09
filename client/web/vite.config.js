import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
// https://vitejs.dev/config/
export default defineConfig({
    build: {
        // Ensure assets are hashed
        assetsDir: "assets", // Directory to store static assets
        rollupOptions: {
            output: {
                assetFileNames: `assets/[name].[hash][extname]`, // Hash in asset filenames
                chunkFileNames: `assets/[name].[hash].js`, // Hash in chunk filenames
                entryFileNames: `assets/[name].[hash].js`, // Hash in entry filenames
            },
        },
    },
    plugins: [react()],
});
