<template>
  <div>
    <vega-chart id="factures-chart" :spec="spec"></vega-chart>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import VegaChart from "@/components/charts/VegaChart.vue";

export default defineComponent({
  components: {
    VegaChart,
  },
  data() {
    return {
      spec: {
        data: {
          url: "/mocks/charts/factures.json",
        },
        mark: "bar",
        encoding: {
          x: { field: "mois", timeUnit: "month", axis: { grid: false } },
          y: { field: "ca", type: "quantitative" },
        },
        width: 500,
        height: 200,
      },
    };
  },
});
</script>