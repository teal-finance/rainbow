import path from 'path'
import { defineConfig } from 'vite'
import typescript2 from "rollup-plugin-typescript2"
import vue from '@vitejs/plugin-vue'
import Components from 'unplugin-vue-components/vite'
import Icons from 'unplugin-icons/vite'
import IconsResolver from 'unplugin-icons/resolver'

// https://vitejs.dev/config/ 
export default defineConfig({
  // https://rollupjs.org/guide/en/#outputentryfilenames
  // https://rollupjs.org/guide/en/#outputchunkfilenames
  // https://rollupjs.org/guide/en/#outputassetfilenames
  build: {
    rollupOptions: {
      output: {
        // The [hash] permits the use of "Cache-Control: max-age=604800"
        // because [hash] will be different each time its content changes.
        entryFileNames: `assets/js/[name]-[hash].js`,
        chunkFileNames: `assets/js/[name]-[hash].js`,
        assetFileNames: `assets/css/[name]-[hash].[ext]`,
      },
    },
  },

  plugins: [
    typescript2({
      check: false,
      tsconfig: path.resolve(__dirname, 'tsconfig.json'),
      clean: true
    }),
    vue(),
    Components({
      resolvers: IconsResolver(),
    }),
    Icons({
      scale: 1.2,
      defaultClass: 'inline-block align-middle',
      compiler: 'vue3',
    }),
  ],

  resolve: {
    alias: [
      { find: '@/', replacement: '/src/' }
    ]
  },
})
