import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import process from "process";

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  define: {
    BACKEND_PORT: process.env.BACKEND_PORT,
  },
  server: {
    host: true,
    strictPort: true,
    port: process.env.FRONTEND_PORT
      ? parseInt(process.env.FRONTEND_PORT)
      : 5173,
  },
});
