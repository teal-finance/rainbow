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
  tr:nth-child(2)
    th:nth-child(2), th:nth-child(5)
      @apply bg-primary dark:bg-primaryTable text-primary-r p-3 rounded-tl-lg
    th:nth-child(3), th:nth-child(6)
      @apply dark:bg-primaryTable-dark text-primary-r p-3 rounded-tr-lg
  tr:nth-child(3)
    th
      @apply bg-primaryTable dark:bg-primaryTable-dark text-primary-r
    th:nth-child(4), th:nth-child(5), th:nth-child(9), th:nth-child(10)
      @apply bg-primary bg-opacity-10 text-primary-dark p-3 dark:text-primary-r
    th:nth-child(6), th:nth-child(7), th:nth-child(11), th:nth-child(12)
      @apply bg-primaryTable dark:bg-primaryTable-dark dark:bg-opacity-60 bg-opacity-10 text-primary-dark dark:text-primary-r
  tbody
    @apply xl:bg-background xl:dark:bg-background-dark
    td
      @apply px-3 py-1
      &#provider, &#asset, &expiry
        @apply bg-background dark:bg-background-dark
      &#strike
        @apply bg-background dark:bg-background-dark text-left
      &#callBidSize, &#putBidSize, &#callBidPrice, &#putBidPrice
        @apply bg-primary bg-opacity-10 dark:text-primary-r
      &#callAskSize, &#putAskSize, &#callAskPrice, &#putAskPrice
        @apply bg-primaryTable dark:bg-primaryTable-dark bg-opacity-10 dark:text-primary-r dark:bg-opacity-60
      &#callBidPrice, &#putBidPrice
        @apply text-success dark:text-success-dark font-bold text-right
      &#callAskPrice, &#putAskPrice
        @apply text-danger dark:text-danger-dark font-bold text-left
    tr
      @apply border-gray-200 border-b dark:border-gray-700
</style>
