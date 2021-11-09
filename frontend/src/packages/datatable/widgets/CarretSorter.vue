<template>
  <div v-if="isActive">
    <div v-if="model.sortFilter?.direction == 'desc'">
      <carret-up :class="carretClass" @click="deactivate()"></carret-up>
    </div>
    <div v-else>
      <carret-down :class="carretClass" @click="sortDesc()"></carret-down>
    </div>
  </div>
  <div v-else>
    <carret-down :class="carretClass" @click="sortAsc()"></carret-down>
  </div>
</template>

<script lang="ts">
import { defineComponent, computed, toRefs } from "vue";
import SwDatatableModel from "../models/datatable";
import CarretUp from "./CarretUp.vue";
import CarretDown from "./CarretDown.vue";

export default defineComponent({
  components: {
    CarretUp,
    CarretDown,
  },
  props: {
    col: {
      type: String,
      required: true
    },
    model: {
      type: Object as () => SwDatatableModel,
      required: true,
    },
  },
  setup(props) {
    const { col, model } = toRefs(props);

    const isActive = computed<boolean>(() => model.value.sortFilter?.col == col.value);

    const carretClass = computed<string>(() => {
      let s = "sort-neutral";
      if (isActive.value === true) {
        s = "sort-active";
      }
      return s;
    });

    const sortAsc = () => {
      model.value.addSortFilterAsc(col.value);
    }
    const sortDesc = () => {
      model.value.addSortFilterDesc(col.value)
    }
    const deactivate = () => {
      model.value.removeSortFilter()
    }

    return {
      carretClass,
      sortAsc,
      sortDesc,
      deactivate,
      isActive,
    }
  }
});
</script>

<style lang="sass" scoped>
.sort-neutral
  @apply text-neutral dark:text-neutral-dark
.sort-active
  @apply bg-success
</style>