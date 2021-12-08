<template>
  <div>Landing page in construction</div>
  <button class="mt-5 btn" @click="$router.push('/options')">Dashboard page</button>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import api from '@/api';
import OptionsDatatable from '@/components/OptionsDatatable.vue';

export default defineComponent({
  components: {
    OptionsDatatable
  },
  setup() {
    async function fetchData() {
      const uri = "v0/options/cp"
      const data = await api.get<Record<string, any>>(uri);
      console.log("DATA", data)
      return data;
    }

    return {
      fetchData
    }
  }
})
</script>
