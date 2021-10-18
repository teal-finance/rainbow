<template>
  <div :class="{ 'dark': user.isDarkMode.value == true }">
    <div
      class="w-full bg-background dark:bg-background-dark text-foreground dark:text-foreground-dark"
    >
      <div
        class="flex items-center h-20 bg-primary dark:bg-primary-dark text-primary-r dark:text-primary-r-dark"
      >
        <div class="flex items-center flex-grow">
          <div class="inline-block mx-3">
            <img alt="logo" src="./assets/logo-transparent.png" height="68" width="68" />
          </div>
          <div class="inline-block text-xl tracking-widest">
            <!-- img alt="Rainbow" src="./assets/rainbow-chancery.png" height="49" width="185" / -->
            Rainbow
          </div>
        </div>
        <div>
          <div class="mr-5 text-lg cursor-pointer" @click="user.toggleDarkMode()">
            <i-fa-solid-moon v-if="user.isDarkMode.value == false"></i-fa-solid-moon>
            <i-fa-solid-sun v-else></i-fa-solid-sun>
          </div>
        </div>
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
import { user } from "@/state";

export default defineComponent({
  components: {
    RainbowDatatable,
  },
  setup() {
    const datatable = ref(new SwDataTableModel<Option>());

    async function fetchData() {
      const data = await api.get<Array<Record<string, string | number | Array<Record<string, string | number>>>>>("/v0/options");
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

<style lang="css">
html,
body {
  margin: 0;
  padding: 0;
}
</style>
