<script setup lang="ts">
import { ref } from 'vue'

import PoliticalKeywordCard from '@/components/pages/keywords/PoliticalKeywordCard.vue'
import politicalKeywords from '@/data/political-keywords.json'

type Keyword = {
  id: number
  name: string
  description: string
  name_kana: string
}

const keywords = ref<Keyword[]>(politicalKeywords)
// 選択されたキーワードを状態として管理
const selectedKeyword = ref<Keyword | null>(null)
// キーワード選択時に実行する関数
const selectKeyword = (keyword: Keyword) => {
  selectedKeyword.value = keyword
}
// キーワード詳細を閉じる関数
const clearSelection = () => {
  selectedKeyword.value = null
}
</script>

<template>
  <div class="p-5">
    <h1 class="text-3xl font-bold text-center">政治に関する言葉の概要を理解する</h1>
  </div>
  <div class="grid grid-cols-2 h-screen">
    <div class="p-5">
      <div class="flex flex-wrap gap-4">
        <PoliticalKeywordCard
          v-for="keyword in keywords"
          :key="keyword.id"
          :name="keyword.name"
          :description="keyword.description"
          @click="selectKeyword(keyword)"
        />
      </div>
    </div>

    <!-- 右側：詳細表示 -->
    <div class="p-5 bg-white-100 dark:bg-gray-800 border-l border-gray-300 dark:border-gray-700">
      <div v-if="selectedKeyword" class="space-y-4">
        <h2 class="text-2xl font-bold text-gray-900 dark:text-white">{{ selectedKeyword.name }}</h2>
        <p class="text-gray-700 dark:text-gray-300">{{ selectedKeyword.description }}</p>
        <button
          class="px-4 py-2 bg-red-500 text-white rounded hover:bg-red-600"
          @click="clearSelection"
        >
          閉じる
        </button>
      </div>
      <div v-else class="text-gray-500 dark:text-gray-400 text-center">
        キーワードを選択すると、ここに詳細が表示されます。
      </div>
    </div>
  </div>
</template>
