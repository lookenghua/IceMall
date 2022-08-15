<template>
  <div class="editor">
    <div class="toolbar-container" ref="toolbarRef"><!-- 工具栏 --></div>
    <div class="editor-container" ref="editorRef"><!-- 编辑器 --></div>
  </div>
</template>

<script setup>
import { onMounted, ref, watch } from "vue";
import { createEditor, createToolbar } from "@wangeditor/editor";
import "@wangeditor/editor/dist/css/style.css";

const props = defineProps({
  modelValue: {
    type: Array,
    default: () => [],
  },
  placeholder: {
    type: String,
    default: "请输入内容~",
  },
  customUploadImage: {
    type: Function,
  },
  customUploadVideo: {
    type: Function,
  },
});
const emit = defineEmits(["update:modelValue"]);
const toolbarRef = ref();
const editorRef = ref();
let editor, toolbar;
watch(
  () => props.modelValue,
  (val) => {
    if (val.length === 0) {
      editor.clear();
    }
  }
);
onMounted(() => {
  editor = createEditor({
    content: props.modelValue,
    selector: editorRef.value,
    config: {
      placeholder: props.placeholder,
      onChange(editor) {
        emit("update:modelValue", editor.children);
      },
      MENU_CONF: {
        uploadImage: {
          async customUpload(file, insertFn) {
            if (!props.customUploadImage) {
              alert("请自定义上传图片方法");
              return;
            }
            props.customUploadImage(file).then(({ url, alt, href }) => {
              insertFn(url, alt, href);
            });
          },
        },
        uploadVideo: {
          async customUpload(file, insertFn) {
            // TS 语法
            // async customUpload(file, insertFn) {                   // JS 语法
            // file 即选中的文件
            // 自己实现上传，并得到视频 url poster
            // 最后插入视频
            if (!props.customUploadVideo) {
              alert("请自定义上传图片方法");
              return;
            }
            props.customUploadVideo(file).then(({ url, poster }) => {
              insertFn(url, poster);
            });
          },
        },
      },
    },
  });
  toolbar = createToolbar({
    editor,
    selector: toolbarRef.value,
    config: {
      excludeKeys: ["codeBlock", "code"],
    },
    mode: "default", // or 'simple'
  });
});
</script>

<style scoped lang="scss">
.editor {
  width: 100%;
  border: 3px solid #0d0d0d;
  border-radius: 0.75rem;
  .toolbar-container {
    border-bottom: 3px solid #0d0d0d;
    border-top-left-radius: 0.75rem;
    border-top-right-radius: 0.75rem;
    padding: 3px;
  }
  .editor-container {
    min-height: 500px;
  }
}
</style>
