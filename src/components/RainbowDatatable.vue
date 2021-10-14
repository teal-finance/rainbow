<template>
  <div>
    <div>
      Providers:
      <values-filter :model="model" col="provider"></values-filter>
    </div>
    <div>
      Type:
      <values-filter :model="model" col="type"></values-filter>
    </div>
    <div>
      Providers:
      <!--small-badge
        v-for="provider in initial.providers"
        :text="provider"
        :is-active="state.providers.has(provider)"
        @click="filters.toggleProvider(provider)"
      ></small-badge-->
    </div>
    <sw-datatable
      :model="model"
      :renderers="renderers"
      class="mt-3 border border-collapse table-auto border-light"
      id="rtable"
    ></sw-datatable>
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
      "bids": BidAskRenderer,
      "asks": BidAskRenderer,
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
    @apply bg-light
  td
    @apply px-3 py-1
</style>
