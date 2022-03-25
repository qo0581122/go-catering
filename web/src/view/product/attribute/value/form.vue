<template>
    <div>
    <el-dialog
        :title="dialogTitle"
        center
        v-model="dialogVisiable"
        width="500px"
        :before-close="handleClose">
        <el-form ref="dialogForm" :model="dialogForm" label-width="80px" label-position="right" :rules="formRules">
            <div v-if="isEdit">
                <el-form-item label="id">
                    <el-input v-model="dialogForm.id"  :disabled="isEdit"></el-input>
                </el-form-item>
            </div>
            <el-form-item label="属性" prop="attribute_id">
                <el-select v-model="dialogForm.attribute_id" placeholder="请选择">
                    <el-option
                    v-for="item in attributes"
                    :key="item.id"
                    :label="item.attribute_name"
                    :value="item.id">
                    </el-option>
                </el-select>
            </el-form-item>
            <el-form-item label="属性值" prop="attribute_value">
                <el-input v-model="dialogForm.attribute_value"></el-input>
            </el-form-item>
            <el-form-item label="状态">
                  <el-switch v-model="dialogForm.status"></el-switch>
            </el-form-item>
        </el-form>
        <template #footer>
        <span class="dialog-footer">
          <el-button type="primary" @click="handleSubmit('dialogForm')"
            >提交</el-button
          >
          <el-button @click="handleClose">取消</el-button>
        </span>
      </template>
    </el-dialog>
    </div>
</template>

<script>
import {fetchProductAttributeAll} from '@/api/product_attribute.js'
const defaultForm = {
    id: "",
    attribute_value: "",
    status: true,
    attribute_id: null,
}
export default {
    name: 'Dform',
    props: {
        dialogVisiable: {
            type: Boolean,
            required: true,
        },
        isEdit: {
            type: Boolean,
            required: true,
            default: false,
        },
        form: {
            type: Object,
            default: Object.assign({}, defaultForm)
        }
    },
    computed: {
        dialogTitle: function() {
            return this.isEdit ? this.defaultDialogTitles[1] : this.defaultDialogTitles[0]
        },
    },
    watch: {
        form() {
            this.dialogForm = Object.assign({}, this.form)
            this.dialogForm.status = this.dialogForm.status == 1 ? true : false
        },
    },
    created() {
        fetchProductAttributeAll().then(res => {
            this.attributes = res.data.list
        })
    },
    data() {
        return {
            defaultDialogTitles: ["创建", "编辑"],
            formRules: {
                attribute_value: [
                    {required: true, message: '请输入属性值', trigger: 'blur'}
                ],
                attribute_id: [
                    {required: true, type: 'number', message: '请选择属性', trigger: 'blur'}
                ]
            },
            dialogForm: Object.assign({}, defaultForm),
            attributes: [],
        }
    },
    methods: {
        handleClose() {
            this.dialogForm = Object.assign({}, defaultForm)
            this.$emit("visableChange", false)
         },
        handleSubmit(formName) {
            this.$refs[formName].validate((valid) => {
            if (valid) {
                this.$emit("submit", this.dialogForm)
            } else {
                console.log('error submit!!');
                return false;
            }
            });
        },
    }
}
</script>

<style>

</style>