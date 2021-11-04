<template>
  <div>
    <vega-chart id="ventes-chart" :spec="spec"></vega-chart>
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
          url: "/mocks/charts/ventes.json",
        },
        mark: "arc",
        encoding: {
          theta: { field: "pourcentage", type: "quantitative" },
          color: { field: "transformation", type: "nominal" },
        },
        view: { stroke: null },
        width: 240,
        height: 240,
      },
    };
  },
});
</script>