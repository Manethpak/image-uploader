import { defineConfig, loadEnv } from "vite";
import { svelte } from "@sveltejs/vite-plugin-svelte";

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
  return {
    plugins: [svelte()],
    server: {
      proxy: {
        "/api": {
          target:
            mode === "development"
              ? "http://localhost:3000"
              : "https://go-server-p05m.onrender.com",
          changeOrigin: true,
        },
      },
    },
  };
});
