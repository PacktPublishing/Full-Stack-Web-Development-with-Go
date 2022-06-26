import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path';


// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  // Add the '@' resolver
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'src'),
    },
  },
})
