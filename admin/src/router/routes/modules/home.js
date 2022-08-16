export default {
  path: "/",
  component: () => import("@/layout/index.vue"),
  name: "Home",
  meta: {
    title: "首页",
    icon: "dashboard",
    sort: 100,
  },
  children: [
    {
      path: "/dashboard",
      name: "Dashboard",
      meta: {
        title: "Dashboard",
        icon: "dashboard",
        sort: 101,
        affix: true,
      },
      component: () => import("@/pages/home/dashboard/dashboard.vue"),
    },
  ],
};
