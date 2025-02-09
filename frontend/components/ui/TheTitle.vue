<script setup lang="ts">
import { defineProps } from 'vue'

type TextAlign = 'text-left' | 'text-center' | 'text-right'

defineProps<{
  title: string
  titleAlign: TextAlign
  showDescription?: boolean
  description?: string
  descriptionAlign?: TextAlign
}>()
</script>

<template>
  <div class="p-5">
    <h1 :class="['text-3xl', 'font-bold', titleAlign]">{{ title }}</h1>
    <p v-if="showDescription" :class="['p-4', descriptionAlign]">{{ description }}</p>
  </div>
</template>
