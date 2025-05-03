import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'
import tailwindcss from '@tailwindcss/vite'


// https://vite.dev/config/
export default defineConfig({
  plugins: [react(), tailwindcss()],
  resolve: {
    alias: {
      '@assets': '/src/assets',
      '@components': '/src/components',
      '@constants': '/src/constants',
      '@hooks': '/src/hooks',
      '@interfaces': '/src/interfaces',
      '@pages': '/src/pages',
      '@redux': '/src/redux',
      '@utils': '/src/utils',
    },
  },
  optimizeDeps: {
    exclude: ['js-big-decimal'],
  },
  css: {
    preprocessorOptions: {
      scss: {
        api: 'modern',
        quietDeps: true,
      },
    },
  },
  build: {
    rollupOptions: {
      output: {
        manualChunks: {
          vendor: ['react', 'react-dom'],
        },
      },
    },
  },
})
