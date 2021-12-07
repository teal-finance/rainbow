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

  tr:nth-child(1)

    // Bid headers for CALL and PUT
    th:nth-child(2)
      @apply bg-secondary dark:bg-secondary-dark bg-opacity-20 p-1 rounded-t-lg

    // Ask headers for CALL and PUT
    th:nth-child(4)
      @apply bg-secondary dark:bg-secondary-dark bg-opacity-20 p-1 rounded-t-lg

  tr:nth-child(2)

    // Bid headers for CALL and PUT
    th:nth-child(4), th:nth-child(5), th:nth-child(9), th:nth-child(10)
      @apply bg-bid dark:bg-bid-dark bg-opacity-20 dark:bg-opacity-80

    // Ask headers for CALL and PUT
    th:nth-child(6), th:nth-child(7), th:nth-child(11), th:nth-child(12)
      @apply bg-ask dark:bg-ask-dark bg-opacity-20 dark:bg-opacity-50

    // Price Bid headers
    th:nth-child(5), th:nth-child(10)
      @apply text-bid-dark dark:text-bid

    // Price Ask headers
    th:nth-child(6), th:nth-child(11)
      @apply text-ask-dark dark:text-ask

    th
      @apply bg-secondary dark:bg-secondary-dark bg-opacity-20

  tbody
    @apply xl:bg-background xl:dark:bg-background-dark
    td
      @apply px-3 py-1
      &.provider, &.asset, &expiry
        @apply bg-background dark:bg-background-dark
      &.strike
        @apply bg-background dark:bg-background-dark text-right

      // Bid values with same background
      &.callBidSize, &.putBidSize, &.callBidPrice, &.putBidPrice
        @apply bg-bid bg-opacity-5

      // Ask values with another background
      &.callAskSize, &.putAskSize, &.callAskPrice, &.putAskPrice
        @apply bg-ask bg-opacity-5

      // Bid prices with same foreground color
      &.callBidPrice, &.putBidPrice
        @apply text-bid-dark dark:text-bid text-right

      // Ask prices with another foreground color
      &.callAskPrice, &.putAskPrice
        @apply text-ask-dark dark:text-ask text-right

    tr
      @apply border-gray-200 border-b dark:border-gray-700
</style>
