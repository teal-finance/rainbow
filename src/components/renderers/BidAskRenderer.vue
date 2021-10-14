<template>
  <div>{{ val.size }}</div>
</template>

<script lang="ts">
import { defineComponent, ref, toRefs } from "vue";
import { guidGenerator } from "@/datatable/utils";
import Bid from "@/models/bid";
import Ask from "@/models/ask";

export default defineComponent({
  props: {
    k: {
      type: String,
      required: true,
    },
    v: {
      type: Object as () => Set<Bid> | Set<Ask>,
      required: true,
    },
  },
  emits: ["update:v"],
  setup(props, { emit }) {
    const { k, v } = toRefs(props);

    const val = ref<Set<Bid> | Set<Ask>>(v.value);

    const id = guidGenerator();

    function onChange(event: boolean) {
      console.log("Change", val.value, event)
      emit("update:v", { k: k.value, v: val.value });
    }

    return {
      id,
      onChange,
      val
    };
  },
});
</script>
