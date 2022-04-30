<template>
  <component
    :is="renderer"
    :col="col"
    :values="entries"
    @include="include($event)"
    @exclude="exclude($event)"
  ></component>
</template>

<script lang="ts">
import { defineComponent, toRefs, ref, onMounted, reactive } from 'vue'
import SwDatatableModel from '../models/datatable'
import ValuesFilterSwitchRender from './ValuesFilterSwitchRender.vue';

export default defineComponent({
  components: {
    ValuesFilterSwitchRender
  },
  props: {
    model: {
      type: Object as () => SwDatatableModel,
      required: true,
    },
    col: {
      type: String,
      required: true
    },
    renderer: {
      type: Object,
      default: () => ValuesFilterSwitchRender
    },
    filtersState: {
      type: Object as () => Record<any, boolean>,
      default: {}
    },
  },
  setup(props) {
    const { model, col, filtersState } = toRefs(props);
    const entries = reactive<Record<any, boolean>>({});

    function distinctValues() {
      const vals = new Set();
      model.value.state.rows.forEach((row) => {
        vals.add(row[col.value]);
      });
      vals.forEach((v) => {
        if (`${v}` in filtersState.value) {
          entries[`${v}`] = filtersState.value[`${v}`]
        } else {
          entries[`${v}`] = true
        }
      });
      console.log("entries", JSON.stringify(entries, null, "  "))
    }

    function exclude(evt: any) {
      model.value.addExcludeFilter(col.value, evt)
    }

    function include(evt: any) {
      try {
        model.value.removeExcludeFilter(col.value, evt)
      } catch (e) {
        console.log("Error removing filter", e)
      }
    }

    onMounted(() => {
      distinctValues();
    })

    return {
      entries,
      include,
      exclude,
    }
  }
});
</script>