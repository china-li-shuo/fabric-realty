<template>
   <div>不良记录列表</div>
</template>

<script>
    import { mapGetters } from 'vuex'
    import { queryAccountList } from '@/api/account'
    import { queryRealEstateList } from '@/api/realEstate'
    import { createSelling } from '@/api/selling'
    import { createDonating } from '@/api/donating'

    export default {
        name: 'BadRecordAll',
        data() {
           return {}
        },
        methods: {
        }
    }

</script>

<style>
</style>
