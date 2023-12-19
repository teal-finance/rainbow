<template>
  <div class="flex flex-col min-w-[14em]">
    <div class="mb-3 flex justify-center w-full flex-row text-sm" v-if="toggleAllButtons">
      <button class="btn rounded-none txt-light flex flex-row items-center" @click="enableAll()">
        <div class="mr-2">
          <i-mdi:checkbox-multiple-marked-outline class="text-xl"></i-mdi:checkbox-multiple-marked-outline>
        </div>
        <div>All</div>
      </button>
      <button class="btn rounded-none txt-light flex flex-row items-center" @click="clearAll()">
        <div class="mr-2">
          <i-mdi:close-box-multiple-outline class="text-xl"></i-mdi:close-box-multiple-outline>
        </div>
        <div>Clear</div>
      </button>
    </div>
    <div v-for="(k, i) in Object.keys(values)" :key="i" class="flex flex-row items-center mb-3 ml-5">
      <sw-switch :id="k" v-model:value="values[k]" class="mr-2 switch-secondary"
        @update:value="toggleActivate(k, $event)">
        <div class="ml-2" v-html="k"></div>
      </sw-switch>
    </div>
  </div>
</template>

<script setup lang="ts">
import SwSwitch from "@snowind/switch";

const props = defineProps({
  col: {
    type: String,
    required: true,
  },
  values: {
    type: Object as () => Record<any, boolean>,
    required: true,
  },
  toggleAllButtons: {
    type: Boolean,
    default: true,
  }
});

const emit = defineEmits(["include", "exclude"]);

function clearAll() {
  for (const [k, v] of Object.entries(props.values)) {
    if (v) {
      toggleActivate(k, false)
    }
  }
}

function enableAll() {
  for (const [k, v] of Object.entries(props.values)) {
    if (!v) {
      toggleActivate(k, true)
    }
  }
}

function toggleActivate(k: any, v: boolean) {
  //console.log("Toggle", k, v)
  if (v) {
    emit("include", k)
  } else {
    emit("exclude", k)
  }
}
</script>
