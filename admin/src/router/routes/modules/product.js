export default {
  path: "/product",
  component: () => import("@/layout/index.vue"),
  name: "Product",
  meta: {
    title: "商品管理",
    sort: 200,
    icon: "product",
    alwaysShow: true,
  },
  children: [
    {
      path: "list",
      name: "ProductList",
      meta: {
        title: "商品列表",
        icon: "product-list",
      },
      component: () =>
        import("@/pages/product-manage/product-list/product-list.vue"),
    },
    {
      path: "create",
      name: "CreateProduct",
      meta: {
        title: "创建商品",
        hidden: true,
      },
      component: () =>
        import("@/pages/product-manage/create-product/create-product.vue"),
    },
  ],
};
