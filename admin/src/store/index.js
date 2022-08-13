import { createPinia } from "pinia";
import Colada, { PiniaColadaPlugin } from "colada-plugin";
const store = createPinia();

export function setupStore(app) {
  app.use(store);
  // 时间穿梭机
  store.use(PiniaColadaPlugin);
  app.use(Colada);
}

export { store };
