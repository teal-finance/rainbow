import { createHtmlPlugin } from 'vite-plugin-html'
import { defineConfig, loadEnv } from 'vite'
import Components from 'unplugin-vue-components/vite'
import Icons from 'unplugin-icons/vite'
import IconsResolver from 'unplugin-icons/resolver'
import path from 'path'
import typescript2 from "rollup-plugin-typescript2"
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig(({ command, mode }) => {
  // Load env file based on `mode` https://vitejs.dev/config/#environment-variables
  const env = loadEnv(mode, process.cwd(), '')

  return {
    build: {
      rollupOptions: {
        output: {
          // The [hash] permits the use of "Cache-Control: max-age=604800"
          // because [hash] will be different each time its content changes.
          // https://rollupjs.org/guide/en/#outputentryfilenames
          // https://rollupjs.org/guide/en/#outputchunkfilenames
          // https://rollupjs.org/guide/en/#outputassetfilenames
          entryFileNames: `js/[name]-[hash].js`,
          chunkFileNames: `js/[name]-[hash].js`,
          assetFileNames: `assets/[name]-[hash].[ext]`,
          // CSS and images go to assetFileNames folder.
          // To put images somewhere else use the public folder:
          // https://vitejs.dev/guide/assets.html#the-public-directory
        },
      },
    },

    plugins: [
      // https://github.com/vbenjs/vite-plugin-html
      // https://stackoverflow.com/q/68180648
      createHtmlPlugin({
        minify: true,
        inject: {
          data: {
            VersionInformation: env.VITE_VERS,
            VersionEndpoint: env.VITE_ADDR + "/version"
          },
        },
      }),

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
  }
})
