<template>
   <div class="app-container" style="width: 400px;">
      <el-form ref="ruleForm" v-loading="loading" :model="ruleForm" :rules="rules" label-width="120px">
         <el-form-item label="姓名" prop="name">
            <el-input v-model="ruleForm.name" />
         </el-form-item>
         <el-form-item label="身份证号" prop="idCard">
            <el-input v-model="ruleForm.idCard" />
         </el-form-item>
         <el-form-item label="是否有不良记录" prop="proprietor">
            <el-select v-model="ruleForm.isLock" placeholder="请选择…" @change="selectGet">
               <el-option :label="'是'" :value="true"></el-option>
               <el-option :label="'否'" :value="false"></el-option>
            </el-select>
         </el-form-item>
         <el-form-item>
            <el-button type="primary" @click="submitForm('ruleForm')">立即创建</el-button>
            <el-button @click="resetForm('ruleForm')">重置</el-button>
         </el-form-item>
      </el-form>
   </div>
</template>


<script>
   import { mapGetters } from 'vuex'
   import { queryAccountList } from '@/api/account'
   import { badRecordAdd } from '@/api/badRecord'

   export default {
      name: 'BadRecordAdd',
      data() {
         return {
            ruleForm: {
               name: '',
               idCard: '',
               isLock: '',
            },
            accountList: [],
            rules: {
               name: [
                  { required: true, trigger: 'blur' }
               ],
               idCard: [
                  { required: true, trigger: 'blur' }
               ],
               isLock: [
                  { required: true, message: '请选择不良记录', trigger: 'change' }
               ],
            },
            loading: false
         }
      },
      methods: {
         submitForm(formName) {
            this.$refs[formName].validate((valid) => {
               if (valid) {
                  this.$confirm('是否立即创建?', '提示', {
                     confirmButtonText: '确定',
                     cancelButtonText: '取消',
                     type: 'success'
                  }).then((action) => {
                     console.log('Action:', action); // Log the action to see if it's being correctly captured
                     if (action === 'confirm') {
                        console.log('Action is confirm'); // Log to verify if the condition is met
                        // Move loading to start of the confirmation
                        this.loading = true;
                        // Correct API call parameters
                        badRecordAdd({
                           name: this.ruleForm.name,
                           id_card: this.ruleForm.idCard,
                           is_lock: `${this.ruleForm.isLock}`,
                        }).then(response => {
                           console.log('Response:', response); // Log the response to see if it's being received
                           this.loading = false;
                           if (response !== null) {
                              this.$message({
                                 type: 'success',
                                 message: '创建成功!'
                              });
                           } else {
                              this.$message({
                                 type: 'error',
                                 message: '创建失败!'
                              });
                           }
                        }).catch(error => {
                           console.error('Error:', error); // Log any errors from the API call
                           this.loading = false;
                        });
                     } else {
                        console.log('Action is not confirm');
                        // This part will execute when the '取消' button is clicked
                        console.log('创建取消');
                     }
                  }).catch(() => {
                     console.log('Action is fail');
                     // This part will execute when the dialog is dismissed
                     console.log('创建取消');
                  });
               } else {
                  return false;
               }
            });
         },


         resetForm(formName) {
            this.$refs[formName].resetFields()
         },
         selectGet(isLock) {
            this.ruleForm.isLock = isLock
         }
      }
   }
</script>

<style scoped>
</style>
