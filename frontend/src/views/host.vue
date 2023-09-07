<template>
    <host-box v-for="(item, index) in connect" v-show="select == index" :host-id="index" />
</template>
  
<script>
import HostBox from "../components/hostBox.vue";

export default {
    name: "Host",
    components: { HostBox },
    data: () => ({
        select: null,
        connect: {}
    }),
    methods: {
        openConnect(id) {
            if (this.connect[id] == undefined) {
                this.connect[id] = new Date().getTime()
                this.select = id
            } else this.select = id
            console.log('Show host -> '+this.select)
        }
    },
    mounted() {
        this.openConnect(this.$route.params.id)
    },
    watch: {
        $route: {
            handler(to) {
                this.openConnect(to.params.id)
            },
            deep: true,
        },
    },
};
</script>
<style scoped></style>