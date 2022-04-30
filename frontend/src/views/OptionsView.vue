<template>
  <div v-if="isReady">
    <div class="flex flex-row" v-if="!isMobile">
      <div class="w-4/5">
        <options-datatable :model="datatable"></options-datatable>
      </div>
      <div class="pl-5 pr-3">
        <div>
          <select class="form-select px-4 py-3">
            <option
              v-for="preset in Object.keys(filterPresets)"
              v-html="preset"
              @click="mutatePreset(preset)"
            ></option>
          </select>
        </div>
        <div>
          <div class="text-xl mt-3">Asset</div>
          <hr class="my-3" />
          <values-filter :model="datatable" col="asset" :filters-state="filterConf.assets"></values-filter>
        </div>
        <div class="mt-2">
          <div class="text-xl">Provider</div>
          <hr class="my-3" />
          <values-filter :model="datatable" col="provider" :filters-state="filterConf.providers"></values-filter>
        </div>
        <div class="mt-2">
          <div class="text-xl">Expiry</div>
          <hr class="my-3" />
          <values-filter :model="datatable" col="expiry"></values-filter>
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


<script setup lang="ts">
import { ref, onMounted, reactive } from 'vue'
import Option from '@/models/options/option';
import SwDataTableModel from "@/packages/datatable/models/datatable";
import ValuesFilter from '@/packages/datatable/filters/ValuesFilter.vue'
import { OptionsJsonDataset, OptionsTable } from '@/models/options/types';
import OptionsDatatable from '@/components/OptionsDatatable.vue';
import LoadingIndicator from '@/components/widgets/LoadingIndicator.vue';
import { isMobile } from '@/state';
import { query } from '@/api/graphql';
import ValuesFilterBadgeRender from '@/packages/datatable/filters/ValuesFilterBadgeRender.vue';
import filterPresets from "@/const/filter_presets";

const datatable = ref(new SwDataTableModel<OptionsTable>());
const isReady = ref(false);
const filterConf = reactive<Record<string, Record<string, boolean>>>({
  assets: { 'defaultValue': true },
  providers: { 'defaultValue': true },
});

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
}

function mutatePreset(presetname: string) {
  const preset = filterPresets[presetname];
  filterConf.assets = preset.assets;
  filterConf.providers = preset.providers;
}

onMounted(() => {
  console.log("FConf", JSON.stringify(filterConf.assets))
  query().then((d) => {
    loadData(d);
    isReady.value = true;
  });
})
</script>