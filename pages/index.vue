<script setup lang="ts">
import PoliticalTopicCard from '@/components/PoliticalTopicCard.vue'

const json = [
  {
    title: '経済',
    description: '経済に関する情報です',
  },
  {
    title: '財政・税',
    description: '財政・税に関する情報です',
  },
  {
    title: '行政・政治改革',
    description: '行政・政治改革に関する情報です',
  },
  {
    title: '働き方',
    description: '働き方に関する情報です',
  },
  {
    title: '子育て・教育',
    description: '子育てに関する情報です',
  },
  {
    title: '社会保障',
    description: '社会保障に関する情報です',
  },
  {
    title: '教育',
    description: '教育に関する情報です',
  },
  {
    title: '医療',
    description: '医療に関する情報です',
  },
]

</script>

<template>
  <div v-for="item in json" :key="item.title">
    <PoliticalTopicCard :title="item.title" :description="item.description" />
  </div>
</template>
