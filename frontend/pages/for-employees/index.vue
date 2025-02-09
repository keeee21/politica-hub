<script setup lang="ts">
import TheTitle from '@/components/ui/TheTitle.vue'
</script>

<template>
  <TheTitle
    title="会社員の生活に関わるあれこれを理解する"
    title-align="text-center"
    show-description
    description="正確さよりも、まずはざっくり分かりやすく理解できるように"
    description-align="text-center"
  />

  <div class="grid grid-cols-1 h-screen justify-items-center">
    <div class="p-5 w-full max-w-lg">
      <div class="flex-col gap-4">
        <NuxtLink to="/for-employees/salary-details">
          <div class="my-2 bg-white border border-secondary-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700">
            <div class="p-5">
              <h5 class="mb-2 text-l text-center font-bold tracking-tight text-gray-900 dark:text-white">給与明細の読み方</h5>
              <p class="mb-3 font-normal text-md text-center text-gray-700 dark:text-gray-400">給料から色々控除されているが、これらは何なのか</p>
            </div>
          </div>
        </NuxtLink>

        <NuxtLink to="/for-employees/shakaihoken-overview">
          <div class="my-2 bg-white border border-secondary-200 rounded-lg shadow dark:bg-gray-800 dark:border-gray-700">
            <div class="p-5">
              <h5 class="mb-2 text-l font-bold tracking-tight text-gray-900 dark:text-white">天引きされた社会保険は実際どのように活用されているのか</h5>
              <p class="mb-3 font-normal text-md text-gray-700 dark:text-gray-400">知らなかっただけで、こんな形で活用されている</p>
            </div>
          </div>
        </NuxtLink>
      </div>
    </div>
  </div>
</template>
