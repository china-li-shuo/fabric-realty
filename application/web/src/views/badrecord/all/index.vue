<template>
   <div>
      <el-table :data="tableData" style="width: 100%">
         <!-- el-table-column components -->
         <el-table-column prop="accountId" label="账户ID" width="180"></el-table-column>
         <el-table-column prop="userName" label="用户名" width="180"></el-table-column>
         <el-table-column prop="balance" label="余额"></el-table-column>
         <el-table-column label="操作">
            <!-- slot-scope="scope" -->
            <template slot-scope="scope">
               <el-button size="mini" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
               <el-popover placement="right" width="400" trigger="click">
                  <el-table :data="gridData">
                     <el-table-column width="100" property="name" label="姓名"></el-table-column>
                     <el-table-column width="150" property="id_card" label="身份证号"></el-table-column>
                     <el-table-column width="300" property="is_lock" label="是否有不良记录">
                        <template slot-scope="scope">
                           {{ scope.row.is_lock == 'true' ? '是' : '否' }}
                        </template>
                     </el-table-column>
                  </el-table>
                  <el-button size="mini" type="danger" slot="reference" @click="handleBadRecordDetail(scope.$index, scope.row)">不良记录详情</el-button>
               </el-popover>
            </template>
         </el-table-column>
      </el-table>


   </div>
</template>

<script>
    import { queryAccountList } from '@/api/account'
    import { queryBadRecordListByIdCard } from '@/api/badRecord'

    export default {
        name: 'BadRecordAll',
        data() {
           return {
              tableData: [],
              gridData: []
           }
        },
       //生命周期，页面创建事件
    created() {
       queryAccountList().then(response => {
          console.log(response)
          if (response !== null) {
             this.tableData = response
          }
          this.loading = false
       }).catch(_ => {
          this.loading = false
       })
    },
        methods: {
           handleEdit(index, row) {
              console.log(index, row);
           },
           handleBadRecordDetail(index, row) {

              let data  = {
                 name:row.userName,
                 id_card:row.accountId,
              }

              queryBadRecordListByIdCard(data).then(response => {
                 console.log(response)
                 if (response !== null) {
                    this.gridData = response
                 }
                 this.loading = false
              }).catch(_ => {
                 this.loading = false
              })
           }
        }
    }

</script>

<style>
</style>
