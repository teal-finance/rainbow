<template>
  <div v-if="isReady">
    <div class="flex flex-row" v-if="!isMobile">
      <div class="w-4/5">
        <options-datatable :model="datatable"></options-datatable>
      </div>
      <div class="pl-5 pr-3 w-1/5">
        <presets-select v-if="isReady" @changepreset="mutatePreset($event)"></presets-select>
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
import { isMobile, user } from '@/state';
import { query } from '@/api/graphql';
//import ValuesFilterBadgeRender from '@/packages/datatable/filters/ValuesFilterBadgeRender.vue';
import filterPresets from "@/const/filter_presets";
import PresetsSelect from "@/components/PresetsSelect.vue";

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
    "callBidIv": "IV",
    "callBidSize": "Size",
    "callBidPrice": "  BID",
    "callAskPrice": " ASK",
    "callAskSize": "Size",
    "callAskIv":"IV",
    "strike": "Strike",
    "putBidIv": "IV",
    "putBidSize": "Size",
    "putBidPrice": "  BID",
    "putAskPrice": " ASK",
    "putAskSize": "Size",
    "putAskIv": "IV",
  }
  datatable.value = new SwDataTableModel<OptionsTable>({ columns: columns, rows: Array.from(options) });
}

function mutatePreset(presetname: string) {
  const preset = filterPresets[presetname];
  filterConf.assets = preset.assets;
  filterConf.providers = preset.providers;
  user.currentPreset.value = presetname;
  //console.log("Filterconf mutation", JSON.stringify(filterConf, null, "  "))
}

onMounted(() => {
  //console.log("FConf", JSON.stringify(filterConf.assets))
  query().then((d) => {
    mutatePreset(user.currentPreset.value)
    loadData(d);
    isReady.value = true;
  });
})
</script>

<style scoped lang="sass">
#presets-selector
  @apply block w-full bg-transparent border bord-lighter focus:outline-primary px-4 py-2 pr-8 rounded shadow leading-tight focus:outline-none
  & option
    @apply background
</style>
