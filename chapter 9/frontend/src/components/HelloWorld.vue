<script setup>
import { ref } from 'vue';
import * as demoAPI from '@/api/demo';

// Sample to show how we can inspect mode 
// and import env variables
const deploymentMode = import.meta.env.MODE;
const myBaseURL = import.meta.env.VITE_BASE_API_URL;

async function getData() {
  const { data } = await demoAPI.getFromServer()
  // console.log(data.Message)
  result.value.push(data.Message)
}

async function postData() {
  const { data } = await demoAPI.postToServer({ Inbound: msg.value})
  // console.log(data.OutBound)
  result.value.push(data.OutBound)
}

const result = ref([])
const msg = ref("")

defineProps({
  sampleProp: String,
});

</script>

<template>
  <div class="flex space-2 justify-center">
    <button
      @click="getData()" 
      type="button"
      class="inline-block px-6 py-2.5 bg-blue-600 text-white font-medium text-lg leading-tight normal-case rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out"
    >
      Click to Get
    </button>
  </div>
  <div class="flex mt-4 space-2 justify-center">
    <input type="text" 
      class="inline-block px-6 py-2.5 text-blue-600 font-medium text-lg leading-tight rounded shadow-md border-2 border-solid border-black focus:shadow-lg  focus:ring-1 "
   v-model="msg" />
    <button
      @click="postData()"
      type="button"
      class="inline-block px-6 py-2.5 bg-blue-600 text-white font-medium text-lg leading-tight normal-case rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out"
    >
      Click to Post
    </button>
  </div>
  <p>You are in {{ deploymentMode }} mode</p>
  <p>Your API is at {{ myBaseURL }}</p>
  <li v-for="(r, index) in result">
    {{ r }}
  </li>
  
</template>

<style scoped></style>
