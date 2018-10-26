
import Vue from 'vue/dist/vue.js'
import MuseUI from 'muse-ui'
import 'muse-ui/dist/muse-ui.css'
import './gallery.scss'

// date
const today = new Date()
const beforeDate = new Date(new Date().setDate(today.getDate() - 10))
const dayList = ['日', '月', '火', '水', '木', '金', '土']
const customDateFormat = {
    formatDisplay(date) {
        return `${date.getMonth() + 1}月 ${date.getDate()}日, ${dayList[date.getDay()]}`
    },
    formatMonth(date) {
        return `${date.getFullYear()}年 ${date.getMonth() + 1}月`
    },
    getWeekDayArray(firstDayOfWeek) {
        const beforeArray = []
        const afterArray = []
        dayList.forEach((day, index) => {
            if (index < firstDayOfWeek) {
                afterArray.push(day)
            } else {
                beforeArray.push(day)
            }
        })
        return beforeArray.concat(afterArray)
    },
    getMonthList() {
        return [...Array(11).keys()].map(_ => `${_ + 1}月`)
    }
}

function toDBdate(date) {
    return `${date.getFullYear()}-${date.getMonth() + 1}-${date.getDate()}`
}
function toThumb(url) {
    return url.replace('cdn.discordapp.com', 'media.discordapp.net')
}
function toOrigin(url) {
    return url.replace('media.discordapp.net', 'cdn.discordapp.com')
}

// app
Vue.use(MuseUI)
const app = new Vue({
    el: '#app',
    data: {
        uploader: 'All',
        users: [],

        images: [],

        startDate: beforeDate,
        endDate: today,
        customDateFormat,
    },
    methods: {
        getImages() {
            const start = toDBdate(this.startDate)
            const end = toDBdate(this.endDate)

            fetch(`/images?startDate=${start}&endDate=${end}`)
            .then(res => res.json())
            .then(list => {
                this.images = list.map(_ => toThumb(_.ORIGIN_URL) + '?width=400&height=225')
            })
        }
    },
})

app.getImages()