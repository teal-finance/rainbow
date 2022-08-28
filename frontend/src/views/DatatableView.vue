<template>
  <div v-if="isReady">
    <div class="flex flex-row" v-if="!isMobile">
      <div class="w-4/5">
        <generic-datatable :model="datatable" :type="type"></generic-datatable>
      </div>
      <div class="w-1/5 pl-5 pr-3">
        <presets-select v-if="isReady && type == 'classic'" @changepreset="mutatePreset($event)"></presets-select>
        <div>
          <div class="mt-3 text-xl">Asset</div>
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
      <generic-datatable :model="datatable" :type="type"></generic-datatable>
    </div>
  </div>
  <div v-else>
    <loading-indicator></loading-indicator>
  </div>
</template>


<script setup lang="ts">
import { ref, reactive, onBeforeMount } from 'vue'
import Option from '@/models/options/option';
import SwDataTableModel from "@/packages/datatable/models/datatable";
import ValuesFilter from '@/packages/datatable/filters/ValuesFilter.vue'
import { OptionsJsonDataset, OptionsTable } from '@/models/options/types';
import GenericDatatable from '@/components/GenericDatatable.vue';
import LoadingIndicator from '@/components/widgets/LoadingIndicator.vue';
import { isMobile, user } from '@/state';
import { classicQuery, exoticQuery } from '@/api/graphql';
import filterPresets from "@/const/filter_presets";
import PresetsSelect from "@/components/PresetsSelect.vue";
import { OptionType } from '@/interfaces';
import { onBeforeRouteLeave, onBeforeRouteUpdate } from 'vue-router';
import router from "@/router";

const props = defineProps({
  type: {
    type: String as () => OptionType,
    required: true
  }
});

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

async function init(t: OptionType) {
  let d: OptionsJsonDataset;
  if (t == "classic") {
    d = await classicQuery()
    mutatePreset(user.currentPreset.value)
  } else {
    d = await exoticQuery()
  }
  loadData(d);
  isReady.value = true;
}

onBeforeMount(() => {
  console.log("Mount table", props.type)
  init("classic")
})

onBeforeRouteLeave((to, from) => {
  //console.log("Update table", props.type)
  const t = props.type;
  if (to.path.replace("/", "") != t.toString()) {
    //console.log("Route change")
    isReady.value = false;
    init(t).then(() => isReady.value = true)
  }
})
</script>

<style scoped lang="sass">
#presets-selector
  @apply block w-full bg-transparent border bord-lighter focus:outline-primary px-4 py-2 pr-8 rounded shadow leading-tight focus:outline-none
  & option
    @apply background
</style>
