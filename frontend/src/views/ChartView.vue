<template>
  <div class="flex flex-col space-y-5">
    <commercial-chart></commercial-chart>
    <factures-chart></factures-chart>
    <ventes-chart></ventes-chart>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import VentesChart from '@/components/charts/VentesChart.vue'
import CommercialChart from '@/components/charts/CommercialChart.vue'
import FacturesChart from '@/components/charts/FacturesChart.vue'

export default defineComponent({
  components: {
    VentesChart,
    CommercialChart,
    FacturesChart,
  },
})
</script>