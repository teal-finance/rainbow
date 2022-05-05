<template>
  <small-badge
    v-for="[k, v] in Object.entries(values)"
    :key="k"
    :text="`${k}`"
    :is-active="v === true"
    @click="toggleActivate(k, !v)"
  ></small-badge>
</template>

<script setup lang="ts">
import SmallBadge from "../widgets/SmallBadge.vue";

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
