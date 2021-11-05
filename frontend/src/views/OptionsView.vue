<template>
  <options-datatable :model="datatable"></options-datatable>
</template>


<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue'
import api from '@/api';
import Option from '@/models/options/option';
import SwDataTableModel from "@/packages/datatable/models/datatable";
import { user } from "@/state";
import { OptionsJsonDataset, OptionsTable } from '@/models/options/types';
import OptionsDatatable from '@/components/OptionsDatatable.vue';

export default defineComponent({
  components: {
    OptionsDatatable
  },
  setup() {
    const datatable = ref(new SwDataTableModel<OptionsTable>());

    async function fetchData() {
      const uri = "/mocks/data/options_table.json"
      const data = await api.get<OptionsJsonDataset>(uri, true);
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
      //console.log("OPTIONS", options);
      const columns = {
        "provider": "Provider",
        "asset": "Asset",
        "expiry": "Expiry",
        "putBidPrice": "Bid",
        "putBidSize": "Bid size",
        "putAskPrice": "Ask",
        "putAskSize": "Ask size",
        "strike": "strike",
        "callBidPrice": "Bid",
        "callBidSize": "Bid size",
        "callAskPrice": "Ask",
        "callAskSize": "Ask size",
      }
      datatable.value = new SwDataTableModel<OptionsTable>({ columns: columns, rows: Array.from(options) });
      //datatable.value.setColumnsFromData();
    }

    onMounted(() => {
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