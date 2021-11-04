<template>
  <div>
    <div :id="id"></div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import { Config, TopLevelSpec, compile } from "vega-lite";
import embed from "vega-embed";

export default defineComponent({
  props: {
    id: {
      type: String,
      required: true,
    },
    spec: {
      type: Object,
      required: true,
    },
    config: {
      type: Object,
      default: () => {}, // eslint-disable-line
    },
  },
  methods: {
    async draw() {
      const config: Config = this.config;
      const spec = this.spec as TopLevelSpec;
      spec.$schema = "https://vega.github.io/schema/vega-lite/v5.json";
      const vegaSpec = compile(spec, { config }).spec;
      await embed(`#${this.id}`, vegaSpec, { actions: false });
    },
  },
  mounted() {
    this.draw();
  },
});
</script>