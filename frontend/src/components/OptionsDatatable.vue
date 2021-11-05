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

export default defineComponent({
  components: {
    SwDatatable,
  },
  props: {
    model: {
      type: Object as () => SwDatatableModel<Record<string, number | string>>,
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
  tbody
    @apply bg-gray-50 dark:bg-gray-700
</style>
