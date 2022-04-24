<!--example from https://buefy.org/documentation-->
<template>
  <section>
    <div class="container">
      <b-carousel>
        <b-carousel-item v-for="(carousel, i) in carousels" :key="i">
          <section :class="`hero is-medium is-${carousel.color}`">
            <div class="hero-body has-text-centered">
              <h1 class="title">{{ carousel.text }}</h1>
            </div>
          </section>
        </b-carousel-item>
      </b-carousel>
    </div>

    <div class="container">
      <b-table :data="data" :columns="columns"></b-table>
    </div>

    <div class="container">
      <section class="b-tooltips">
        <b-tooltip label="Tooltip right"
                   position="is-right">
          <b-button label="Right" type="is-dark"/>
        </b-tooltip>

        <b-tooltip label="Tooltip top">
          <b-button label="Top (default)" type="is-dark"/>
        </b-tooltip>

        <b-tooltip label="Tooltip bottom"
                   position="is-bottom">
          <b-button label="Bottom" type="is-dark"/>
        </b-tooltip>

        <b-tooltip label="Tooltip left"
                   position="is-left">
          <b-button label="Left" type="is-dark"/>
        </b-tooltip>

        <b-tooltip label="delayed by 1000ms"
                   :delay="1000">
          <b-button label="Delayed" type="is-dark"/>
        </b-tooltip>
      </section>
    </div>

    <div id="app" class="container">

      <section>
        <b-breadcrumb align="is-centered">
          <b-breadcrumb-item href="/">Home</b-breadcrumb-item>
          <b-breadcrumb-item>Docs</b-breadcrumb-item>
          <b-breadcrumb-item active>Breadcrumb</b-breadcrumb-item>
        </b-breadcrumb>
      </section>
    </div>


    <div class="container">
      <section>

        <b-collapse :open="false" aria-id="contentIdForA11y1">
          <template #trigger="props">
            <b-button
                label="Click me!"
                type="is-primary"
                aria-controls="contentIdForA11y1"
                :aria-expanded="props.open"/>
          </template>
          <div class="notification">
            <div class="content">
              <h3>
                Subtitle
              </h3>
              <p>
                Lorem ipsum dolor sit amet, consectetur adipiscing elit. <br/>
                Nulla accumsan, metus ultrices eleifend gravida, nulla nunc varius lectus, nec rutrum justo nibh eu
                lectus. <br/>
                Ut vulputate semper dui. Fusce erat odio, sollicitudin vel erat vel, interdum mattis neque.
              </p>
            </div>
          </div>
        </b-collapse>

      </section>


    </div>

  </section>
</template>

<script>
export default {
  methods: {
    clickMe() {
      this.$buefy.notification.open("Clicked!!");
    }
  },
  data() {
    return {
      carousels: [
        {text: 'Slide 1', color: 'primary'},
        {text: 'Slide 2', color: 'info'},
        {text: 'Slide 3', color: 'success'},
        {text: 'Slide 4', color: 'warning'},
        {text: 'Slide 5', color: 'danger'}
      ],
      data: [
        {'id': 1, 'first_name': 'Jesse', 'last_name': 'Simmons', 'date': '2016-10-15 13:43:27', 'gender': 'Male'},
        {'id': 2, 'first_name': 'John', 'last_name': 'Jacobs', 'date': '2016-12-15 06:00:53', 'gender': 'Male'},
        {'id': 3, 'first_name': 'Tina', 'last_name': 'Gilbert', 'date': '2016-04-26 06:26:28', 'gender': 'Female'},
        {'id': 4, 'first_name': 'Clarence', 'last_name': 'Flores', 'date': '2016-04-10 10:28:46', 'gender': 'Male'},
        {'id': 5, 'first_name': 'Anne', 'last_name': 'Lee', 'date': '2016-12-06 14:38:38', 'gender': 'Female'}
      ],
      columns: [
        {
          field: 'id',
          label: 'ID',
          width: '40',
          numeric: true
        },
        {
          field: 'first_name',
          label: 'First Name',
        },
        {
          field: 'last_name',
          label: 'Last Name',
        },
        {
          field: 'date',
          label: 'Date',
          centered: true
        },
        {
          field: 'gender',
          label: 'Gender',
        }
      ]
    }
  }
};
</script>