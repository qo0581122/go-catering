<template>
  <div>
    <el-container style="margin: 15px 20px 20px 20px">
      <el-main>
        <el-form :inline="true">
          <el-form-item label="Uid">
            <el-input v-model="listQuery.uid" placeholder="Uid"></el-input>
          </el-form-item>
          <el-form-item label="升级前等级id">
            <el-input
              v-model="listQuery.before_level_id"
              placeholder="升级前等级id"
            ></el-input>
          </el-form-item>
          <el-form-item label="升级后等级id">
            <el-input
              v-model="listQuery.after_level_id"
              placeholder="升级后等级id"
            ></el-input>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="handleQuery">查询</el-button>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="handleClear">清空</el-button>
          </el-form-item>
        </el-form>

        <el-table :data="data" style="width: 100%" :loading="loading" border>
          <el-table-column
            prop="id"
            label="id"
            align="center"
            width="200px"
          >
          </el-table-column>
          <el-table-column
            prop="uid"
            label="Uid"
            align="center"
            width="200px"
          >
          </el-table-column>
          <el-table-column
            prop="before_level_id"
            label="升级前等级id"
            align="center"
            width="200px"
          >
          </el-table-column>
          <el-table-column
            prop="before_level"
            label="升级前等级"
            align="center"
          >
             <template #default="scope">
              Lv{{scope.row.before_level}}
            </template>
          </el-table-column>
          <el-table-column
            prop="before_level_name"
            label="升级前等级名称"
            align="center"
            width="200px"
          >
          </el-table-column>
          <el-table-column
            prop="after_level_id"
            label="升级后等级id"
            align="center"
            width="200px"
          >
          </el-table-column>
          <el-table-column
            prop="after_level"
            label="升级后等级"
            align="center"
          >
             <template #default="scope">
              Lv{{scope.row.after_level}}
            </template>
          </el-table-column>
          <el-table-column
            prop="after_level_name"
            label="升级后等级名称"
            align="center"
            width="200px"
          >
          </el-table-column>
          <el-table-column
            prop="up_time"
            align="center"
            label="升级时间"
            width="200px"
          >
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
import { fetchUserVipLevelLogList } from "@/api/user_vip_level.js";
const defaultPageSize = [10, 20, 50, 100, 200];
const defaultListQuery = {
  pageSize: 20,
  pageNum: 1,
  uid: null,
  before_level_id: null,
  after_level_id: null,
};
export default {
  name: "Voucher",
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
    };
  },
  created() {
    this.handleListData();
  },
  methods: {
    handleListData() {
      this.loading = true;
      fetchUserVipLevelLogList(this.listQuery)
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