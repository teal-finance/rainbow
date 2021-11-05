<template>
  <table>
    <thead v-if="!isSmallScreen">
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
        <td
          v-for="(cell, ii) in Object.keys(model.state.columns)"
          :key="ii"
          v-if="!isSmallScreen"
          @click="onClick(row, cell)"
        >
          <component
            :key="i.toString() + cell"
            :is="getRenderer(cell)"
            :k="cell"
            :v="row[cell]"
            @update:v="onChange($event, parseInt(i.toString()))"
          ></component>
        </td>
        <td v-else :colspan="model.numCols">
          <component
            :key="i.toString() + 'm'"
            :is="mobileRenderer"
            :row="row"
            @update:v="onChange($event, parseInt(i.toString()))"
          ></component>
        </td>
      </tr>
    </tbody>
  </table>
</template>

<script lang="ts">
import { defineComponent, toRefs, ref, toRaw } from "vue";
import { useBreakpoints, breakpointsTailwind } from '@vueuse/core';
import SwDatatableModel from "./models/datatable";
import DefaultCellRenderer from "./renderers/DefaultCellRenderer.vue";

export default defineComponent({
  props: {
    model: {
      type: Object as () => SwDatatableModel,
      required: true,
    },
    renderers: {
      type: Object,
      default: () => { }, // eslint-disable-line
    },
    mobileRenderer: {
      type: Object,
      default: () => null, // eslint-disable-line
    },
    breakpoint: {
      type: String as () => "sm" | "md" | "lg" | "xl" | "2xl",
      default: () => "sm",
    },
  },
  emits: ["onClickCell"],
  setup(props, { emit }) {
    const breakpoints = useBreakpoints(breakpointsTailwind)
    const { model, renderers, mobileRenderer, breakpoint } = toRefs(props);
    const isResponsive = ref(mobileRenderer.value !== null);
    let isSmallScreen = ref(false);
    if (isResponsive.value === true) {
      //console.log("BP", breakpoint.value);
      isSmallScreen.value = breakpoints.isSmaller(breakpoint.value);
    }


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

    function onClick(row: any, cell: string) {
      if (model.value.idCol) {
        emit("onClickCell", { rowId: toRaw(row)[model.value.idCol], cell: cell })
      }
      else {
        throw new Error("Please provide an idCol in model")
      }
    }

    return {
      getRenderer,
      onChange,
      onClick,
      isResponsive,
      isSmallScreen,
    };
  },
});
</script>