import { fileURLToPath, URL } from 'node:url';

import { defineConfig, loadEnv } from 'vite';
import vue from '@vitejs/plugin-vue';
import vueDevTools from 'vite-plugin-vue-devtools';
import Components from 'unplugin-vue-components/vite';
import { AntDesignVueResolver } from 'unplugin-vue-components/resolvers';

// https://vite.dev/config/
export default defineConfig(({ mode }) => {
  // nạp biến môi trường theo mode
  const env = loadEnv(mode, process.cwd(), '');
  return {
    base: env.VITE_BASE_PATH,
    plugins: [
      vue(),
      vueDevTools(),
      Components({
        resolvers: [
          AntDesignVueResolver({
            importStyle: false,
          }),
        ],
      }),
    ],
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url)),
      },
    },
    experimental: {
      renderBuiltUrl(fileName, ctx) {
        if (['js', 'css'].includes(ctx.hostType)) {
          return `window._buildChunkUrl('${fileName}')`;
        }
      },
    },
  };
});
