<template>
  <div>
    <el-container style="margin: 15px 20px 20px 20px">
      <el-main>
        <el-form :inline="true">
          <el-form-item label="UID">
            <el-input v-model="listQuery.uid" placeholder="UID"></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="handleQuery">查询</el-button>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="handleClear">清空</el-button>
          </el-form-item>
        </el-form>
        <el-table :data="data" style="width: 100%" :loading="loading" border>
          <el-table-column prop="id" label="ID" align="center">
            <template #default="scope">
              {{ scope.row.id }}
            </template>
          </el-table-column>
          <el-table-column prop="uid" width="120px" label="UID" align="center">
            <template #default="scope">
              {{ scope.row.uid }}
            </template>
          </el-table-column>
          <el-table-column label="收货人姓名" align="center">
            <template #default="scope">
              {{ scope.row.receive_name }}
            </template>
          </el-table-column>
          <el-table-column
            label="收获人性别"
            align="center"
            width="200px"
          >
            <template #default="scope">
              <div v-if="scope.row.receive_sex == 1">
                男
              </div>
              <div v-else >
                女
              </div>
            </template>
          </el-table-column>
          <el-table-column align="center" label="收获手机号码" width="200px">
            <template #default="scope">
              {{ scope.row.receive_phone }}
            </template>
          </el-table-column>
          <el-table-column align="center" label="省份" width="200px">
            <template #default="scope">
              {{ scope.row.province.province_name }}
            </template>
          </el-table-column>
          <el-table-column align="center" label="城市" width="200px">
            <template #default="scope">
              {{ scope.row.city.city_name }}
            </template>
          </el-table-column>
          <el-table-column align="center" label="区县" width="200px">
            <template #default="scope">
              {{ scope.row.district.district_name }}
            </template>
          </el-table-column>
          <el-table-column align="center" label="详细地址" width="200px">
            <template #default="scope">
              {{ scope.row.detail_address }}
            </template>
          </el-table-column>
          <el-table-column
            align="center"
            label="默认地址"
            width="200px"
          >
            <template #default="scope">
              <el-tag v-if="scope.row.is_default == 1" type="success" effect="dark">{{"默认"}}</el-tag>
            </template>
          </el-table-column>
          <el-table-column align="center" label="排序" width="200px">
            <template #default="scope">
              {{ scope.row.sort }}
            </template>
          </el-table-column>
          <el-table-column
            prop="get_type"
            label="标签"
            width="200px"
            align="center"
          >
            <template #default="scope">
              <span v-for="(item, i) in scope.row.tags" :key="i"  style="margin-right:5px">
                <el-tag
                  v-bind:style="{
                    border: item.border_color,
                    'background-color': item.background_color,
                    color: item.text_color,
                  }"
                >
                  {{ item.tag_name }}
                </el-tag>
              </span>
            </template>
          </el-table-column>
        </el-table>
      </el-main>
      <el-footer>
        <div style="text-align: center">
          <el-pagination
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
            :current-page="listQuery.pageNum"
            :page-sizes="pageSize"
            :page-size="listQuery.pageSize"
            layout="total, sizes, prev, pager, next, jumper"
            :total="total"
          >
          </el-pagination>
        </div>
      </el-footer>
    </el-container>
  </div>
</template>

<script>
import { fetchUserAddressList } from "@/api/user_address.js";
import global_ from "@/components/Global/global.vue";
const defaultPageSize = [10, 20, 50, 100, 200];
const defaultListQuery = {
  pageSize: 20,
  pageNum: 1,
  uid: null,
  coupon_id: null,
  use_status: null,
  get_type: null,
};
export default {
  name: "Coupon",
  data() {
    return {
      pageSize: defaultPageSize,
      listQuery: Object.assign({}, defaultListQuery),
      total: null,
      dialogVisiable: false,
      isEdit: false,
      form: null,
      loading: false,
      data: null,
      useStatusOptions: global_.UseStatusOptions,
    };
  },
  created() {
    this.handleListData();
  },
  methods: {
    handleListData() {
      this.loading = true;
      fetchUserAddressList(this.listQuery)
        .then((res) => {
          this.data = res.data.list;
          this.total = res.data.total;
          this.$message({
            message: "查询成功",
            type: "success",
          });
          this.loading = false;
        })
        .catch((res) => {
          console.log(res);
          this.$message({
            message: "查询失败",
            type: "error",
          });
          this.loading = false;
        });
    },
    handleSizeChange(val) {
      this.listQuery.pageSize = val;
      this.listQuery.pageNum = 1;
      this.handleListData();
    },
    handleCurrentChange(val) {
      this.listQuery.pageNum = val;
      this.handleListData();
    },
    handleClear() {
      this.listQuery = Object.assign({}, defaultListQuery);
    },
    handleQuery() {
      this.handleListData();
    },
  },
};
</script>

<style>
.el-table .warning-row {
  background: oldlace;
}

.el-table .success-row {
  background: #f0f9eb;
}
</style>