<script setup lang="ts">
import { useHead } from '#imports'
import TheHeader from '@/components/TheHeader.vue'
useHead({
  htmlAttrs: {
    lang: 'ja',
  },
  charset: 'utf-8',
  title: 'Politica-Hub',
  meta: [
    {
      name: 'robots',
      content: 'noindex, nofollow',
    },
    {
      name: 'viewport',
      content: 'width=device-width, initial-scale=1.0',
    },
  ],
})
</script>
<template>
  <div>
    <TheHeader />
    <NuxtPage />
  </div>
</template>
