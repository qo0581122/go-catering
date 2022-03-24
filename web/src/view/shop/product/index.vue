<template>
  <div>
      <el-container style="margin: 15px 20px 20px 20px">
        <el-main>
            <el-form :inline="true" >
                <el-form-item label="分类名称">
                    <el-input v-model="listQuery.category_name" placeholder="分类名称"></el-input>
                </el-form-item>
                <el-form-item label="状态">
                    <el-select v-model="listQuery.status" placeholder="状态">
                    <el-option label="启用" value="1"></el-option>
                    <el-option label="禁用" value="2"></el-option>
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
            <el-table
            :data="data"
            style="width: 100%" :loading="loading" border >
                <el-table-column
                    prop="id"
                    label="ID"
                    >
                </el-table-column>
                <el-table-column
                    label="店铺ID"
                    >
                    <template #default>
                        {{scope.row.shop.id}}
                    </template>
                </el-table-column>
               <el-table-column
                    label="商铺名称"
                    >
                    <template #default>
                        {{scope.row.shop.shop_name}}
                    </template>
                </el-table-column>
                <el-table-column
                    label="商铺地址"
                    >
                    <template #default>
                        {{scope.row.shop.shop_detail_address}}
                    </template>
                </el-table-column>
                <el-table-column
                    label="商铺营业状态"
                    >
                    <template #default="scope">
                        <el-tag
                            v-if="scope.row.business_flag == 1"
                            type="success"
                            effect="dark">
                           营业中
                        </el-tag>
                        <el-tag
                            v-else
                            type="danger"
                            effect="dark">
                            未营业
                        </el-tag>
                    </template>
                </el-table-column>
               <el-table-column
                    label="产品ID"
                    >
                    <template #default>
                        {{scope.row.product.id}}
                    </template>
                </el-table-column>
                <el-table-column
                    label="产品名称"
                    >
                    <template #default>
                        {{scope.row.product.product_name}}
                    </template>
                </el-table-column>
                <el-table-column prop="url" label="产品展示图" width="200px">
                    <template #default="scope">
                    {{ scope.row.product.url }}
                    </template>
                </el-table-column>
                <el-table-column prop="url" label="上架状态" width="200px">
                    <template #default="scope">
                        <e-tag type="success" v-if="scope.row.status == 1">上架</e-tag>
                        <e-tag type="danger" v-if="scope.row.status == 2">下架</e-tag>
                    </template>
                </el-table-column>
                <el-table-column
                    fixed="right"
                    label="操作"
                    width="200px">
                    <template #default="scope">
                        <el-button  v-if="scope.row.status == 1" @click="handleUpdate(scope.row)" type="success" size="small">下架</el-button>
                        <el-button  v-if="scope.row.status == 2" @click="handleUpdate(scope.row)" type="success" size="small">上架</el-button>
                        <el-button type="success" size="small" @click="handleDelete(scope.row)">删除</el-button>
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
                    :total="total">
                </el-pagination>
            </div>
        </el-footer>
    </el-container>
    <dform :dialogVisiable="dialogVisiable" :is-edit="isEdit" :form="form" @visableChange="handleCloseDialog" @submit="handleSubmit"></dform>
  </div>
</template>

<script>
import Dform from './form.vue'
import {fetchShopCategoryPage, createShopCategory, updateShopCategory, deleteShopCategory} from '@/api/shop_category.js'
const defaultPageSize = [10, 20, 50, 100, 200]
const defaultListQuery = {
    pageSize: 20,
    pageNum: 1,
    category_name: '',
    status: null
}
const defaultForm = {
    id: null,
    category_name: '',
    status: true
}
export default {
    components: { Dform },
    name: 'Provinces',
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
            fetchShopCategoryPage(this.listQuery).then(res => {
                this.data = res.data.list
                this.total = res.data.total
                this.$message({
                    message: '查询成功',
                    type: 'success'
                });
                this.loading = false
            }).catch(res => {
                console.log(res)
                this.$message({
                    message: '查询失败',
                    type: 'error'
                });
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
            let form = Object.assign({}, data)
            console.log(data)
            form.status = form.status == true ? 1 : 2
            if (this.isEdit) {
                //更新
                updateShopCategory(form).then(res => {
                    this.$message({
                        message: res.msg,
                        type: 'success'
                    });
                    this.handleListData()
                }).catch(res => {
                    console.log(res)
                    this.$message({
                        message: res.msg,
                        type: 'error'
                    });
                })
            }else {
                //增加
                createShopCategory(form).then( res=> {
                    this.$message({
                        message: res.msg,
                        type: 'success'
                    });
                     this.handleListData()
                }).catch(res => {
                    this.$message({
                        message: res.msg,
                        type: 'error'
                    });
                })
            }
            this.handleCloseDialog()
        },
        handleDelete(data) {

        }
    }
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