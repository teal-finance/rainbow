// Copyright 2021 Teal.Finance/Rainbow contributors
// This file is part of Teal.Finance/Rainbow,
// a screener for DeFi options under the MIT License.
// SPDX-License-Identifier: MIT

import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router"

const baseTitle = "Rainbow"

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    component: () => import("./views/ClassicView.vue"),
    meta: {
      title: "DeFi options comparator for trading opportunities"
    }
  },
  {
    path: "/about",
    component: () => import("./views/pages/AboutView.vue"),
    meta: {
      title: "DeFi options comparator for trading opportunities"
    }
  },
  {
    path: "/exotic",
    component: () => import("./views/ExoticView.vue"),
    meta: {
      title: "Exotic options comparator for trading opportunities"
    }
  },
  {
    path: "/classic",
    component: () => import("./views/ClassicView.vue"),
    meta: {
      title: "Classic options comparator for trading opportunities"
    }
  }
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes
});

router.afterEach((to, from) => { // eslint-disable-line
  if (to.meta?.title) {
    document.title = `${baseTitle} - ${to.meta?.title}`
  }
});

export default router
