<template>
  <div>
    <vega-chart id="com-pipeline" :spec="spec"></vega-chart>
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
          url: "/mocks/charts/commercial.csv",
        },
        mark: {
          type: "area",
          point: {
            filled: false,
            fill: "white",
          },
        },
        encoding: {
          x: { field: "date", type: "temporal", timeUnit: "date" },
          y: {
            field: "valeur",
            type: "quantitative",
            scale: { zero: false },
          },
          color: { field: "type", type: "nominal" },
        },
        width: 850,
        height: 300,
      },
    };
  },
});
</script>