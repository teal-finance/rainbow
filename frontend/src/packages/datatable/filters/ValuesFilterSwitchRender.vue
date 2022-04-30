<template>
  <div v-for="(k, i) in Object.keys(values)" :key="i" class="flex flex-row items-center mb-3">
    <sw-switch
      v-model:value="values[k]"
      class="mr-2 switch-secondary"
      @update:value="toggleActivate(k, $event)"
    >
      <div class="ml-2" v-html="k"></div>
    </sw-switch>
  </div>
</template>

<script setup lang="ts">
import SwSwitch from "@snowind/switch";

defineProps({
  col: {
    type: String,
    required: true,
  },
  values: {
    type: Object as () => Record<any, boolean>,
    required: true,
  },
});

const emit = defineEmits(["include", "exclude"])

function toggleActivate(k: any, v: boolean) {
  console.log("Toggle", k, v)
  if (v) {
    emit("include", k)
  } else {
    emit("exclude", k)
  }
}
</script>
