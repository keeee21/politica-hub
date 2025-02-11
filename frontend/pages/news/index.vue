<script setup lang="ts">
import { ref } from 'vue'
import { convISOStringToJST } from '@/utils/date'
import TheTitle from '@/components/ui/TheTitle.vue'
import { useGetNews } from '@/composables/news/useGetNews'

const { data, error } = useGetNews()
const selectedNews = ref<string | null>(null)

const toggleNews = (url: string) => {
  selectedNews.value = selectedNews.value === url ? null : url
}
</script>

<template>
  <TheTitle
    title="news"
    title-align="text-left"
  />

  <div v-if="error">
    <p class="text-red-500">データの取得に失敗しました。</p>
  </div>

  <div v-else-if="!data">
    <p>Loading...</p>
  </div>

  <div v-else class="space-y-6">
    <div
      v-for="news in data"
      :key="news.id"
      class="bg-white rounded-lg shadow-lg mx-6 p-6 border border-gray-200 transition hover:shadow-xl cursor-pointer"
      @click="toggleNews(news.url)"
    >
      <div class="flex justify-between items-center text-gray-500 text-sm">
        <span>{{ convISOStringToJST(news.published_at) }}</span>
        <font-awesome-icon
          :icon="selectedNews === news.url ? ['fas', 'caret-up'] : ['fas', 'caret-down']"
          class="ml-2 text-gray-400 text-sm"
        />
      </div>
      <p
        class="block mt-2 text-md font-bold text-blue-600 hover:underline">
        {{ news.title }}
      </p>
      <p class="text-sm text-gray-700 mt-2">{{ news.description }}</p>

      <div v-if="selectedNews === news.url" class="mt-4">
        <iframe :src="news.url" class="w-full h-[80vh] border rounded-lg" frameborder="0"/>
      </div>
    </div>
  </div>
</template>
