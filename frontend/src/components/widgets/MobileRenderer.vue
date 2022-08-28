<template>
  <div class="flex flex-col p-3 space-y-3 border rounded-sm shadow">
    <div class="flex flex-row">{{ row.provider }} {{ row.asset }} Expiry: {{ row.expiry }}</div>
  </div>
</template>

<script setup lang="ts">
defineProps({
  row: {
    type: Object,
    required: true,
  },
});
</script>