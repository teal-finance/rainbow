<template>
  <div class="relative z-10 w-full min-h-[2rem] text-sm">
    <div class="absolute top-0 right-0 flex flex-col justify-start background">
      <div id="presets-select" class="flex flex-row w-full lighter" @click="collapse = !collapse">
        <div class="flex-grow">
          <button class="w-full text-left btn">{{ user.currentPreset.value }}</button>
        </div>
        <div class="flex items-center justify-center mr-2 text-lg">
          <carret-down v-if="collapse"></carret-down>
          <carret-up v-else></carret-up>
        </div>
      </div>
      <div :class="collapse === true ? ['slideup'] : ['slidedown']" class="duration-200 border slide-y bord-lighter">
        <div v-for="(preset, i) in Object.keys(filterPresets)" :key="i" class="w-full background hover:secondary">
          <button :id="'preset-' + preset" class="w-full text-left btn" v-if="preset != user.currentPreset.value"
            @click="onChangePreset(preset, $event)">{{ preset }}</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { user } from '@/state';
import filterPresets from "@/const/filter_presets";
import CarretDown from '@/packages/datatable/widgets/CarretDown.vue';
import CarretUp from '@/packages/datatable/widgets/CarretUp.vue';

const collapse = ref(true);

const emit = defineEmits(["changepreset"]);

function onChangePreset(presetname: string, event?: any) {
  //console.log("Change preset", presetname)
  emit("changepreset", presetname)
  if (event) {
    collapse.value = !collapse.value
    event.target.blur()
  }
}
</script>