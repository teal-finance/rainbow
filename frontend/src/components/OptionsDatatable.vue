<template>
  <sw-datatable
    :model="model"
    :sortable-cols="['expiry', 'strike']"
    :extra-header="ExtraHeader"
    :mobile-renderer="MobileRenderer"
    :tablet-renderer="MobileRenderer"
    :default-renderer="HtmlCellRenderer"
    tablet-breakpoint="xl"
    class="w-full border border-collapse table-auto dark:border-neutral border-light"
    id="rtable"
  ></sw-datatable>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import SwDatatableModel from '@/packages/datatable/models/datatable'
import SwDatatable from '@/packages/datatable/SwDatatable.vue'
import HtmlCellRenderer from '@/packages/datatable/renderers/HtmlCellRenderer.vue';
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
      HtmlCellRenderer,
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
    th:nth-child(4), th:nth-child(5), th:nth-child(9), th:nth-child(10)
      @apply bg-secondary text-secondary-r p-3
    th:nth-child(6), th:nth-child(7), th:nth-child(11), th:nth-child(12)
      @apply bg-rainbow-800 text-secondary-r
  tbody
    @apply xl:bg-gray-50 xl:dark:bg-gray-700
    td
      @apply px-3 py-1
      &#provider, &#asset, &expiry
        @apply bg-gray-100 dark:bg-gray-700
      &#strike
        @apply bg-gray-200 dark:bg-gray-900 text-right
      &#callBidSize, &#putBidSize, &#callBidPrice, &#putBidPrice
        @apply bg-rainbow-400 dark:text-foreground
      &#callAskSize, &#putAskSize, &#callAskPrice, &#putAskPrice
        @apply bg-rainbow-300 dark:text-foreground
      &#callBidPrice, &#putBidPrice
        @apply text-success dark:text-success-dark font-bold text-right
      &#callAskPrice, &#putAskPrice
        @apply text-danger dark:text-danger-dark font-bold text-right
</style>
