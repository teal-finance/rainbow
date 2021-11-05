import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router"

const baseTitle = "Rainbow"

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    component: () => import("./views/HomeView.vue")
  },
  {
    path: "/options",
    component: () => import("./views/OptionsView.vue"),
    meta: {
      title: "Options"
    }
  },
  {
    path: "/chart",
    component: () => import("./views/ChartView.vue"),
    meta: {
      title: "Chart proto"
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