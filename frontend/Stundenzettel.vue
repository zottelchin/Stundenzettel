<template>
    <div class="container">
        <h1 class="title is-1 has-text-centered no-print" style="display: flex; justify-content: center;align-items: center; margin-top: 1rem">
            <img src="https://img.icons8.com/dusk/58/000000/clock.png" style="margin-right: 1rem">
            Stundenzettel
        </h1>
        <h1 class="print">
            Stundenzettel: {{ monat(this.$route.params.month) }} {{ this.$route.params.year }}
        </h1>
        <div class="level no-print">
            <div class="level-item level-left">
                <router-link class="button" :to="prev()">
                    vorheriger Monat
                </router-link>
            </div>
            <div class="level-item">
                <h3>{{ monat($route.params.month) }} {{ $route.params.year }}</h3>
            </div>
            <div class="level-item level-right">
                <router-link class="button" :to="next()">
                    nächster Monat
                </router-link>
            </div>
        </div>
        <span style="float: right; font-size: 6pt !important; margin-bottom: 7px;" class="print">Stand: {{ (new Date()).toLocaleString() }}</span>
        <table class="table is-fullwidth">
            <thead>
                <tr>
                    <th width="170px">Datum</th>
                    <th>Arbeitsbegin</th>
                    <th>Arbeitsende</th>
                    <th>Pause</th>
                    <th>Arbeitszeit</th>
                    <th width="80px" class="no-print"></th>
                </tr>
            </thead>
                <tag v-for="day in tageImMonat" :tag="day" :key="day.day" />
        </table>
    </div>
</template>

<script>
import tag from "./Tag.vue";

export default {
    components: {
        tag
    },
    methods: {
        monat(zahl) {
            var month = new Array();
            month["1"] = "Januar";
            month["2"] = "Februar";
            month["3"] = "März";
            month["4"] = "April";
            month["5"] = "Mai";
            month["6"] = "Juni";
            month["7"] = "Juli";
            month["8"] = "August";
            month["9"] = "September";
            month["10"] = "Oktober";
            month["11"] = "November";
            month["12"] = "Dezember";
            return month[zahl];
        },
        wochentag(num) {
            var tag = ["So", "Mo", "Di", "Mi", "Do", "Fr", "Sa"];
            return tag[num];
        },
        async save(day) {
            console.log("save.." +day);
            let r = await api.POST("/" + this.$route.params.year + "/" + this.$route.params.month + "/" + day, this.stunden[day])
        },
        prev() {
            if (this.$route.params.month == "1"){
                return "/" + (+this.$route.params.year - 1) + "/12"
            }else{
                return "/" + this.$route.params.year + "/" + (this.$route.params.month - 1)
            }
        },
        next() {
            if (this.$route.params.month == "12"){
                return "/" + (Number(this.$route.params.year) + 1) + "/1"
            }else{
                return "/" + this.$route.params.year + "/" + (+this.$route.params.month + 1)
            }
        }
    },
    computed: {
        tageImMonat() {
            let days = new Array();
            let max = new Date(this.$route.params.year, this.$route.params.month, 0).getDate()
            for (var i = 1; i <= max; i++) {
                days.push({
                    "day": i, 
                    "month": this.$route.params.month, 
                    "year": this.$route.params.year,
                    "weekday": this.wochentag(new Date(this.$route.params.year, this.$route.params.month-1, i).getDay()),
                    "weekend": [0,6].indexOf(new Date(this.$route.params.year, this.$route.params.month-1, i).getDay()) > -1 ? true : false
                });
            }
            return days;
        }
    }
}
</script>