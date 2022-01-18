<template>
  <div v-for="(v,i) in values" :key="i" class="flex flex-row items-center mb-3">
    <sw-switch v-model:value="state[`${v}`]" @change="toggleActivate(v)" class="mr-2 primary"></sw-switch>
    <div v-html="v"></div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs, watchEffect } from "vue";
import SwSwitch from "@snowind/switch";

export default defineComponent({
  components: {
    SwSwitch
  },
  props: {
    col: {
      type: String,
      required: true,
    },
    values: {
      type: Set,
      required: true,
    },
  },
  emits: ["include", "exclude"],
  setup(props, { emit }) {
    const { values } = toRefs(props);
    let state = reactive<Record<any, boolean>>({});

    function toggleActivate(v: any) {
      if (state[v]) {
        // deactivate
        state[v] = false;
        emit("exclude", v)
        return
      } else {
        // activate
        state[`${v}`] = true
      }
      emit("include", v);
    }

    function _setState() {
      for (const v of values.value) {
        state[`${v}`] = true
      }
    }

    watchEffect(() => {
      _setState()
    });

    return {
      state,
      toggleActivate
    };
  },
});
</script>
