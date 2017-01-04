<template>
  <md-layout md-gutter v-if="!isLoading">
    <md-layout md-flex-small="10" md-flex-medium="10" md-flex-large="10" md-flex-xlarge="20"></md-layout>
    <md-layout md-flex-small="80" md-flex-medium="80" md-flex-large="80" md-flex-xlarge="60">
      <md-card class="activity-card" v-for="act in activities">
        <md-toolbar class="md-dense">
          <h2 class="md-title">
            <md-icon>event</md-icon>
            <span>{{act.name}}</span>
            <md-tooltip md-direction="bottom">{{act.name}} - [{{act.id}}]</md-tooltip>
          </h2>
          <span class="gutter"></span>
          <countdown v-bind:expired="act.closed_at"></countdown>
          <div class="home-toolbar-menu">
            <!--<md-button-toggle class="md-primary" md-single>
              <md-button class="md-icon-button md-toggle" @click="viewTable()">
                <md-icon>view_module</md-icon>
              </md-button>
              <md-button class="md-icon-button" @click="viewList()">
                <md-icon>view_list</md-icon>
              </md-button>
            </md-button-toggle>-->
            <md-button v-if="$auth.check('admin')" @click="exportExcel(act)" class="md-icon-button">
              <md-icon>import_export</md-icon>
            </md-button>
            <div>
        </md-toolbar>
        <md-card-content>
          <order v-if="view==='table'" :activity="act" ></order>
        </md-card-content>
      </md-card>
    </md-layout>
    <md-layout md-flex-small="10" md-flex-medium="10" md-flex-large="10" md-flex-xlarge="20"></md-layout>
  </md-layout>
</template>

<script>
  import Order from './Order';
  import _ from 'lodash';
  export default {
    name: 'home',
    data() {
      return {
        view: 'table'
      };
    },
    computed: {
      activities() {
        return this.$store.state.openActivities;
      },
      isLoading() {
        return this.$store.getters.isLoading;
      },
    },
    methods: {
      viewList() {
        this.view = 'list';
      },
      viewTable() {
        this.view = 'table';
      },
      fetchData() {
        this.$store.dispatch('getHomeActivities', 1);
      },
      getClosedDate(closedDate) {
        try {
          const date = new Date(closedDate);
          console.log(date, date.getTime());
          return date.getTime();
        } catch (ex) {
          console.log(ex);
        }
        return new Date().getTime();
      },
      menuOrderList(restaurant) {
        const menuOrderBy = _.reduce(restaurant.menus,(menus,menu)=>{
          const orderBy = _.reduce(restaurant.orders,(orders,order)=>{
            if ( order.menu === menu.name) {
              orders.push(this.lazyUserInfo(order.user_id));
            }
            return orders;
          },[]);
          menu.orderBy = orderBy;
          menus.push(menu);
          return menus;
        },[]);

        menuOrderBy.sort((a,b) => {
          const sa = a.orderBy.length;
          const sb = b.orderBy.length;

           if (sa >= sb ) {
              return -1;
            }
            if (sa < sb) {
              return 1;
            }
            return 0;
        });
        return menuOrderBy;
      },
      lazyUserInfo(userID) {
        const users = this.$store.state.users.filter(user => user.id === userID);
        return users.length > 0 ? users[0].username : "";
      },
      exportExcel(activity) {
        if (!activity || !activity.restaurants) {
          return;
        }
        const data = (!activity.order_items) ? [] : _.reduce(activity.order_items, (_orders, order) => {
          activity.restaurants.forEach((restaurant) => {
            const menus = restaurant.menus;
            menus.forEach((menu) => {
              if (menu.id === order.menu_id) {
                _orders.push([this.lazyUserInfo(order.user_id), menu.name, menu.price]);
              }
            })
          });
          return _orders;
        }, []);
        exportToExcel("activity_orders", _.concat([["ชื่อ", "รายการที่สั่ง", "ราคา"]], data));
        // window.exportToExcel("activity_orders",data);
      },
    },
    created() {
      this.fetchData();
    },
    components: {
      Order,
    },
  };
</script>
<style>
.md-tab-header-container {
  min-width: 250px;
  font-size: 1.5em;
}
.people-info {
  display: flex;
  flex-direction: row;
}
.home-toolbar-menu {
  display: flex;
  flex-direction: row;
}
.menus {
  width: 100%;
}
.activity-button-bar{
  text-align: left;
}

.people-list{
  display:flex;
  flex-direction: column;
}
.people-list-container {
  display: flex;
  flex-direction: row;
}
</style>
<style scoped>
.activity-card {
  width: 100%;
  margin-bottom: 20px;
}
.gutter {
  flex: 1;
}
</style>
