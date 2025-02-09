<script setup lang="ts">
import { ref } from 'vue'

import TheTitle from '@/components/ui/TheTitle.vue'
import PoliticalTopicCard from '~/components/pages/topic/PoliticalTopicCard.vue'
import politicalTopics from '@/data/master-political-topics.json'

type Topic = {
  id: number
  title: string
  description: string
  detail: string
}

const topics = ref<Topic[]>(politicalTopics)
const selectedTopic = ref<Topic | null>(null)
const selectTopic = (topic: Topic) => {
  selectedTopic.value = topic
}
const clearSelection = () => {
  selectedTopic.value = null
}
</script>

<template>
  <TheTitle
    title="政治の争点を理解する"
    title-align="text-center"
  />

  <div class="grid grid-cols-2 h-screen">
    <div class="p-5">
      <div class="flex flex-wrap gap-4">
        <PoliticalTopicCard
          v-for="topic in topics"
          :key="topic.id"
          :topic="topic"
          @click="selectTopic(topic)"
        />
      </div>
    </div>

    <!-- 右側：詳細表示 -->
    <div class="p-5 bg-white-100 dark:bg-gray-800 border-l border-gray-300 dark:border-gray-700">
      <div v-if="selectedTopic" class="space-y-4">
        <h2 class="text-2xl font-bold text-gray-900 dark:text-white">{{ selectedTopic.title }}</h2>
        <p class="text-gray-700 dark:text-gray-300">{{ selectedTopic.description }}</p>
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
