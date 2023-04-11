<template>
  <div class="filter_history">
    <h3>Фильтр</h3>
    <div class="filter_history__inputs">
      <div class="select_input_block">
        <p>Сортировать список</p>
        <select v-model="filterData.sortList">
          <option disabled>Сортировать список</option>
          <option
            value="ascending"
            selected
            @click="filterSortList('sortList')"
          >
            По возрастанию
          </option>
          <option value="descending" @click="filterSortList('sortList')">
            По убыванию
          </option>
        </select>
      </div>
      <div class="select_input_block">
        <p>Сортировать по дате</p>
        <select v-model="filterData.sortListByDate">
          <option disabled>Сортировать по дате</option>
          <option
            value="ascending"
            selected
            @click="filterSortList('sortListByDate')"
          >
            По возрастанию
          </option>
          <option value="descending" @click="filterSortList('sortListByDate')">
            По убыванию
          </option>
        </select>
      </div>
    </div>
  </div>
</template>

<script>
import { mapActions } from "vuex";
export default {
  name: "FilterVue",
  props: {
    name: String,
  },
  data() {
    return {
      filterData: {
        sortList: "ascending",
        sortListByDate: "ascending",
        nameList: "",
      },
    };
  },
  methods: {
    ...mapActions(["sortList"]),

    async filterSortList(key) {
      Object.keys(this.filterData).forEach((k) => {
        if (k === "nameList") this.filterData.nameList = this.name;
        if (key !== k) {
          this.filterData[k] = "ascending";
        }
        return;
      });

      await this.sortList(this.filterData);
    },
  },
};
</script>

<style scoped>
.filter_history {
  padding-bottom: 20px;
}

.filter_history__inputs {
  display: flex;
  flex-wrap: wrap;
}

.select_input_block {
  display: flex;
  flex-direction: column;
  margin-right: 10px;
}

.select_input_block p {
  margin: 0;
}

.select_input_block select {
  border: 2px solid #eee;
  padding: 5px;
  min-width: 150px;
  height: 30px;
}
</style>