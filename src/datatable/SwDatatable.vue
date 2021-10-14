<template>
  <table>
    <thead>
      <tr>
        <th
          v-for="(colslug, i) in Object.keys(model.state.columns)"
          :key="i"
          v-html="model.state.columns[colslug]"
        ></th>
      </tr>
    </thead>
    <tbody>
      <tr v-for="(row, i) in model.state.rows" :key="i.toString()">
        <td v-for="(cell, ii) in Object.keys(model.state.columns)" :key="ii">
          <component
            :key="i.toString() + cell"
            :is="getRenderer(cell)"
            :k="cell"
            :v="row[cell]"
            @update:v="onChange($event, parseInt(i.toString()))"
          ></component>
        </td>
      </tr>
    </tbody>
  </table>
</template>

<script lang="ts">
import { defineComponent, toRefs } from "vue";
import SwDatatableModel from "./models/datatable";
import DefaultCellRenderer from "./renderers/DefaultCellRenderer.vue";
import Option from "@/models/option";

export default defineComponent({
  props: {
    model: {
      type: Object as () => SwDatatableModel<Option>,
      required: true,
    },
    renderers: {
      type: Object,
      default: () => { }, // eslint-disable-line
    },
  },
  setup(props) {
    const { model, renderers } = toRefs(props);

    function getRenderer(k: string) {
      if (renderers.value !== undefined) {
        const rendererNames = new Set<string>(Object.keys(renderers.value));
        if (rendererNames.has(k)) {
          return renderers.value[k];
        }
      }
      return DefaultCellRenderer;
    }

    // eslint-disable-next-line
    function onChange(event: Record<string, any>, i: number) {
      console.log("ON CHG", event)
      model.value.state.rows[i][event.k] = event.v;
    }

    return {
      getRenderer,
      onChange,
    };
  },
});
</script>