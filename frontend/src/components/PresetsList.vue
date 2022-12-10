<template>
  <div>
    <div v-for="(preset, i) in Object.keys(filterPresets)" :key="i" class="w-full background hover:secondary">
      <button :id="'preset-' + preset" class="w-full text-left btn"
        :class="preset == user.currentPreset.value ? 'primary dark:success rounded-none' : ''"
        @click="onChangePreset(preset, $event)">{{ preset }}</button>
    </div>
  </div>
</template>

<script setup lang="ts">
import filterPresets from "@/const/filter_presets";
import { user, filterConf } from '@/state';

const emit = defineEmits(['select']);

function onChangePreset(presetname: string, event?: any) {
  emit('select')
  const preset = filterPresets[presetname];
  filterConf.assets = preset.assets;
  filterConf.providers = preset.providers;
  user.currentPreset.value = presetname;
}
</script>