<script setup lang="ts">
import { ref } from 'vue'

import TheTitle from '@/components/ui/TheTitle.vue'
import PoliticalPartyCard from '~/components/pages/party/PoliticalPartyCard.vue'
import politicalParty from '@/data/master-political-party.json'

type Party = {
  id: number
  name: string
}

const parties = ref<Party[]>(politicalParty)
const selectedParty = ref<Party | null>(null)
const selectParty = (party: Party) => {
  selectedParty.value = party
}
const clearSelection = () => {
  selectedParty.value = null
}
</script>

<template>
  <TheTitle
    title="政党について理解する"
    title-align="text-center"
  />
  <div class="grid grid-cols-2 h-screen">
    <div class="p-5">
      <div class="flex flex-wrap gap-4">
        <PoliticalPartyCard
          v-for="party in parties"
          :key="party.id"
          :party="party"
          @click="selectParty(party)"
        />
      </div>
    </div>

    <!-- 右側：詳細表示 -->
    <div class="p-5 bg-white-100 dark:bg-gray-800 border-l border-gray-300 dark:border-gray-700">
      <div v-if="selectedParty" class="space-y-4">
        <h2 class="text-2xl font-bold text-gray-900 dark:text-white">{{ selectedParty.name }}</h2>
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
