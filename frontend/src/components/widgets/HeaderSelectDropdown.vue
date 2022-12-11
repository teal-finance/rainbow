<template>
  <div class="relative">
    <button ref="btn" class="h-10 secondary btn " @click="toggle()">{{ name }}</button>
    <div ref="target" :class="isCollapsed === true ? ['slideup'] : ['slidedown']"
      class="absolute duration-200 top-10 slide-y bord-light lighter dark:primary min-w-[10em] rounded-b-xl rounded-tr-md">
      <div class="p-3">
        <slot></slot>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { onClickOutside } from '@vueuse/core';

defineProps({
  name: {
    type: String,
    required: true
  },
  isCollapsed: {
    type: Boolean,
    default: true
  }
});

const emit = defineEmits(['close', 'toggle']);

const target = ref(null);
const btn = ref(null);

onClickOutside(target, (event) => emit('close'), {
  ignore: [btn]
});

function toggle() {
  emit('toggle')
}
</script>