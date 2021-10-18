<template>
  <Popper :hover="true" arrow placement="right">
    <button>{{ v.size }}</button>
    <template #content="{ close }">
      <div class="p-3 bg-neutral text-neutral-r dark:bg-neutral-dark dark:text-neutral-r-dark">
        <div
          v-for="(line,i) in v"
          :key="i"
        >Quote: {{ line.quoteCurrency }} Price: {{ line.price }} Quantity {{ line.quantity }}</div>
      </div>
    </template>
  </Popper>
</template>

<script>
import { defineComponent, ref, toRefs } from "vue";
import Popper from "vue3-popper";
import { guidGenerator } from "@/datatable/utils";

export default defineComponent({
  components: {
    Popper,
  },
  props: {
    k: {
      type: String,
      required: true,
    },
    v: {
      type: Set,
      required: true,
    },
  },
  emits: ["update:v"],
  setup(props, { emit }) {
    const { k, v } = toRefs(props);

    const val = ref(v.value);

    const id = guidGenerator();

    function onChange(event) {
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
