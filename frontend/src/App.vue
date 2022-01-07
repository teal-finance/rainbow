<template>
  <div :class="{ 'dark': user.isDarkMode.value == true }">
    <div
      class="w-full bg-background dark:bg-background-dark text-foreground dark:text-foreground-dark"
    >
      <the-header></the-header>
      <div class="p-5">
        <router-view></router-view>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted } from 'vue'
import { user } from "@/state";
import TheHeader from './components/TheHeader.vue';
import { query } from './api/graphql';

export default defineComponent({
  components: {
    TheHeader,
  },
  setup() {
    onMounted(() => {
      query()
    })
    return {
      user,
    }
  }
})
</script>

<style lang="css">
html,
body {
  margin: 0;
  padding: 0;
}
</style>
