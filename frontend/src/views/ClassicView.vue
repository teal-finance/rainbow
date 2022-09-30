<template>
  <datatable-view type="classic" :preset="_preset"></datatable-view>
</template>

<script setup lang="ts">
import { onBeforeMount, ref } from 'vue';
import DatatableView from './DatatableView.vue';
import router from "@/router";

const _preset = ref("");

onBeforeMount(() => {
  const path = router.currentRoute.value.path;
  const preset = path.replace("/classic", "").replaceAll("/", "");
  if (preset.length > 0) {
    _preset.value = decodeURIComponent(preset)
  }
})
</script>