<template>
  <div>
    <el-container style="margin: 15px 20px 20px 20px">
      <el-main>
        <el-form :inline="true">
          <el-form-item label="区域名称">
            <el-input
              v-model="listQuery.district_name"
              placeholder="区域名称"
            />
          </el-form-item>
          <el-form-item label="状态">
            <el-select v-model="listQuery.status" placeholder="状态">
              <el-option label="启用" value="1" />
              <el-option label="禁用" value="2" />
            </el-select>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="handleQuery">查询</el-button>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="handleClear">清空</el-button>
          </el-form-item>
          <el-form-item>
            <el-button type="primary" @click="handleAdd">添加</el-button>
          </el-form-item>
        </el-form>
        <el-table :data="data" style="width: 100%" :loading="loading" border>
          <el-table-column prop="id" label="id" />
          <el-table-column prop="shop_name" label="店铺名称" width="100px" />
          <el-table-column prop="category_name" label="店铺分类" width="100px" />
          <el-table-column prop="shop_address" label="店铺地址" width="180px" />
          <el-table-column
            prop="shop_detail_address"
            label="店铺详细地址"
            width="250px"
          />
          <el-table-column prop="contact_number" label="联系电话" width="120px" />
          <el-table-column prop="business_flag" label="营业">
            <template #default="scope">
              <el-tag
                v-if="scope.row.business_flag == 1"
                type="success"
                effect="dark"
              >
                营业中
              </el-tag>
              <el-tag v-else type="danger" effect="dark"> 未营业 </el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="status" label="状态">
            <template #default="scope">
              <el-tag v-if="scope.row.status == 1" type="success" effect="dark">
                启用
              </el-tag>
              <el-tag v-else type="danger" effect="dark"> 禁用 </el-tag>
            </template>
          </el-table-column>
          <el-table-column prop="created_time" label="创建时间" width="200px" />
          <el-table-column prop="updated_time" label="更新时间" width="200px" />
          <el-table-column fixed="right" label="操作" width="200px">
            <template #default="scope">
              <el-button
                type="success"
                size="small"
                @click="handleUpdate(scope.row)"
              >编辑</el-button>
              <el-button
                type="success"
                size="small"
                @click="handleDelete(scope.row)"
              >删除</el-button>
            </template>
          </el-table-column>
        </el-table>
      </el-main>
      <el-footer>
        <div style="text-align: center">
          <el-pagination
            :current-page="listQuery.pageNum"
            :page-sizes="pageSize"
            :page-size="listQuery.pageSize"
            layout="total, sizes, prev, pager, next, jumper"
            :total="total"
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
          />
        </div>
      </el-footer>
    </el-container>
    <dform
      :dialog-visiable="dialogVisiable"
      :is-edit="isEdit"
      :form="form"
      @visableChange="handleCloseDialog"
      @submit="handleSubmit"
    />
  </div>
</template>

<script>
import Dform from './form.vue'
import {
  fetchShopPage,
  createShop,
  updateShop,
  deleteShop,
} from '@/api/shop.js'
const defaultPageSize = [10, 20, 50, 100, 200]
const defaultListQuery = {
  pageSize: 20,
  pageNum: 1,
  district_name: '',
  status: null,
}
const defaultForm = {
  id: '',
  shop_name: '',
  shop_address: '',
  logintude: null,
  latitude: null,
  contact_number: null,
  business_flag: 2,
  status: 2,
  category_id: null,
  district_id: null,
  province_id: null,
  city_id: null,
}
export default {
  name: 'Districts',
  components: { Dform },
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
    }
  },
  created() {
    this.handleListData()
  },
  methods: {
    handleListData() {
      this.loading = true
      fetchShopPage(this.listQuery)
        .then((res) => {
          this.data = res.data.list
          this.total = res.data.total
          this.$message({
            message: '查询成功',
            type: 'success',
          })
          this.loading = false
        })
        .catch((res) => {
          console.log(res)
          this.$message({
            message: '查询失败',
            type: 'error',
          })
          this.loading = false
        })
    },
    handleCloseDialog(value) {
      this.dialogVisiable = false
      this.form = Object.assign({}, defaultForm)
    },
    handleUpdate(data) {
      this.dialogVisiable = true
      this.isEdit = true
      this.form = data
    },
    handleAdd() {
      this.dialogVisiable = true
      this.isEdit = false
      this.form = Object.assign({}, defaultForm)
    },
    handleSizeChange(val) {
      this.listQuery.pageSize = val
      this.listQuery.pageNum = 1
      this.handleListData()
    },
    handleCurrentChange(val) {
      this.listQuery.pageNum = val
      this.handleListData()
    },
    handleClear() {
      this.listQuery = Object.assign({}, defaultListQuery)
    },
    handleQuery() {
      this.handleListData()
    },
    handleSubmit(data) {
      const form = Object.assign({}, data)
      form.status = Number(form.status)
      form.business_flag = Number(form.business_flag)
      form.contact_number = String(form.contact_number)
      if (this.isEdit) {
        // 更新
        updateShop(form)
          .then((res) => {
            this.$message({
              message: res.msg,
              type: 'success',
            })
            this.handleListData()
          })
          .catch((res) => {
            console.log(res)
            this.$message({
              message: res.msg,
              type: 'error',
            })
          })
      } else {
        // 增加
        createShop(form)
          .then((res) => {
            this.$message({
              message: res.msg,
              type: 'success',
            })
            this.handleListData()
          })
          .catch((res) => {
            this.$message({
              message: res.msg,
              type: 'error',
            })
          })
      }
      this.handleCloseDialog()
    },
    handleDelete(data) {
      this.$confirm('此操作将永久删除该记录, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
        center: true,
      })
        .then(() => {
          deleteShop(data.id)
            .then((res) => {
              this.$message({
                message: res.data,
                type: 'success',
              })
              this.handleListData()
            })
            .catch((res) => {
              this.$message({
                message: res.data,
                type: 'error',
              })
            })
        })
        .catch(() => {
          this.$message({
            type: 'info',
            message: '已取消删除',
          })
        })
    },
  },
}
</script>

<style>
.el-table .warning-row {
  background: oldlace;
}

.el-table .success-row {
  background: #f0f9eb;
}
</style>
