<template>
  <div id="routerdiv">
    <table>
      <tr>
        <td>
          <router-link :to="{ name: 'Home'}">Home</router-link>
        </td>
        <td>
          <router-link :to="{ name: 'Login'}">Login</router-link>
        </td>
      </tr>
    </table>
    <router-view></router-view>
  </div>
</template>
