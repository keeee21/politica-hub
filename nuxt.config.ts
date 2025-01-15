import { defineNuxtConfig } from 'nuxt/config'
// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },
  modules: [
    '@nuxt/eslint',
    '@nuxtjs/tailwindcss',
  ],
  css: [
    '@fortawesome/fontawesome-svg-core/styles.css',
    '@/assets/css/tailwind.css',
  ],
})
