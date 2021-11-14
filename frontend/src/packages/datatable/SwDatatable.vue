<template>
  <table>
    <thead v-if="!isResponsive || isDesktop">
      <component v-if="extraHeader != null" :is="extraHeader"></component>
      <tr>
        <th v-for="(colslug, i) in Object.keys(model.state.columns)" :key="i">
          {{ model.state.columns[colslug] }}
          <div
            class="inline-block align-middle cursor-pointer"
            v-if="sortableCols.includes(colslug)"
          >
            <carret-sorter :model="model" :col="colslug"></carret-sorter>
          </div>
        </th>
      </tr>
    </thead>
    <tbody v-if="isResponsive">
      <tr v-for="(row, i) in model.state.rows" :key="i.toString()">
        <td
          v-for="(cell, ii) in Object.keys(model.state.columns)"
          v-if="isDesktop"
          :key="ii"
          :class="cell"
          @click="onClick(row, cell)"
        >
          <component
            :key="i.toString() + cell"
            :is="getRenderer(cell)"
            :class="cell"
            :k="cell"
            :v="row[cell]"
            @update:v="onChange($event, parseInt(i.toString()))"
          ></component>
        </td>
        <td v-else-if="isTablet" :colspan="model.numCols">
          <component
            :key="i.toString() + 'm'"
            :is="tabletRenderer"
            :row="row"
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
import { TwBreakpoint, useScreenSize } from "@snowind/state";
import SwDatatableModel from "./models/datatable";
import DefaultCellRenderer from "./renderers/DefaultCellRenderer.vue";
import CarretSorter from "./widgets/CarretSorter.vue";

export default defineComponent({
  components: {
    CarretSorter,
  },
  props: {
    model: {
      type: Object as () => SwDatatableModel,
      required: true,
    },
    renderers: {
      type: Object,
      default: () => { },
    },
    mobileRenderer: {
      type: Object,
      default: () => null,
    },
    mobileBreakpoint: {
      type: String as () => TwBreakpoint,
      default: () => "sm",
    },
    tabletRenderer: {
      type: Object,
      default: () => null,
    },
    tabletBreakpoint: {
      type: String as () => TwBreakpoint,
      default: () => "md",
    },
    sortableCols: {
      type: Array as () => Array<string>,
      default: () => []
    },
    extraHeader: {
      type: Object,
      default: () => null,
    },
    defaultRenderer: {
      type: Object,
      default: () => null,
    }
  },
  emits: ["onClickCell"],
  setup(props, { emit }) {
    const {
      model,
      renderers,
      mobileRenderer,
      mobileBreakpoint,
      tabletRenderer,
      tabletBreakpoint,
      defaultRenderer,
    } = toRefs(props);
    const isResponsive = ref(mobileRenderer.value !== null || tabletRenderer.value !== null);
    const { isMobile, isTablet, isDesktop } = useScreenSize(mobileBreakpoint.value, tabletBreakpoint.value);
    const _defaultRenderer = defaultRenderer.value ?? DefaultCellRenderer;

    function getRenderer(k: string) {
      if (renderers.value !== undefined) {
        const rendererNames = new Set<string>(Object.keys(renderers.value));
        if (rendererNames.has(k)) {
          return renderers.value[k];
        }
      }
      return _defaultRenderer;
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
      isMobile,
      isTablet,
      isDesktop,
    };
  },
});
</script>