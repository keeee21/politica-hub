<script setup lang="ts">
import { convISOStringToJST } from '@/utils/date';
import TheTitle from '@/components/ui/TheTitle.vue';
import { useGetNews } from '@/composables/news/useGetNews';

const { data, status, error, refresh, clear } = useGetNews();
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
        class="bg-white rounded-lg shadow-lg mx-6 p-6 border border-gray-200 transition hover:shadow-xl"
      >
        <!-- 日付 -->
        <div class="flex justify-between items-center text-gray-500 text-sm">
          <span>{{ convISOStringToJST(news.published_at) }}</span>
          {{ typeof news.published_at }}
        </div>

        <!-- タイトル -->
        <a
          :href="news.url"
          target="_blank"
          rel="noopener noreferrer"
          class="block mt-2 text-md font-bold text-blue-600 hover:underline"
        >
          {{ news.title }}
        </a>

        <!-- 説明 -->
        <p class="text-sm text-gray-700 mt-2">
          {{ news.description }}
        </p>
      </div>
    </div>
</template>
