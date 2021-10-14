<template>
  <div :class="{ 'dark': isDarkMode }">
    <div
      class="w-full bg-background dark:bg-background-dark text-foreground dark:text-foreground-dark"
    >
      <div class="h-20 bg-primary dark:bg-primary-dark text-primary-r dark:text-primary-r-dark">
          <img alt="logo" src="/public/favicon.png" height="16" />
          <img alt="Rainbow" src="/public/rainbow-chancery.png" height="16" />
      </div>
      <div class="p-5">
        <rainbow-datatable :model="datatable" v-if="datatable.hasData"></rainbow-datatable>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue'
import RainbowDatatable from '@/components/RainbowDatatable.vue'
import api from '@/api';
import Option from './models/option';
import { OptionData } from './types';
import SwDataTableModel from "@/datatable/models/datatable";

export default defineComponent({
  components: {
    RainbowDatatable,
  },
  setup() {
    const isDarkMode = ref(false);
    const datatable = ref(new SwDataTableModel<Option>());

    async function fetchData() {
      const data = await api.get<Array<Record<string, string | number | Array<Record<string, string | number>>>>>("/v0/options");
      return data;
    }

    function loadData(dataset: Array<Record<string, string | number | Array<Record<string, string | number>>>>) {
      const options = new Set<Option>();
      for (const line of dataset) {
        options.add(new Option(line as OptionData))
      }
      //console.log("OPTS", options);
      const columns = {
        "provider": "Provider",
        "asset": "Asset",
        "type": "Type",
        "bids": "Bids size",
        "strike": "Strike",
        "asks": "Asks size",
        "chain": "Chain",
        "name": "Instrument",

      }
      datatable.value = new SwDataTableModel<Option>({ columns: columns, rows: Array.from(options) });
      //datatable.value.setColumnsFromData();
      //console.log(datatable.value.state.columns)
    }

    onMounted(() => {
      fetchData().then((d) => loadData(d));
    })

    return {
      isDarkMode,
      datatable,
    }
  }
})
</script>

<style lang="css">
html,
body {
  margin: 0;
  padding: 0;
}
</style>
