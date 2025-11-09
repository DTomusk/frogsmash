import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    host: '0.0.0.0',  // ensures Vite is reachable from Docker
    port: 5173,
    proxy: {
      '/api': {
        target: 'http://api:8080',  // <--- Docker Compose service name
        changeOrigin: true,
        secure: false,
      },
    },
  },
})
