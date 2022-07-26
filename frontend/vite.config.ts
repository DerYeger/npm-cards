import react from '@vitejs/plugin-react'
import { presetUno } from 'unocss'
import unocss from 'unocss/vite'
import { defineConfig } from 'vite'
import tsconfigPaths from 'vite-tsconfig-paths'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react(), tsconfigPaths(), unocss({ presets: [presetUno()] })],
  server: {
    port: 3000,
  },
})
