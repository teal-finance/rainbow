<template>
  <div class="flex flex-row mt-3">
    <div class="w-4/5">
      <sw-datatable
        :model="model"
        :renderers="renderers"
        class="w-full border border-collapse table-auto dark:border-neutral border-light"
        id="rtable"
      ></sw-datatable>
    </div>
    <div class="pl-5 pr-3">
      <div>
        <div class="text-xl">Asset</div>
        <hr class="my-3" />
        <values-filter :model="model" col="asset"></values-filter>
      </div>
      <div class="mt-2">
        <div class="text-xl">Type</div>
        <hr class="my-3" />
        <values-filter :model="model" col="type"></values-filter>
      </div>
      <div class="mt-2">
        <div class="text-xl">Providers</div>
        <hr class="my-3" />
        <values-filter :model="model" col="provider"></values-filter>
      </div>
      <div class="mt-2">
        <div class="text-xl">Chain</div>
        <hr class="my-3" />
        <values-filter :model="model" col="chain"></values-filter>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import SwDatatableModel from '@/datatable/models/datatable'
import { defineComponent, toRefs, onMounted, reactive } from 'vue'
import SwDatatable from '@/datatable/SwDatatable.vue'
import Option from '@/models/option'
import BidAskRenderer from '@/components/renderers/BidAskRenderer.vue'
//import SmallBadge from '@/components/widgets/SmallBadge.vue'
import ValuesFilter from '@/datatable/filters/ValuesFilter.vue'

export default defineComponent({
  components: {
    SwDatatable,
    //SmallBadge,
    ValuesFilter,
  },
  props: {
    model: {
      type: Object as () => SwDatatableModel<Option>,
      required: true
    }
  },
  setup: (props) => {
    const { model } = toRefs(props);
    const renderers = {
      "bid": BidAskRenderer,
      "ask": BidAskRenderer,
    };
    const initial = reactive({
      rows: new Array<Option>(),
      providers: new Set<string>([])
    });
    const state = reactive({
      providers: new Set<string>([])
    })

    const filters = {
      toggleProvider: function toggleProvider(p: string) {
        // console.log("Toggle provider", p);
        const rows = new Array<Option>();
        if (state.providers.has(p)) {
          if (state.providers.size === 1) {
            console.log("Keep at least one filter");
            return
          }
          state.providers.delete(p);
        } else {
          state.providers.add(p)
        }
        for (const row of initial.rows) {
          if (state.providers.has(row.provider)) {
            rows.push(row)
          }
        }
        model.value.state.rows = rows;
      }
    }

    return { renderers, state, initial, filters }

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
