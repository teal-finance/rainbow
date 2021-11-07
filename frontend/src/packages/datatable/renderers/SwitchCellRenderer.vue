<template>
  <sw-switch :id="id" :checked="v" @change="onChange($event)" v-model:value="val"></sw-switch>
</template>

<script lang="ts">
import { defineComponent, ref, toRefs } from "vue";
import { guidGenerator } from "../utils";
import SwSwitch from "@snowind/switch";

export default defineComponent({
  components: {
    SwSwitch
  },
  props: {
    k: {
      type: String,
      required: true,
    },
    v: {
      type: Boolean,
      required: true,
    },
  },
  emits: ["update:v"],
  setup(props, { emit }) {
    const { k, v } = toRefs(props);

    const val = ref<boolean>(v.value);

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
