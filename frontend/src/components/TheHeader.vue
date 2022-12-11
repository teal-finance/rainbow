<template>
  <div>
    <sw-header class="h-20 bg-primary dark:bg-header-dark text-primary-r dark:text-primary-r-dark"
      @togglemenu="isMenuVisible = !isMenuVisible">
      <template #branding>
        <div class="flex flex-row items-center ml-5 text-lg cursor-pointer" @click="$router.push(titleLink)">
          <div class="inline-block mx-3">
            <img alt="logo representing a pixelized cloud over a rainbow" :src="`${logoURL}`" height="68" width="68" />
          </div>
          <div class="inline-block text-xl tracking-widest">
            <!-- img alt="Rainbow" src="./assets/rainbow-chancery.png" height="49" width="185" / -->
            {{ titleText }}
          </div>
        </div>
      </template>
      <template #mobile-branding>
        <div class="ml-5 text-lg cursor-pointer" @click="$router.push('/')">{{ mobilePageTitle }}</div>
      </template>
      <template #mobile-back>
        <i-ion-arrow-back-outline v-if="!isHome" class="inline-flex ml-2 text-3xl"></i-ion-arrow-back-outline>
        <img v-else alt="logo representing a pixelized cloud over a rainbow" :src="`${logoURL}`" height="68" width="68"
          class="ml-3" />
      </template>
      <template #menu>
        <div class="flex flex-row items-center justify-end h-full space-x-1">
          <div class="flex flex-row items-center justify-end h-full space-x-3" v-if="isDatatableReady">
            <header-select-dropdown name="Presets" :is-collapsed="isPresetCollapsed" @close="isPresetCollapsed = true"
              @toggle="isPresetCollapsed = !isPresetCollapsed">
              <presets-list @select="isPresetCollapsed = true"></presets-list>
            </header-select-dropdown>
            <header-dropdown name="Assets">
              <values-filter :model="datatable" col="asset" :filters-state="filterConf.assets"></values-filter>
            </header-dropdown>
            <header-dropdown name="Providers">
              <values-filter :model="datatable" col="provider" :filters-state="filterConf.providers"></values-filter>
            </header-dropdown>
            <header-dropdown name="Expiry">
              <values-filter :model="datatable" col="expiry"></values-filter>
            </header-dropdown>
            <div class="w-8"></div>
          </div>
          <!-- button class="border-none btn" @click="openView('/options')">Options</button -->
          <button class="border-none btn" @click="$router.push(linkURL)">{{ linkText }}</button>
          <button class="border-none btn" @click="$router.push('/about')">About</button>
          <button class="text-xl border-none btn" @click="openSourceCode()">
            <i-ant-design-github-filled></i-ant-design-github-filled>
          </button>
          <div class="px-5 text-lg cursor-pointer" @click="user.toggleDarkMode()" id="dark-mode-switch">
            <i-fa-solid-moon v-if="user.isDarkMode.value == false"></i-fa-solid-moon>
            <i-fa-solid-sun v-else></i-fa-solid-sun>
          </div>
        </div>
      </template>
    </sw-header>
    <sw-mobile-menu :is-visible="isMenuVisible"
      class="bg-neutral text-neutral-r dark:bg-background-dark dark:text-neutral-r">
      <div class="flex flex-col items-center p-3 space-y-5">
        <!-- button class="border-none btn" @click="openView('/options')">Options</button -->
        <div class="w-full border-b border-foreground">
          <button class="w-full border-none btn" @click="$router.push('/about')">About</button>
        </div>
        <div class="w-full border-b border-foreground">
          <button class="w-full border-none btn">
            Source code
            <i-ant-design-github-filled></i-ant-design-github-filled>
          </button>
        </div>
        <div class="w-full px-5 text-lg text-center border-b cursor-pointer border-foreground"
          @click="user.toggleDarkMode(); isMenuVisible = false">
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
import HeaderDropdown from './widgets/HeaderDropdown.vue';
import HeaderSelectDropdown from './widgets/HeaderSelectDropdown.vue';
import PresetsList from '@/components/PresetsList.vue';
import { user, datatable, isDatatableReady, filterConf } from "@/state";
import ValuesFilter from '@/packages/datatable/filters/ValuesFilter.vue'
import router from '@/router';

export default defineComponent({
  components: {
    SwHeader,
    SwMobileMenu,
    HeaderDropdown,
    HeaderSelectDropdown,
    ValuesFilter,
    PresetsList
  },
  setup() {
    const isMenuVisible = ref(false);
    const isPresetCollapsed = ref(true);

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

    function openSourceCode() {
      window.location.href = 'https://github.com/teal-finance/rainbow'
    }

    const logoURL = computed<string>(() => {
      const logo = import.meta.env.BASE_URL + "img/logo-transparent.png"
      const exotic = import.meta.env.BASE_URL + "img/exotic-transparent.png"
      return router.currentRoute.value.path.startsWith("/exotic") ? exotic : logo
    })

    const titleText = computed<"Exotic" | "Rainbow">(() => {
      return router.currentRoute.value.path.startsWith("/exotic") ? "Exotic" : "Rainbow"
    })

    const titleLink = computed<"/exotic" | "/">(() => {
      return router.currentRoute.value.path.startsWith("/exotic") ? "/exotic" : "/"
    })

    const linkText = computed<"Classic" | "Exotic">(() => {
      let mode: "Classic" | "Exotic" = "Exotic"
      if (router.currentRoute.value.path.startsWith("/exotic")) {
        mode = "Classic"
      }
      return mode
    })

    const linkURL = computed<"/" | "/exotic">(() => {
      return router.currentRoute.value.path.startsWith("/exotic") ? "/" : "/exotic"
    })

    return {
      isMenuVisible,
      closeMenu,
      user,
      isHome,
      openView,
      mobilePageTitle,
      openSourceCode,
      logoURL,
      titleText,
      titleLink,
      linkText,
      linkURL,
      datatable,
      isDatatableReady,
      filterConf,
      isPresetCollapsed
    }
  }
});
</script>
