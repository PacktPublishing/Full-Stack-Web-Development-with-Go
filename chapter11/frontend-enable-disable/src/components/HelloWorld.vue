<script setup>
import { ref } from 'vue';
const deploymentMode = import.meta.env.MODE;
const myBaseURL = import.meta.env.VITE_BASE_API_URL;


defineProps({
  sampleProp: String,
});
</script>


<script>
import axios from 'axios';

export default {
  data() {
    return {
      enabled: true
    }
  },
  mounted() {
    axios({method: "GET", "url": "http://localhost:8080/features/disable_get"}).then(result => {
      this.enabled = result.data.enabled
      console.log(result);
    }, error => {
      console.error(error);
    });
  }
}
</script>

<template>
  <div  v-if="enabled" class="flex space-2 justify-center">
    <button
      type="button"
      class="inline-block px-6 py-2.5 bg-blue-600 text-white font-medium text-lg leading-tight normal-case rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out"
    >
      Click to Get
    </button>
  </div>
  <div class="flex mt-4 space-2 justify-center">
    <input type="text"
      class="inline-block px-6 py-2.5 text-blue-600 font-medium text-lg leading-tight rounded shadow-md border-2 border-solid border-black focus:shadow-lg  focus:ring-1 "
 />
    <button
      type="button"
      class="inline-block px-6 py-2.5 bg-blue-600 text-white font-medium text-lg leading-tight normal-case rounded shadow-md hover:bg-blue-700 hover:shadow-lg focus:bg-blue-700 focus:shadow-lg focus:outline-none focus:ring-0 active:bg-blue-800 active:shadow-lg transition duration-150 ease-in-out"
    >
      Click to Post
    </button>
  </div>
  <p>You are in {{ deploymentMode }} mode</p>
  <p>Your API is at {{ myBaseURL }}</p>

</template>

<style scoped></style>
