<script setup>
import { ref } from 'vue';
import * as demoAPI from '@/api/demo';

// Sample to show how we can inspect mode
// and import env variables
const deploymentMode = import.meta.env.MODE;
const myBaseURL = import.meta.env.VITE_BASE_API_URL;

async function getData() {
  const { data } = await demoAPI.getFromServer();
  console.log(data);
  result.value.push(data.message);
}

async function postData() {
  const { data } = await demoAPI.postToServer({ Inbound: msg.value });
  console.log(data);
  result.value.push(data.outbound);
}

const result = ref([]);
const msg = ref('');

defineProps({
  sampleProp: String,
});
</script>

<template>
  <div class="flex min-h-screen bg-gray-800 justify-center">
    <div class="bg-gray-800 flex flex-col items-center justify-center">
      <div class="flex flex-col space-x-2 items-start justify-start">
        <p class="text-2xl font-semibold text-white">FullyStacked 💪</p>
        <p class="text-xs text-gray-50">Login to your account</p>
      </div>
      <form class="mt-12">
        <!-- email -->
        <div class="pt-4 w-full flex space-x-2 items-center justify-end">
          <p class="text-xs font-bold text-white">Email Address</p>
          <input
            type="text"
            class="border focus-within:border-red-600 focus:outline-2 border-red-500 rounded-md p-2 text-sm"
            placeholder="you@somewhere.com"
          />
        </div>

        <!-- password -->
        <div class="pt-4 w-full flex space-x-2 items-center justify-end">
          <p class="text-xs font-bold text-white">Password</p>
          <input
            type="password"
            class="border focus-within:border-red-600 focus:outline-2 border-red-500 rounded-md p-2 text-sm"
            placeholder="🔐"
          />
        </div>

        <div class="flex items-center justify-center flex-1 mt-12">
          <button class="px-4 pt-2 pb-2.5 w-full rounded-lg bg-red-500 hover:bg-red-600">
            <span class="flex-1 text-sm font-bold text-center text-white">Login</span>
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<style scoped></style>
