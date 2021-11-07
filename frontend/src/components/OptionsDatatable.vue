<template>
  <sw-datatable
    :model="model"
    class="w-full border border-collapse table-auto dark:border-neutral border-light"
    id="rtable"
  ></sw-datatable>
</template>

<script lang="ts">
import SwDatatableModel from '@/packages/datatable/models/datatable'
import { defineComponent } from 'vue'
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
    }
  },
  setup: (props) => {
  },
})
</script>

<style lang="sass">
#rtable
  th
    @apply bg-light text-light-r dark:bg-light-dark dark:text-light-r-dark
  td
    @apply px-3 py-1
  td:nth-child(1), td:nth-child(2), td:nth-child(3)
    @apply bg-gray-300 dark:bg-gray-600
  td:nth-child(8)
    @apply bg-gray-400 dark:bg-gray-900
  tbody
    @apply bg-gray-50 dark:bg-gray-700
</style>
