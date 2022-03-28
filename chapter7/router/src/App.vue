<template>
  <div id="routerdiv">
    <table>
      <tr>
        <td>
          <router-link to="/">Home</router-link>
        </td>
        <td>
          <router-link to="/login">Login</router-link>
        </td>
      </tr>
    </table>
    <router-view></router-view>
  </div>
</template>
