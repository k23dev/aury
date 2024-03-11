import { defineConfig } from 'vite'
import { svelte } from '@sveltejs/vite-plugin-svelte'
import { postcss } from 'tailwindcss'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [svelte(),postcss],
})
