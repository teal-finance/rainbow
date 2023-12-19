<template>
  <div>
    <div v-if="isDatatableReady" class="w-full">
      <generic-datatable :model="datatable" :type="type"></generic-datatable>
    </div>
    <div v-else>
      <loading-indicator></loading-indicator>
    </div>
  </div>
</template>


<script setup lang="ts">
import { onBeforeMount, onBeforeUnmount } from 'vue'
import Option from '@/models/options/option';
import SwDataTableModel from "@/packages/datatable/models/datatable";
import { OptionsJsonDataset, OptionsTable } from '@/models/options/types';
import GenericDatatable from '@/components/GenericDatatable.vue';
import LoadingIndicator from '@/components/widgets/LoadingIndicator.vue';
import { user, datatable, isDatatableReady, filterConf } from '@/state';
import { classicQuery, exoticQuery } from '@/api/graphql';
import filterPresets from "@/const/filter_presets";
import { OptionType } from '@/interfaces';

const props = defineProps({
  type: {
    type: String as () => OptionType,
    required: true
  }
});

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
  user.currentPreset.value = presetname;
}

async function init() {
  let d: OptionsJsonDataset;
  if (props.type == "classic") {
    d = await classicQuery()
    mutatePreset(user.currentPreset.value)
  } else {
    d = await exoticQuery()
  }
  loadData(d);
  isDatatableReady.value = true;
}

onBeforeMount(() => {
  init()
});

onBeforeUnmount(() => isDatatableReady.value = false)
</script>

<style scoped lang="sass">
#presets-selector
  @apply block w-full bg-transparent border bord-lighter focus:outline-primary px-4 py-2 pr-8 rounded shadow leading-tight focus:outline-none
  & option
    @apply background
</style>
