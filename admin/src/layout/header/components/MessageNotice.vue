<template>
  <n-popover trigger="hover" placement="bottom">
    <template #trigger>
      <n-badge :value="todoCount" :show-zero="true">
        <Icon
          icon="ant-design:bell-outlined"
          width="24"
          height="24"
          :inline="true"
          class="icon"
          style="color: rgba(0, 0, 0, 0.85)"
        />
      </n-badge>
    </template>
    <div style="width: 250px">
      <n-tabs type="segment" v-model:value="activeKey">
        <n-tab-pane name="1" tab="消息">
          我这辈子最疯狂的事，发生在我在 Amazon
          当软件工程师的时候，故事是这样的：<br /><br />
          那时我和女朋友住在一起，正在家里远程工作。忽然同事给我发来了紧急消息：”我们的服务出现了
          SEV 2 级别的故障！需要所有的人马上协助！“我们组的应用全挂掉了。<br /><br />
          当我还在费力的寻找修复方法的时候，忽然闻到隔壁房间的的焦味，防火报警器开始鸣叫。
        </n-tab-pane>
        <n-tab-pane name="2" tab="评论">
          “威尔！着火了！快来帮忙！”我听到女朋友大喊。现在一个难题在我面前——是恢复一个重要的
          Amazon 服务，还是救公寓的火。<br /><br />
          我的脑海中忽然出现了 Amazon
          著名的领导力准则”客户至上“，有很多的客户还依赖我们的服务，我不能让他们失望！所以着火也不管了，女朋友喊我也无所谓，我开始
          debug 这个线上问题。
        </n-tab-pane>
        <n-tab-pane name="3" tab="其他">
          但是忽然，公寓的烟味消失，火警也停了。我的女朋友走进了房间，让我震惊的是，她摘下了自己的假发，她是
          Jeff Bezos（Amazon 老板）假扮的！<br /><br />
          “我对你坚持顾客至上的原则感到十分骄傲”，说完，他递给我一张五美金的亚马逊礼品卡，从我家窗户翻了出去，跳上了一辆
          Amazon 会员服务的小货车，一溜烟离开了。<br /><br />虽然现在我已不在
          Amazon
          工作，但我还是非常感激在哪里学的到的经验，这些经验我终身难忘。你们同意么？
        </n-tab-pane>
      </n-tabs>
    </div>
  </n-popover>
</template>

<script setup>
import { ref } from "vue";

let activeKey = ref("2");
const todoCount = ref(0);
</script>

<style scoped>
.icon {
  font-size: 20px;
  vertical-align: middle !important;
}
</style>
