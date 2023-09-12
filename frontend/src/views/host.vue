<template>
    <div v-for="(item, index) in connect">
        <host-box  v-show="select == index" :key="item" :host-id="index" />
    </div>
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
                setTimeout(() => {
                    this.connect[id] = new Date().getTime()
                    this.select = id
                }, 100)
            } else this.select = id
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