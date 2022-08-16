import { h } from "vue";
import { RouterLink } from "vue-router";
import SvgIcon from "@/components/SvgIcon/SvgIcon.vue";
export function generatorMenu(routerMap) {
  const routes = [];
  routerMap.forEach((it) => {
    let route = {
      label: renderLabel(it.meta?.title, it.name, it.children),
      key: it.name,
      icon: renderIcon(it.meta?.icon),
    };
    if (it.children) {
      const children = getChildrenRouter(it.children);
      if (
        children.length > 1 ||
        (children.length === 1 && it.meta?.alwaysShow)
      ) {
        route.children = children;
      } else if (children.length === 1 && !it.meta?.alwaysShow) {
        route = {
          label: renderLabel(
            children[0].meta?.title,
            children[0].name,
            children[0].children
          ),
          key: children[0].name,
          icon: renderIcon(children[0].meta?.icon),
        };
      }
    }
    routes.push(route);
  });
  return routes;
}

/**
 * 递归组装子菜单
 * */
export function getChildrenRouter(routerMap) {
  const routes = [];
  routerMap.forEach((it) => {
    if (!it.meta.hidden) {
      let route = {
        label: renderLabel(it.meta?.title, it.name, it.children),
        key: it.name,
        icon: renderIcon(it.meta?.icon),
        ...it,
      };
      if (it.children) {
        route.children = getChildrenRouter(it.children);
      }
      routes.push(route);
    }
  });
  return routes;
}
// 渲染图标
function renderIcon(icon) {
  return () => h(SvgIcon, { name: icon }, null);
}
// 渲染名称
function renderLabel(title, name, children) {
  if (children) {
    return () => h("div", null, title);
  } else {
    return () =>
      h(
        RouterLink,
        {
          to: {
            name,
          },
        },
        { default: () => title }
      );
  }
}
