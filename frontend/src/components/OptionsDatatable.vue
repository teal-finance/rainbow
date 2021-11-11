<template>
  <sw-datatable
    :model="model"
    :sortable-cols="['expiry', 'strike']"
    :extra-header="ExtraHeader"
    :mobile-renderer="MobileRenderer"
    :tablet-renderer="MobileRenderer"
    tablet-breakpoint="xl"
    class="w-full border border-collapse table-auto dark:border-neutral border-light"
    id="rtable"
  ></sw-datatable>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import SwDatatableModel from '@/packages/datatable/models/datatable'
import SwDatatable from '@/packages/datatable/SwDatatable.vue'
import { OptionsTable } from '@/models/options/types';
import ExtraHeader from './widgets/ExtraHeader.vue';
import MobileRenderer from './widgets/MobileRenderer.vue';

export default defineComponent({
  components: {
    SwDatatable,
    MobileRenderer,
  },
  props: {
    model: {
      type: Object as () => SwDatatableModel<OptionsTable>,
      required: true
    },
  },
  setup: () => {
    return {
      ExtraHeader,
      MobileRenderer,
    }
  },
})
</script>

<style lang="sass">
#rtable
  @apply border-none
  tr:nth-child(3)
    th
      @apply bg-rainbow-300 dark:bg-rainbow-700-dark
    th:nth-child(1), th:nth-child(2)
      @apply bg-light dark:bg-light-dark
    th:nth-child(4), th:nth-child(5), th:nth-child(9), th:nth-child(10)
      @apply bg-secondary text-secondary-r p-3
    th:nth-child(6), th:nth-child(7), th:nth-child(11), th:nth-child(12)
      @apply bg-rainbow-800 text-secondary-r
  tbody
    @apply xl:bg-gray-50 xl:dark:bg-gray-700
    td
      @apply px-3 py-1
    td:nth-child(1), td:nth-child(2), td:nth-child(3),td:nth-child(8)
      @apply bg-gray-100 dark:bg-gray-700
    td:nth-child(4), td:nth-child(5), td:nth-child(9), td:nth-child(10)
      @apply bg-rainbow-300 dark:text-foreground
    td:nth-child(6), td:nth-child(7), td:nth-child(11), td:nth-child(12)
      @apply bg-rainbow-400 dark:text-foreground
    td:nth-child(5), td:nth-child(10)
      @apply text-success dark:text-success-dark font-bold
    td:nth-child(6), td:nth-child(11)
      @apply text-danger dark:text-danger-dark font-bold
    td:nth-child(8)
      @apply bg-gray-400 dark:bg-gray-900
</style>
