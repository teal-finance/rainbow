import path from 'path'
import { defineConfig } from 'vite'
import typescript2 from "rollup-plugin-typescript2"
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    typescript2({
      check: false,
      tsconfig: path.resolve(__dirname, 'tsconfig.json'),
      clean: true
    }),
    vue(),
  ],
  resolve: {
    alias: [
      { find: '@/', replacement: '/src/' }
    ]
  },
})