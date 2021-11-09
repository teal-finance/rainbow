<template>
  <sw-datatable
    :model="model"
    :sortable-cols="Object.keys(model.state.columns)"
    :extra-header="ExtraHeader"
    :mobile-renderer="MobileRenderer"
    class="w-full border border-collapse table-auto dark:border-neutral border-light"
    id="rtable"
  ></sw-datatable>
</template>

<script lang="ts">
import { defineComponent, toRef, onMounted } from 'vue'
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
  setup: (props) => {
    const model = toRef(props, "model");

    onMounted(() => {
      //model.value.addSortFilterDesc("strike")
    });

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
  th
    @apply sm:bg-light sm:text-light-r sm:dark:bg-gray-500 sm:dark:text-white
  tbody
    @apply sm:bg-gray-50 sm:dark:bg-gray-700
    td
      @apply px-3 py-1
    td:nth-child(1), td:nth-child(2), td:nth-child(3)
      @apply sm:bg-gray-300 sm:dark:bg-gray-600
    td:nth-child(8)
      @apply sm:bg-gray-400 sm:dark:bg-gray-900
</style>
