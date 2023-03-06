import { defineConfig, loadEnv } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
  return {
    plugins: [svelte()],
    server: {
      proxy: {
        "/api": {
          target: "https://go-server-p05m.onrender.com",
          changeOrigin: true,
        },
      },
    },
  };
});
