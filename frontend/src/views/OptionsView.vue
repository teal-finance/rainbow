<template>
  <options-datatable :model="datatable"></options-datatable>
</template>


<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue'
import api from '@/api';
import Option from '@/models/options/option';
import SwDataTableModel from "@/packages/datatable/models/datatable";
import { user } from "@/state";
import { OptionsJsonDataset } from '@/models/options/types';
import OptionsDatatable from '@/components/OptionsDatatable.vue';

export default defineComponent({
  components: {
    OptionsDatatable
  },
  setup() {
    const datatable = ref(new SwDataTableModel<Record<string, number | string>>());

    async function fetchData() {
      const uri = "/mocks/data/options.json"
      const data = await api.get<OptionsJsonDataset>(uri, true);
      //console.log("DATA", data)
      return data;
    }

    function loadData(dataset: OptionsJsonDataset) {
      // console.log("DATA", dataset)
      const options = new Set<Record<string, number | string>>();
      for (const line of dataset.table) {
        const opt = new Option(line).toRow();
        options.add(opt)
      }
      console.log("OPTIONS", options);
      const columns = {
        "provider": "Provider",
        "putBid": "Bid",
        "putAsk": "Ask",
        "putIv": "Iv",
        "putLastPrice": "Last price",
        "putChange": "Change",
        "putOpen": "Open",
        "putVolume": "Volume",
        "strike": "strike",
        "callBid": "Bid",
        "callAsk": "Ask",
        "callIv": "Iv",
        "callLastPrice": "Last price",
        "callChange": "Change",
        "callOpen": "Open",
        "callVolume": "Volume",
      }
      datatable.value = new SwDataTableModel<Record<string, number | string>>({ columns: columns, rows: Array.from(options) });
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