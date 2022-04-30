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
import { defineComponent, toRefs, ref, onMounted, reactive, watchEffect, watch } from 'vue'
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
      default: { 'defaultValue': true }
    }
  },
  setup(props) {
    const { model, col, filtersState } = toRefs(props);
    const entries = reactive<Record<any, boolean>>({});
    const vals = new Set();

    function distinctValues() {
      model.value.state.rows.forEach((row) => {
        vals.add(row[col.value]);
      });
      //console.log("entries", JSON.stringify(entries, null, "  "))
    }

    function setFilters() {
      vals.forEach((v) => {
        if (`${v}` in filtersState.value) {
          entries[`${v}`] = filtersState.value[`${v}`]
        } else {
          entries[`${v}`] = filtersState.value.defaultValue
        }
        if (entries[`${v}`] == false) {
          exclude(`${v}`)
        } else {
          try {
            include(`${v}`)
          } catch (e) { }
        }
      });
    }

    function exclude(evt: any) {
      //console.log("exclude filter", col.value, evt)
      model.value.addExcludeFilter(col.value, evt)
    }

    function include(evt: any) {
      //console.log("include filter", col.value, evt)
      model.value.removeExcludeFilter(col.value, evt)
    }

    onMounted(() => {
      distinctValues();
      setFilters()
    })

    watch(() => props.filtersState, (newval: Record<any, boolean>, oldval: Record<any, boolean>) => {
      setFilters()
    });

    return {
      entries,
      include,
      exclude,
    }
  }
});
</script>