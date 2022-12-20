<template>
  <v-app>
    <v-container>

      <v-row>
        <v-text-field
            label="Test "
            solo
        >
        </v-text-field>
      </v-row>

      <v-row
          align="right"
          justify="center">
        <v-btn color="red lighten-1">
          <span>Recents</span>
          <v-icon>mdi-history</v-icon>
        </v-btn>

        <v-btn>
          <span>Favorites</span>

          <v-icon>mdi-heart</v-icon>
        </v-btn>

        <v-btn color="teal lighten-4">
          <span>Nearby</span>
          <v-icon>mdi-map-marker</v-icon>
        </v-btn>

      </v-row>

      <div style="height:50px"></div>

      <v-row
          align="right"
          justify="center"
      >
        <v-badge
            bordered
            color="error"
            icon="mdi-lock"
            overlap
        >
          <v-btn dark rounded>
            Lock Account
          </v-btn>
        </v-badge>

        <div class="mx-3"></div>

        <v-badge
            bordered
            bottom
            color="deep-purple accent-4"
            dot
            offset-x="10"
            offset-y="10"

        >
          <v-avatar size="40">
            <v-img src="https://cdn.vuetifyjs.com/images/lists/2.jpg"></v-img>
          </v-avatar>
        </v-badge>

        <div class="mx-3"></div>

        <v-badge
            avatar
            bordered
            overlap
            sm="6"
        >
          <template v-slot:badge>
            <v-avatar>
              <v-img src="https://cdn.vuetifyjs.com/images/logos/v.png"></v-img>
            </v-avatar>
          </template>

          <v-avatar size="40">
            <v-img src="https://cdn.vuetifyjs.com/images/john.png"></v-img>
          </v-avatar>
        </v-badge>
      </v-row>

      <v-row>
        <v-speed-dial
            v-model="fab"
            absolute
            bottom
            right
            direction="top"
        >
          <template v-slot:activator>
            <v-btn
                v-model="fab"
                color="pink accent-3"
                dark
                fab>
              <v-icon>
                mdi-close
              </v-icon>
            </v-btn>
          </template>
          <v-btn
              fab
              dark
              small
              @click="$emit('click:timer')">
            <v-icon>mdi-timer-outline</v-icon>
          </v-btn>
        </v-speed-dial>
      </v-row>

      <v-row>
        <v-col>
          <v-btn-toggle v-model="toggle_exclusive">
            <v-btn>
              <v-icon>mdi-format-align-left</v-icon>
            </v-btn>

            <v-btn>
              <v-icon>mdi-format-align-center</v-icon>
            </v-btn>

            <v-btn>
              <v-icon>mdi-format-align-right</v-icon>
            </v-btn>

            <v-btn>
              <v-icon>mdi-format-align-justify</v-icon>
            </v-btn>
          </v-btn-toggle>
        </v-col>

        <v-col>
          <v-btn-toggle
              v-model="toggle_multiple"
              dense
              background-color="primary"
              dark
              multiple
          >
            <v-btn>
              <v-icon>mdi-format-bold</v-icon>
            </v-btn>

            <v-btn>
              <v-icon>mdi-format-italic</v-icon>
            </v-btn>

            <v-btn>
              <v-icon>mdi-format-underline</v-icon>
            </v-btn>

            <v-btn>
              <v-icon>mdi-format-color-fill</v-icon>
            </v-btn>
          </v-btn-toggle>
        </v-col>
      </v-row>

      <v-row>
        <v-radio-group v-model="radioButton">
          <v-radio color="primary" label="Radio 1" value="Radio-1"></v-radio>
          <v-radio color="primary" label="Radio 2" value="Radio-2"></v-radio>
        </v-radio-group>
      </v-row>

      <v-row>
        <v-col>
          Color Picker
        </v-col>
        <v-col>
          <v-color-picker/>
        </v-col>
      </v-row>
    </v-container>
  </v-app>
</template>

<style lang="scss"></style>

<script>
export default {
  data: () => ({
    sticky: false,
  }),
}
</script>

<script>
export default {
  data: () => ({value: 1}),
}
</script>