<template>
  <div v-if="isReady">
    <div class="flex flex-row" v-if="!isMobile">
      <div class="w-4/5">
        <options-datatable :model="datatable"></options-datatable>
      </div>
      <div class="pl-5 pr-3">
        <div>
          <div class="text-xl">Asset</div>
          <hr class="my-3" />
          <values-filter :model="datatable" col="asset"></values-filter>
        </div>
        <div class="mt-2">
          <div class="text-xl">Provider</div>
          <hr class="my-3" />
          <values-filter :model="datatable" col="provider"></values-filter>
        </div>
      </div>
    </div>
    <div v-else>
      <options-datatable :model="datatable"></options-datatable>
    </div>
  </div>
  <div v-else>
    <loading-indicator></loading-indicator>
  </div>
</template>


<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue'
import api from '@/api';
import Option from '@/models/options/option';
import SwDataTableModel from "@/packages/datatable/models/datatable";
import ValuesFilter from '@/packages/datatable/filters/ValuesFilter.vue'
import { OptionsJsonDataset, OptionsTable } from '@/models/options/types';
import OptionsDatatable from '@/components/OptionsDatatable.vue';
import LoadingIndicator from '@/components/widgets/LoadingIndicator.vue';
import { isMobile } from '@/state';

export default defineComponent({
  components: {
    OptionsDatatable,
    ValuesFilter,
    LoadingIndicator,
  },
  setup() {
    const datatable = ref(new SwDataTableModel<OptionsTable>());
    const isReady = ref(false);

    async function fetchData() {
      const uri = "v0/options/cp"
      const data = await api.get<OptionsJsonDataset>(uri);
      //console.log("DATA", data)
      return data;
    }

    function loadData(dataset: OptionsJsonDataset) {
      // console.log("DATA", dataset)
      const options = new Set<OptionsTable>();
      for (const line of dataset.rows) {
        const opt = new Option(line).toRow();
        options.add(opt)
      }
      console.log("OPTIONS", options);
      const columns = {
        "provider": "Provider",
        "asset": "Asset",
        "expiry": "Expiry",
        "callBidSize": "Size",
        "callBidPrice": "  BID",
        "callAskPrice": " ASK",
        "callAskSize": "Size",
        "strike": "Strike",
        "putBidSize": "Size",
        "putBidPrice": "  BID",
        "putAskPrice": " ASK",
        "putAskSize": "Size",
      }
      datatable.value = new SwDataTableModel<OptionsTable>({ columns: columns, rows: Array.from(options) });
      //datatable.value.setColumnsFromData();
    }

    onMounted(() => {
      fetchData().then((d) => {
        loadData(d);
        isReady.value = true;
      });
      //loadData(data)
    })

    return {
      isReady,
      datatable,
      isMobile,
    }
  }
})
</script>