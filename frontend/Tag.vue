<template>
  <tbody v-if="stunden.length != 0" style="border-top: 2px solid #888888">
    <tr :class="{ weekend: tag.weekend}" v-for="(zeit, index) in stunden" :key="zeit.ID">
      <td style="letter-spacing: 1px; vertical-align: middle;" v-if="index == 0" :rowspan="stunden.length + 1">
        <span style="width: 30px; display: inline-block;">{{ tag.weekday }}.</span>
        {{ tag.day }}.{{ tag.month }}.{{tag.year}}
      </td>
      <td>{{ zeit.Beginn}} Uhr</td>
      <td>{{ zeit.Ende }} Uhr</td>
      <td>{{ zeit.Pause }} h</td>
      <td>{{ ((+zeit.Ende.split(':')[0] * 60) + zeit.Ende.split(':')[1] - (+zeit.Beginn.split(':')[0] * 60) - +zeit.Beginn.split(':')[1] - (+zeit.Pause.split(':')[0] * 60) - +zeit.Pause.split(':')[1]) /60 }} h</td>
      <td class="no-print">
          <i class="ri-delete-bin-line"></i>
      </td>
    </tr>
    <tr :class="{ weekend: tag.weekend}" class="no-print">
      <td v-show="edit"><input type="time" v-model="beginn"> Uhr</td>
      <td v-show="edit"><input type="time" v-model="ende"> Uhr</td>
      <td v-show="edit"><input type="time" v-model="pause"> Stunden</td>
      <td v-show="edit">{{ isNaN(+ende.split(':')[0] + (+ende.split(':')[1] / 60) - beginn.split(':')[0] - (+beginn.split(':')[1] / 60) - pause.split(':')[0] - (+pause.split(':')[1] / 60)) ? "": +ende.split(':')[0] + (+ende.split(':')[1] / 60) - beginn.split(':')[0] - (+beginn.split(':')[1] / 60) - pause.split(':')[0] - (+pause.split(':')[1] / 60) }}</td>
      <td v-show="edit" class="no-print">
          <i class="ri-save-2-line" @click="save()"></i>
          <i class="ri-close-circle-line" @click="edit = false"></i>
      </td>
      <td v-show="!edit"></td>
      <td v-show="!edit"></td>
      <td v-show="!edit"></td>
      <td v-show="!edit"></td>
      <td v-show="!edit" @click="edit = true" class="no-print">
        <i class="ri-add-circle-line" title="Zeit erfassen"></i>
      </td>
    </tr>
  </tbody>
  <tbody v-else style="border-top: 2px solid #888888">
    <tr :class="{ weekend: tag.weekend}">
      <td style="letter-spacing: 1px;">
        <span style="width: 30px; display: inline-block;">{{ tag.weekday }}.</span>
        {{ tag.day }}.{{ tag.month }}.{{tag.year}}
      </td>
      <td v-show="edit"><input type="time" v-model="beginn"> Uhr</td>
      <td v-show="edit"><input type="time" v-model="ende"> Uhr</td>
      <td v-show="edit"><input type="time" v-model="pause"> Stunden</td>
      <td v-show="edit">{{ isNaN(+ende.split(':')[0] + (+ende.split(':')[1] / 60) - beginn.split(':')[0] - (+beginn.split(':')[1] / 60) - pause.split(':')[0] - (+pause.split(':')[1] / 60)) ? "": +ende.split(':')[0] + (+ende.split(':')[1] / 60) - beginn.split(':')[0] - (+beginn.split(':')[1] / 60) - pause.split(':')[0] - (+pause.split(':')[1] / 60) }}</td>
      <td v-show="edit" class="no-print">
          <i class="ri-save-2-line" @click="save()"></i>
          <i class="ri-close-circle-line" @click="edit = false"></i>
      </td>
      <td v-show="!edit"></td>
      <td v-show="!edit"></td>
      <td v-show="!edit"></td>
      <td v-show="!edit"></td>
      <td v-show="!edit" @click="edit = true" class="no-print">
        <i class="ri-add-circle-line no-print" title="Zeit erfassen"></i>
      </td>
    </tr>
  </tbody>
</template>

<script>
export default {
  props: {
    tag: Object
  },
  data() {
      setTimeout(() => this.load());
    return {
      edit: false,
      beginn: "07:00",
      ende: "16:00",
      pause: "00:30",
      stunden: []
    };
  },
  methods: {
      async save() {
            console.log("Speichere Arbeitszeit vom " + this.tag.day + "." + this.tag.month + "." + this.tag.year + 
                "\nZeit: " + this.beginn + " Uhr - " + this.ende + " Uhr mit " + this.pause + " h Pause.");
            let r = await api.POST("/" + encodeURIComponent(this.tag.year) + "/" + encodeURIComponent(this.tag.month) + "/" + encodeURIComponent(this.tag.day), {
                "Beginn": this.beginn,
                "Ende": this.ende,
                "Pause": this.pause
            });
            this.edit = false;
            this.load()
        },
        async load() {
            console.log("Lade Daten für " + this.tag.day + "." + this.tag.month + "." + this.tag.year);
            let r = await api.GET("/" + encodeURIComponent(this.tag.year) + "/" + encodeURIComponent(this.tag.month) + "/" + encodeURIComponent(this.tag.day));
            this.stunden = (r.ok ? r.content : [])
            this.stunden.length != 0 ? console.log(this.stunden) : "";
        }
  }
};
</script>