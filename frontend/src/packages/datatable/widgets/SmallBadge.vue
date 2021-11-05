<template>
  <button
    class="inline-block px-2 py-1 mb-2 mr-1 text-xs font-bold rounded-lg btn primary"
    :class="isActive ? '' : 'light'"
    v-html="text"
  ></button>
</template>

<script lang="ts">
import { defineComponent } from 'vue'

export default defineComponent({
  props: {
    text: {
      type: String,
      required: true
    },
    isActive: {
      type: Boolean,
      default: true
    }
  }
});
</script>