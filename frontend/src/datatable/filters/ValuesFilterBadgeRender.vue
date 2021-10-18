<template>
  <small-badge
    v-for="(v,i) in values"
    :key="i"
    :text="`${v}`"
    :is-active="state[`${v}`] === true"
    @click="toggleActivate(v)"
  ></small-badge>
</template>

<script lang="ts">
import { defineComponent, reactive, toRefs, watchEffect } from "vue";
import SmallBadge from "../widgets/SmallBadge.vue";

export default defineComponent({
  components: {
    SmallBadge
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
    const { col, values } = toRefs(props);
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
