<script setup>
import { ref } from 'vue';

defineProps({
  msg: String,
});

const count = ref(0);
</script>

<template>
  <div class="flex space-2 justify-center">
    <button
      @click="count++"
      type="button"
      class="inline-block px-6 py-2.5 bg-blue-600 text-white font-medium text-lg leading-tight normal-case rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out"
    >
      Click me - my count is {{ count }}
    </button>
  </div>
</template>

<style scoped></style>
