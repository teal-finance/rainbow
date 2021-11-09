<template>
  <div>
    <sw-header
      class="h-20 bg-primary dark:bg-primary-dark text-primary-r dark:text-primary-r-dark"
      @togglemenu="isMenuVisible = !isMenuVisible"
    >
      <template #branding>
        <div
          class="flex flex-row items-center ml-5 text-lg cursor-pointer"
          @click="$router.push('/')"
        >
          <div class="inline-block mx-3">
            <img alt="logo" src="@/assets/logo-transparent.png" height="68" width="68" />
          </div>
          <div class="inline-block text-xl tracking-widest">
            <!-- img alt="Rainbow" src="./assets/rainbow-chancery.png" height="49" width="185" / -->
            Rainbow
          </div>
        </div>
      </template>
      <template #mobile-branding>
        <div class="ml-5 text-lg cursor-pointer" @click="$router.push('/')">{{ mobilePageTitle }}</div>
      </template>
      <template #mobile-back>
        <i-ion-arrow-back-outline v-if="!isHome" class="inline-flex ml-2 text-3xl"></i-ion-arrow-back-outline>
        <img
          v-else
          alt="logo"
          src="@/assets/logo-transparent.png"
          height="68"
          width="68"
          class="ml-3"
        />
      </template>
      <template #menu>
        <div class="flex flex-row items-center justify-end h-full space-x-1">
          <!-- button class="border-none btn" @click="openView('/options')">Options</button -->
          <div class="px-5 text-lg cursor-pointer" @click="user.toggleDarkMode()">
            <i-fa-solid-moon v-if="user.isDarkMode.value == false"></i-fa-solid-moon>
            <i-fa-solid-sun v-else></i-fa-solid-sun>
          </div>
        </div>
      </template>
    </sw-header>
    <sw-mobile-menu
      :is-visible="isMenuVisible"
      class="bg-neutral text-neutral-r dark:bg-neutral-dark dark:text-neutral-r-dark"
    >
      <div class="flex flex-col items-center p-3 space-y-5">
        <!-- button class="border-none btn" @click="openView('/options')">Options</button -->
        <div
          class="px-5 text-lg cursor-pointer"
          @click="user.toggleDarkMode(); isMenuVisible = false"
        >
          <div v-if="user.isDarkMode.value == false">
            <i-fa-solid-moon></i-fa-solid-moon>&nbsp;Dark mode
          </div>
          <div v-else>
            <i-fa-solid-sun></i-fa-solid-sun>&nbsp;Light mode
          </div>
        </div>
      </div>
    </sw-mobile-menu>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, computed } from 'vue'
import { SwHeader, SwMobileMenu } from "@snowind/header";
import { user } from "@/state";
import router from '@/router';

export default defineComponent({
  components: {
    SwHeader,
    SwMobileMenu,
  },
  setup() {
    const isMenuVisible = ref(false);

    const isHome = computed<boolean>(() => router.currentRoute.value.path == "/");
    const mobilePageTitle = computed<string>(() => {
      if (router.currentRoute.value.meta?.title) {
        return `${router.currentRoute.value.meta.title}`;
      }
      return "Rainbow"
    });

    function closeMenu() {
      isMenuVisible.value = false;
    }

    function openView(url: string) {
      router.push(url);
      closeMenu();
    }

    return {
      isMenuVisible,
      closeMenu,
      user,
      isHome,
      openView,
      mobilePageTitle,
    }
  }
});
</script>