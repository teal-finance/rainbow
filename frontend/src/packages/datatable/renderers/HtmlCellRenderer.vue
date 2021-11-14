<template>
  <div v-html="v"></div>
</template>

<script lang="ts">
import { defineComponent } from "vue";

export default defineComponent({
  props: {
    k: {
      type: String,
      required: true,
    },
    v: {
      required: true,
    },
  },
  setup() {
    return {};
  },
});
</script>
