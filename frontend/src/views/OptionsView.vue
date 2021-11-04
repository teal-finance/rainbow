<template>
  <rainbow-datatable :model="datatable" v-if="datatable.hasData"></rainbow-datatable>
</template>


<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue'
import RainbowDatatable from '@/components/RainbowDatatable.vue'
import api from '@/api';
import Option from '@/models/option';
import { OptionData } from '@/types';
import SwDataTableModel from "@/datatable/models/datatable";
import { user } from "@/state";

export default defineComponent({
  components: {
    RainbowDatatable,
  },
  setup() {
    const datatable = ref(new SwDataTableModel<Option>());

    async function fetchData() {
      const data = await api.get<Array<Record<string, string | number | Array<Record<string, string | number>>>>>("/v0/tables/default");
      return data;
    }

    function loadData(dataset: Array<Record<string, string | number | Array<Record<string, string | number>>>>) {
      // console.log("DATA", dataset)
      const options = new Set<Option>();
      for (const line of dataset) {
        const opt = new Option(line as OptionData);
        options.add(opt)
      }
      //console.log("OPTS", options);
      const columns = {
        "provider": "Provider",
        "asset": "Asset",
        "type": "Type",
        "bid": "Bid",
        "bestBidQty": "Bid size",
        "bestBidPx": "Bid price",
        "strike": "Strike",
        "ask": "Ask",
        "bestAskPx": "Ask price",
        "bestAskQty": "Ask size",
        "chain": "Chain",
        "name": "Instrument",
      }
      datatable.value = new SwDataTableModel<Option>({ columns: columns, rows: Array.from(options) });
      //datatable.value.setColumnsFromData();
      //console.log(datatable.value.state.columns)
    }

    onMounted(() => {
      //console.log("DATA", data)
      fetchData().then((d) => loadData(d));
      //loadData(data)
    })

    return {
      user,
      datatable,
    }
  }
})
</script>