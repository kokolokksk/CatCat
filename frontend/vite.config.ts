import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

// Import the requireTransform function from the appropriate module
import requireTransform from 'vite-plugin-require-transform';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react(), requireTransform()],
  server: {
    open: true,
    host: '0.0.0.0',
    proxy: {
      '/client': 'https://api.ririra.com',
      '/repos':"https://api.github.com",
      "/qrcode":"https://passport.bilibili.com"
    }
  }


});

 
