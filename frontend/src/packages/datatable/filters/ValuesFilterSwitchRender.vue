<template>
  <div
    v-for="[k, v] in Object.entries(state)"
    :key="k.toString()"
    class="flex flex-row items-center mb-3"
  >
    {{ k }} / {{ v }}
    <sw-switch :checked="state[`${v}`]" class="mr-2 switch-secondary" @change="toggleActivate(v)">
      <div class="ml-2" v-html="state[`${k}`]"></div>
    </sw-switch>
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
      type: Object as () => Record<any, boolean>,
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
      for (const [k, v] of Object.entries(values.value)) {
        state[`${k}`] = v
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
