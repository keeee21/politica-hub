<script setup lang="ts">
type Party = {
  id: number
  name: string
}

defineProps<{
  party: Party
}>()
</script>

<template>
  <div class="max-w-sm w-full sm:w-1/3 bg-white border border-secondary-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700">
    <div class="p-5">
      <h5 class="mb-2 text-l font-bold tracking-tight text-gray-900 dark:text-white">{{ party.name }}</h5>
    </div>
  </div>
</template>
