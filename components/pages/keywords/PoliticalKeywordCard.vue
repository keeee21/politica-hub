<script setup lang="ts">
defineProps({
  name: {
    type: String,
    required: true,
  },
  description: {
    type: String,
    required: true,
  },
})
</script>

<template>
  <div class="max-w-sm w-full sm:w-1/3 bg-white border border-secondary-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700">
    <div class="p-5">
      <h5 class="mb-2 text-l font-bold tracking-tight text-gray-900 dark:text-white">{{ name }}</h5>
      <p class="mb-3 font-normal text-md text-gray-700 dark:text-gray-400">{{ description }}</p>
    </div>
  </div>
</template>
