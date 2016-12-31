<template>
  <div>
    <md-dialog md-open-from="#fab" md-close-to="#fab" ref="menuModal">
      <md-dialog-title>Add {{ restaurant.name }}'s menu</md-dialog-title>
      <md-dialog-content class="add-order-dialog">
        <form v-on:submit.prevent="addMenu">
          <md-input-container>
            <label>Menu</label>
            <md-select placeholder="Name" name="name" v-model="newMenu.id">
              <md-subheader>Select Menu</md-subheader>
              <md-option v-for="(value, rowIndex) in restaurant.menus" :key="rowIndex" :value="value.id">{{ value.name }} ( {{ value.price }} baht )</md-option>
            </md-select>
          </md-input-container>
          <md-input-container>
            <label>Quantity</label>
            <md-select placeholder="Quantity" name="quantity" v-model="newMenu.quantity">
              <md-option v-for="(n, rowIndex) in 10" :key="rowIndex" :value="n">{{ n }}</md-option>
            </md-select>
          </md-input-container>
        </form>
      </md-dialog-content>

      <md-dialog-actions>
        <md-button class="md-primary" @click="addMenu('menuModal')">Add</md-button>
        <md-button class="md-primary" @click="closeModal('menuModal')">Cancel</md-button>
      </md-dialog-actions>
    </md-dialog>
    <md-button class="md-raised md-button md-warn" id="addMenu" @click="openModal('menuModal')">
      <md-icon>add</md-icon><span> Add order</span>
    </md-button>
  </div>
</template>
<script>
  // import vSelect from "vue-select"
  export default {
    // components: { vSelect },
    props: ['restaurant','activityID'],
    data() {
      return {
        orders: [],
        newMenu: {
          id:''
        },
      }
    },
    computed: {

    },
    methods: {
      openModal(ref) {
        this.$refs[ref].open();
      },
      closeModal(ref) {
        this.$refs[ref].close();
      },
      addMenu(ref) {
        const order = {
                        user_id:this.$auth.user().id,
                        menu_id: +this.newMenu.id,
                        activity_id: +this.activityID,
                      };
        this.$store.dispatch('addOrder', order );
        this.$refs[ref].close();
        this.newMenu.id = '';
      },
    }
  }
</script>
