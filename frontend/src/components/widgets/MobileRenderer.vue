<template>
  <div class="flex flex-col p-3 space-y-3 border rounded-sm shadow">
    <div class="flex flex-row">{{ row.provider }} {{ row.asset }} Expiry: {{ row.expiry }}</div>
  </div>
</template>

<script lang="ts">
import { defineComponent, toRef } from "vue";

export default defineComponent({
  props: {
    row: {
      type: Object,
      required: true,
    },
  },
});
</script>