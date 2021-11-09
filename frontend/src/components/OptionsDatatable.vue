<template>
  <sw-datatable
    :model="model"
    :sortable-cols="Object.keys(model.state.columns)"
    class="w-full border border-collapse table-auto dark:border-neutral border-light"
    id="rtable"
  ></sw-datatable>
</template>

<script lang="ts">
import { defineComponent, toRef, onMounted } from 'vue'
import SwDatatableModel from '@/packages/datatable/models/datatable'
import SwDatatable from '@/packages/datatable/SwDatatable.vue'
import { OptionsTable } from '@/models/options/types'

export default defineComponent({
  components: {
    SwDatatable,
  },
  props: {
    model: {
      type: Object as () => SwDatatableModel<OptionsTable>,
      required: true
    },
  },
  setup: (props) => {
    const model = toRef(props, "model");

    onMounted(() => {
      model.value.addSortFilterDesc("strike")
    })
  },
})
</script>

<style lang="sass">
#rtable
  @apply border-none
  th
    @apply bg-light text-light-r dark:bg-gray-500 dark:text-white
  tbody
    @apply bg-gray-50 dark:bg-gray-700
    td
      @apply px-3 py-1
    td:nth-child(1), td:nth-child(2), td:nth-child(3)
      @apply bg-gray-300 dark:bg-gray-600
    td:nth-child(8)
      @apply bg-gray-400 dark:bg-gray-900
</style>
