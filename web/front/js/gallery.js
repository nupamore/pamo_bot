
import Vue from 'vue/dist/vue.js'
import MuseUI from 'muse-ui'
import VueLoadImage from 'vue-load-image'
import 'muse-ui/dist/muse-ui.css'
import '../css/gallery.scss'
import Cookies from 'js-cookie'


// date
const today = new Date()
const beforeDate = new Date(new Date().setDate(today.getDate() - 30))
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
    components: {
        'vue-load-image': VueLoadImage
    },
    data: {
        // guild
        userid: '',
        username: '',
        avatar: '',
        guilds: [],

        // gallery
        uploader: 'All',
        uploaders: [{
            owner: 'All',
            amount: ''
        }],

        images: [],
        imagesTotal: 0,
        imagesCurrentPage: 1,

        startDate: beforeDate,
        endDate: today,
        customDateFormat,
    },
    methods: {
        getImages() {
            const galleryId = Cookies.get('galleryId')
            const start = toDBdate(this.startDate)
            const end = toDBdate(this.endDate)

            fetch(`/images?galleryId=${galleryId}&owner=${this.uploader}&page=1`)
            .then(res => res.json())
            .then(list => {
                this.images = list.images.map(_ => {
                    _.thumb = toThumb(_.ORIGIN_URL) + '?width=400&height=225'
                    return _
                })
                this.imagesTotal = list.total
                this.imagesCurrentPage = 1
            })
        },
        getImagesPage() {
            const galleryId = Cookies.get('galleryId')
            const start = toDBdate(this.startDate)
            const end = toDBdate(this.endDate)

            fetch(`/images?galleryId=${galleryId}&owner=${this.uploader}&page=${this.imagesCurrentPage}`)
            .then(res => res.json())
            .then(list => {
                this.images = list.images.map(_ => {
                    _.thumb = toThumb(_.ORIGIN_URL) + '?width=400&height=225'
                    return _
                })
            })
        },
        getUploaders() {
            const galleryId = Cookies.get('galleryId')
            fetch(`/uploaders?galleryId=${galleryId}`)
            .then(res => res.json())
            .then(list => {
                this.uploaders[0].amount = list.reduce((p, n) => p + n.amount, 0)
                this.uploaders.push(...list)
            })
        },
        getProfile() {
            fetch(`/profile`)
            .then(res => res.json())
            .then(profile => {
                this.userid = profile.id
                this.username = profile.username
                this.discriminator = profile.discriminator
                this.avatar = profile.avatar
                this.guilds = profile.guilds.sort((x, y) => (x.hasBot == y.hasBot) ? 0 : x ? -1 : 1)
            })
        },
        toGallery(guild) {
            Cookies.set('galleryId', guild.id)
            location.href = '/gallery.html'
        }
    },
})