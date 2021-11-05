<template>
  <component
    :is="renderer"
    :col="col"
    :values="vals"
    @include="include($event)"
    @exclude="exclude($event)"
  ></component>
</template>

<script lang="ts">
import { defineComponent, toRefs, ref, onMounted } from 'vue'
import SwDatatableModel from '../models/datatable'
import ValuesFilterBadgeRender from './ValuesFilterBadgeRender.vue';

export default defineComponent({
  components: {
    ValuesFilterBadgeRender
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
      default: () => ValuesFilterBadgeRender
    },
  },
  setup(props) {
    const { model, col } = toRefs(props);
    const vals = ref(new Set());

    function distinctValues() {
      model.value.state.rows.forEach((row) => {
        vals.value.add(row[col.value]);
      });
    }

    function exclude(evt: any) {
      console.log("exclude", evt);
      //model.value.filterInclude(evt)
      model.value.addExcludeFilter(col.value, evt)
    }

    function include(evt: any) {
      console.log("include", evt);
      //model.value.filterInclude(evt)
      model.value.removeExcludeFilter(col.value, evt)
    }

    onMounted(() => {
      distinctValues();
    })

    return {
      vals,
      include,
      exclude,
    }
  }
});
</script>