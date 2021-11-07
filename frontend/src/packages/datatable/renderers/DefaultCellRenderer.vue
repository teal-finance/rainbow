<template>{{ v }}</template>

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