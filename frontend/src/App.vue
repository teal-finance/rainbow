<template>
  <div id="toplevel" :class="{ 'dark': user.isDarkMode.value == true }">
    <div class="w-full background">
      <the-header></the-header>
      <div class="p-5">
        <router-view></router-view>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { user } from "@/state";
import TheHeader from './components/TheHeader.vue';
</script>

<style lang="sass">

html
  overflow-y: scroll
</style>
